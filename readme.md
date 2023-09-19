# Funtory Interview Test:

***Author: Mehrshad Darvish***

-----

This repository contains a Go application designed to connect to whatsapp

## Prerequisites:

1. PostgreSQL
2. Go Modules (Initialized with "go mod init")
3. Dependencies Managed ("go mod tidy")

-----

### Sample Data:

1- Use make file is a first option:

    - make migration-up
-----

### CLI Commands:

The CLI commands are provided to manage and interact with the application:

- make migrate-up -> Migrates the database schema up.
- make migrate-down -> Rolls back the database schema migration.
- make serv -> Starts the HTTP server.

### Endpoints:

The application exposes the following endpoints:

- GET http://localhost:8601/health
- GET http://localhost:8601/api/v1/users/register 
- GET http://localhost:8601/api/v1/users/login 
- GET http://localhost:8601/api/v1/users/connect 

***Please make sure to adjust the URLs and port numbers and the Database Config based on your deployment environment.***

-----

Feel free to modify the code and extend the application based on your requirements.
