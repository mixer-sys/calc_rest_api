SWAG = swag
SWAG_DIR = ./api/openapi-spec/v1
GENERALINFO = ./cmd/app/main.go

.PHONY: all clean swag

all: swag

swag:
	$(SWAG) init --output $(SWAG_DIR) --generalInfo $(GENERALINFO)

clean:
	rm -rf $(SWAG_DIR)/*