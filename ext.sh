#!/bin/bash

rm -rf ext/sig ext/ext ext/fn local-testfn-latest.scale

mkdir ext/sig

echo "Creating signature"
./cmd/cmd signature new -d ext/sig

cp ext/scale.signature ext/sig/scale.signature

echo "Generating signature"
./cmd/cmd signature generate -d ext/sig testsig:latest

mkdir ext/ext

echo "Creating extension"
./cmd/cmd extension new -d ext/ext

cp ext/scale.extension ext/ext/scale.extension

echo "Generating extension"
./cmd/cmd extension generate testext:latest -d ext/ext

mkdir ext/fn

echo "Creating function"
./cmd/cmd function new -d ext/fn -s local/testsig:latest -e local/testext:latest testfn:latest

cat ext/fn_code.go > ext/fn/main.go

echo "Building function"
./cmd/cmd function build -d ext/fn

echo "Exporting function"
./cmd/cmd function export local/testfn:latest ext/

echo "Running..."

cd ext/runner

go run .
