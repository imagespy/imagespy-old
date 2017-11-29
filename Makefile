build: build_ui build_server build_image

build_server:
	docker run --rm -v "$(PWD):/go/src/github.com/wndhydrnt/imagespy" golang:1.9.1-alpine sh -c 'cd /go/src/github.com/wndhydrnt/imagespy/app && go build -v -o imagespy'

build_ui:
	docker run --rm -v "$(PWD)/ui:/ui" --workdir /ui node:9.2.0 bash -c 'PUBLIC_URL=/app yarn build'

build_image:
	docker build -t imagespy .
