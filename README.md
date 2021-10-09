# Instgram API
## _Submission for appointy tech_

instagram-api is an api which is capable of creating  
users and posts relationally using go and mongodb.

## Features

- Create/Get a user
- Create/Get a post
- Get all posts of user

## Libraries

instagram-api uses 2 libraries.

- [std lib] - The standard go library!
- [mongodb driver] - The MongoDB supported driver for Go.

#### Setup

Configure mongodb uri in main.go (line:28)
```sh
uri := "<mongo_connection_uri>"
```


To run:

```sh
go run .
```

#### Models
- user
    - id
    - name
    - email
    - password (sha512 hashed)
- post
    - id
    - caption
    - imageUrl
    - userId
    - timeCreated

#### Endpoints
- user
    - /users/<id> - GET
    - /users/     - POST
- post
    - /posts/<id> - GET
    - /posts/     - POST
    - /posts/users/<user_id> - GET

   [std lib]: <https://pkg.go.dev/std>
   [mongodb driver]: <https://pkg.go.dev/go.mongodb.org/mongo-driver>

