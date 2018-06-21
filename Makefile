all: generate

fmt:
	go fmt ./...

install-deps:
	go get github.com/jteeuwen/go-bindata/...
	go get github.com/elazarl/go-bindata-assetfs/...

glide-install:
	glide install --force

dep-ensure:
	dep ensure

logrus-fix:
	rm -fr vendor/github.com/Sirupsen
	find vendor -type f -exec sed -i 's/Sirupsen/sirupsen/g' {} +

generate: clean generate-proto

generate-proto:
	protoc --gogofaster_out=. -Iproto onnx.proto3 

clean-proto:
	rm -fr *pb.go

clean: clean-proto

travis: install-deps dep-ensure logrus-fix shared
	echo "building..."
	go build

shared:
	mkdir -p build
	go build -buildmode=c-shared -o build/onnx_go.so main.go
	# go build -buildmode=c-archive -o build/onnx_go.a main.go
