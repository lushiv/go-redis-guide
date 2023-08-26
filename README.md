# Getting Started
To get started, make sure you have Go installed on your system. Additionally, you'll need a Redis server up and running. If you don't have Redis installed, you can download it from redis.io/download.

    Clone this repository to your local machine using:

    bash

git clone https://github.com/lushiv/go-redis-guide.git


Change these values with your :
```
var redisHost = ""
var redisPassword = ""
var redisDB = 0

```

Navigate to the repository's directory:

    bash
    `cd go-redis-guide`
    `go mod init go-redis-guide`
    `go get github.com/go-redis/redis/v8`
    `go run main.go`


