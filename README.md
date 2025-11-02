# Airline Voucher App

This project consists of a backend service and a frontend application, both containerized using Docker.

## Prerequisites

- Docker
- Docker Compose

## Running the Application

1. Clone the repository:
```bash
git clone git@github.com:harri88/airline-voucher-app.git
cd airline-voucher-app
```

2. Start the services using Docker Compose:
```bash
docker-compose up -d
```

This will start both the backend and frontend services.

- Frontend will be available at: `http://localhost:3000`
- Backend will be available at: `http://localhost:8080`

To stop the services:
```bash
docker-compose down
```

## Services

- **Frontend**: React application
- **Backend**: Golang API service

## Environment Variables

Create a `.env` file in the root directory with the following variables:

```env
BACKEND_PORT=8080
FRONTEND_PORT=3000
```

## Monitoring

You can check the status of your containers using:
```bash
docker-compose ps
```

For logs:
```bash
docker-compose logs -f
```
