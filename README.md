# RE Partners Coding Challenge - Golang Pack Shipment Calculator (gRPC)

This Golang application calculates the optimal number of packs to ship to customers based on their order quantity and available pack sizes. It adheres to the following rules:

1. **Whole Packs Only:** Only complete packs can be shipped; no breaking packs open.
2. **Minimize Items:** Fulfill the order with the least number of items possible.
3. **Minimize Packs:** While adhering to rule #2, use the fewest number of packs.

## Features

* **Flexible Pack Sizes:** Pack sizes can be easily configured and modified without code changes.
* **gRPC Service:** Provides a gRPC service to calculate pack shipments for given order quantities.
* **User Interface:** Includes a simple UI to interact with the gRPC service.
* **Unit Tests:** Comprehensive unit tests ensure the correctness of the pack calculation logic.

## Getting Started with Docker Compose

 **Build and run with Docker Compose:** 
 ```bash
 docker-compose up --build
```

This will:

* Build the Go application image.
* Start the gRPC server container.
* Start the UI client container.

4. **Access the UI:** Open your web browser and navigate to `http://localhost:8080` (or the port specified in your `docker-compose.yml`).

## Configuration

* **Pack Sizes:** Modify the `packSizes` array in the configuration file or database to adjust available pack sizes.


## Testing

* **Unit Tests:** Run the unit tests using
```bash
  go test ./...
```
