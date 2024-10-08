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

**Prerequisites**

* **Configure Environment Variables:** Before building, create a `.env` file in the project root with the following content:

```
SERVER_PORT=20000
CLIENT_PORT=9090
```

* **Configure Pack Sizes:** Create a `config.yaml` file in the project root with the desired pack sizes:

```yaml
pack_sizes:
  - 500
  - 900
```

**Build and run with Docker Compose:**

```bash
docker-compose up --build
```

This will:

* Build the Go application image.
* Start the gRPC server container.
* Start the UI client container.

**Access the UI:** Open your web browser and navigate to `http://localhost:9090` (or the port specified in your `docker-compose.yml` under `CLIENT_PORT`).

## Configuration

* **Pack Sizes:** Modify the `pack_sizes` list in the `config.yaml` file to customize the available pack sizes.

## Testing

* **Unit Tests:** Run the unit tests using

```bash
go test ./...
```

**Important:** Ensure you have Docker and Docker Compose installed on your system before running the application.

Please let me know if you have any other questions or requests. 
