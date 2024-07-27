#!/bin/bash
set -eux

go generate --tags wireinject ./...
