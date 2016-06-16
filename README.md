# Flag 

Flag service with fqdn go.micro.srv.flag
Flag API with fqdn go.micro.api.flag

Data is stored in Elasticsearch.


## Getting Started

### Prerequisites

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

### Create flag
 
```
micro query go.micro.srv.flag Flag.Create '{"key": "my-unique-flag-key", "description": "You know, for UI feed", "value": true}'
{}
```


### Read flag
 
```
micro query go.micro.srv.flag Flag.Read '{"key": "my-unique-flag-key"}'
{
	"key": "my-unique-flag-key",
	"description": "You know, for UI feed",
	"value": true
}

```


### Flip flag
 
```
micro query go.micro.srv.flag Flag.Flip '{"key": "my-unique-flag-key"}'
{}
```


### Delete flag
 
```
micro query go.micro.srv.flag Flag.Delete '{"key": "my-unique-flag-key"}'
{}
```


### List flags
 
```
micro query go.micro.srv.flag Flag.List '{}'
{
	"result": [
		{
			"key": "4",
			"description": "asdfasdfasdf"
		},
		{
			"key": "somekey-asdfasdf",
			"description": "asdfasdfasdf",
			"value": true
		},
		{
			"key": "3",
			"description": "asdfasdfasdf",
			"value": true
		},
		{
			"key": "34",
			"description": "asdfasdfafffffffffffffffffffffffsdf",
			"value": true
		}
	]
}
```


