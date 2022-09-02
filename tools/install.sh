#!/bin/bash

echo "start tool install" &&

go install github.com/google/wire/cmd/wire@latest
go install github.com/schemalex/schemalex/cmd/schemadiff@latest &&
go install github.com/schemalex/schemalex/cmd/schemalex@latest &&
go install github.com/schemalex/schemalex/cmd/schemalint@latest &&

echo "finish"
