# Software-Market Go API

SoftwareMarketGoAPI is a RESTful API built with Golang that provides endpoints to manage software products in a marketplace. It allows users to perform operations such as creating, retrieving, updating, and deleting software products.

## Features

- User authentication and authorization using JSON Web Tokens (JWT)
- CRUD operations for blog posts
- Error handling and response formatting
- Input validation and data sanitization
- Logging and request/response middleware
- Database integration using MYSQL gorm orm

## Installation

Make sure you have Go installed on your system. You can download and install it from the official Go website: https://golang.org

Clone the repository:

```shell
git clone https://github.com/BaseMax/SoftwareMarketGoAPI.git
```


Navigate to the project directory:

```shell
cd SoftwareMarketGoAPI
```

Install the project dependencies:

```bash
go mod download
```

Set up the environment variables:

- Create a `.env` file in the root directory of the project.

- Define the following environment variables in the `.env` file:

```plaintext
PORT=8080
DB_HOST=your_database_host
DB_PORT=your_database_port
DB_USER=your_database_user
DB_PASSWORD=your_database_password
DB_NAME=your_database_name
JWT_SECRET=your_jwt_secret
```

Run the application:

```bash
go run main.og
```

The API will be accessible at `http://localhost:8080` by default. You can change the port in the .env file if necessary.

## API Endpoints

The following endpoints are available in the panel:

| Method | 	Endpoint | 	Description |
| ---- | -------- | -------- |
| POST |	/panel/software/	| Add new software |
| PUT |	/panel/software/:id	| Update software |
| DELETE |	/panel/software/:id	| Delete software |
| POST |	/panel/users/admin	| Add new admin |
| POST |	/panel/users/login	| Login admin |
| PUT |	/panel/users/admin/:id	| Update admin |

The following endpoints are available in the api:

| Method | 	Endpoint | 	Description |
| ---- | -------- | -------- |
| GET |	/api/software/	| Get all softwares |
| GET |	/api/software/:id	| Get a software by id |
| GET |	/api/software/category/:category	| Get software by category |
| GET |	/api/software/search/:query	| Search softwares |
| GET |	/api/software/top-rated	| Get top rated softwares |
| GET |	/api/software/recommended	| Get recommended softwares |
| GET |	/api/software/:id/reviews	| Get software's reviews |
| POST |	/api/software/:id/ratings	| Add rate to software |
| POST |	/api/software/:id/reviews	| Add reviews to software |
| GET |	/api/users/reviews	| Get user's reviews |
| GET |	/api/users/:id/software	| Get associated software |
| POST |	/api/users/login	| Login user |
| POST |	/api/users/register	| Register uesr |



## Database Schema

The application uses a GORM orm MYSQL and have these data modles and use auto migrate:

```go

type Software struct {
	ID          uuid.UUID `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Creator     string    `json:"creator"`
	Version     string    `json:"version"`
	Category    string    `json:"category_list" `
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Rate struct {
	ID         uuid.UUID `json:"id" gorm:"primaryKey"`
	Value      uint8     `json:"value"` // value can be between 1 and 5
	SoftwareId uuid.UUID `json:"software_id"`
	UserId     uuid.UUID `json:"user_id" `
	RatedAt    time.Time `json:"Rated_at"`
}

type Review struct {
	ID         uuid.UUID `json:"id" gorm:"primaryKey"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	SoftwareId uuid.UUID `json:"software_id" `
	UserId     uuid.UUID `json:"user_id" `
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type User struct {
	ID                 uuid.UUID `json:"id" gorm:"primaryKey"`
	Name               string    `json:"name"`
	Email              string    `json:"email"`
	Password           string    `json:"password"`
	IsAdmin            bool      `json:"is_admin"`
	FavoriteCategories string    `json:"favorite_categories"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

```

## Dependencies

The project utilizes the following third-party libraries:

- `go-playground/validator/v10`: Struct field validation
- `golang-jwt/jwt/v4`: JWT implementation 
- `google/uuid`: UUID generation
- `joho/godotenv`: Environment variable loading
- `labstack/echo/v4`: HTTP web framework
- `x/crypto`: Hashing password library
- `driver/mysql`: Struct field validation
- `gorm`: Database orm


Make sure to run go mod download as mentioned in the installation steps to fetch these dependencies.

## Contributing

Contributions to this project are welcome. Feel free to open issues and submit pull requests.

## License

This project is licensed under the GPL-3.0 License.

## Postman collection

Software-Market-Go-API/software-market.postman_collection.json

## Authors

- AmirAtashghah
- Max Base

Copyright 2023, Max Base

Copyright 2023, Asrez group
