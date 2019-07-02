#
# This only works in the docker image because of go compiler version differences
#
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o cli .

shell:
	docker run --rm -it --user "${UID}:${UID}" -v "${PWD}/.cache/pkg":/go/pkg -v "${PWD}":/go/src/cli -w /go/src/cli -e GO111MODULE=on golang:1.12 bash
