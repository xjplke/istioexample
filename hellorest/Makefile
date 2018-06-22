
version = v1

all: docker

clean:
	@- rm hellorest
	@- docker rmi hellorest:$(version)

service: clean
	CGO_ENABLED=0 GOOS=linux go build -o hellorest

docker: service
	docker build -t hellorest:$(version) .

