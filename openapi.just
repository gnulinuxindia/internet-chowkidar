validate:
    @echo "Validating OpenAPI spec..."
    bash ./spec/bin/validate.sh

build: validate
    @echo "Building OpenAPI spec..."
    bash ./spec/bin/build.sh
