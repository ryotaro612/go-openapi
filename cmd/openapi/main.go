package main

import (
	"fmt"
	"github.com/ryotaro612/go-openapi/internal/handler"
	"github.com/ryotaro612/go-openapi/internal/log"
	"github.com/swaggest/openapi-go/openapi3"
	"github.com/swaggest/rest/gorillamux"
)

func main() {
	l := log.NewLogger()

	router := handler.NewRouter()
	refl := openapi3.NewReflector()
	refl.SpecSchema().SetTitle("Sample API")
	refl.SpecSchema().SetVersion("v1.2.3")
	refl.SpecSchema().SetDescription("This is an example.")
	// Walk the router with OpenAPI collector.
	c := gorillamux.NewOpenAPICollector(refl)

	err := router.Walk(c.Walker)
	if err != nil {
		l.Error("failed to traverse the defined routing tree.", "error", err)
	}
	// Get the resulting schema.
	yml, err := refl.Spec.MarshalYAML()
	if err != nil {
		l.Error("failed to save the OpenAPI definition to a file in YAML format.", "error", err)
	}
	fmt.Println(string(yml))
}
