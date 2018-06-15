# Exploring Go for Web
Building a REST API server with Golang

# How to build
Use Docker and docker-compose
```
  docker-compose build
  docker-compose up
```
TODO: you will need to create a database, migration files are not yet implemented

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
  fresh
```

# Known issues
If you encountered an error while executing `dep ensure`, then you probably clone the repo in the wrong place.

Solution: make sure you clone the repo under your `$GOPATH/src`. See [dep new project](https://github.com/golang/dep/blob/master/docs/new-project.md) for more information.
