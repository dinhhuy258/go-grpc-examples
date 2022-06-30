# Zuul

## How to start

### Start service registry

1. cd service-registry
2. ./gradlew bootRun

### Start customer service

1. cd customer-service
2. ./gradlew bootRun

### Start zuul proxy

1. cd zuul proxy
2. ./gradlew bootRun

## List of server

- Eureka server: http://localhost:8761/
- Customer service: http://localhost:8081/
- Zuul proxy: http://localhost:8080/

**Get list of API gateway routes**

http://localhost:8080/actuator/routes

**Access customer service via api gateway**

http://localhost:8080/customer-service/customers
