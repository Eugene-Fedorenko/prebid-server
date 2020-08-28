package valueimpression

import (
	"encoding/json"
	"testing"

	"github.com/eugene-fedorenko/prebid-server/openrtb_ext"
)

// This file actually intends to test static/bidder-params/valueimpression.json
// These also validate the format of the external API: request.imp[i].ext.valueimpression
// TestValidParams makes sure that the ValueImpression schema accepts all imp.ext fields which we intend to support.

func TestValidParams(t *testing.T) {
	validator, err := openrtb_ext.NewBidderParamsValidator("../../static/bidder-params")
	if err != nil {
		t.Fatalf("Failed to fetch the json-schemas. %v", err)
	}

	for _, validParam := range validParams {
		if err := validator.Validate(openrtb_ext.BidderValueImpression, json.RawMessage(validParam)); err != nil {
			t.Errorf("Schema rejected ValueImpression params: %s", validParam)
		}
	}
}

// TestInvalidParams makes sure that the ValueImpression schema rejects all the imp.ext fields we don't support.
func TestInvalidParams(t *testing.T) {
	validator, err := openrtb_ext.NewBidderParamsValidator("../../static/bidder-params")
	if err != nil {
		t.Fatalf("Failed to fetch the json-schemas. %v", err)
	}

	for _, invalidParam := range invalidParams {
		if err := validator.Validate(openrtb_ext.BidderValueImpression, json.RawMessage(invalidParam)); err == nil {
			t.Errorf("Schema allowed unexpected params: %s", invalidParam)
		}
	}
}

var validParams = []string{
	`{"siteId": "123"}`,
}

var invalidParams = []string{
	`{}`,
	`null`,
	`true`,
	`154`,
	`{"siteId": 123}`, // siteId should be string
	`{"invalid_param": "123"}`,
}
