#!/bin/bash
set -eux

# build spec into a single yaml
bash spec/bin/build.sh

# go server codegen from spec yaml
ogen --package genapi --target api/gen --clean spec/build/internet-chowkidar/index.yaml
