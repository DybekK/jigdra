package model

import (
	"context"
	"errors"
	"regexp"
	"strings"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       primitive.ObjectID `bson:"_id" json:"id"`
	Username *string            `json:"username" validate:"required"`
	Email    *string            `json:"email" validate:"required"`
	Password *string            `json:"password" validate:"required"`
}

type LoginUser struct {
	Email    *string `json:"email" validate: "required"`
	Password *string `json:"password" validate:"required"`
}

type GetUserStruct struct {
	Id       primitive.ObjectID `bson:"_id" json:"id"`
	Username *string            `json:"username" validate:"required"`
	Email    *string            `json:"email" validate:"required"`
}

func (d *Database) ValidateEmail(email string) error {
	match := email_regex.MatchString(email)
	if match {
		return nil
	} else {
		return errors.New("invalid email")
	}
}

func (d *Database) GetUser(login *LoginUser, ctx context.Context) (*User, error) {
	coll := d.GetCollection("users")

	res := coll.FindOne(ctx, bson.M{
		"email": login.Email,
	})

	var result User
	decode_err := res.Decode(&result)
	if decode_err != nil {
		return nil, errors.New("decode error")
	}
	if verifyPassword(*login.Password, *result.Password) {
		return &result, nil
	} else {
		return nil, errors.New("invalid password")
	}
}

func (d *Database) CreateUser(req_user *User, ctx context.Context) (*mongo.InsertOneResult, error) {
	var user User
	user.Id = primitive.NewObjectID()
	email_err := d.ValidateEmail(*req_user.Email)
	if email_err != nil {
		return nil, email_err
	}
	user.Username = req_user.Username
	user.Email = req_user.Email

	var hash, hash_error = hashPassword(*req_user.Password)
	if hash_error != nil {
		return nil, errors.New("hashing error")
	}

	user.Password = &hash
	result, insertError := d.GetCollection("users").InsertOne(ctx, user)
	if insertError != nil {
		matched, _ := regexp.MatchString(`duplicate key`, insertError.Error())
		if matched {
			return nil, errors.New("409")

		} else {
			return nil, errors.New(insertError.Error())
		}
	}
	return result, nil
}

var validate = validator.New()

func (d *Database) GetMiddleWare() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm: "jigdra",
		Key:   []byte("test"),
		//SigningAlgorithm: "RS256",
		SendCookie: true,
		// SecureCookie:   true,
		// CookieHTTPOnly: true,
		// TokenLookup:    "cookie:token",
		// CookieDomain:   "localhost",
		// CookieSameSite: http.SameSiteDefaultMode,
		Timeout:     time.Minute * 2,
		MaxRefresh:  time.Minute * 10,
		IdentityKey: "_id",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					"identityKey": v.Id.Hex(),
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			objid, _ := primitive.ObjectIDFromHex(claims["_id"].(string))
			return &User{
				Id: objid,
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			q := c.Query("redirect")
			var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
			if q != "" {
				objid, err := primitive.ObjectIDFromHex(strings.Split(q, "\"")[1])
				if err != nil {
					return nil, err
				}
				res := d.GetCollection("users").FindOne(ctx, bson.M{"_id": objid})
				var user User
				decode_err := res.Decode(&user)
				if decode_err != nil {
					return nil, decode_err
				}
				defer cancel()
				return user, nil
			} else {
				var req_login LoginUser
				if err := c.BindJSON(&req_login); err != nil {
					return nil, err
				}
				validation_error := validate.Struct(req_login)
				if validation_error != nil {
					return nil, validation_error
				}
				defer cancel()

				user, user_err := Interface.GetUser(&req_login, ctx)
				if user_err != nil {
					return nil, jwt.ErrFailedAuthentication
				}
				return user, nil
			}

		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
	}
}

func (d *Database) GetUserById(id string, ctx context.Context) (*GetUserStruct, error) {
	coll := d.GetCollection("users")
	objId, _ := primitive.ObjectIDFromHex(id)
	res := coll.FindOne(ctx, bson.M{"_id": objId})
	var user GetUserStruct
	decode_err := res.Decode(&user)
	if decode_err != nil {
		return nil, decode_err
	}

	return &user, nil

}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func verifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
