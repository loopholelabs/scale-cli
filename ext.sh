#!/bin/bash

rm -rf ext/sig ext/ext ext/fn local-testfn-latest.scale

# First lets create a signature
mkdir ext/sig
echo "Creating signature"
./cmd/cmd signature new -d ext/sig
cp ext/scale.signature ext/sig/scale.signature
echo "Generating signature"
./cmd/cmd signature generate -d ext/sig testsig:latest

# Now create the extension
mkdir ext/ext
echo "Creating extension"
./cmd/cmd extension new -d ext/ext
cp ext/scale.extension ext/ext/scale.extension
echo "Generating extension"
./cmd/cmd extension generate testext:latest -d ext/ext

# Create a function using the extension in GO
mkdir ext/fn-go
echo "Creating function"
./cmd/cmd function new -d ext/fn-go -s local/testsig:latest -e local/testext:latest testfngo:latest
cat ext/fn_code.go > ext/fn-go/main.go
echo "Building function"
./cmd/cmd function build -d ext/fn-go
echo "Exporting function"
./cmd/cmd function export local/testfngo:latest ext/

# Create a function using the extension in Rust
mkdir ext/fn-rs
echo "Creating function"
./cmd/cmd function new -d ext/fn-rs -s local/testsig:latest -e local/testext:latest -l rust testfnrs:latest
# cat ext/fn_code.rs > ext/fn-rs/main.rs
echo "Building function"
./cmd/cmd function build -d ext/fn-rs
echo "Exporting function"
./cmd/cmd function export local/testfnrs:latest ext/

# Create a function using the extension in Typescript
mkdir ext/fn-ts
echo "Creating function"
./cmd/cmd function new -d ext/fn-ts -s local/testsig:latest -e local/testext:latest -l rust testfnts:latest
# cat ext/fn_code.ts > ext/fn-ts/main.ts
echo "Building function"
./cmd/cmd function build -d ext/fn-ts
echo "Exporting function"
./cmd/cmd function export local/testfnts:latest ext/

# Sort out runner
echo "Updating runner go.mod..."
sig=`cat ext/fn-go/go.mod | grep testsig_latest | sed -e "s/guest/host/"`
ext=`cat ext/fn-go/go.mod | grep testext_latest | sed -e "s/guest/host/"`
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
