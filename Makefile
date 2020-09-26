IMAGE_NAME := "touchify/cert-manager-webhook-scaleway"
IMAGE_TAG := "latest"

OUT := $(shell pwd)/_out

$(shell mkdir -p "$(OUT)")

update-client:
	./scripts/update-client.sh

verify:
	./scripts/fetch-test-binaries.sh
	go test -v .

build:
	docker build -t "$(IMAGE_NAME):$(IMAGE_TAG)" .

.PHONY: rendered-manifest.yaml
rendered-manifest.yaml:
	helm template \
	    --name cert-manager-webhook-scaleway \
        --set image.repository=$(IMAGE_NAME) \
        --set image.tag=$(IMAGE_TAG) \
        deploy/cert-manager-webhook-scaleway > "$(OUT)/rendered-manifest.yaml"
