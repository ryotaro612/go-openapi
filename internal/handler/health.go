package handler

import (
	"encoding/json"
	"github.com/swaggest/openapi-go"
	"net/http"
)

func (h *healthHandler) SetupOpenAPIOperation(oc openapi.OperationContext) error {
	oc.SetTags("My Tag")
	oc.SetSummary("My Summary")
	oc.SetDescription("This endpoint aggregates request in structured way.")
	oc.AddRespStructure(healthResponse{})
	oc.AddRespStructure(nil, openapi.WithContentType("text/html"), openapi.WithHTTPStatus(http.StatusBadRequest))
	oc.AddRespStructure(nil, openapi.WithContentType("text/html"), openapi.WithHTTPStatus(http.StatusInternalServerError))
	return nil
}

type healthResponse struct {
	Status string `json:"status"`
}

// // https://pkg.go.dev/github.com/swaggest/rest@v0.2.67/gorillamux
// func temp() {
// 	decoderFactory := request.NewDecoderFactory()
// 	decoderFactory.ApplyDefaults = true
// 	decoderFactory.SetDecoderFunc(rest.ParamInPath, gorillamux.PathToURLValues)

// 	return &status
// }

func (h *healthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(healthResponse{Status: "ok"})
}

func newHealthHandler() *healthHandler {
	return &healthHandler{}
}

type healthHandler struct {
}
