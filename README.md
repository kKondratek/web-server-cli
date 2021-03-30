# web-server-cli
Simple CLI application to run local web server
## How to use
### Run
`$ go run web-server-cli.go [commands]`
### Build and run
`$ go build`

`$ ./web-server-cli [commands]`
### Commands
```
run, r      Run local http server

version, v  Print the version

help, h     Shows a list of commands or help for one command
```
### Run tests
`$ go test`
### Try with sample
`$ web-server-cli run --file ./sample/index.html`
