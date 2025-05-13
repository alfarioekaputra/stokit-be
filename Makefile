up: dev-nodemon

dev-nodemon: $(nodemon) ## start nodemon ( Continous Development app)
	nodemon --exec go run cmd/main.go --signal SIGTERM

build: ## Builds binary
	@ echo "Building aplication... "
	@ GOOS=linux GOARCH=amd64
	@ go build \
		-ldflags "-s -w" \
		-trimpath  \
		-o engine \
		./cms/
	@ echo "done"
