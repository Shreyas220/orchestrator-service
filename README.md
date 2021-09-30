# orchestrator-service

### Description
The following project has 3 RPC services 
- GetUserByName    (on port:9000)
- GetUser          (on port:9001)
- GetMockUserData  (on port:10000)

##### The client inputs name from the user then sends it GetUserByName and so on, the diagram below explains the sequence of methods called  


 ```
    client ---RPC--> (GetUserByName)orchestrator_1(:9000) ---RPC--> (GetUser)orchestrator_2(:9001) ---RPC--> (GetMockUserData)mock_data_service(:10000)
 ```
 
 ##### The project has 4 folders  
 - protos -> the proto file and the generated code are present here
 - client -> the client function is present here 
 - logic -> orchestration logic resides
 - datamock -> dummy data service resides here

## How to use 

### Downloading the repo in your local system 

 ```
    git clone https://github.com/Shreyas220/orchestrator-service.git
 ```
  
  enter the folder
  
 ```
    cd orchestrator-service
 ```
### Starting the server 
There is a server.go file present which starts the three RPC services  
 
 ```
    go run server.go
 ```
 
 ### Sending the name from client 
 open another terminal in orchestrator-service folder
 then enter client folder and run client.go 
 
  ```
    cd client
 ```
 
  ```
    go run client.go
 ```
 you will now be asked to enter a name and will be returned with a response or an error if is length is less than 6 

