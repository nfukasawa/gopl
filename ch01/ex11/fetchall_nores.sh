#!/bin/bash
set -eux

go run noressrv/main.go &
go run main.go http://localhost:8989
rm results_*