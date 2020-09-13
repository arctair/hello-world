# hello-world
Sample golang service for testing an [automated deployment utility](https://github.com/arctair/quarky).
## Run the tests
```
$ go test arctair.com/hello-world/v1
$ go test -tags acceptance
```
or
```
$ nodemon
```
### Run the tests against a deployment
```
$ BASE_URL=https://hello-world.arctair.com go test
```
## Run the server
```
$ go run .
$ curl localhost:5000
```
## Build a docker image
```
$ go build -o bin/hello-world
$ docker build -t arctair/hello-world .
```
