package unruly

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/mxmCherry/openrtb"
	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/adapters/adapterstest"
	"github.com/eugene-fedorenko/prebid-server/config"
	"github.com/eugene-fedorenko/prebid-server/errortypes"
	"github.com/eugene-fedorenko/prebid-server/openrtb_ext"
	"github.com/stretchr/testify/assert"
)

func TestJsonSamples(t *testing.T) {
	bidder, buildErr := Builder(openrtb_ext.BidderUnruly, config.Adapter{
		Endpoint: "http://targeting.unrulymedia.com/openrtb/2.2"})

	if buildErr != nil {
		t.Fatalf("Builder returned unexpected error %v", buildErr)
	}

	adapterstest.RunJSONBidderTest(t, "unrulytest", bidder)
}

func TestReturnsNewUnrulyBidderWithParams(t *testing.T) {
	bidder, buildErr := Builder(openrtb_ext.BidderUnruly, config.Adapter{
		Endpoint: "http://mockEndpoint.com"})

	if buildErr != nil {
		t.Fatalf("Builder returned unexpected error %v", buildErr)
	}

	bidderUnruly := bidder.(*UnrulyAdapter)

	assert.Equal(t, "http://mockEndpoint.com", bidderUnruly.URI)
}

func TestBuildRequest(t *testing.T) {
	request := openrtb.BidRequest{}
	expectedJson, _ := json.Marshal(request)
	mockHeaders := http.Header{}
	mockHeaders.Add("Content-Type", "application/json;charset=utf-8")
	mockHeaders.Add("Accept", "application/json")
	mockHeaders.Add("X-Unruly-Origin", "Prebid-Server")
	data := adapters.RequestData{
		Method:  "POST",
		Uri:     "http://mockEndpoint.com",
		Body:    expectedJson,
		Headers: mockHeaders,
	}

	adapter := UnrulyAdapter{URI: "http://mockEndpoint.com"}

	actual, _ := adapter.BuildRequest(&request)
	expected := data
	if !reflect.DeepEqual(expected, *actual) {
		t.Errorf("actual = %v expected = %v", actual, expected)
	}

}

func TestReplaceImp(t *testing.T) {
	imp1 := openrtb.Imp{ID: "imp1"}
	imp2 := openrtb.Imp{ID: "imp2"}
	imp3 := openrtb.Imp{ID: "imp3"}
	newImp := openrtb.Imp{ID: "imp4"}
	request := openrtb.BidRequest{Imp: []openrtb.Imp{imp1, imp2, imp3}}
	adapter := UnrulyAdapter{URI: "http://mockEndpoint.com"}
	newRequest := adapter.ReplaceImp(newImp, &request)

	if len(newRequest.Imp) != 1 {
		t.Errorf("Size of Imp Array should be 1")
	}
	if !reflect.DeepEqual(request, openrtb.BidRequest{Imp: []openrtb.Imp{imp1, imp2, imp3}}) {
		t.Errorf("actual = %v expected = %v", request, openrtb.BidRequest{Imp: []openrtb.Imp{imp1, imp2, imp3}})
	}
	if !reflect.DeepEqual(newImp, newRequest.Imp[0]) {
		t.Errorf("actual = %v expected = %v", newRequest.Imp[0], newImp)
	}
}

func TestConvertBidderNameInExt(t *testing.T) {
	imp := openrtb.Imp{Ext: json.RawMessage(`{"bidder": {"uuid": "1234", "siteid": "aSiteID"}}`)}

	actualImp, _, err := convertBidderNameInExt(&imp)

	if err != nil {
		t.Errorf("actual = %v expected = %v", err, nil)
	}

	var unrulyExt ImpExtUnruly
	err = json.Unmarshal(actualImp.Ext, &unrulyExt)

	if err != nil {
		t.Errorf("actual = %v expected = %v", err, nil)
	}

	if unrulyExt.Unruly.UUID != "1234" {
		t.Errorf("actual = %v expected = %v", unrulyExt.Unruly.UUID, "1234")
	}

	if unrulyExt.Unruly.SiteID != "aSiteID" {
		t.Errorf("actual = %v expected = %v", unrulyExt.Unruly.SiteID, "aSiteID")
	}
}

