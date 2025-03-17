# 🔊 Echoserver

Hey there! Welcome to Echoserver - your friendly neighborhood HTTP parrot! 🦜

## 🎯 What's This?

Echoserver is a super simple HTTP server that does one thing really well - it echoes back whatever you send to it! Perfect for testing your HTTP clients, checking network connectivity, or just having a digital conversation partner that always agrees with you! 😉

## ✨ Cool Stuff It Does

- 🔄 Mirrors back any HTTP request body you throw at it
- 🏷️ Keeps your Content-Type header intact (it respects your style!)
- 📝 Keeps a log of who's talking to it (method and URL path)
- 🤷‍♂️ Handles empty requests without complaining
- 🧩 Zero external dependencies - just pure Go standard library goodness!

## 🚀 Let's Get Started!

### Fire It Up

```bash
# Run it on the fly
go run main.go

# Or build yourself a shiny executable
go build -o echoserver
./echoserver
```

Your echo chamber will be ready and waiting on port 8089! 🎉

### Talk To It

Chat with your new echo friend using any HTTP client:

```bash
# Send some JSON love
curl -X POST -H "Content-Type: application/json" -d '{"message": "Hello, Echo!"}' http://localhost:8089

# Or just plain text if you're feeling old-school
curl -X POST -H "Content-Type: text/plain" -d "Echo, echo, echo..." http://localhost:8089
```

Whatever you say, Echoserver says it right back! It's like having a conversation with yourself, but more technical! 🤓

## 🔧 How It Works

Under the hood, this little marvel is powered by Go and has two main parts:

1. `echoHandler`: The friendly receptionist that greets your requests, keeps your Content-Type, and copies your message to the reply
2. `main`: The manager that sets everything up and makes sure the server is listening on port 8089

## 📜 License

This project is free as in freedom! Available under the [MIT License](https://opensource.org/licenses/MIT). Go wild!

## 🤝 Join The Echo Chamber!

Got ideas? Found a bug? Want to make Echoserver even more echo-y? Contributions are super welcome! Open issues, send PRs, or just shout really loud and see if it echoes back! 😄
