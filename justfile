mod migration
mod openapi
mod docker

all:
    bash ./bin/allgen.sh

mocks:
    bash ./bin/mockgen.sh

lint:
    bash ./bin/lint.sh

providers:
    bash ./bin/providergen.sh

ent:
    bash ./bin/dbgen.sh

api:
    bash ./bin/apigen.sh

wire:
    bash ./bin/gogen.sh

run:
    go run .

build:
    GOOS=linux GOARCH=amd64 go build -o main .
