# go-easy-qr

![go-easy-qr-code](https://github.com/ziyacivan/go-easy-qr/blob/master/pr-qr-code.png)

🎱 This repository includes a QR Code generator Web API written in Golang.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [License](#license)

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/ziyacivan/go-easy-qr.git
    cd go-easy-qr
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

## Usage

1. Run the application:
    ```sh
    go run main.go
    ```

2. The server will start on `http://localhost:3000`.

## API Endpoints

### General Route

- **URL:** `/api/v1/`
- **Method:** `GET`
- **Response:**
  - Content-Type: `application/json`
  - Body: `{"message": "Welcome to Easy QR API", "routes": {"generate": "http://localhost:3000/api/v1/generate"}}`

- **Example:**
    ```sh
    curl "http://localhost:3000/api/v1/"
    ```

### Generate QR Code

- **URL:** `/api/v1/generate`
- **Method:** `POST`
- **Query Parameter:**
  - `data` (string): The data to encode in the QR code.

- **Response:**
  - Content-Type: `image/png`
  - Body: The generated QR code image.

- **Example:**
    ```sh
    curl -X POST "http://localhost:3000/api/v1/generate?data=https://example.com"
    ```

## Docker

### Build the Docker image

1. Build the Docker image:
    ```sh
    docker build -t go-easy-qr .
    ```

### Run the Docker container

2. Run the Docker container:
    ```sh
    docker run -p 3000:3000 go-easy-qr
    ```

The server will start on `http://localhost:3000`.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
