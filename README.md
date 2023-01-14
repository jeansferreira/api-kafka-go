# events-poc-api

A proof of concept application.

## How to run?

### Requirements

- Docker
- Docker compose

### Running the application

To run the application, run:

```sh
docker compose up -d
```

This will create the following services:

- Zookeeper (2181)
- Kafka (9092)
- Kafka UI (8080)
- Server (3000)

### Creating topics

For the application to work, you need to create the following topics:

- `events_home` - for the home events;
- `events_product` - for the user events;

> This can be done using the Kafka UI.
