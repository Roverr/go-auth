# go-auth
### Standard authentication API written in Go.
[ ![Codeship Status for Roverr/go-auth](https://codeship.com/projects/88e72ab0-0b34-0134-b18c-129a07c0a376/status?branch=master)](https://codeship.com/projects/155801)

## Why?
This project is like a TODO MVC for me in Golang, except it is just a basic authentication API where users can register, login, get information about themselves, and finally delete their profiles if they wish to.

### Attention! This project might contain anti-pattern implementations!

## Project structure
This project has kinda it's own structure, since the dependencies in the files are referring to go-auth, and not like github.com/Roverr/go-auth.

You can test recursively by typing
```
go test ./...
```
If you would like to see the tests one by one, you should add -v
flag to the end of the command.

Wait for it script is provided by https://github.com/vishnubob/wait-for-it
