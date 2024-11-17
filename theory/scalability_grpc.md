# Scalability on gRPC
### Server: everything is Async, 
so the main thread not blocked and thus server can handle many requests in parallel
###  Client: Async or Blocking
freedom of async or not, but when response is critical for your app, and expect sync result. But, when it isn't go with Async response