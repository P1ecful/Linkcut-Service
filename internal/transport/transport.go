package transport

import (
	"context"
	"encoding/json"
	"linkcut/internal/service"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type cutRequest struct {
	Link string `json:"LinkRequest"`
}

type cutResponse struct {
	Result string `json:"ResultResponse"`
}

func MakeLinkCutEndpoint(lcs service.LinkCutService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(cutRequest)
		return cutResponse{Result: lcs.Cut(req.Link)}, nil
	}
}

func DecodeCutRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request cutRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
