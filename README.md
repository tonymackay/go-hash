# Overview
`go-hash` is a command-line utility written in Go that hashes a password using the `bcrypt` hashing algorithm.

## Building
Clone the repo then run the following commands:

```
clone https://github.com/tonymackay/go-hash.git
cd go-hash
go build
```

## Using

```
go-hash -password YOURPASSWORD -cost 10
```

## License
[MIT License](LICENSE)