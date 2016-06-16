default: clean  build

clean:
	rm -rf api

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo  .
