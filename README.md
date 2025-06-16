# GraphQL Federation Example

This repository demonstrates a GraphQL Federation setup using Apollo Router and three distinct subgraphs: `Product`, `Review`, and `Shipping`. It showcases how to combine multiple independent GraphQL services into a single, unified supergraph.

## Architecture

The project consists of the following components:

*   **Apollo Router**: The entry point for all GraphQL queries. It orchestrates requests across the different subgraphs.
*   **Product Service**: A subgraph responsible for product-related data (e.g., `id`, `name`, `price`, `weight`).
*   **Review Service**: A subgraph handling product reviews (e.g., `id`, `body`).
*   **Shipping Service**: A subgraph providing shipping information, including estimated delivery times, which depends on data from the Product service.

## Getting Started

To get this project up and running, follow these steps:

### Prerequisites

*   [Docker](https://www.docker.com/get-started) (Docker Compose included)
*   [Go](https://golang.org/doc/install) (for building the Go services, though Docker handles this)

### 1. Compose the Supergraph

The supergraph schema is composed using the Apollo Rover CLI. This step generates the `supergraph.graphql` file that the Apollo Router uses.

```bash
make compose-supergraph
```

This command will:
*   Pull the `worksome/rover:latest` Docker image if not already present.
*   Run the Rover CLI in a Docker container to compose the supergraph based on `router/supergraph-config.yaml`.
*   Output the combined schema to `router/supergraph.graphql`.

### 2. Run the Services

Once the supergraph is composed, you can start all services using Docker Compose:

```bash
make start
```

This command will:
*   Build the Docker images for the `product`, `review`, and `shipping` services.
*   Start all four services (`router`, `product`, `review`, `shipping`) in detached mode.

## Services Overview

*   **Product Service**: Exposes product details.
    *   Endpoint: `http://localhost:8081/query`
*   **Review Service**: Manages product reviews.
    *   Endpoint: `http://localhost:8082/query`
*   **Shipping Service**: Provides shipping estimates.
    *   Endpoint: `http://localhost:8083/query`

## Usage

Once all services are running, the GraphQL supergraph endpoint will be available at:

```
http://localhost:4000/
```

You can use a GraphQL client (like Apollo Sandbox, GraphQL Playground, or Insomnia) to send queries to this endpoint and interact with the federated graph.

Example Query:

```graphql
query GetProductList {
  productList {
    id
    name
    price
    weight
    reviews {
      id
      body
    }
    estimatedDeliveryTime
  }
}
```

Alternatively, you can access `http://localhost:4000/graphql` to enter the router sandbox mode, where instrospection can be done, and sample queries can be constructed more easily before running.