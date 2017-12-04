GOPACKAGES = `go list -f '{{join .Deps "\n"}}' | grep -v vendor/ | xargs go list -f '{{if not .Standard}}{{.ImportPath}}{{end}}'`
CWD = app

.PHONY: all

project_id := ${PROJECT_ID}
version := ${GAE_VERSION}

help: ## Print this help
		@echo 'Usage: make [target]'
		@echo ''
		@echo 'Targets:'
		@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

setup: ## setup
		dep ensure

serve: ## ローカルサーバ実行
		dev_appserver.py $(CWD)

deploy: ## gaeへデプロイ OPTIONS: project_id=${PROJECT_ID} version=${VERSION}
		goapp deploy -application ${project_id} -version ${version} $(CWD)

test: ## Test実行
		cd $(CWD) && goapp test -cover -v $(GOPACKAGES)
