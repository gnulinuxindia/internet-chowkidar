#!/bin/sh
docker run -t --rm -v $(pwd):/app -v ~/.cache/golangci-lint/v1.57.2:/root/.cache -w /app golangci/golangci-lint:v1.57.2 golangci-lint run -v
