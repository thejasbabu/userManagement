# User Management Service 

A service which stores and retrieves user data from mongoDb.

## Dependencies  

1. Install Go
2. Install Dep
3. Install Docker
4. Install MongoDb

## Configuration

The application requires setting the following env variables  
| Env Variable | Desc |
|---------------|-------|
| UM_APPPORT | Port at which the application will run | 
| UM_MGDBURL | URL of mongoDB |
| UM_MGDBNAME | Database name |
| UM_MGUSERCOLLECTION | Name of the user collection |

## Set-up

1. Run the below command to install dependencies   
`dep ensure`   

2. Run the below command to compile    
`make build` or `make linux-build`    

3. Run the mongoDb using    
`mongod`

4. Run the executable to run the application  
`./main`

5. Build the docker image by running the following command   
`docker build -t um-service .`

6. Run the container using   
`docker run -p 8000:8000 -d um-service`

7. Hit `localhost:8000/user` to get a list of users
