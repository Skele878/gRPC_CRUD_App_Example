# gRPC_CRUD_App_Example
A gRPC CRUD Application example. 
1. gRPC uses protocol buffers by default and the first step when working with protocol buffers is to define the structure of the data that want to serialize in a <proto> file. 

To generate client and server code using proto file use in terminal below commands:

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/movie.proto

 this will genetare server, client code int he proto folder. Files  as proto.pb.go and proto_grpc.pb.go
