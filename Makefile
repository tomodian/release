.PHONY: build run test

BUILD := ./build

install:
	go install github.com/mitchellh/gox@latest
	curl -sSfL https://raw.githubusercontent.com/anchore/quill/main/install.sh | sh -s -- -b . v0.4.2

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

	@echo "Notarizing MacOS binary.."
	./quill sign-and-notarize ./build/release_darwin_amd64 || true
	./quill sign-and-notarize ./build/release_darwin_arm64 || true

	@echo "Bundling.."
	$(MAKE) bundle-nix
	$(MAKE) bundle-windows

run:
	go run main.go

test:
	go test -cover -count 1 -race -coverprofile=coverage.txt -covermode=atomic ./...

cover:
	go test -cover -coverprofile coverage.log ./... && go tool cover -html=coverage.log
