# Student CRUD API

## Building the Docker Image

To build the Docker image for the Student CRUD API, use the following command:

```bash
docker build -t sre:1.0.0 .
```
# Student CRUD API

## Prerequisites

Ensure the following tools are installed on your local machine:
- Docker
- Docker Compose
- Make

You can install these tools by following the official documentation:
- [Docker Installation](https://docs.docker.com/get-docker/)
- [Docker Compose Installation](https://docs.docker.com/compose/install/)
- [Make Installation](https://www.gnu.org/software/make/)

## Building the Docker Image

To build the Docker image for the Student CRUD API, use the following command:

```bash
docker build -t ghcr.io/dileepkushwaha/sre-bootcamp:1.0.0 .
```

## Setup and Running

To set up and run the API and its dependent services, follow these steps:

1. **Start the Database Container**

   ```bash
   make db-start
2. Start the API and Other Services

```bash
docker-compose up -d

```
## API Endpoints
Once everything is set up, you can interact with the API using the following endpoints:

Add a New Student
```
curl -X POST http://localhost:8080/api/v1/students -H "Content-Type: application/json" -d '{"name": "John Doe", "age": 21, "grade": "A"}'

```

Get All Students
```
curl http://localhost:8080/api/v1/students
```
Get a Student by ID
```
curl http://localhost:8080/api/v1/students/1
```
Update Existing Student Information
```
curl -X PUT http://localhost:8080/api/v1/students/1 -H "Content-Type: application/json" -d '{"name": "John Smith", "age": 22, "grade": "B"}'

```
Delete a Student Record
```
curl -X DELETE http://localhost:8080/api/v1/students/1

```
Health Check
```
curl http://localhost:8080/healthcheck
```
