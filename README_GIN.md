# 🔊 Echoserver (Gin Version)

Welcome to the Gin-powered version of Echoserver - your friendly neighborhood HTTP parrot! 🦜

## 🎯 What's This?

This is a port of the original Echoserver to the Gin framework. It maintains all the functionality of the original server but leverages Gin's features for better performance and maintainability.

## ✨ Features

- 🔄 Mirrors back any HTTP request body you throw at it
- 🏷️ Keeps your Content-Type header intact
- 📝 Logs request method and URL path
- 🤷‍♂️ Handles empty requests gracefully
- 🚀 Powered by the high-performance Gin framework
- 🧪 Includes comprehensive tests

## 🚀 Getting Started

### Installation

```bash
# Clone the repository
git clone https://github.com/adlternative/echoserver.git
cd echoserver

# Install dependencies
go mod tidy
```

### Running the Server

```bash
# Run the Gin version
go run main_gin.go

# Or build and run
go build -o echoserver_gin main_gin.go
./echoserver_gin
```

The server will start on port 8089.

### Testing

```bash
# Run the tests
go test -v
```

### Usage Examples

```bash
# Send JSON data
curl -X POST -H "Content-Type: application/json" -d '{"message": "Hello, Gin Echo!"}' http://localhost:8089

# Send plain text
curl -X POST -H "Content-Type: text/plain" -d "Echo, echo, echo..." http://localhost:8089
```

## 🔧 How It Works

The Gin version of Echoserver has these main components:

1. `echoHandlerGin`: Handles incoming requests, preserves Content-Type, and echoes back the request body
2. `setupRouter`: Configures the Gin router with appropriate middleware and routes
3. `main`: Sets up the server and starts listening on port 8089

## 📜 License

This project is available under the [MIT License](https://opensource.org/licenses/MIT).

## 🤝 Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.
