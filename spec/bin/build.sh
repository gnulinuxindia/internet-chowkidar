#!/bin/bash

# Params
set -eu
script_dir=$(cd $(dirname $0); pwd)

# Execution
cd ${script_dir}/../

for file in `find . -maxdepth 2 -name index.yaml -type f`; do
    site=$(dirname $file | sed -e 's/\.\///g')
    build_dir="../build/${site}"

    pushd $site
        mkdir -p ${build_dir}/
		snippet-enricher-cli --input=index.yaml > ${build_dir}/index.enriched.json
        redocly bundle ${build_dir}/index.enriched.json -o ${build_dir}/index.json --ext json
		# File will be in spec/index.html
        redocly build-docs ${build_dir}/index.json -o ${build_dir}/../../index.html
    popd
done
