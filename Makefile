.PHONY: gen help
.DEFAULT_GOAL := help

gen: ## Generate code. e.g. make gen entity=user
	go run cmd/driver/main.go -d domain -o ${entity}_gen -t templates/idriver.tmpl -e ${entity}
	go run cmd/driver/main.go -d infrastructure -o ${entity}_gen -t templates/driver.tmpl -e ${entity}
	go fmt ./...

example: ## Generate example code
	go run _examples/main.go
	go fmt ./...

help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
