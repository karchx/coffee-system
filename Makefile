run-product:
	cd apps/backend/cmd/product && go mod tidy && go mod download && \
		CGO_ENABLED=0 go run main.go

clean:
	go clean
