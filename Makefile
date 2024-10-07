
PROJECT_ROOT := $(dir $(lastword $(MAKEFILE_LIST)))

## Display this message
help:
	@go run github.com/Songmu/make2help/cmd/make2help $(MAKEFILE_LIST)

OPENAPI_HTML = dist/openapi.html
## Generate the Open API specification
doc: $(OPENAPI_HTML)

## Remove the intermediate files
clean:
	rm -rf $(PROJECT_ROOT)dist

OPENAPI_YML = dist/openapi.yml

dist/openapi.html: $(OPENAPI_YML)
	#npx @redocly/cli lint $(OPENAPI_YML)
	npx @redocly/cli build-docs $(OPENAPI_YML) --output $(OPENAPI_HTML)

$(OPENAPI_YML):
	mkdir -p dist
	go run cmd/openapi/main.go > $(OPENAPI_YML)


.PHONY: help clean
