.PHONY: build install clean run build-all

BINARY_NAME=gwa

# Detect operating system
ifeq ($(OS),Windows_NT)
	# Windows
	BINARY_EXT=.exe
	INSTALL_PATH=$(HOME)/go/bin
else
	# Unix-like (Mac/Linux)
	BINARY_EXT=
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Darwin)
		# macOS - try /usr/local/bin, fallback to ~/go/bin
		INSTALL_PATH=/usr/local/bin
	else
		# Linux - try /usr/local/bin, fallback to ~/go/bin  
		INSTALL_PATH=/usr/local/bin
	endif
endif

build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p bin
	go build -o bin/$(BINARY_NAME)$(BINARY_EXT) cmd/gwa/main.go
	@echo "✓ Built bin/$(BINARY_NAME)$(BINARY_EXT)"

install: build
	@echo "Installing $(BINARY_NAME) to $(INSTALL_PATH)..."
ifeq ($(OS),Windows_NT)
	@mkdir -p $(INSTALL_PATH)
	cp bin/$(BINARY_NAME)$(BINARY_EXT) $(INSTALL_PATH)/
else
	@if [ -w $(INSTALL_PATH) ]; then \
		cp bin/$(BINARY_NAME) $(INSTALL_PATH)/; \
		echo "✓ Installed to $(INSTALL_PATH)/$(BINARY_NAME)"; \
	else \
		echo "Need sudo for $(INSTALL_PATH)..."; \
		sudo cp bin/$(BINARY_NAME) $(INSTALL_PATH)/; \
		echo "✓ Installed to $(INSTALL_PATH)/$(BINARY_NAME)"; \
	fi
endif
	@echo "✓ Installation complete!"

clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/
	@echo "✓ Cleaned"

run:
	go run cmd/gwa/main.go $(ARGS)

build-all:
	@echo "Building for all platforms..."
	@mkdir -p bin
	GOOS=darwin GOARCH=amd64 go build -o bin/$(BINARY_NAME)-mac-amd64 cmd/gwa/main.go
	GOOS=darwin GOARCH=arm64 go build -o bin/$(BINARY_NAME)-mac-arm64 cmd/gwa/main.go
	GOOS=linux GOARCH=amd64 go build -o bin/$(BINARY_NAME)-linux cmd/gwa/main.go
	GOOS=windows GOARCH=amd64 go build -o bin/$(BINARY_NAME)-windows.exe cmd/gwa/main.go
	@echo "✓ Built all platform binaries:"
	@ls -lh bin/