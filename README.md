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

Copyright 2023, Max Base
