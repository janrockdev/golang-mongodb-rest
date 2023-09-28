GoLang + MongoDB RESTFul
----------------------------------------------------

[![Go Report Card](https://goreportcard.com/badge/github.com/janrockdev/golang-mongodb-rest)](https://goreportcard.com/report/github.com/janrockdev/golang-mongodb-rest)

This project is a 'Boilerplate' or 'Starter' to build RESTful
Applications and microservices using
GoLang ([Gin](https://github.com/gin-gonic/gin) HTTP web framework), MongoDB.


Table of Contents
-----------------

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Project Structure](#project-structure)
- [Future Work](#future-work)

Features
--------

- [Gin](https://github.com/gin-gonic/gin)
- [JWT Authentication](https://github.com/golang-jwt/jwt)
- DotEnv with [Viper](https://github.com/spf13/viper)
- Request Validation with [Ozzo Validation](https://github.com/go-ozzo/ozzo-validation)
- MongoDB ODM ([mgm](https://github.com/Kamva/mgm))
- Docker support
- Swagger with [gin-swagger](https://github.com/swaggo/gin-swagger)

Prerequisites
-------------

- Git
- Docker
- [Go 1.20](https://go.dev/doc/install) and above

_If you are not using Docker, you should also have:_

- MongoDB

Getting Started
---------------
#### Repository Clone

```bash
# Clone Project
git clone https://github.com/janrockdev/golang-mongodb-rest.git go-rest

# Change Directory
cd go-rest
```

#### Using Docker

```bash
# Build & Create Docker Containers
docker-compose up -d
```

#### Using Local Environment

```bash
# Copy Example Env file
cp ./env.example .env

# Change MongoDB URI and Database Name

# MONGO_URI=<mongo_uri>
# MONGO_DATABASE=<db_name>

# Download Modules
go mod download

# Build Project
go build -o go-starter .

# Run the Project
./go-starter
```

The application starts at port 8080:

- `GET /v1/ping` Health check endpoint, returns 'pong' message

---

- `POST /v1/auth/register` Creates a user and tokens
- `POST /v1/auth/refresh` Refresh expired tokens
- `POST /v1/auth/login` Login a user

---

- `POST /v1/notes` Create a new note
- `GET /v1/notes` Get paginated list of notes
- `GET /v1/notes/:id` Get a one note details
- `PUT /v1/notes/:id` Update a note
- `DELETE /v1/notes/:id` Delete a note

---

- `GET /swagger/*` Auto created swagger endpoint

You can also see: http://localhost:8080/swagger/index.html

> If you want to add new routes and swagger docs, you should run ```swag init```
> See: [Gin Swagger](https://github.com/swaggo/gin-swagger)

Project Structure
-----------------

```
├── controllers         # contains api functions and main business logic
├── docs                # swagger files 
├── logs
├── middlewares         # request/response middlewares
│   └── validators      # data/request validators
├── models              
│   └── db              # collection models
├── routes              # router initialization
└── services            # general service & database actions
```

Future Work
-----------

- Rate Limiting
- Testing

License
-------

[MIT License](LICENSE) - [Jan Rock](https://github.com/janrockdev)
