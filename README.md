# Flag Srv

This is the Flag service with fqdn go.micro.srv.flag.

## Getting Started

### Prerequisites

Install Consul
[https://www.consul.io/intro/getting-started/install.html](https://www.consul.io/intro/getting-started/install.html)

Run Consul
```
$ consul agent -dev -advertise=127.0.0.1
```

### Run Service

```
$ go run main.go
```

### Building a container

If you would like to build the docker container do the following
```
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o flag-srv ./main.go
docker build -t flag-srv .

```


{"key": "somekey-asdfasdf",
"description": "asdfasdfasdf",
"value": true}
