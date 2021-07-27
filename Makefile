ifeq ($(CICD_MODE),true)
	ENVFILE_BASE=.env.cicd
else
	ENVFILE_BASE=.env.example
endif

.env:
	cp $(ENVFILE_BASE) .env

.PHONY: deps
deps: .env
	docker-compose run --rm golang go mod tidy

.PHONY: test
test: .env
	docker-compose run --rm golang go test -v ./...

.PHONY: build
build: .env
	docker-compose run --rm golang make _build

.PHONY: _build
_build:
	rm -rf bin
	@for dir in $(wildcard functions/*/) ; do \
		fxn=$$(basename $$dir) ; \
		GOOS=linux go build -ldflags="-s -w" -o bin/$$fxn functions/$$fxn/*.go ; \
	done

.PHONY: deploy
deploy: .env bin
	docker-compose run --rm serverless sls deploy

.PHONY: fmt
fmt: .env
	docker-compose run --rm golang go fmt ./...

.PHONY: genMocks
genMocks: .env
	docker-compose run --rm mockery --all --inpackage --keeptree
