#!/bin/bash
# TODO: consider rewrite


# Params
set -eu
script_dir=$(cd $(dirname $0); pwd)


# Execution
cd ${script_dir}/../

for file in `find . -maxdepth 2 -name index.yaml -type f`; do
    site=$(dirname $file | sed -e 's/\.\///g')
    pushd $site
        redocly lint --extends=minimal index.yaml
    popd
done
