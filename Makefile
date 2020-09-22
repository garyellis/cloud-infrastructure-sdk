VERSION=v0.2.8
NAME=cloud-infrastructure-sdk


.PHONY: help
	.DEFAULT_GOAL := help

help: ## show this message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'



build: ## build the binary
	docker build --build-arg VERSION=$(VERSION) -t build/$(NAME) .
	docker create --name $(NAME)  build/$(NAME)
	docker cp $(NAME):/release/ .
	docker rm $(NAME)

release: ## release the binary and docker image
	echo 'release to somewhere'
