# CI tests and build

source ./venv/bin/activate

GO111MODULE=on go vet ./...
GO111MODULE=on go test ./...
errcheck .
GO111MODULE=on CGO_ENABLED=0 go build -a -installsuffix cgo .
behave
