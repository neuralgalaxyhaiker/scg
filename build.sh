#!/usr/bin/env bash
set -e

pushd ./nodejs
echo "build nodejs"
npm i
npm run webpack
popd

echo "lancher"
pushd ./lancher
go mod download
go build -o dist/lancher
make dist
popd
