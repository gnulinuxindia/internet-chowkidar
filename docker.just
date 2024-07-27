build tag:
    @echo "Building docker container with tag {{tag}}..."
    docker build -t tag . -f docker/Dockerfile
