# Why gRPC use Protocol Buffer?
gRPC uses protocl buffer for communication, the reason cause protocolbuffer **has smaller payload size**

## Efficiency over JSON
- with protocol buffer, we can save a lot of bandwidth or storage space because the messages are smaller
- less CPU Intesive: because it very close ith the actual data represent in memory
- Faster
- more efficient communication with devices are bit weaker, such as mobile or microcontroller