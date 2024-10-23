# Go Enigma Laundry

A simple Go application for managing laundry transactions for the admin side. This project demonstrates various aspects of Go programming, including handling HTTP requests, using the Gin framework, implementing clean architecture, connecting to a database, handling transactions, unit testing, and authentication.

## Branches Overview

The project is organized into several branches, each focusing on a different aspect of the application:

1. **http**
2. **gin**
3. **clean-arch**
4. **connect-db**
5. **transaction**
6. **unit-testing**
7. **auth**

### Branch Details

#### 1. http

This branch covers the basics of handling HTTP requests and responses in Go. It provides a simple implementation of CRUD operations for managing users and transactions without using any external framework.

#### 2. gin

This branch integrates the Gin framework, a popular web framework for Go, to handle routing and middleware more efficiently. It includes improvements over the `http` branch by leveraging Gin's features for better request handling and response formatting.

#### 3. clean-arch

This branch implements the clean architecture principles, separating the application into layers: domain, repository, service, and delivery. It promotes maintainability and testability by organizing the code into distinct modules with clear responsibilities.

#### 4. connect-db

This branch connects the application to a database, using a specific database driver (e.g., PostgreSQL, MySQL) to perform CRUD operations. It includes database configuration, migration scripts, and examples of querying and updating data in the database.

#### 5. transaction

This branch focuses on implementing transaction handling within the application. It includes models and services for managing laundry transactions, such as creating, updating, and retrieving transaction records, along with calculating total prices and managing transaction statuses.

#### 6. unit-testing

This branch introduces unit testing to the application. It includes tests for various components, such as models, services, and handlers, using Go's testing package and other testing libraries to ensure the application's functionality and reliability.

#### 7. auth

This branch adds authentication features to the application. It implements JWT-based authentication, user login, and token generation. It secures endpoints and ensures that only authenticated users can access certain parts of the application.

### Models

- **Customer**: Represents a customer with ID, name, address, phone number, and email.
- **Service**: Represents a service with ID, description, and price.
- **Transaction**: Represents a transaction with details such as customer, date, pickup date, status, and transaction details.
- **User**: Represents a user with ID, full name, email, username, password, and role.

## Build and Run

### Prerequisites

