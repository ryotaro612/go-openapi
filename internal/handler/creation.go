package handler

import (
	"encoding/json"
	"fmt"
	"github.com/swaggest/openapi-go"
	"github.com/swaggest/rest"
	"github.com/swaggest/rest/gorillamux"
	"github.com/swaggest/rest/jsonschema"
	"github.com/swaggest/rest/nethttp"
	o "github.com/swaggest/rest/openapi"
	"github.com/swaggest/rest/request"
	"net/http"
)

func (h *userCreationHandler) SetupOpenAPIOperation(oc openapi.OperationContext) error {
	oc.SetTags("My Tag")
	oc.SetSummary("My Summary")
	oc.SetDescription("This endpoint aggregates request in structured way.")
	oc.AddReqStructure(createUserRequest{})
	oc.AddRespStructure(createUserResponse{}, openapi.WithHTTPStatus(http.StatusCreated))
	oc.AddRespStructure(nil, openapi.WithContentType("text/html"), openapi.WithHTTPStatus(http.StatusBadRequest))
	oc.AddRespStructure(nil, openapi.WithContentType("text/html"), openapi.WithHTTPStatus(http.StatusInternalServerError))
	return nil
}

// https://manual.iij.jp/iid/iidapi/19001059.html#id-%E3%83%A6%E3%83%BC%E3%82%B6API-%E3%83%A6%E3%83%BC%E3%82%B6%E3%81%AE%E4%BD%9C%E6%88%90
type createUserRequest struct {
	UserName string `json:"userName" required:"true"`
}

type createUserResponse struct {
	DisplayName string `json:"displayName"`
}

func (h *userCreationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var in createUserRequest
	fmt.Println("test")
	if err := h.dec.Decode(r, &in, h.valid); err != nil {
		fmt.Println("fail")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("user: %v\n", in)
	fmt.Println("succ")
	json.NewEncoder(w).Encode(createUserResponse{DisplayName: "jane.doe"})
}

func newUserCreationHandler() *userCreationHandler {

	decoderFactory := request.NewDecoderFactory()
	decoderFactory.ApplyDefaults = true
	decoderFactory.SetDecoderFunc(rest.ParamInPath, gorillamux.PathToURLValues)
	validator := jsonschema.NewFactory(&o.Collector{}, &o.Collector{}).MakeRequestValidator(http.MethodPost, &createUserRequest{}, nil)
	return &userCreationHandler{
		dec:   decoderFactory.MakeDecoder(http.MethodPost, createUserRequest{}, nil),
		valid: validator,
	}

}

type userCreationHandler struct {
	dec   nethttp.RequestDecoder
	valid rest.Validator
}
