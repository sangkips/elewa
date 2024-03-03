# Book API
Welcome to the README file for the Book API written in Golang. This API provides endpoints to manage a collection of books, allowing users to perform operations such as creating, reading, updating, and deleting books. This document will guide you through the setup and usage of the API.
## Setup
Before running the Book API, ensure that you have Go installed on your system. You can download and install Go by following the instructions provided on the official Go website.
Also, make sure you have a MongoDB database, you can download MongoDB from the official MongoDB website
For this example, I am running Mongo Atlas, it's easy to operate, you can get it from the official MongoDB website 

- Once Go is installed, follow these steps to set up the Book API:
- Clone the repository to your local machine:
- Navigate to the project directory:
- Install dependencies using `go mod tidy`
- Build the application using `make build`
- Run the application `make run`

## Usage
The Book API exposes several endpoints to interact with the book collection. You can use tools like cURL, Postman, or your favorite HTTP client to interact with the API.

### Endpoints
Book
- `GET /books`: Retrieve all books.
- `GET /books/{id}`: Retrieve a specific book by ID.
- `POST /book`: Create a new book.
- `PATCH /books/{id}`: Update an existing book by ID.
- `DELETE /books/{id}`: Delete a book by ID.

User Authentication 
- `GET /users`: Retrieve all users.
- `GET /user/{id}`: Retrieve a specific user by ID.
- `PATCH /user/{id}`: Update an existing user by ID.
- `POST /register`: signup new users
- `POST /login`: Login to the API
Only `register` and `login` endpoints don't require authentication.
All other endpoints require you to provide a token:
- Use Postman for easy convenience. Place your token on the `Headers` use `token` as `key` and copy `tokenvalue` after successful login

### Example requests
Retrieve book
```
{
            "id": "65e439d17f2ca10de65f312e",
            "name": "Portia",
            "price": 500,
            "book_id": "65e439d17f2ca10de65f312e",
            "category_id": "65e429b0d2560149a0ec8948",
            "Author": {
                "first_name": "julian",
                "last_name": "mongo"
            }
}
```
Create a book
```
{
    "success": {
        "InsertedID": "65e4947a0d3df732a0992535"
    }
}
```

### Deployment notes on AWS

- Configure VPC
- Configure EC2 instance
- Copy go-code to the EC2 instance
- Push the code to Docker registry / ECR
- Pull from docker / ECR on EC2
- Run the container
