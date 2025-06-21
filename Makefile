# Name of the binary
BINARY_NAME = studylight

# Output directory
BUILD_DIR = dist

# Your token to bake into the binary
TOKEN ?= ""

$(shell mkdir -p $(BUILD_DIR))

.PHONY: all clean

# Build all targets
all: linux-amd64 macos-arm64 rpi-arm windows

linux-amd64:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-X 'main.builtInToken=$(TOKEN)'" -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 main.go

macos-arm64:
	GOOS=darwin GOARCH=arm64 go build -ldflags="-X 'main.builtInToken=$(TOKEN)'" -o $(BUILD_DIR)/$(BINARY_NAME)-macos-arm64 main.go

rpi-arm:
	GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=0 go build -ldflags="-X 'main.builtInToken=$(TOKEN)'" -o $(BUILD_DIR)/$(BINARY_NAME)-rpi-arm main.go

windows:
	GOOS=windows GOARCH=amd64 go build -ldflags="-X 'main.builtInToken=$(TOKEN)'" -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe main.go

clean:
	rm -rf $(BUILD_DIR)
