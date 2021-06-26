## Task tools and packages

- ✨Gin✨ 
- ✨gorilla/websocket✨ 
- ✨Gorm✨
- ✨Viber✨
- [✨gocron✨](https://golangrepo.com/repo/go-co-op-gocron-go-cron-jobs)
## Task technology information

- dataBase driver : postgres
- running server machine : docker images in docker-compose
- Testing : mock testing for user repo

## Task architecture and model management information

- we try to use clean architecture .
- user directory will contain the User APIs , usecase and repo implementation
- infrastructure directory will contain the main configurations for the app (database connection , seeder ,cron job ...etc)
- entity directory will have the main blueprint for the user table in the database
- model directory will contain the structs for the req and res Json objects
- transformer will transform between models and entities
- we use `go mod` and `go vendor` for dependency management
- we use `gorm migration` for creating the schema 

## Task Running

- to run the app via docker composer run `docker-compose run`
- to run the app in your machine 
   - make sure you have postgresql database server is installed and run
   - create database and user with same configurations in `config.yml` file 
    - make sure that you have redis database server is installed and run
    - run `go mod vendor` 
    - run `go mod tidy`
    - run `go run .`

## Task Test
 - run `cd user/repository/`
 - run `go test`

## Additional Info
I have created a new `Postman collection` [here](https://www.getpostman.com/collections/750dbda433c0eb695f7c) for your test . with `documentation` [here](https://documenter.getpostman.com/view/6696943/Tzedi53f)
also to test the websocket I use `Chrome extension` called `simple ws client` you can find it [here](https://chrome.google.com/webstore/detail/simple-websocket-client/pfdhoblngboilpfeibdedpjgfnlcodoo/related?hl=en)
