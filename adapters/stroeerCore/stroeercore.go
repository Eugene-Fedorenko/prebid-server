package stroeerCore

import (
	"encoding/json"
	"fmt"
	"github.com/mxmCherry/openrtb"
	"github.com/prebid/prebid-server/adapters"
	"github.com/prebid/prebid-server/config"
	"github.com/prebid/prebid-server/errortypes"
	"github.com/prebid/prebid-server/openrtb_ext"
	"net/http"
)

type StroeerCoreBidder struct {
	Url string `json:"url"`
}

type StroeerRootResponse struct {
	Bids []StroeerBidResponse `json:"bids"`
}

type StroeerBidResponse struct {
	Id     string  `json:"id"`
	BidId  string  `json:"bidId"`
	Cpm    float64 `json:"cpm"`
	Width  uint64  `json:"width"`
	Height uint64  `json:"height"`
	Ad     string  `json:"ad"`
	CrId   string  `json:"crid"`
}

func (a *StroeerCoreBidder) MakeBids(internalRequest *openrtb.BidRequest, externalRequest *adapters.RequestData, response *adapters.ResponseData) (*adapters.BidderResponse, []error) {
	if response.StatusCode != http.StatusOK {
		return nil, []error{&errortypes.BadServerResponse{
			Message: fmt.Sprintf("Unexpected http status code: %d.", response.StatusCode),
		}}
	}

	var errors []error
	stroeerResponse := StroeerRootResponse{}

	if err := json.Unmarshal(response.Body, &stroeerResponse); err != nil {
		errors = append(errors, err)
		return nil, errors
	}

	bidderResponse := adapters.NewBidderResponseWithBidsCapacity(len(stroeerResponse.Bids))
	bidderResponse.Currency = "EUR"

	for _, bid := range stroeerResponse.Bids {
		openRtbBid := openrtb.Bid{
			ID:    bid.Id,
			ImpID: bid.BidId,
			W:     bid.Width,
			H:     bid.Height,
			Price: bid.Cpm,
			AdM:   bid.Ad,
			CrID:  bid.CrId,
		}

		bidderResponse.Bids = append(bidderResponse.Bids, &adapters.TypedBid{
			Bid:     &openRtbBid,
			BidType: openrtb_ext.BidTypeBanner,
		})
	}

	return bidderResponse, errors
}

func (b *StroeerCoreBidder) MakeRequests(internalRequest *openrtb.BidRequest, reqInfo *adapters.ExtraRequestInfo) ([]*adapters.RequestData, []error) {
	errors := make([]error, 0, len(internalRequest.Imp))
	endpoint := ""

	for idx := range internalRequest.Imp {
		imp := &internalRequest.Imp[idx]
		var bidderExt adapters.ExtImpBidder
		if err := json.Unmarshal(imp.Ext, &bidderExt); err != nil {
			errors = append(errors, err)
			continue
		}

		var stroeerExt openrtb_ext.ExtImpStroeercore
		if err := json.Unmarshal(bidderExt.Bidder, &stroeerExt); err != nil {
			errors = append(errors, err)
			continue
		}

		imp.TagID = stroeerExt.Sid

		if endpoint == "" {
			endpoint = stroeerExt.Endpoint
		}
	}

	if len(endpoint) == 0 {
		endpoint = b.Url
	}

	if internalRequest.Device != nil {
		if internalRequest.Device.Geo != nil {
			internalRequest.Device.Geo.Type = openrtb.LocationType(1)
		}
	}

	reqJSON, err := json.Marshal(*internalRequest)
	if err != nil {
		errors = append(errors, err)
		return nil, errors
	}

	headers := http.Header{}
	headers.Add("Content-Type", "application/json;charset=utf-8")
	headers.Add("Accept", "application/json")

	return []*adapters.RequestData{{
		Method:  "POST",
		Uri:     endpoint,
		Body:    reqJSON,
		Headers: headers,
	}}, errors
}

// Builder builds a new instance of the StroeerCore adapter for the given bidder with the given config.
func Builder(bidderName openrtb_ext.BidderName, config config.Adapter) (adapters.Bidder, error) {
	bidder := &StroeerCoreBidder{
		Url: config.Endpoint,
	}
	return bidder, nil
}
