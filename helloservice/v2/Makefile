
version = v2

all: docker

clean:
	@- rm service
	@- docker rmi helloservice:$(version)

service: clean
	CGO_ENABLED=0 GOOS=linux go build -o service

docker: service
	docker build -t helloservice:$(version) .

