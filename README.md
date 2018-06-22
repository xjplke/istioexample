# istioexample



1、make at hellorest  helloservice/v1   helloservice/v2
    when you change code, remove container before make
2、docker run test for image
    docker run -d -p 50051:50051 --name=helloservice_v1 helloservice:v1
    docker run -d -p 8123:8123 -e HELLO_PORT=50051 -e HELLO_SERVICE=192.168.100.100 -e PORT=8123  --name=hellorest --add-host helloservice:192.168.100.100 hellorest:v1

    for a long time no call， the first call will failed.
    https://github.com/grpc/grpc/issues/5468




