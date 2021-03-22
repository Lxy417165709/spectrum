#!/bin/sh
if [ !  $1 ]; then
  go run ./service/mvp/main.go
else
  go run ./service/mvp/main.go -env $1
fi