#!/usr/bin/sh

echo "== Building GophEngine =="
if ! make IS_RELEASE=false; then
    echo ""
    echo "Build failed with exit code: $?"
    exit 1
fi

GOOS=$(go env GOOS)
GOARCH=$(go env GOARCH)
GOEXE=$(go env GOEXE)

logLevel=debug

echo "== Running GophEngine with log-level=${logLevel} =="
if ! ./bin/${GOOS}-${GOARCH}/gophengine${GOEXE} --log-level=${logLevel}; then
    echo ""
    echo "Execution failed with exit code: $?"
    exit 1
fi
