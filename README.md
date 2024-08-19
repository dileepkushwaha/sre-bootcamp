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

## Setup and Running

To set up and run the API and its dependent services, follow these steps:

1. **Start the Database Container**

   ```bash
   make db-start

