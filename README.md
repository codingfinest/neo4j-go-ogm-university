# Hilly Fields Technical College

This demo application is inspired by the university [example](https://github.com/neo4j-examples/neo4j-ogm-university) for the official Neo4j Java OGM. 

Stack:

* Neo4j go OGM v1.0.0
* Echo v4.1.14
* AngularJS 1.3
* Bootstrap 3.3

Note the client side of the project was copied from https://github.com/neo4j-examples/neo4j-ogm-university/tree/master/src/assets with slight modification to work with this project.


Getting started:

* Download and install Neo4j
* Setup the environment for the Neo4j Go driver as described [here](https://github.com/neo4j/neo4j-go-driver#requirements)
* clone git@github.com:codingfinest/neo4j-go-ogm-university.git
* Update OGM config in `store.go` with the appropriate `username` and `password`
* Set the env variable `OGM_EAMPLE_SCHOOL_CQL_PATH` to put to the file path pointing to the file for the initial database query
* cd neo4j-go-ogm-university
* run the program: `go run main.go`


The application runs on at http://localhost:9000. On the browser go to http://localhost:9000/api/relaod to initilize the databse. Note this wipes out the databse.


Play with the application.