- Go (https://golang.org/dl/)
- PostgreSQL (https://www.postgresql.org/)

### Initials Database

```sql
create database db_laundry;

/c db_laundry;

create table mst_customers(
    id varchar(255) primary key,
    name varchar(150),
    address TEXT,
    phone_number varchar(100) unique,
    email varchar(100) unique
);

create table mst_services(
    id varchar(255) primary key,
    description varchar(150),
    price float
);

CREATE TABLE transactions (
    id VARCHAR(100) PRIMARY KEY,
    customer_id VARCHAR(100),
    order_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    pickup_date DATE,    
    status VARCHAR(50),
    is_pickup BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (customer_id) REFERENCES mst_customers(id)
);

CREATE TABLE transaction_details (    
    id VARCHAR(100) PRIMARY KEY,
    transaction_id VARCHAR(100),
    service_id VARCHAR(100),
    qty INT DEFAULT 1,
    FOREIGN KEY (transaction_id) REFERENCES transactions(id),
    FOREIGN KEY (service_id) REFERENCES mst_services(id)
);

```


### Library

This project is using the following modules :

- [gin-gonic](https://pkg.go.dev/github.com/gin-gonic/gin)
- [google-uuid](https://github.com/google/uuid)
- [godotenv](https://github.com/joho/godotenv)
- [driver-pq](https://github.com/lib/pq)

### Installing Dependencies

After cloning the project, navigate to the project directory and run the following command to download and install the required dependencies:

```bash
$ go mod tidy
```

### Env File

You have to set value on env file before run the project

```ini
DB_HOST=localhost
DB_PORT=5432
DB_NAME=db_laundry
DB_USER=postgres
DB_PASS=password
DB_DRIVER=postgres
ISSUER_NAME=incubation-golang-22
SIGNATURE=secretkey
TOKEN_LIFE_TIME=20
```

### Running the Project

```bash
$ go run .
```

### Building the Project

```bash
$ go build -o enigma-laundry
$ ./enigma-laundry
```

### EndPoints

Below are examples of some request endpoints. For the remaining endpoints, please refer to the controllers.

#### Master Customer

1. **Register New user**

You can shoose Role only between **ADMIN** or **USER**

- `POST` -> localhost:8085/auth/register
- Request : 
```json
{
    "fullName" : "Dian Anggraeni",
    "email" : "dian@gmail.com",
    "username" : "dian21",
    "password" : "dian123",
    "role" : "ADMIN"
}
```
- Response : 
```json
{
    "status": {
        "code": 201,
        "description": "Success Register new User"
    },
    "data": {
        "id": "fa2d4c9e-3c0b-4f30-a477-a6bd5c84b219",
        "fullName": "Dian Anggraeni",
        "email": "dian@gmail.com",
        "username": "dian21",
        "role": "ADMIN"
    }
}
```

2. **Login**

- `POST` -> localhost:8085/auth/register
- Request : 
```json
{
    "username" : "dian21",
    "password" : "dian123"
}
```
- Response : 
```json
{
    "status": {
        "code": 201,
        "description": "Success Login as User"
    },
    "data": {
        "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTYyODQ0ODEsImlzcyI6ImluY3ViYXRpb24tZ29sYW5nLTIyIiwidXNlcklkIjoiZmEyZDRjOWUtM2MwYi00ZjMwLWE0NzctYTZiZDVjODRiMjE5IiwiZW1haWwiOiJkaWFuQGdtYWlsLmNvbSIsInJvbGUiOiJBRE1JTiJ9.2XuYUJTr6BG7VPY8m9rDumWog3zRGhoE27lpIdxhZfE",
        "userId": "fa2d4c9e-3c0b-4f30-a477-a6bd5c84b219"
    }
}
```

2. **Register New Customer**

- `POST` -> localhost:8085/customers
- Request : 
```json
{   
    "name" : "Rosa",
    "address" : "Bali",
    "phoneNumber" : "8907896723",
    "email" : "rosa@gmail.com"
}
```
- Response : 
```json
{
    "status": {
        "code": 201,
        "description": "Success Register new Customer"
    },
    "data": {
        "id": "f7a628f0-adeb-47c4-8eba-b1248089d71a",
        "name": "Rosa",
        "address": "Bali",
        "phoneNumber": "8907896723",
        "email": "rosa@gmail.com"
    }
}
```

3. **Get All Customer**

- `GET` -> localhost:8085/customers?page=1&size=10
- Response : 
```json
{
    "status": {
        "code": 200,
        "description": "Success Get List Customer"
    },
    "data": [
        [
            {
                "id": "f7a628f0-adeb-47c4-8eba-b1248089d71a",
                "name": "Rosa",
                "address": "8907896723",
                "phoneNumber": "Bali",
                "email": "rosa@gmail.com"
            }
        ]
    ],
    "paging": {
        "Page": 1,
        "RowsPerPage": 10,
        "TotalRows": 1,
        "TotalPages": 1
    }
}
```

4. **Register New Service**

- `POST` -> localhost:8085/services
- Request : 
```json
{
    "description" : "CUCI KERING SETRIKA",
    "price" : 35000
}
```
- Response : 
```json 
{
    "data": {
        "id": "07d9b877-038c-4edc-99c2-8bdefdd592c6",
        "description": "CUCI KERING SETRIKA",
        "price": 35000
    },
    "message": "Success Register new Service"
}
```

- `POST` -> localhost:8085/services
- Request : 
```json
{
    "description" : "CUCI KERING LIPAT",
    "price" : 25000
}
```
- Response : 
```json 
{
    "data": {
        "id": "005e1790-db7e-4d58-b9ee-08b449b7505e",
        "description": "CUCI KERING LIPAT",
        "price": 25000
    },
    "message": "Success Register new Service"
}
```

5. **Get Service**

- `GET` -> localhost:8085/services
- Response : 
```json
{
  "data": [
    {
      "id": "f22cc587-2e5a-478f-b0c6-779c6a5d7bc3",
      "description": "Cuci Kering",
      "price": 20000
    },
    {
      "id": "801e29b1-048a-4b53-9233-2c9be9b1999e",
      "description": "Kering",
      "price": 10000
    },
    {
      "id": "2722a000-7fba-4312-b0ea-5f16b54d0463",
      "description": "CUCI KERING SETRIKA",
      "price": 35000
    }
  ],
  "message": "Success Find All Service"
}
```

6. **Create Transaction**

- `POST` -> localhost:8085/transactions
- Request : 
```json
{
    "customerId" : "15584b89-35d0-4d72-bb37-5913f630f46b",
    "transactionDetails" : [
        {
            "serviceId" : "f22cc587-2e5a-478f-b0c6-779c6a5d7bc3",
            "qty" : 1
        },
        {
            "serviceId" : "2722a000-7fba-4312-b0ea-5f16b54d0463",
            "qty" : 3
        }
    ]
}
```
- Response : 
```json
{
  "data": {
    "id": "70a5cb60-9c71-4847-8dc8-b4aabe5e314c",
    "customer": {
      "id": "15584b89-35d0-4d72-bb37-5913f630f46b",
      "name": "ridho",
      "address": "",
      "phoneNumber": "bumi",
      "email": "eak@gmail.com"
    },
    "date": "0001-01-01T00:00:00Z",
    "pickupDate": "2024-10-25T16:47:36.2326747+07:00",
    "status": "UnPaid",
    "isPickedUp": false,
    "transactionDetails": [
      {
        "id": "5d6e2141-6c06-450f-98e1-5ccf265d2af6",
        "transactionId": "70a5cb60-9c71-4847-8dc8-b4aabe5e314c",
        "service": {
          "id": "f22cc587-2e5a-478f-b0c6-779c6a5d7bc3",
          "description": "Cuci Kering",
          "price": 20000
        },
        "qty": 1
      },
      {
        "id": "a403112f-21d0-4af2-9c6d-919b25f937c6",
        "transactionId": "70a5cb60-9c71-4847-8dc8-b4aabe5e314c",
        "service": {
          "id": "2722a000-7fba-4312-b0ea-5f16b54d0463",
          "description": "CUCI KERING SETRIKA",
          "price": 35000
        },
        "qty": 3
      }
    ],
    "totalPrice": 125000
  },
  "message": "Success Create Transaction"
}
```

7. **Get All Transaction**

- `POST` -> localhost:8085/transactions
- Response : 
```json
{
  "data": [
    {
      "id": "70a5cb60-9c71-4847-8dc8-b4aabe5e314c",
      "customer": {
        "id": "15584b89-35d0-4d72-bb37-5913f630f46b",
        "name": "ridho",
        "address": "",
        "phoneNumber": "bumi",
        "email": "eak@gmail.com"
      },
      "date": "2024-10-23T16:47:36.234031+07:00",
      "pickupDate": "2024-10-25T00:00:00Z",
      "status": "UnPaid",
      "isPickedUp": false,
      "transactionDetails": [
        {
          "id": "5d6e2141-6c06-450f-98e1-5ccf265d2af6",
          "transactionId": "70a5cb60-9c71-4847-8dc8-b4aabe5e314c",
          "service": {
            "id": "f22cc587-2e5a-478f-b0c6-779c6a5d7bc3",
            "description": "Cuci Kering",
            "price": 20000
          },
          "qty": 1
        },
        {
          "id": "a403112f-21d0-4af2-9c6d-919b25f937c6",
          "transactionId": "70a5cb60-9c71-4847-8dc8-b4aabe5e314c",
          "service": {
            "id": "2722a000-7fba-4312-b0ea-5f16b54d0463",
            "description": "CUCI KERING SETRIKA",
            "price": 35000
          },
          "qty": 3
        }
      ]
    },
    {
      "id": "8ace70ba-bf54-4eb6-853a-6aa443040e43",
      "customer": {
        "id": "0f977e59-a2cf-413e-a9a2-a07eafcae829",
        "name": "tes123",
        "address": "08172387",
        "phoneNumber": "tes@gmail.com",
        "email": "eak"
      },
      "date": "2024-10-23T17:43:48.765426+07:00",
      "pickupDate": "2024-10-25T00:00:00Z",
      "status": "UnPaid",
      "isPickedUp": false,
      "transactionDetails": [
        {
          "id": "decad476-f2f6-456d-b694-0f1ef562ab75",
          "transactionId": "8ace70ba-bf54-4eb6-853a-6aa443040e43",
          "service": {
            "id": "f22cc587-2e5a-478f-b0c6-779c6a5d7bc3",
            "description": "Cuci Kering",
            "price": 20000
          },
          "qty": 1
        },
        {
          "id": "54cd0936-58e5-4daa-a26b-1c45b58b44b1",
          "transactionId": "8ace70ba-bf54-4eb6-853a-6aa443040e43",
          "service": {
            "id": "2722a000-7fba-4312-b0ea-5f16b54d0463",
            "description": "CUCI KERING SETRIKA",
            "price": 35000
          },
          "qty": 1
        }
      ]
    }
  ],
  "message": "Success get list Transaction"
}
```
