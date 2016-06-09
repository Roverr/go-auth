# go-auth
### Standard authentication API written in Go.
[ ![Codeship Status for Roverr/go-auth](https://codeship.com/projects/88e72ab0-0b34-0134-b18c-129a07c0a376/status?branch=master)](https://codeship.com/projects/155801)

## Why?
This project is like a TODO MVC for me in Golang, except it is just a basic authentication API where users can register, login, get information about themselves, and finally delete their profiles if they wish to.

### Attention! This project might contain anti-pattern implementations!

## Install project
If you would like to install the project, you should have atleast Go 1.6 environment ready. After that you can install the project by typing:
```
go get github.com/Roverr/go-auth
```

## Docker
If you do not want to have Go environment or you just prefer running the application as a container you can use the docker-compose file which is starting a MySQL service and this Go application in a Go1.6 container.

Use docker-compose command to run this services:
```
docker-compose up
```

This application is using [wait-for-it](https://github.com/vishnubob/wait-for-it) script to wait for the MySQL container to come up.
Unfortunately this is a known [issue.](https://github.com/docker-library/mysql/issues/81)

Also I know that docker-compose can use .env files and there is no need for public MySQL root password, however, this container is **not for production use**. It is for **developing purposes only**.

# About the application
## API
#### POST - /auth/register
Endpoint used for registering a new user.
***Notice here***, that if you delete a user, you will not be able to register, login, or basicly do anything anymore with that user.

Request JSON body should look like:
```
{
  userName: "John",
  realName: "John Doe", // optional
  password: "macilaci",
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
  password: "macilaci",
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

Wait for it script is provided by https://github.com/vishnubob/wait-for-it
