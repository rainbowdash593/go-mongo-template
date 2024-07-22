## Description:
- Go clean architecture template
- Mongo DB
- Gin HTTP Server

## Requirements:
- docker
- docker-compose

## Start application:

1. Copy environment variables


    cp .env.example .env.local
3. Run docker-compose command


    docker-compose -f docker/docker-compose-local.yml up --build

## Run in debug mode:
1. Check `GOARCH` variable. It should be `arm64` for apple m1 chips
2. Run docker-compose command


    docker-compose -f docker/docker-compose-debug.yml up --build
3. Connect to delve in your IDE, using port in `DELVE_PORT` variable