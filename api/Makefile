#
#
#

all: build 


build:
	#docker run --rm -it --user "${UID}:${UID}" -v "${PWD}/.cache/pkg":/go/pkg -v "${PWD}":/go/src/api -w /go/src/api -e GO111MODULE=on golang:1.12 go build -v
	docker build -t api-server .

test:
	docker run -p 8080:8888 api-server

shell:
	docker run --rm -it --user "${UID}:${UID}" -v "${PWD}/.cache/pkg":/go/pkg -v "${PWD}":/go/src/api -e GO111MODULE=on golang:1.12