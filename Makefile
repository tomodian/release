.PHONY: build run test

BUILD := ./build

install:
	go get github.com/mitchellh/gox

clean:
	mkdir -p $(BUILD)
	rm -Rf $(BUILD)/*

bundle-nix:
	cd $(BUILD) && find . -type f ! -name '*.exe' | xargs -I % sh -c "mv % release && zip %.zip release && rm -f release"

bundle-windows:
	cd $(BUILD) && find . -type f -name '*.exe' | xargs -I % sh -c "mv % release.exe && zip %.zip release.exe && rm -f release.exe"

build: clean
	@echo "Building.."
	gox -output="$(BUILD)/{{.Dir}}_{{.OS}}_{{.Arch}}" \
		-osarch="darwin/amd64" \
		-osarch="darwin/arm64" \
		-osarch="linux/arm" \
		-osarch="linux/amd64" \
		-osarch="windows/amd64"
	@echo "Bundling.."
	$(MAKE) bundle-nix
	$(MAKE) bundle-windows

run:
	go run main.go

test:
	go test -cover -count 1 -race -coverprofile=coverage.txt -covermode=atomic ./...

cover:
	go test -cover -coverprofile coverage.log ./... && go tool cover -html=coverage.log
