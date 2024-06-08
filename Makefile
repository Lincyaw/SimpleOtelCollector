# Makefile

# Variables for paths
PROTO_DIR := opentelemetry-proto
GEN_DIR := opentelemetry-proto/gen/go/go.opentelemetry.io/proto
DEST_DIR := ./
NEW_PATH := github.com/Lincyaw/SimpleOtelCollector/proto/otlp

# Define the target for generating Go code
gen-go:
	cd $(PROTO_DIR) && \
	make gen-go

# Define the target for copying the generated code to the destination directory
copy-go:
	cp -r $(GEN_DIR) $(DEST_DIR)

replace-paths:
	find $(DEST_DIR)/proto/otlp -type f -name "*.go" -exec sed -i 's|go.opentelemetry.io/proto/otlp|$(NEW_PATH)|g' {} +

# Define the default target that runs both steps
all: gen-go copy-go replace-paths
