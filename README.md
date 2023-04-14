# hello-grpc
Quick POC to learn GRPC

Userful link: 


To generate Proto


protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    hello.proto


(Change the above folder as required)


To run the application use 

 go run server/server.go 

 go run client/greeter_client.go