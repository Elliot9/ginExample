# GinExample 

An example of gin contains basic features

## Installation

```shell
go get github.com/Elliot9/ginExample
```

## How to start
Copy from `.env.example` file to `.env` file in the root of your project

```shell
cp .env.example .env
```
If you're even lazier than that, you can just use default setting from `./config/loader`

### Migration
```shell
cd scripts && ./migration.sh
```

## Run
```
$ cd $GOPATH/src/go-gin-example

$ go run main.go 
```

## Features
- App configurable
- Gin
- Gorm
- Redis
- Mail
- Oauth
- JWT

##
![image](https://i.imgur.com/RUyFCJL.jpeg)
![image](https://i.imgur.com/x2hXWx3.jpeg)
![image](https://i.imgur.com/n2ZhBSH.jpeg)


## Todo
- Docs
- Swagger
- Rate Limiting
- RabbitMQ
- Websocket
- File Upload
- API Document
- CI/CD
- K8S
- Test