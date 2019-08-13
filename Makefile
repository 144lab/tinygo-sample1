PWD := $(shell pwd)
PACKAGE_PATH := $(shell go mod why | tail -1)
PROJECT_PATH := /go/src/$(PACKAGE_PATH)
nRF_SDK := "/Users/nobo/Documents/SEGGER Embedded Studio for ARM Projects/nRF5_SDK"
nRF_SDK_PATH := /opt/nRF5_SDK

build:
	docker run -it --rm -v $(PWD):$(PROJECT_PATH) \
	-v $(nRF_SDK):$(nRF_SDK_PATH) \
	-e GOPATH=/go \
	-w $(PROJECT_PATH) tinygo/tinygo tinygo build -target custom.json -o app.bin .
	uf2conv.py app.bin -c -b 0x26000 -f 0xADA52840 -o app.uf2

deploy:
	cp app.uf2 /Volumes/NRF52BOOT/

version:
	docker run -it --rm tinygo/tinygo tinygo version