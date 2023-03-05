run-product:
	cd apps/backend/cmd/product && go mod tidy -compat=1.17 && go mod download && \
		CGO_ENABLED=0 go run main.go

clean:
	go clean
