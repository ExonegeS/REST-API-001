# REST-API-001
My implementation of a simple REST API backend using Golang &amp; Postgresql to get/create/update users

# Project Setup Instructions

## Prerequisites

- Docker
- Docker Compose
- Make (optional but recommended for easier setup)

### Setting Up the Project

1. **Clone the repository:**
    
    If you haven't already cloned the repository, do so with the following command:

    ```shell
    git clone <repository-url> cd <project-folder>
    ```

2. **Set up environment variables:**

    Copy the `.env.example` to `.env` for local development:

    ```shell
    cp .env.example .env
    ```

3. **Build the project:**
    
    To build the project, run the following `make` command:

    ```shell
    make build
    ```
    This will compile the Go code and place the resulting binary in the `bin/app` directory.

4. **Run the project:**

    Start the project locally by using:

    ```shell
    make run
    ```

    This will start the Go backend server on port `8888` using the configuration from `.env`.

### Running the Project with Docker Compose

To set up and run the project using Docker Compose, follow these steps:

1. **Build and run the services:**

    First, use Docker Compose to build and run the services (backend and PostgreSQL database) by executing:

    ```shell
    docker-compose up --build
    ```


    This will start both the backend API server and the PostgreSQL database. The backend will be available at `http://localhost:4004`.

2. **Stop the services:**

    To stop the services, run:

    ```shell
    docker-compose down
    ```

    This stops the containers and removes them.

### Database Configuration

1. **PostgreSQL Setup:**
    The project uses PostgreSQL as the database. Make sure that you have set up your `.env` or `.env.production` correctly with the appropriate database credentials. The relevant environment variables for PostgreSQL are:

    ```env
    DATABASE_HOST=localhost
    DATABASE_PORT=5444
    DATABASE_USER=admin
    DATABASE_PASS=adminadmin
    DATABASE_DB=users
    ```

2. **Exporting the Database Schema:**

    If you want to export your PostgreSQL database schema, you can run:

    ```shell
    pg_dump -U admin -h localhost -p 5444 --no-data --schema=public users > schema.sql
    ```

    This will create a file `schema.sql` containing the structure of your database.

3. **Recreating the Database Schema:**

    If you need to set up the schema on a fresh PostgreSQL instance, run the following command:

    ```shell
    psql -U admin -h localhost -p 5444 -d users -f schema.sql
    ```

    Make sure the PostgreSQL server is running before executing this command.


### Docker

The project includes Docker and Docker Compose configuration. The backend service is built using the Dockerfile, and PostgreSQL is used as the database.

#### Dockerfile

The `Dockerfile` is used to create the container for the Go backend. The backend service is built from the Go code, and the binary is copied into a minimal Alpine Linux image. It listens on port `8888` and is ready to run when the container starts.

#### docker-compose.yml

This file defines the `backend` and `db` services:

- **backend**: This service is built from the `Dockerfile` and runs the Go binary. It listens on port `8888`.
- **db**: This service runs a PostgreSQL container on port `5444`.

To run the services with Docker Compose, run the following command:

```shell
docker-compose up --build
```

### Troubleshooting

- **Error: Duplicate email during user creation**
If you encounter an error about a duplicate email, make sure the email is unique in the database. The `email` field is set to be unique in the database schema.

- **Error: User not found**
If you attempt to update or delete a user that does not exist, you will receive a "user not found" error. Ensure the user exists in the database before performing operations.


