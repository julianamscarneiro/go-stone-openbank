package openbank

import (
"fmt"
"net/http"

"github.com/stone-co/go-stone-openbank/types"
)

// PIXService handles communication with Stone Openbank API
type PIXService struct {
	client *Client
}

// GetOutboundPix is a service used to retrieve information details from a Pix.
func (s *PIXService) GetOutboundPix(id string) (*types.PIXOutBoundOutput, *Response, error) {

	path := fmt.Sprintf("/api/v1/pix/outbound_pix_payments/%s", id)

	req, err := s.client.NewAPIRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var pix types.PIXOutBoundOutput
	resp, err := s.client.Do(req, &pix)
	if err != nil {
		return nil, resp, err
	}

	return &pix, resp, err
}

// GetQRCodeData is a service used to retrieve information details from a Pix QRCode.
func (s *PIXService) GetQRCodeData(input types.GetQRCodeInput) (*types.QRCode, *Response, error) {
	const path = "/api/v1/pix/outbound_pix_payments/brcodes"

	req, err := s.client.NewAPIRequest(http.MethodGet, path, input)
	if err != nil {
		return nil, nil, err
	}

	var qrcode types.QRCode
	resp, err := s.client.Do(req, &qrcode)
	if err != nil {
		return nil, resp, err
	}

	return &qrcode, resp, err
}
