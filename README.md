# Software-Market Go API

SoftwareMarketGoAPI is a RESTful API built with Golang that provides endpoints to manage software products in a marketplace. It allows users to perform operations such as creating, retrieving, updating, and deleting software products.

## Installation

Make sure you have Go installed. You can download it from the official website: https://golang.org/dl/

Clone the repository:

```shell
git clone https://github.com/BaseMax/SoftwareMarketGoAPI.git
```

Navigate to the project directory:

```shell
cd SoftwareMarketGoAPI
```

Install the dependencies:

```shell
go get -u github.com/.../...
go get -u github.com/.../...
go get -u github.com/.../...
```

Set up the database connection by modifying the config.go file according to your MySQL database credentials.

Run the migrations to create the necessary tables:

```shell
go run migrations/migrate.go
```

## Usage

Start the API server:

```shell
go run main.go
```

The server will start running on http://localhost:8080.

Use a REST client (e.g., cURL, Postman) to interact with the API endpoints.

Copyright 2023, Max Base
