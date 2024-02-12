# Webservice with golang PoC  

## Description

This project is a simple Go web service that connects to a PostgreSQL database. It allows users to add and read messages through HTTP requests.

## Prerequisites

- Docker
- Docker Compose
- Make (optional, for convenience)

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Running the Application

1. **Clone the repository**

2. **Start the PostgreSQL database and the Go application:**

   ```bash
   docker-compose up -d
   ```

   This command builds and starts the containers defined in your `docker-compose.yml`. The `-d` flag runs them in detached mode.

3. **Create the database and table(s):**

   - To create the database, run:

     ```bash
     make create-db
     ```

   - To create the table, run:

     ```bash
     make create-table
     ```

### Interacting with the Application

#### Adding a Message

To add a message to the database, send a POST request to the `/add` endpoint with the message content. You can use `curl` for this purpose:

```bash
curl -X POST http://localhost:8080/add -d "content=Your message here"
```

This will add a new message with the content "Your message here" to the database.

#### Reading Messages

To read messages from the database, send a GET request to the `/read` endpoint. You can simply use `curl` or visit the URL in a web browser:

```bash
curl http://localhost:8080/read
```

This will retrieve all messages stored in the database and display them.

#### Updating a Message

To update an existing message, send a POST request to the `/update` endpoint with the message ID and the new content. Use `curl` like this:

```bash
curl -X POST http://localhost:8080/update -d "id=1&content=Updated message content"
```

This will update the message with ID 1 to have the new content "Updated message content".
