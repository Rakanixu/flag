default: build-docker

install:
	go install

clean:
	go clean

build-docker:
	cd srv && make && cd ..
	cd api && make && cd ..
	docker-compose -f docker-compose-build-minimal.yml build

protoc:
		protoc -I$$GOPATH/src --go_out=plugins=micro:$$GOPATH/src $$PWD/srv/proto/**/*.proto

