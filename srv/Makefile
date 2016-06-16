default: clean	build

clean:
	rm -rf srv

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo  .
