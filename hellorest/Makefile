
version = v1

all: docker

clean:
	@- rm hellorest
	@- docker rmi hellorest:$(version)

hellorest: clean
	CGO_ENABLED=0 GOOS=linux go build -o hellorest

docker: hellorest 
	docker build -t hellorest:$(version) .

