#!/usr/bin/env bash
set -e

pushd ./lancher/dist
./lancher server.out.js
popd 