pre-commit:
	pre-commit autoupdate
	pre-commit run --all-files

bump:
	cz bump
	git push
	git push --tags

snapshot:
	goreleaser --snapshot --clean

release:
	goreleaser --verbose --clean

build:
	go mod tidy
	CGO_ENABLED=0 go build -ldflags '-w -s' -v ./cmd/ec2-compliance-report/...

test: build
	./ec2-compliance-report
