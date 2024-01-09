# basic-contract-test-pact

Steps to run this project-->

Install pact tool if u want to execute via cli commands (optional):
curl -LO https://github.com/pact-foundation/pact-ruby-standalone/releases/download/v2.0.12/pact-2.0.12-linux-x86_64.tar.gz
tar xzf pact-2.0.12-linux-x86_64.tar.gz
export PATH=$PATH:/home/rachit/pact/bin (Add to path)
pact-broker version

Run the service-->
simply run the command go run cmd/main.go on 1 terminal to start the service
open the same project in another terminal
run command go test -v ./...
Both the client side contract testing and provider side testing will execute with a contract json file created in pact folder