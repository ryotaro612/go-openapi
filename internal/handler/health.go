package handler

import (
	"encoding/json"
	"github.com/swaggest/openapi-go"
	// "github.com/swaggest/openapi-go/openapi3"
	// "github.com/swaggest/rest"
	// "github.com/swaggest/rest/gorillamux"
	// "github.com/swaggest/rest/request"
	"net/http"
)

type response struct {
	Status string `json:"status"`
}

// // https://pkg.go.dev/github.com/swaggest/rest@v0.2.67/gorillamux
// func temp() {
// 	decoderFactory := request.NewDecoderFactory()
// 	decoderFactory.ApplyDefaults = true
// 	decoderFactory.SetDecoderFunc(rest.ParamInPath, gorillamux.PathToURLValues)

// 	return &status
// }

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(response{Status: "ok"})
}

func (h *Handler) SetupOpenAPIOperation(oc openapi.OperationContext) error {

	oc.SetTags("My Tag")
	oc.SetSummary("My Summary")
	oc.SetDescription("This endpoint aggregates request in structured way.")
	//oc.AddReqStructure(myRequest{})
	oc.AddRespStructure(response{})
	oc.AddRespStructure(nil, openapi.WithContentType("text/html"), openapi.WithHTTPStatus(http.StatusBadRequest))
	oc.AddRespStructure(nil, openapi.WithContentType("text/html"), openapi.WithHTTPStatus(http.StatusInternalServerError))
	return nil
}

type Handler struct {
}

func NewHealthHandler() *Handler {
	return &Handler{}
}
