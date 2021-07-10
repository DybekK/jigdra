package model

import (
	"context"
	"errors"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
			var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
			if c.Request.Method == "GET" {
				q := c.Query("redirect")
				if q != "" {
					id, exists := d.VerifyRedirect(c, q)
					if exists != nil {
						return nil, errors.New("failed redirect")
					}
					objid, err := primitive.ObjectIDFromHex(id)
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
				}
			} else if c.Request.Method == "POST" {
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
			return nil, errors.New("unsupported http method")

		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
	}
}
