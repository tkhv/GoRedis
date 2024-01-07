# Redis in Go

A key value store in Go implementing some of the Redis Serialization Protocol. Built as a practice project.

## Install

This was containerized with Docker, so you would require Docker and a redis client such as [redis-cli](https://redis.io/docs/connect/cli/) installed.

You can build with:

`docker build -t [image_name] .`

followed by running it with the default 6379 port exposed:

`docker run -p 6379:6379 [image_name]`

You should then be able to interact with the server using redis-cli at localhost:6379 (other clients may work but haven't been tested).
