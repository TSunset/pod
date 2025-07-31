# My Go Project

This is a simple Go web server using Echo and GORM.

## Features
- CRUD operations for messages
- PostgreSQL database
- RESTful API

## Installation
- docker terminal paste this: docker run --name my-postgres -e POSTGRES_PASSWORD=1234 -p 5432:5432 -d postgres
- run main.go
Now you can check the program through Postman

## Example
- Post: localhost:8080/messages
  {
    "text": "test3"
  }
- Putch: localhost:8080/messages/1
  {
    "text": "test666"
  }
- Get: localhost:8080/messages
  {
    "text": "test666"
  }
- Delete: localhost:8080/messages/1
  
