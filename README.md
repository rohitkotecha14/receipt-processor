# Receipt Processor

The Receipt Processor is a simple web service built in Go. It processes receipt data to calculate loyalty points based on a set of business rules and exposes two endpoints:

- **POST `/receipts/process`**: Accepts a receipt JSON payload, calculates points, and returns a unique receipt ID.
- **GET `/receipts/{id}/points`**: Returns the number of points awarded for a given receipt ID.

## Prerequisites

- [Go](https://golang.org/dl/)
- [Docker](https://www.docker.com/get-started) (optional, for containerized execution)

## Running the Application Locally

1. **Clone the repository:**

   ```bash
   git clone https://github.com/rohitkotecha14/receipt-processor.git
   cd receipt-processor
   ```
2. **Build and Run**
   ```bash
   go run main.go
   ```
   Or build the binary and run:
   ```bash
   go build -o receipt-processor
   chmod +x receipt-processor
   ./receipt-processor
   ```
   The server will start on port 8080.

## Running the Application with Docker

1. **Build the Docker Image:**

   ```bash
   docker build -t receipt-processor .
   ```
2. **Run the Docker Container**
   ```bash
   docker run -d -p 8080:8080 receipt-processor
   ```

## Testing the Endpoints
1. **Process a Receipt:**
   Send a POST request to /receipts/process:
   ```bash
   curl -X POST http://localhost:8080/receipts/process \
        -H "Content-Type: application/json" \
        -d '{
                "retailer": "Target",
                "purchaseDate": "2022-01-01",
                "purchaseTime": "13:01",
                "items": [
                    {"shortDescription": "Mountain Dew 12PK", "price": "6.49"},
                    {"shortDescription": "Emils Cheese Pizza", "price": "12.25"},
                    {"shortDescription": "Knorr Creamy Chicken", "price": "1.26"},
                    {"shortDescription": "Doritos Nacho Cheese", "price": "3.35"},
                    {"shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ", "price": "12.00"}
                ],
                "total": "35.35"
        }'
   ```
   Expected Response:
   ```bash
   { "id": "uuid" }
   ```
   
2. **Get Receipt Points**:
   ```bash
   curl http://localhost:8080/receipts/<id>/points
   ```