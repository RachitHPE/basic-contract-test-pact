# basic-contract-test-pact

Steps to run this project-->

Install pact tool if u want to execute via cli commands (optional):

1) curl -LO https://github.com/pact-foundation/pact-ruby-standalone/releases/download/v2.0.12/pact-2.0.12-linux-x86_64.tar.gz
2) tar xzf pact-2.0.12-linux-x86_64.tar.gz
3) export PATH=$PATH:/home/rachit/pact/bin (Add to path)
4) pact-broker version



Run the service-->

1) Simply run the command go run cmd/main.go on 1 terminal to start the service

2) Open the same project in another terminal

3) Run command go test -v ./...

4) Both the client side contract testing and provider side testing will execute with a contract json file created in pact folder
