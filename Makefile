.PHONY: build

test:
	go test ./...
build: test
	go build -o terraform-provider-daizy

plan:
	pushd examples &&\
	TF_CLI_CONFIG_FILE=/Users/dan/repos/github/cuotos/terraform-provider-daizy/examples/dev.tfrc terraform plan &&\
	popd

apply:
	pushd examples &&\
	TF_CLI_CONFIG_FILE=/Users/dan/repos/github/cuotos/terraform-provider-daizy/examples/dev.tfrc terraform apply &&\
	popd

init:
	pushd examples &&\
	TF_CLI_CONFIG_FILE=/Users/dan/repos/github/cuotos/terraform-provider-daizy/examples/dev.tfrc terraform init &&\
	popd
