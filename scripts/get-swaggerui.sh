#!/usr/bin/env bash

version=3.25.4

wget https://github.com/swagger-api/swagger-ui/archive/v${version}.tar.gz
tar xvfz v${version}.tar.gz
mkdir ./public
mv ./swagger-ui-${version}/dist/ ./public/swagger-ui
