# API

The API is the entrypoint of the server. 

For now, it can manages 2 routes : 

# Documentation autogénéré Swagger

swagger-codegen generate -i doc.yaml -l html

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


### Password Endpoints
Chaque mot de passe doit remplir ces conditions : 

* At least one lowercase letter ((?=.*[a-z]))
* At least one uppercase letter ((?=.*[A-Z]))
* At least one digit ((?=.*[0-9]))
* At least one special character ((?=.*[!@#$%^&*]))
* At least 8 characters in total ((?=.{8,}))

* /forgetPassword/:email -> 
Ou l'email du compte est attendu, si il la trouve ça va lui envoyer un mail avec un code que
l'utilisateur va devoir renseigné pour prouvé que c'est bien son compte.

* /forgetPassword/:email/:code -> 
L'endpoint qui vérifie si le code renseigné est bien celui envoyé par mail par le précédent endpoint

* /setPassword/:email/:new ->
Cet endpoint permet, après vérification du code de changer le mot de passe.

* /editPassword/:userID/:old/:new ->
Manière classique de mettre à jour son mot de passe quand on connait son ancien mot de passe



# Errors 

errors : ../../../../../../go/src/google.golang.org/api/storage/v1/storage-gen.go:1:1: expected 'package', found 'EOF'
../../../../../../go/src/google.golang.org/api/transport/dial.go:1:1: expected 'package', found 'EOF'
../../../../../../go/src/google.golang.org/api/transport/grpc/dial.go:1:1: expected 'package', found 'EOF'
../../../../../../go/src/google.golang.org/api/transport/http/configure_http2_go116.go:1:1:

FIX : go env -> export GO111MODULE="on"
