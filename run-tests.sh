#!/bin/bash

cd diacritic
go test -v
cd ..

cd email
go test -v
cd ..

cd web
go test -v
cd ..

cd datetime
go test -v
cd ..

cd slice
go test -v
cd ..

cd str
go test -v
cd ..

