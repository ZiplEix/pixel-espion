# API Documentation

## Overview

This API is built using Go and the Fiber framework. It provides endpoints for user authentication and managing spies and their records. The API supports JWT-based authentication and utilizes a PostgreSQL database.

## Features

- User registration and login
- CRUD operations for spies and records
- JWT authentication for protected routes

## Getting Started

### Prerequisites (development)

- Go 1.23+
- PostgreSQL 13+
- air (for live reloading)

### Prerequisites

- Docker
- Docker Compose

### Installation

1. Clone the repository

```sh
git clone git@github.com:ZiplEix/pixel-espion.git
cd pixel-espion/api
```

2. Create a `.env` file in the root directory following the example in `.env.example`

3. Build and run the application with Docker Compose

```sh
docker-compose up --build
```

This will start the API server and PostgreSQL database. The API will be available at `http://localhost:<PORT>`.

3. If you want to run the application without Docker, you can do so by running the following commands:

start the PostgreSQL database

```sh
docker run --name pixel-espion-db -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=pixel_espion -p 5432:5432 -d postgres:13
```

start the API server

```sh
go run main.go

# or with live reloading
air
```

## API Endpoints

For a detailed description of the available API endpoints, please refer to the Swagger Documentation after starting the server at `http://localhost:<PORT>/swagger/index.html`.

## Authentication

- **Login**: Use the /login endpoint to obtain a JWT token.
- **Register**: Use the /register endpoint to create a new user.

> _(The JWT will be stored in a cookie for subsequent requests.)_

## Protected routes

Most routes for managing spies and records require authentication. Ensure to include the JWT token in the Authorization header for these requests.

## Testing

Run the tests with the following command:

```sh
go test ./...
```

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue.

## License

Distributed under the MIT License. See `LICENSE` for more information.
