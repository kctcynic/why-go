

build:
	docker build -t example-server .

test:
	docker run -p 8080:8888 example-server

shell:
	docker run --rm -it --user "${UID}:${UID}" -v "${PWD}/.cache/pkg":/go/pkg -v "${PWD}":/go/src/app -w /go/src/app -e GO111MODULE=on golang:1.12 bash
