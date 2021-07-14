# Golang for jigdra

By default this service starts at port 4201.
In order for it to start without crashing you need to have MongoDB running.

### IMPORTANT

In `.env` in root directory there are four variables set for MongoDB.
```
MONGO_INITDB_ROOT_USERNAME=admin
MONGO_INITDB_ROOT_PASSWORD=password
MONGO_INITDB_DATABASE=jigdra
MONGO_HOST=mongodb
```

The first 3 are self explanatory.

The 4th is the name of the host, in this case,
the name of the Docker container in which MongoDB is running. Preferably you should run both `golang` and `mongodb` together with docker-compose.

```bash
# first, build the golang container
docker-compose build golang
# now run the containers
docker-compose up mongodb golang
```

If you want to run `golang` localy. You must either:

- set `mongodb` in your `/etc/hosts` to `localhost`

```bash
# ...
# contents of /etc/hosts
127.0.0.1       mongodb
```

- set `MONGO_HOST` in `.env` to `localhost`
## Available endpoints and HTTP Methods

### General

If you call a wrong method on an entopoint you get `405 Not Allowed`. If you call an endpoint that doesn't exist you get `404 Not Found`


### POST `/v1/register`
This endpoint is for user registration. If it doesn't throw any errors along the way, it redirects with `302` to GET /v1/login.

Example request:

```bash
curl -L \
  --url http://localhost:4201/v1/register \
  --header 'content-type: application/json' \
  --data '{"name": "John","surname": "Doe","username": "john_doe","email": "joemamma@mail.com","password": "verystrongpasswd"}'
```

Result for that request

```json
{ "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjYzMTUyMjUsImlkZW50aXR5a2V5IjoiNjBlZjBkMzhhOWVhMzdlZTNhYjhiNmRiIn0.sWceqUARxwg2dedxd1byBTQIhVJoF7zM6P9QVto3UuU",
"refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjYzMTUyMjV9.US5C2pf9s8jw5pV6q0lUPYD_7Qok0udDf7UI2zvZTl0"
}
```

Both `username` and `email` are unique.

If username already exists in MongoDB it will return `HTTP 400` with:
```json
{"error":"username taken"}
```

If email already exists in MongoDB it will return `HTTP 409` with:
```json
{"error":"email in use"}
```
If you provide wrong request body you will receive `HTTP 400`:
```json
{"error":"EOF"}
```

**Note**: /v1/register first checks the username then checks the email. So if both already exist in database it will return `username taken`.

### GET `/v1/login`

It shouldn't be called on it's own as `/v1/register` automatically redirects to it with neccessary `redirect` URL query paramater that has a unique hex value.

### POST `/v1/login`

Example request

```bash
curl --request POST \
  --url http://localhost:4201/v1/login \
  --header 'content-type: application/json' \
  --data '{"email": "uwa5@mail.com","password": "potezne"}'
```

Response for that request

```json
{           
"access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjYzMTY0NzMsImlkZW50aXR5a2V5IjoiNjBlZGZkM2RmNTRiYzcxNzVmOGU5MDM4In0.1OmlaTq4UuYyQAmEzrwWNntGZyzttpg_WnLHd3e2Xd4",

"refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjYzMTY0NzN9.-3IZ0hlHB2XDMQu-Uw7x9bIURK8D6FSRT8uENMwCkHo"
}
```

If you provide wrong email and/or password you will receive `HTTP 401`:
```json
{"error":"invalid email or password"}
```

If you provide wrong request body you will receive `HTTP 400`:
```json
{"error":"EOF"}
```


### GET `/v1/user/:id`

This endpoint requires you to pass an `id` as a parameter. `id` is the hex value of `ObjectId`.

Example request:
```bash
curl --request GET \
  --url http://localhost:4201/v1/user/60ef0d38a9ea37ee3ab8b6db \
```

Response for that request:
```json
{
  "id": "60ef0d38a9ea37ee3ab8b6db",
  "username": "john_doe",
  "name": "John",
  "surname": "Doe",
  "email": "joemamma@mail.com",
  "genderId": 0,
  "dateOfBirth": ""
}
```

If the server doesn't find anything with provided `id` then it will return `HTTP 400`:
```json
{"error":"user not found"}
```


### POST `/v1/refresh`

Refreshes `access_token`

Example request:
```bash
curl --request POST \
  --url http://localhost:4201/v1/refresh \
  --header 'content-type: application/json' \
  --data '{"refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjYzMjIzNzYsImlkZW50aXR5a2V5IjoiNjBlZGZkM2RmNTRiYzcxNzVmOGU5MDM4In0.OzMWHDSKJRA4N0yuqxfahR0K5QWrAPmWwmXQToTj0yc"}'
```

Response to that request:
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjYzMjA1MTEsImlkZW50aXR5a2V5IjoiNjBlZGZkM2RmNTRiYzcxNzVmOGU5MDM4In0.cSW-_jqYozIWfQXtuJUboGRVN4nO7VKiiu4JbtSuvvg",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjYzMjA1MTF9.Cume_ciO21wyobQDr1fkLCa4PgOQ5Yy8FTE4mpCYMSo"
}
```

If you don't provide request body you will get `HTTP 400`:

```json
{"error":"invalid request body"}
```