func TestMakeRequests(t *testing.T) {
	adapter := UnrulyAdapter{URI: "http://mockEndpoint.com"}

	imp1 := openrtb.Imp{ID: "imp1", Ext: json.RawMessage(`{"bidder": {"uuid": "uuid1", "siteid": "siteID1"}}`)}
	imp2 := openrtb.Imp{ID: "imp2", Ext: json.RawMessage(`{"bidder": {"uuid": "uuid2", "siteid": "siteID2"}}`)}
	imp3 := openrtb.Imp{ID: "imp3", Ext: json.RawMessage(`{"bidder": {"uuid": "uuid3", "siteid": "siteID3"}}`)}

	expectImp1 := openrtb.Imp{ID: "imp1", Ext: json.RawMessage(`{"unruly": {"uuid": "uuid1", "siteid": "siteID1"}}`)}
	expectImp2 := openrtb.Imp{ID: "imp2", Ext: json.RawMessage(`{"unruly": {"uuid": "uuid2", "siteid": "siteID2"}}`)}
	expectImp3 := openrtb.Imp{ID: "imp3", Ext: json.RawMessage(`{"unruly": {"uuid": "uuid3", "siteid": "siteID3"}}`)}

	expectImps := []openrtb.Imp{expectImp1, expectImp2, expectImp3}

	inputRequest := openrtb.BidRequest{Imp: []openrtb.Imp{imp1, imp2, imp3}}
	actualAdapterRequests, _ := adapter.MakeRequests(&inputRequest, &adapters.ExtraRequestInfo{})
	mockHeaders := http.Header{}
	mockHeaders.Add("Content-Type", "application/json;charset=utf-8")
	mockHeaders.Add("Accept", "application/json")
	mockHeaders.Add("X-Unruly-Origin", "Prebid-Server")
	if len(actualAdapterRequests) != 3 {
		t.Errorf("should have 3 imps")
	}
	for n, imp := range expectImps {
		request := openrtb.BidRequest{Imp: []openrtb.Imp{imp}}
		expectedJson, _ := json.Marshal(request)
		data := adapters.RequestData{
			Method:  "POST",
			Uri:     "http://mockEndpoint.com",
			Body:    expectedJson,
			Headers: mockHeaders,
		}
		if !reflect.DeepEqual(data, *actualAdapterRequests[n]) {
			t.Errorf("actual = %v expected = %v", *actualAdapterRequests[0], data)
		}
	}
}

func TestGetMediaTypeForImpIsVideo(t *testing.T) {
	testID := string("4321")
	testBidMediaType := openrtb_ext.BidTypeVideo
	imp := openrtb.Imp{
		ID:    testID,
		Video: &openrtb.Video{},
	}
	imps := []openrtb.Imp{imp}
	actual, _ := getMediaTypeForImpWithId(testID, imps)

	if actual != "video" {
		t.Errorf("actual = %v expected = %v", actual, testBidMediaType)
	}
}

func TestGetMediaTypeForImpWithNoIDPresent(t *testing.T) {
	imp := openrtb.Imp{
		ID:    "4321",
		Video: &openrtb.Video{},
	}
	imps := []openrtb.Imp{imp}
	_, err := getMediaTypeForImpWithId("1234", imps)
	expected := &errortypes.BadInput{
		Message: fmt.Sprintf("Failed to find impression \"%s\" ", "1234"),
	}
	if !reflect.DeepEqual(expected, err) {
		t.Errorf("actual = %v expected = %v", expected, err)
	}
}

func TestConvertToAdapterBidResponseHasCorrectNumberOfBids(t *testing.T) {
	imp := openrtb.Imp{
		ID:    "1234",
		Video: &openrtb.Video{},
	}
	imp2 := openrtb.Imp{
		ID:    "1235",
		Video: &openrtb.Video{},
	}

	mockResponse := adapters.ResponseData{StatusCode: 200,
		Body: json.RawMessage(`{"seatbid":[{"bid":[{"impid":"1234"}]},{"bid":[{"impid":"1235"}]}]}`)}
	internalRequest := openrtb.BidRequest{Imp: []openrtb.Imp{imp, imp2}}
	mockBidResponse := adapters.NewBidderResponseWithBidsCapacity(5)

	typedBid := &adapters.TypedBid{
		Bid:     &openrtb.Bid{ImpID: "1234"},
		BidType: "Video",
	}
	typedBid2 := &adapters.TypedBid{
		Bid:     &openrtb.Bid{ImpID: "1235"},
		BidType: "Video",
	}

	mockBidResponse.Bids = append(mockBidResponse.Bids, typedBid)
	mockBidResponse.Bids = append(mockBidResponse.Bids, typedBid2)

	actual, _ := convertToAdapterBidResponse(&mockResponse, &internalRequest)
	if !reflect.DeepEqual(*actual.Bids[0].Bid, *mockBidResponse.Bids[0].Bid) {
		t.Errorf("actual = %v expected = %v", *actual.Bids[0].Bid, *mockBidResponse.Bids[0].Bid)
	}
	if !reflect.DeepEqual(*actual.Bids[1].Bid, *mockBidResponse.Bids[1].Bid) {
		t.Errorf("actual = %v expected = %v", *actual.Bids[1].Bid, *mockBidResponse.Bids[1].Bid)
	}
}
