#!/usr/bin/env bash

cd ui && npm start
cd ..
go build -o build/dashboard && ./build/dashboard
