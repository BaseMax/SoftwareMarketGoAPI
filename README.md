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

## API Routes

### `GET /software`

Retrieves all software available in the marketplace.

### `GET /software/:id`

Retrieves a specific software by its ID.

### `POST /software`

Creates a new software listing in the marketplace. Requires the following JSON fields in the request body:

  - name (string): Name of the software.
  - description (string): Description of the software.
  - price (float): Price of the software.

### `PUT /software/:id`

Updates an existing software listing in the marketplace. Requires the following JSON fields in the request body:

  - name (string): Updated name of the software.
  - description (string): Updated description of the software.
  - price (float): Updated price of the software.

### `DELETE /software/:id`

Deletes a software listing from the marketplace based on its ID.

### `GET /software/category/:category`

Retrieves all software listings in a specific category. The :category parameter represents the category of the software.

### `GET /software/search?query=:searchQuery`

Searches for software listings based on a given search query. The :searchQuery parameter represents the search query string.

### `GET /software/top-rated`

Retrieves the top-rated software listings in the marketplace based on user ratings.

### `GET /software/recommended`

Retrieves a list of recommended software listings for the current user. The recommendations can be based on user preferences, previous purchases, or any other relevant criteria.

### `POST /software/:id/ratings`

Adds a user rating for a specific software listing. Requires the following JSON fields in the request body:

rating (integer): The user's rating for the software (e.g., a value between 1 and 5).

### `GET /software/:id/reviews`

Retrieves all reviews for a specific software listing.

### `POST /software/:id/reviews`

Adds a new review for a specific software listing. Requires the following JSON fields in the request body:

  - title (string): Title of the review.
  - content (string): Content of the review.
  - user_id (integer): ID of the user who posted the review.

### `GET /users/:id/software`

Retrieves all software listings associated with a specific user. The :id parameter represents the ID of the user.

### `GET /users/:id/reviews`

Retrieves all reviews posted by a specific user. The :id parameter represents the ID of the user.

## Test Routes

### `GET /software

```shell
curl -X GET http://localhost:8080/software
```

### `GET /software/:id

```shell
curl -X GET http://localhost:8080/software/1
```

### `POST /software

```shell
curl -X POST -H "Content-Type: application/json" -d '{"name": "Software Name", "description": "Software Description", "price": 9.99}' http://localhost:8080/software
```

### `PUT /software/:id

```shell
curl -X PUT -H "Content-Type: application/json" -d '{"name": "Updated Software Name", "description": "Updated Software Description", "price": 14.99}' http://localhost:8080/software/1
```

### `DELETE /software/:id

```shell
curl -X DELETE http://localhost:8080/software/1
```

### `GET /software/category/:category

```shell
curl -X GET http://localhost:8080/software/category/category-name
```

### `GET /software/search?query=:searchQuery

```shell
curl -X GET http://localhost:8080/software/search?query=search-query
```

### `GET /software/top-rated

```shell
curl -X GET http://localhost:8080/software/top-rated
```

### `GET /software/recommended

```shell
curl -X GET http://localhost:8080/software/recommended
```

### `POST /software/:id/ratings

```shell
curl -X POST -H "Content-Type: application/json" -d '{"rating": 5}' http://localhost:8080/software/1/ratings
```

### `GET /software/:id/reviews

```shell
curl -X GET http://localhost:8080/software/1/reviews
```

### `POST /software/:id/reviews

```shell
curl -X POST -H "Content-Type: application/json" -d '{"title": "Review Title", "content": "Review Content", "user_id": 1}' http://localhost:8080/software/1/reviews
```

### `GET /users/:id/software

```shell
curl -X GET http://localhost:8080/users/1/software
```

### `GET /users/:id/reviews`

```shell
curl -X GET http://localhost:8080/users/1/reviews
```

Copyright 2023, Max Base
