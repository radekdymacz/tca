# Random name and a joke. [![Go Report Card](https://goreportcard.com/badge/github.com/radekdymacz/tca)](https://goreportcard.com/report/github.com/radekdymacz/tca)

API to fetch random name and joke and combining two together using go-micro.

## Contents

- **srv/joke** - an RPC joke service
- **srv/name** - an RPC name service
- **api** - API combining two services above 

## Requirements

 - docker
 - docker-compose

## Instalation

Just tested on Mac

```
$ go get-u github.com/radekdymacz/tca
$ cd $GOPATH/src/github.com/radekdymacz/tca
$ go get ./...
$ ./build.sh

```


## Run

```
docker-compose up
```

## Links

Once you have everything running you can use links below to inspect system

- [Consul UI](http://localhost:8500)
- [Micro Registry](http://localhost:8082/registry)
- [Micro Query](http://localhost:8082/query)
- [Micro CLI](http://localhost:8082/cli)
- [Get a joke](http://127.0.0.1:8080/joke/get)


## Test

```
./test.sh
```
or 
```
./test_race.sh
```

## TODO 

- [x] skeleton for the services 
- [x] Build scripts
- [x] Dockerfiles
- [x] Implement external services calls
- [x] basic unit tests 
- [ ] Add test cases
- [ ] Cleanup the code
- [ ] Benchmarks
- [ ] Dependencies mgmt
- [ ] Circle config
- [ ] Push images to DockerHub
- [ ] Deployment

## Example request and response

```
curl http://127.0.0.1:8080/joke/get

Lütfi Üstün doesn't need an OS

```

