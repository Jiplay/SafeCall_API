# API

The API is the entrypoint of the server. 

For now, it can manages 2 routes : 

GET Login :
ex : http://localhost:8080/login?name=Ju&psw=******

POST Register : 
http://localhost:8080/register?name=Ju&psw=******


# Run with docker with 
    $ docker-compose up
Si des modifications sont effectués il faut rajouter le flag `--build`


### Install Mongo with
    $ go get go.mongodb.org/mongo-driver/mongo

### Install protoc with apt or brew(MacOS)
    $ apt install -y protobuf-compiler
    $ brew install protobuf
    $ protoc --version  
Ensure compiler version is 3+



# Errors 

errors : ../../../../../../go/src/google.golang.org/api/storage/v1/storage-gen.go:1:1: expected 'package', found 'EOF'
../../../../../../go/src/google.golang.org/api/transport/dial.go:1:1: expected 'package', found 'EOF'
../../../../../../go/src/google.golang.org/api/transport/grpc/dial.go:1:1: expected 'package', found 'EOF'
../../../../../../go/src/google.golang.org/api/transport/http/configure_http2_go116.go:1:1:

FIX : go env -> export GO111MODULE="on"
