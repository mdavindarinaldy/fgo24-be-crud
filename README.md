# Project Backend Go - CRUD

This project was made by Muhammad Davinda Rinaldy in Kodacademy Training Program. This project uses Go Language (and PostgreSQL for the database) to make a backend application for basic CRUD simulation.

## Prerequisites

Make sure you already install Go to run this project

## How to Run this Project

1. Create a new empty directory for the project and navigate into it
2. Clone this project into the empty current directory:
```
git clone https://github.com/mdavindarinaldy/fgo24-be-crud.git .
``` 
3. Install dependencies
```
go mod tidy
```
4. Run the project
```
go run main.go
```

## Dependencies
This project use:
1. gin-gonic from github.com/gin-gonic/gin : for handling HTTP request/response data (gin.Context), for defining middleware and route handlers (gin.HandlerFunc), for organizing routes into groups (gin.RouterGroup) and for managing HTTP routing and server configuration (gin.Engine)
2. pgx from github.com/jackc/pgx/v5 : for direct database interactions (PostgreSQL)
3. godotenv from github.com/joho/godotenv : for loading environment variables from a .env file into the application

## CRUD Endpoint Overview
| Method    | Endpoint | Description |
| -------- | -------  | ------- |
| GET  | /users | Get all detail users |
| GET | /users/:id    | Get one detail user |
| POST    | /users | Create new user |
| PATCH | /users/:id | Update user's data |
| DELETE | /users/:id | Delete user | 

## Basic Information
This project is part of training in Kodacademy Bootcamp Batch 2 made by Muhammad Davinda Rinaldy