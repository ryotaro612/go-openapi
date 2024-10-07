package main

import (
	//	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ryotaro612/go-openapi/internal/handler"
	"github.com/swaggest/openapi-go/openapi3"
	"github.com/swaggest/rest/gorillamux"
	"net/http"
	"time"
)

type status struct {
	Status string `json:"status"`
}

func main() {
	router := mux.NewRouter()

	// router.HandleFunc("/v1/health", func(w http.ResponseWriter, r *http.Request) {
	// 	// an example API handler
	// 	w.Header().Set("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(map[string]bool{"ok": true})

	// }).Methods(http.MethodGet)

	router.Handle("/v1/health", handler.NewHealthHandler()).Methods(http.MethodGet)

	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      router,
	}
	refl := openapi3.NewReflector()
	refl.SpecSchema().SetTitle("Sample API")
	refl.SpecSchema().SetVersion("v1.2.3")
	refl.SpecSchema().SetDescription("This is an example.")

	// Walk the router with OpenAPI collector.
	c := gorillamux.NewOpenAPICollector(refl)

	_ = router.Walk(c.Walker)

	// Get the resulting schema.
	yml, _ := refl.Spec.MarshalYAML()
	fmt.Println(string(yml))
	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}
