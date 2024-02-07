prisma-migrate-dev:
	go run github.com/steebchen/prisma-client-go migrate dev


prisma-generate:
	go run github.com/steebchen/prisma-client-go generate

tests:
	go test -v ./...


server:
	go run ./cmd/main.go