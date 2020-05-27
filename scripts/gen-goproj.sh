#!/usr/bin/env bash

set -u

proj=$1

mkdir -p $proj
cd $proj
go mod init $proj
