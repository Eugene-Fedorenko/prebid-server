package mobilefuse

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/mxmCherry/openrtb"
	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/config"
	"github.com/eugene-fedorenko/prebid-server/errortypes"
	"github.com/eugene-fedorenko/prebid-server/macros"
	"github.com/eugene-fedorenko/prebid-server/openrtb_ext"
)

type MobileFuseAdapter struct {
	EndpointTemplate template.Template
}

// Builder builds a new instance of the MobileFuse adapter for the given bidder with the given config.
func Builder(bidderName openrtb_ext.BidderName, config config.Adapter) (adapters.Bidder, error) {
	template, err := template.New("endpointTemplate").Parse(config.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("unable to parse endpoint url template: %v", err)
	}

	bidder := &MobileFuseAdapter{
		EndpointTemplate: *template,
	}
	return bidder, nil
}

func (adapter *MobileFuseAdapter) MakeRequests(request *openrtb.BidRequest, reqInfo *adapters.ExtraRequestInfo) ([]*adapters.RequestData, []error) {
	var adapterRequests []*adapters.RequestData

	adapterRequest, errs := adapter.makeRequest(request)

	if errs == nil {
		adapterRequests = append(adapterRequests, adapterRequest)
	}

	return adapterRequests, errs
}

func (adapter *MobileFuseAdapter) MakeBids(incomingRequest *openrtb.BidRequest, outgoingRequest *adapters.RequestData, response *adapters.ResponseData) (*adapters.BidderResponse, []error) {
	if response.StatusCode == http.StatusNoContent {
		return nil, nil
	}

	if response.StatusCode == http.StatusBadRequest {
		return nil, []error{&errortypes.BadInput{
			Message: fmt.Sprintf("Unexpected status code: %d.", response.StatusCode),
		}}
	}

	if response.StatusCode != http.StatusOK {
		return nil, []error{&errortypes.BadServerResponse{
			Message: fmt.Sprintf("Unexpected status code: %d.", response.StatusCode),
		}}
	}

	var incomingBidResponse openrtb.BidResponse

	if err := json.Unmarshal(response.Body, &incomingBidResponse); err != nil {
		return nil, []error{err}
	}

	outgoingBidResponse := adapters.NewBidderResponseWithBidsCapacity(1)

	for _, seatbid := range incomingBidResponse.SeatBid {
		for i := range seatbid.Bid {
			outgoingBidResponse.Bids = append(outgoingBidResponse.Bids, &adapters.TypedBid{
				Bid:     &seatbid.Bid[i],
				BidType: adapter.getBidType(seatbid.Bid[i].ImpID, incomingRequest.Imp),
			})
		}
	}

	return outgoingBidResponse, nil
}

func (adapter *MobileFuseAdapter) makeRequest(bidRequest *openrtb.BidRequest) (*adapters.RequestData, []error) {
	var errs []error

	mobileFuseExtension, errs := adapter.getFirstMobileFuseExtension(bidRequest)

	if errs != nil {
		return nil, errs
	}

	endpoint, err := adapter.getEndpoint(mobileFuseExtension)

	if err != nil {
		return nil, append(errs, err)
	}

	validImps := adapter.getValidImps(bidRequest, mobileFuseExtension)

	if len(validImps) == 0 {
		err := fmt.Errorf("No valid imps")
		errs = append(errs, err)
		return nil, errs
	}

	mobileFuseBidRequest := *bidRequest
	mobileFuseBidRequest.Imp = validImps
	body, err := json.Marshal(mobileFuseBidRequest)

	if err != nil {
		return nil, append(errs, err)
	}

	headers := http.Header{}
	headers.Add("Content-Type", "application/json;charset=utf-8")
	headers.Add("Accept", "application/json")

	return &adapters.RequestData{
		Method:  "POST",
		Uri:     endpoint,
		Body:    body,
		Headers: headers,
	}, errs
}

func (adapter *MobileFuseAdapter) getFirstMobileFuseExtension(request *openrtb.BidRequest) (*openrtb_ext.ExtImpMobileFuse, []error) {
	var mobileFuseImpExtension openrtb_ext.ExtImpMobileFuse
	var errs []error

	for _, imp := range request.Imp {
		var bidder_imp_extension adapters.ExtImpBidder

		err := json.Unmarshal(imp.Ext, &bidder_imp_extension)

		if err != nil {
			errs = append(errs, err)
			continue
		}

		err = json.Unmarshal(bidder_imp_extension.Bidder, &mobileFuseImpExtension)

		if err != nil {
			errs = append(errs, err)
			continue
		}

		break
	}

	return &mobileFuseImpExtension, errs
}

func (adapter *MobileFuseAdapter) getEndpoint(ext *openrtb_ext.ExtImpMobileFuse) (string, error) {
	publisher_id := strconv.Itoa(ext.PublisherId)

	url, errs := macros.ResolveMacros(adapter.EndpointTemplate, macros.EndpointTemplateParams{PublisherID: publisher_id})

	if errs != nil {
		return "", errs
	}

	if ext.TagidSrc == "ext" {
		url += "&tagid_src=ext"
	}

	return url, nil
}

func (adapter *MobileFuseAdapter) getValidImps(bidRequest *openrtb.BidRequest, ext *openrtb_ext.ExtImpMobileFuse) []openrtb.Imp {
	var validImps []openrtb.Imp

	for _, imp := range bidRequest.Imp {
		if imp.Banner != nil || imp.Video != nil {
			if imp.Banner != nil && imp.Video != nil {
				imp.Video = nil
			}

			imp.TagID = strconv.Itoa(ext.PlacementId)
			imp.Ext = nil
			validImps = append(validImps, imp)

			break
		}
	}

	return validImps
}

func (adapter *MobileFuseAdapter) getBidType(imp_id string, imps []openrtb.Imp) openrtb_ext.BidType {
	if imps[0].Video != nil {
		return openrtb_ext.BidTypeVideo
	}

	return openrtb_ext.BidTypeBanner
}
