# SimpleMQ

A minimal, easy-to-understand message queue implementation written in GO intended for learning and small projects. Provides a lightweight broker with producer/consumer APIs and optional persistence.

## Overview

SimpleMQ is a simple message queue service that allows producers to publish messages to topics and consumers to retrieve messages from those topics. It uses a circular buffer-based queue data structure for efficient message storage.


## Features

- **Topic-based messaging**: Messages are organized by topics
- **Circular queue implementation**: Fixed-size queues with O(1) operations
- **In-memory storage**: Fast message access (persistence support planned)
- **RESTful API**: HTTP endpoints for producing and consuming messages
- **Error handling**: Comprehensive error responses for queue full/empty conditions

## Architecture

### Components

- **Queue Service** (`internal/service/queue.go`): Core queue management logic
- **Models** (`internal/models/models.go`): Data structures and context definitions
- **Message Handler** (`internal/handlers/message_handler.go`): HTTP request/response handling
- **Gin Framework**: Used for HTTP routing and request binding

## API Endpoints

### Produce Message

``` bash
curl --location 'localhost:8080/produce' \
--header 'Content-Type: application/json' \
--data '{
    "payload":"1234",
    "topic":"12334"
}'
```

### Consume Message
``` bash
curl --location 'localhost:8080/consume' \
--header 'Content-Type: application/json' \
--data '{
    "topic":"12334"
}'
```