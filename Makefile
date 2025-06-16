SWAG = swag
SWAG_DIR = ./api/docs
GENERALINFO = ./cmd/myapp/main.go

.PHONY: all clean swag

all: swag

swag:
	$(SWAG) init --output $(SWAG_DIR) --generalInfo $(GENERALINFO)

clean:
	rm -rf $(SWAG_DIR)/*