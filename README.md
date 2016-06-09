# go-auth
### Standard authentication API written in Go.
[ ![Codeship Status for Roverr/go-auth](https://codeship.com/projects/88e72ab0-0b34-0134-b18c-129a07c0a376/status?branch=master)](https://codeship.com/projects/155801)

## Why?
This project is like a TODO MVC for me in Golang, except it is just a basic authentication API where users can register, login, get information about themselves, and finally delete their profiles if they wish to.

### Attention! This project might (and certainly will) contains anti-pattern implementations!

## Install project
If you would like to install the project, you should have atleast Go 1.6 environment ready. After that you can install the project by typing:
```
go get github.com/Roverr/go-auth
```

## Docker
If you do not want to have Go environment or you just prefer running the application as a container you can use the docker-compose file which is starting a MySQL service and this Go application in a Go 1.6 container.

Use docker-compose command to run this services:
```
docker-compose up
```

This application is using [wait-for-it](https://github.com/vishnubob/wait-for-it) script to wait for the MySQL container to come up.
Unfortunately this is a known [issue.](https://github.com/docker-library/mysql/issues/81)

Also I know that docker-compose can use .env files and there is no need for public MySQL root password, however, this container is **not for production use**. It is for **developing purposes only**.

# About the application
## Configuration
The application is using envconfig which would be optimal in production environment as well. Also you can use .env as well, but it will be required to be in the bin directory (same where application binary is), which I did not find really useful now, so this feature probably will be removed, or reworked.
Configuration can be found [here.](https://github.com/Roverr/go-auth/blob/master/config/config.go)

***Notice here*** that envconfig tag describes that how you should export your environment variable. Also you have to use GOAUTH tag before that, so for example:
```
export GOAUTH_DB_NAME=nameOfMyDatabase
```
Will be matching for DbName in the Config struct. This can be useful if you are planning to run multiple applications with the same environment configuration requirements. (Like database names..)

## API
#### POST - /auth/register
Endpoint used for registering a new user.
***Notice here***, that if you delete a user, you will not be able to register, login, or basicly do anything anymore with that user.

Request JSON body should look like:
```
{
  userName: "John",
  realName: "John Doe", // optional
  password: "macilaci"
}
```

If everything was correct, you get an empty response with 200 from the server, structured like this:
```
{
  data: {
    item: null,
  },
  error: null
}
```
#### POST - /auth/register
Endpoint used for logging in.

Request JSON body should look like:
```
{
  userName: "John",
  password: "macilaci"
}
```

If everything was correct, you get an empty response with 200 from the server, structured like this:
```
{
  data: {
    item: null,
  },
  error: null
}
```
Also after that, you will get your JWT token in the header (which can be set in the config of the server, but default is X-Goauth)

#### GET - /user/me
Endpoint used for getting public information to the client side about the user.

Request should have a valid JWT token in header with the same key as it has been in the response of the login endpoint.

If everything was correct, you get a response looking like this:
```
{
  data: {
    item: {
      realName: "John Doe",
      userName: "john",
      id: 1
    }
  },
  error: null,
}
```
Also you get back your JWT token as well. (To be stateless)

#### DELETE - /user/delete
Endpoint used to delete user account.
Once a user has been deleted, currently there is no way to restore it from any endpoint. (Use MySQL branch or restore database container)

Request should have a valid JWT token in header with the same key as it has been in the response of the login endpoint.

If everything was correct, you get an empty response with 200 from the server, structured like this:
```
{
  data: {
    item: null,
  },
  error: null
}
```

## Tests
#### API Tests
API tests are separated from the main package where the endpoints are, because there can be huge number of tests for endpoints, and if they are separated correctly into files, the main directory where the endpoints are, can get really hard to read in a long term.


Based on the project, this can be very different. Since this project has almost no unit-tests, I think this would be the correct separation in a long term.

You can test recursively by typing
```
go test ./...
```
If you would like to see the tests one by one with logging included you can use
```
go test -v ./...
```

Endpoints have logging which will log out information to the Stdout. This way you can easily follow what is happening in the application, however, this is **not the best practice for production use**.

## CI
#### Codeship
Codeship is a really cool CI in my opinion, however, their documentation says that they have Go 1.4 by default, so if you would like to test your code written in other version, you have to add a custom script for it, which can be found [here.](https://github.com/codeship/scripts/blob/master/languages/go.sh)

I don't really think it is the best solution ever, but atleast it is in the documentation!


## Future
I might improve some part of the project as time goes on, I have some random ideas in my mind, but unfortunately I do not have enough time for it now.
