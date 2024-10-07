package handler

import (
	"encoding/json"
	"github.com/swaggest/openapi-go"
	"net/http"
)

func (h *userCreationHandler) SetupOpenAPIOperation(oc openapi.OperationContext) error {
	oc.SetTags("My Tag")
	oc.SetSummary("My Summary")
	oc.SetDescription("This endpoint aggregates request in structured way.")
	oc.AddReqStructure(createUserRequest{})
	oc.AddRespStructure(createUserResponse{})
	oc.AddRespStructure(nil, openapi.WithContentType("text/html"), openapi.WithHTTPStatus(http.StatusBadRequest))
	oc.AddRespStructure(nil, openapi.WithContentType("text/html"), openapi.WithHTTPStatus(http.StatusInternalServerError))
	return nil
}

// https://manual.iij.jp/iid/iidapi/19001059.html#id-%E3%83%A6%E3%83%BC%E3%82%B6API-%E3%83%A6%E3%83%BC%E3%82%B6%E3%81%AE%E4%BD%9C%E6%88%90
type createUserRequest struct {
	UserName string `json:"userName"`
}

type createUserResponse struct {
	DisplayName string `json:"displayName"`
}

// // https://pkg.go.dev/github.com/swaggest/rest@v0.2.67/gorillamux
// func temp() {
// 	decoderFactory := request.NewDecoderFactory()
// 	decoderFactory.ApplyDefaults = true
// 	decoderFactory.SetDecoderFunc(rest.ParamInPath, gorillamux.PathToURLValues)

// 	return &status
// }

func (h *userCreationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(createUserResponse{DisplayName: "jane.doe"})
}

func newUserCreationHandler() *userCreationHandler {
	return &userCreationHandler{}
}

type userCreationHandler struct {
}
