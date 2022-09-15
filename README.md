# bisq-rpc-gateway

Simple GO gRPC Gateway

Routes HTTP 1.1 requests to a remote server's single gRPC service method, which is responsible for detecting if the 
request is from this gateway, and wrapping the response in json.  

This gateway takes care of mapping gRPC error status codes to HTTP status codes, but the remote server is responsible
for translating core api exceptions into gRPC StatusRuntimeExceptions with appropriate Status values.   

There is a single gRPC Service called MessageService with a single method Call(String params).

There is a single HTTP 1.1 endpoint (http://host:port/v1/call) which accepts a request body only; extra
path elements or query strings are ignored, and will likely cause errors.

Here is an example getversion method call from a rest client: 

curl -v -X POST http://localhost:8080/v1/call -H "Authorization:xyz"  -H "Content-Type: application/json" -d '{"params": "getversion"}'

Here is an example setwalletpassword method call from a rest client: 

curl -v -X POST http://localhost:8080/v1/call -H "Authorization:xyz"  -H "Content-Type: application/json" -d '{"params": "setwalletpassword newpassword"}'

## TODO 
Show response examples.

## TODO
Example go runtime setup, bisq-grpc-gateway install and run procedure.