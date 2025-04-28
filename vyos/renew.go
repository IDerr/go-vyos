package vyos

import (
	"context"
	"strings"
)

type RenewService service

type RenewResponse struct {
	*RawResponse
}

// Do sends a request to the VyOS API and returns the response.
func (s *RenewService) Do(ctx context.Context, path string) (*RenewResponse, *Response, error) {

	u := "/renew"
	p := strings.Split(path, " ")

	if p == nil {
		return nil, nil, ErrEmptyPath
	}

	// Create a new request.
	request := Request{
		OPMode: OPMode("renew"),
		Path:   p,
	}

	// Create the HTTP request.
	req, err := s.client.NewRequest(u, &request)
	if err != nil {
		return nil, nil, err
	}

	// Create the Response struct & send the request.
	v := new(RenewResponse)
	resp, err := s.client.Do(ctx, req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}
