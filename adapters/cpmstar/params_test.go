package cpmstar

import (
	"encoding/json"
	"testing"

	"github.com/eugene-fedorenko/prebid-server/openrtb_ext"
)

// This file actually intends to test static/bidder-params/cpmstar.json
// These also validate the format of the external API: request.imp[i].ext.cpmstar
// TestValidParams makes sure that the Cpmstar schema accepts all imp.ext fields which we intend to support.

func TestValidParams(t *testing.T) {
	validator, err := openrtb_ext.NewBidderParamsValidator("../../static/bidder-params")
	if err != nil {
		t.Fatalf("Failed to fetch the json-schemas. %v", err)
	}

	for _, validParam := range validParams {
		if err := validator.Validate(openrtb_ext.BidderCpmstar, json.RawMessage(validParam)); err != nil {
			t.Errorf("Schema rejected Cpmstar params: %s", validParam)
		}
	}
}

// TestInvalidParams makes sure that the Cpmstar schema rejects all the imp.ext fields we don't support.
func TestInvalidParams(t *testing.T) {
	validator, err := openrtb_ext.NewBidderParamsValidator("../../static/bidder-params")
	if err != nil {
		t.Fatalf("Failed to fetch the json-schemas. %v", err)
	}

	for _, invalidParam := range invalidParams {
		if err := validator.Validate(openrtb_ext.BidderCpmstar, json.RawMessage(invalidParam)); err == nil {
			t.Errorf("Schema allowed unexpected params: %s", invalidParam)
		}
	}
}

var validParams = []string{
	`{"placementId": 154}`,
	`{"placementId": 154, "subpoolId": 123}`,
}

var invalidParams = []string{
	`{}`,
	`null`,
	`true`,
	`154`,
	`{"placementId": "154"}`, // placementId should be numeric
	`{"placementId": 154, "subpoolId": "123"}`, // placementId and subpoolId should both be numeric
	`{"invalid_param": 123}`,
}
