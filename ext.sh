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

echo "Updating runner go.mod..."

sig=`cat ext/fn/go.mod | grep testsig_latest | sed -e "s/guest/host/"`
ext=`cat ext/fn/go.mod | grep testext_latest | sed -e "s/guest/host/"`

echo "Signature is " ${sig}
echo "Extension is " ${ext}

# Get rid of any old or previous replacements
cat ext/runner/go.mod | grep -v testsig_latest | grep -v testext_latest > ext/runner/go.mod.new
mv ext/runner/go.mod.new ext/runner/go.mod

# Now insert the correct locations
echo "" >> ext/runner/go.mod
echo ${sig} >> ext/runner/go.mod
echo ${ext} >> ext/runner/go.mod

echo "Running..."

cd ext/runner

go run .
