#!/usr/bin/sh

go run github.com/hajimehoshi/wasmserve@latest ./cmd/gophengine &

# Give the server some time to start and open in default browser
sleep 2
xdg-open "http://localhost:8080" || open "http://localhost:8080" || start "http://localhost:8080" || echo "Please open http://localhost:8080 in your browser"

wait
