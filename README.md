# Exploring Go for Web
Building a REST API server with Golang

# How to build
Use Docker and docker-compose
```
  docker-compose build
  docker-compose up
```
You may have to do a `docker-compose down` first after building and running the app for the first time. Then, execute `docker-compose up` again.

# Other build methods
Use Go in your machine (_I haven't test this!_)

Requirements: you will need `dep` and `fresh`
```
  # install dep
  curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

  # install fresh
  go get -u github.com/pilu/fresh
```
Install dependencies
```
  dep ensure
```
Then run fresh for monitoring changes
```
  fresh -c runner.conf
```
# API Reference
App will be running on `PORT=8080`. The app currently has 1 resource.
```
  GET     /user/
  POST    /user/
  GET     /user/{id:[0-9]+}
  PATCH   /user/{id:[0-9]+}
  DELETE  /user/{id:[0-9]+}
```

# Known issues
If you encountered an error while executing `dep ensure`, then you probably clone the repo in the wrong place.

Solution: make sure you clone the repo under your `$GOPATH/src`. See [dep new project](https://github.com/golang/dep/blob/master/docs/new-project.md) for more information.

# TODO
- create authentication with JWT and sessions
- refactor handlers and the database connection
