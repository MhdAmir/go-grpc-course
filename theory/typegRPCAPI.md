# 1. Unary 
process of the rest API
- client send 1 request -> server
- server send 1 response -> client

# 2. Server Streaming
enabled by HTTP/2
- Clients will send 1 request -> server
- Server can return more than one responses depending on the need

- if client want to get update in real time

# 3. Client Streaming
- client Multiple Requests -> server
- server return one response

- use on uploading or updating information.

# 4. Bi directional Streaming
- Client can Multiple Requests -> server
- Server can return Multiple responses
- trick, after first request, reponses could arrive in order you decide


``` protobuf
service GreetService {
    //unary
    rpc Greet(Greet Request) returns (GreetResponses) {}

    //server streaming
    rpc GreetManyTimes(Greet Request) returns (stream GreetResponses) {}

    // client streaming
    rpc LongGreet(stream Greet Request) returns (GreetResponses) {}

    //Bi directional streaming
    rpc GreetEveryone(stream Greet Request) returns (stream GreetResponses) {}

}
```