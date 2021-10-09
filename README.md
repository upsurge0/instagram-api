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

## Unit Tests

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/16555841-3291efc3-8ef4-4396-a30d-b9fe85a21485?action=collection%2Ffork&collection-url=entityId%3D16555841-3291efc3-8ef4-4396-a30d-b9fe85a21485%26entityType%3Dcollection%26workspaceId%3Da8fc03bc-ae59-4dd3-bdb1-107549cc2e74)

   [std lib]: <https://pkg.go.dev/std>
   [mongodb driver]: <https://pkg.go.dev/go.mongodb.org/mongo-driver>

