# API Server

This project is a simple API server written in Go. It provides functionality for generating short URLs from long URLs and retrieving all stored URLs. The server is structured with versioned APIs (`v1`) for better scalability and maintainability.

## Features

- **Generate Short URLs**: Convert long URLs into short, unique identifiers.
- **Retrieve All URLs**: Fetch all stored URLs in JSON format.
- **Versioned API**: Organized under `/v1` for future extensibility.

## Endpoints

### `/v1/shorturl/add`
- **Method**: `POST`
- **Description**: Adds a new URL to the server.
- **Request Body**:
  ```json
  {
    "longurl": "https://example.com"
  }
  ```
- **Response**:
  ```json
  {
    "message": "The LongURL is https://example.com and the ShortURL is abc123!"
  }
  ```

### `/v1/shorturl/get`
- **Method**: `GET`
- **Description**: Retrieves all stored URLs.
- **Response**:
  ```json
  {
    "data": [
      {
        "longurl": "https://example.com",
        "shorturl": "abc123!"
      }
    ]
  }
  ```

### `/v1/shorturl/get/:id`
- **Method**: `GET`
- **Description**: Retrieve stored URLS based of the given shorturl
- **Request**: 
```bash
curl -X GET http://localhost:8080/v1/shorturl/get/hhrtpsd
```  
- **Response**:
```json
 {
  "longurl": "https://example.com",
  "shorturl": "hhrtpsd",
 }
```
### `/v1/shorturl/delete/:id`
- **Method**: `DELETE`
- **Descritpion**: Deletes the URL based on the given shorturl
- **Request**:
```bash
curl -X DELETE http://localhost:8080/v1/shorturl/DELETE/hhrtpsd
``` 
- **Response**:
```
URL was correctly Deleted
```

"URL was correctly" 
## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/FriendlyJJC/api_server.git
   cd api_server
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the server:
   ```bash
   go run main.go
   ```

## Usage

Start the server and access the endpoints using tools like `curl`, Postman, or a browser.

Example:
```bash
curl -X POST -H "Content-Type: application/json" -d '{"longurl":"https://example.com"}' http://localhost:8080/v1/shorturl/add
```

## Project Structure

```
api_server/
├── apiv1/
│   ├── apiv1.go       # API version 1 implementation
├── main.go            # Entry point for the server
├── go.mod             # Go module file
└── README.md          # Project documentation
```

## Future Improvements

- Add database integration for persistent storage.
- Implement authentication and authorization.
- Add support for custom short URLs.
- Improve error handling and logging.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Author

Created by **FriendlyJJC**.

---

Feel free to reach out for any questions or suggestions!