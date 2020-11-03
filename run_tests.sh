#!/bin/bash

set -e

#cd /go/src/inventory
#go test -v -run TestAPI

cd /go/src/rbac
go test -v -run TestAPI

cd /go/src/vmaas
go test -v -run TestAPI
