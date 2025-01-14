# Simple REST API with Golang and Gin

This project is a simple REST API built with [Golang](https://golang.org/) using the [Gin](https://gin-gonic.com/) framework. It connects to a PostgreSQL database and provides CRUD operations for a product model.

## Project Structure

```plaintext
├── cmd/
│   └── main.go        # Main entry point for the server
├── controller/
│   └── product_controller.go # Controllers for API endpoints
├── db/
│   └── conn.go        # Database connection and setup
├── model/
│   ├── product.go     # Data model for product
│   └── response.go    # Response structure for messages
├── repository/
│   └── product_repository.go # Database access for products
├── usecase/
│   └── product_usecase.go    # Business logic for products
├── Dockerfile         # Docker container configuration
├── docker-compose.yml # Docker Compose orchestration for app and database
├── .gitignore         # Files ignored by Git
└── README.md          # Project documentation
```

## Prerequisites

Before starting, you need to have:

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Setup

### Clone the repository

```bash
git clone <REPOSITORY_URL>
cd simple-rest-api-golang-gin
```

### Build and start the containers

Run the following command to build and start the Docker containers:

```bash
docker-compose up --build
```

### Access

The API will be available at: `http://localhost:8000`

### Database

The PostgreSQL database will be accessible at `localhost:5432` with the following credentials:

- **User**: postgres
- **Password**: 1234
- **Database Name**: postgres

## Endpoints

### Available Endpoints

| Method   | Route               | Description                       |
|----------|---------------------|-----------------------------------|
| `GET`    | `/ping`             | Health check endpoint             |
| `GET`    | `/products`         | Lists all products                |
| `POST`   | `/product`          | Creates a new product             |
| `GET`    | `/product/:id`      | Retrieves a product by its ID     |

### Request Examples

#### Create a Product

**POST** `/product`

Body (JSON):
```json
{
  "name": "Example Product",
  "price": 123.45
}
```

#### Get All Products

**GET** `/products`

Response:
```json
[
  {
    "id_product": 1,
    "name": "Example Product",
    "price": 123.45
  }
]
```

## Technologies Used

- **Golang**: The main language of the project
- **Gin**: Framework for HTTP APIs
- **PostgreSQL**: Relational database
- **Docker**: Containerization of the application
- **Docker Compose**: Container orchestration

## Architecture

The project follows the clean architecture principles, dividing responsibilities into different layers:

1. **Controller**: Layer defining API endpoints.
2. **Usecase**: Implementation of business rules.
3. **Repository**: Database access.
4. **Model**: Definition of data structures and domain objects.

## Contributing

Contributions are welcome! To contribute, follow these steps:

1. Fork the project.
2. Create a branch for your changes: `git checkout -b my-feature`.
3. Commit your changes: `git commit -m 'Add some feature'`.
4. Push to the remote branch: `git push origin my-feature`.
5. Open a pull request.

## License

This project is licensed under the [MIT License](LICENSE).

---

**Developer:** [Davi Speck](https://github.com/DaviSpeck)