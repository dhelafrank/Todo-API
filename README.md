
# Todo API
A simple Todo API built with Go and Gin framework as part of the <b>[0x0_Golang learning](https://github.com/dhelafrank/0x0_Golang) </b>pathway.

## Endpoints
### GET /todos
Retrieve a list of all todos.

### GET /todos/:id
Retrieve a single todo by ID.

### PATCH /todos/:id
Toggle the status of a single todo by ID.

### POST /todos
Create a new todo.

### Running the API
To run the API, simply execute the following command in your terminal:

```go run .```

The API will be available at <http://localhost:9090>.

Note
This project is part of my learning journey in Go, specifically for the 0x0_Golang learning pathway. It's a basic implementation and you may want to add authentication, error handling, and other features depending on your requirements.

## Code Structure
The project consists of the following packages:

<b>main:</b> The main package handles the server execution and api endpoints definition.

<b>controller</b>: This package contains the controller functions that handle the API requests.

<b>types</b>: This package contains type definition for each todo task

<b>data</b>: This package contains the sample todo data and functions for managing todo data

## Learning Objectives
This project aims to cover the following learning objectives:

- Setting up a Gin router and defining API endpoints
- Creating controller functions to handle API requests
- Using Go to build a simple RESTful API

Next Steps
- Explore other features and best practices for building a robust API in Go.
