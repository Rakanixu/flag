# Flag 

Flag service with fqdn go.micro.srv.flag
Flag API with fqdn go.micro.api.flag

Data is stored in Elasticsearch.


## Getting Started

### Prerequisites
Get Micro
[Micro](https://github.com/micro)
```
go get github.com/micro
```


Install Consul
[https://www.consul.io/intro/getting-started/install.html](https://www.consul.io/intro/getting-started/install.html)

Run Consul
```
$ consul agent -dev -advertise=127.0.0.1
```

### Run Service manually

```
$ go run srv/main.go
```

### Run API manually

```
$ go run api/main.go
```


### Run docker containers
Compile Go binaries and build docker image. 
```
make 
```

Run docker container:
```
docker-compose -f docker-compose-build.yml up
```


## Usage
[API](https://github.com/Rakanixu/flag/tree/master/api)

[Microservice](https://github.com/Rakanixu/flag/tree/master/srv)


