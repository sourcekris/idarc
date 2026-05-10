ifeq ($(OS),Windows_NT)
    # Windows cmd.exe syntax
    build_cmd = set GOOS=$(1)&& set GOARCH=$(2)&& go build -o $(BIN_DIR)/$(BINARY_NAME)-$(1)-$(2)$(3) $(MAIN_PATH)
    zip_cmd = powershell -NoProfile -Command "Compress-Archive -Path $(BIN_DIR)\$(BINARY_NAME)-$(1)-$(2)$(3) -DestinationPath $(BIN_DIR)\$(BINARY_NAME)-$(1)-$(2).zip -Force"
    RM_DIR = if exist $(BIN_DIR) rmdir /s /q $(BIN_DIR)
    EXT = .exe
else
    # POSIX shell syntax
    build_cmd = GOOS=$(1) GOARCH=$(2) go build -o $(BIN_DIR)/$(BINARY_NAME)-$(1)-$(2)$(3) $(MAIN_PATH)
    zip_cmd = zip -j $(BIN_DIR)/$(BINARY_NAME)-$(1)-$(2).zip $(BIN_DIR)/$(BINARY_NAME)-$(1)-$(2)$(3)
    RM_DIR = rm -rf $(BIN_DIR)
    EXT =
endif

BINARY_NAME=idarc
MAIN_PATH=.
BIN_DIR=bin

.PHONY: all build clean cross-compile linux windows mac release

all: build

build:
	go build -o $(BIN_DIR)/$(BINARY_NAME)$(EXT) $(MAIN_PATH)

clean:
	$(RM_DIR)

cross-compile: linux windows mac

release: cross-compile
	$(call zip_cmd,linux,amd64,)
	$(call zip_cmd,windows,amd64,.exe)
	$(call zip_cmd,windows,386,.exe)
	$(call zip_cmd,darwin,amd64,)
	$(call zip_cmd,darwin,arm64,)

linux:
	$(call build_cmd,linux,amd64,)

windows:
	$(call build_cmd,windows,amd64,.exe)
	$(call build_cmd,windows,386,.exe)

mac:
	$(call build_cmd,darwin,amd64,)
	$(call build_cmd,darwin,arm64,)
