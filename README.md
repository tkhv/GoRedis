# Redis in Go

[![Build Status](http://54.224.91.178:8080/buildStatus/icon?job=goredis)](http://54.224.91.178:8080/job/goredis/)

A key value store in Go implementing some of the Redis Serialization Protocol. This also comes with a Dockerfile and a Jenkinsfile for easier deployments. Built as a practice project.

## Install & Run

You will require a redis client such as [redis-cli](https://redis.io/docs/connect/cli/) to interact with the server.

You can build and run the server with `go run .`, which launches it at localhost:6379. You should then be able to connect using `redis-cli`.

To run it within a container instead, you can build with:

`docker build -t [image_name] .`

followed by running it with the default 6379 port exposed:

`docker run -p 6379:6379 [image_name]`
