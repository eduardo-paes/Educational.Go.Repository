# Go REST API with Echo

This directory features a simple Go application for building a RESTful API using the [Echo](https://echo.labstack.com/) web framework. The application provides endpoints to create, read, update, and delete user data.

## Table of Contents

- [Go REST API with Echo](#go-rest-api-with-echo)
  - [Table of Contents](#table-of-contents)
  - [1. Introduction](#1-introduction)
  - [2. Prerequisites](#2-prerequisites)
  - [3. Project Structure](#3-project-structure)
  - [4. Getting Started](#4-getting-started)
  - [5. API Endpoints](#5-api-endpoints)

## 1. Introduction

This Go application serves as a basic template for creating a RESTful API with the Echo web framework. It showcases the fundamental setup of the server, middleware, and API endpoints for managing user data.

## 2. Prerequisites

Before using this code, make sure you have the following prerequisites:

- Go installed on your system.
- Familiarity with Go programming.
- Basic knowledge of RESTful APIs.

## 3. Project Structure

The project structure is minimal and consists of the following files:

- `main.go`: The main entry point of the application, containing server setup and API endpoint definitions.

## 4. Getting Started

To get started with this application:

1. Clone or download the repository to your local machine.

2. Install the necessary Go dependencies:

   ```bash
   go mod tidy
   ```

3. Run the Go application:

   ```bash
   go run main.go
   ```

   This will start the server on port `:8080`.

4. Access the API endpoints as described in the next section.

## 5. API Endpoints

The API provides the following endpoints for managing user data:

- **Create User**: `POST /users`
- **Get User by ID**: `GET /users/:id`
- **Get All Users**: `GET /users`
- **Update User by ID**: `PUT /users/:id`
- **Delete User by ID**: `DELETE /users/:id`

You can use tools like `curl`, [Postman](https://www.postman.com/), or your preferred API client to interact with these endpoints.
