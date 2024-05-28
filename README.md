# Golang Developer Assigment

Develop in Go language a service that will provide an API for retrieval of the Last Traded Price of Bitcoin for the following currency pairs:

1. BTC/USD
2. BTC/CHF
3. BTC/EUR


The request path is:
/api/v1/ltp

The response shall constitute JSON of the following structure:
```json
{
  "ltp": [
    {
      "pair": "BTC/CHF",
      "amount": "49000.12"
    },
    {
      "pair": "BTC/EUR",
      "amount": "50000.12"
    },
    {
      "pair": "BTC/USD",
      "amount": "52000.12"
    }
  ]
}

```

You shall provide time accuracy of the data up to the last minute.


# Requirements:
1. Code shall be hosted in a remote public repository
2. readme.md includes clear steps to build and run the app
3. Integration tests
4. Dockerized application

# Docs
The public Kraken API might be used to retrieve the above LTP information
[API Documentation](https://docs.kraken.com/rest/#tag/Spot-Market-Data/operation/getTickerInformation)
(The values of the last traded price is called “last trade closed”)


# goLTPB
Go Last Traded Price of Bitcoin api

## Application layout
main.go <br>
root of the application, here we start the application <br>
- app/<br>
    - api/<br>
        This folder allow the code to start your application, usually the router of the api have in their own directory with their respective purpose or version 
    - biz/<br>
        The biz layer works like a bridge, starting the third part application, our owns pkg, clients like mysql, redis, kafka, etc..; in the start function and using their availables functions that we add in the BizHandle interface
    - client/<br>
        This folder is used to load the services like db, kafka, redis, payments gateways ans so on
    - config/<br>
        This folder is used to load the default configuration and the environment variables that we can add to the application
- build/<br>
    This folder is used to save the builds of the application
- deployment/<br>
    This folder is used to save the yaml files to deploy the application easily, usually have the yaml file for the namespace, secrets, rbca, configmap, ingress, service and deployment|pod|statefulset.
- docs/<br>
    This folder is used to save the swagger documentation<br>
- test/<br>
    This folder is used to save the test required<br>

The pkg folder allows to have the differents services and modules that are used for ours app/client and the app/biz layer
- pkg
    - clients/
        - mysql/<br>
            This able the application to use mysql, this is used for our app/client/db/mysql module

## Start application
In this case we can tun the api using: `go run main.go` <br>

if is needed add environment variables you could use something like this: `DB_TYPE='mysql' DB_MYSQL_IP='localhost:3306' DB_MYSQL_NAME='goiso' DB_MYSQL_USER='root' DB_MYSQL_PASS='12345678' DB_MYSQL_RETRY='5' MIGRATE_DB_USER='root' MIGRATE_DB_PASS='12345678' MIGRATE_DB=true CONTINUE_AFTER_MIGRATE=true go run main.go` <br>

Once is running you can use the swagger page to hit the endpoint <br>
local : http://localhost:8080/swagger/index.html <br>
environment : https://YOUR-URL/swagger/index.html <br>

## Run app with docker-compose
First, create the image of the application: 
- docker build -t goltpb:latest .

Second, run the docker-compose.yaml file. 
This file has all the configuration needed, no require any change. 
- docker-compose up -d

And to turn off the application:
- docker-compose down -v


## Recommendations

### golang-migrate 
For the db use golang-migrate, this [link](https://github.com/golang-migrate/migrate) get you the list of drivers <br>

Installation in ubuntu using WSL<br>
First add the script to migrate for golang:<br>
    curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | sudo bash<br>
Then run the next commands<br>
    sudo apt update <br>
    sudo apt install migrate <br> 

Now you can use migrate. The "create", "goto", "up" and "down" commands are availables to do what you want to.<br>

The create command use the name_of_script that you want to create:<br> 
- migrate create -ext=.sql [name_of_script] <br>

The goto command use the version of the file created by the create command, the version is the prefix of the file until the first "_" character:<br>
- migrate -path=[scripts_folder] -database="mysql://[user]:[password]"@tcp([ip])/[db_name] goto [version]<br>

The up command update the db to the latest version:<br>
- migrate -path=[scripts_folder] -database="mysql://[user]:[password]"@tcp([ip])/[db_name] up <br>

The down command downgrade a specific counts of migrations steps:<br>
- migrate -path=[scripts_folder] -database="mysql://[user]:[password]"@tcp([ip])/[db_name] down [count_steps]<br>

The others parameters are the scripts_folder, this is the path to the folder where you are creating your scripts. user is the user to your database, password is the password to your database, the ip and bd_name should exist.<br>

### Swagger 

Swagger used to test the endpoints an the responses that we need<br>
install command: <br>
- go get -u github.com/swaggo/swag/cmd/swag
- go get -u github.com/swaggo/gin-swagger
- go get -u github.com/swaggo/files

run swagger: <br>
Usually you can use just swag(if the command not found specify the path to the executable)
- $HOME/go/bin/swag init

### mocks to test

mockery for an easy unit testing 
install 
- go install github.com/vektra/mockery/v2@latest

how to generate the files
- mockery --name=ApiHandle --dir=./app/api --output=api --inpackage=true

if fails we can point directly to the executable, for example for the 3 interfaces that we have here: 
- $HOME/go/bin/mockery --name=ApiHandle --dir=./app/api --output=api --inpackage=true
- $HOME/go/bin/mockery --name=Handle --dir=./app/biz --output=biz --inpackage=true
- $HOME/go/bin/mockery --name=Repository --dir=./app/client/db --output=db --inpackage=true

you also can use notations on comments but for leave the code clear I did this steps

### postman

postman <br>
There are 2 files in the test/postman folder. <br>
Import them to the postman app, once you are in the collection use the "local_env" imported. Then run the go app, the most simpliest way is using the simple repo in memory with `go run main.go` and at the end run the request in the collection. <br>
If you are using a production URL, change the value in local_env to the corresponding URL <br>

