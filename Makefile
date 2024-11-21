BIN_NAME := actuator
OUT_PATH := bin
PROGRAM := ./cmd/
Linux64Suffix := linux-64
Win64Suffix := win-64

all: $(Linux64Suffix) $(Win64Suffix)

$(Linux64Suffix): clean-$(Linux64Suffix)
	@GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -v -o $(OUT_PATH)/$(BIN_NAME) $(PROGRAM)

$(Win64Suffix): clean-$(Win64Suffix)
	@GOOS=windows GOARCH=amd64 go build -ldflags="-w -s" -v -o $(OUT_PATH)/$(BIN_NAME)-$(Win64Suffix).exe $(PROGRAM)

.PHONY: clean clean-$(Linux64Suffix) clean-$(Win64Suffix)

clean: clean-$(Linux64Suffix) clean-$(Win64Suffix)

clean-$(Linux64Suffix):
	@rm -rf $(OUT_PATH)/$(BIN_NAME)-$(Linux64Suffix)

clean-$(Win64Suffix):
	@rm -rf $(OUT_PATH)/$(BIN_NAME)-$(Win64Suffix).exe

large-package: $(Linux64Suffix)
	@upx --best --lzma $(OUT_PATH)/$(BIN_NAME)