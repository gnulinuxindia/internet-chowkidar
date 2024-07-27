#!/bin/bash

MOCK_USE_HEADER="//use:mock"

find . -type f -name "*.go" | while read -r file; do
    if [[ $(awk 'NR==1' "$file") == $MOCK_USE_HEADER ]]; then
        dir=$(dirname "$file")
        mock_dir="${dir}/mock"
        mkdir -p "$mock_dir"
        echo "mock(generating): $file"
        mockgen -source="$file" -destination="${mock_dir}/mock_$(basename "$file")"
    fi
done
