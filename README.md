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
 ![image](https://user-images.githubusercontent.com/55346934/135410633-145ec0ee-6f57-460e-9c03-82d99e9d850c.png)

 
 ### Sending the name from client 
 open another terminal in orchestrator-service folder
 then enter client folder and run client.go 
 
  ```
    cd client
 ```
 
  ```
    go run client.go
 ```
 ![image](https://user-images.githubusercontent.com/55346934/135410657-5f2bbdb0-98a4-44f0-b27b-52a887cb5e3d.png)

 you will now be asked to enter a name and will be returned with a response or an error if is length is less than 6 
 
 1. Operating System - Linux (PopOs v21)
2. `GOPATH` = $HOME/go
3. repository path -> home/Desktop

