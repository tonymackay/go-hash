# Overview
`go-hash` is a command-line utility written in Go that hashes a password using the `bcrypt` hashing algorithm.

## Building
Clone the repo then run the following commands:

```
git clone https://github.com/tonymackay/go-hash.git
cd go-hash
go build -o go-hash
go install
```

## Using


```
Usage: go-hash [OPTIONS]

Options:
  -cost int
        Cost between 4 and 31 (default 10)
  -password string
        Password to hash

Examples:
  go-hash -password mypassword
  go-hash -password "hello world" -cost 20
```

## License
[MIT License](LICENSE)