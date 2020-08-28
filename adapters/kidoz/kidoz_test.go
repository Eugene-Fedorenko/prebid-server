package kidoz

import (
	"math"
	"net/http"
	"testing"

	"github.com/mxmCherry/openrtb"
	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/adapters/adapterstest"
	"github.com/eugene-fedorenko/prebid-server/openrtb_ext"
	"github.com/stretchr/testify/assert"
)

func TestJsonSamples(t *testing.T) {
	adapterstest.RunJSONBidderTest(t, "kidoztest", NewKidozBidder("http://example.com/prebid"))
}

func makeBidRequest() *openrtb.BidRequest {
	request := &openrtb.BidRequest{
		ID: "test-req-id-0",
		Imp: []openrtb.Imp{
			{
				ID: "test-imp-id-0",
				Banner: &openrtb.Banner{
					Format: []openrtb.Format{
						{
							W: 320,
							H: 50,
						},
					},
				},
				Ext: []byte(`{"bidder":{"access_token":"token-0","publisher_id":"pub-0"}}`),
			},
		},
	}
	return request
}

func TestMakeRequests(t *testing.T) {
	kidoz := NewKidozBidder("http://example.com/prebid")

	t.Run("Handles Request marshal failure", func(t *testing.T) {
		request := makeBidRequest()
		request.Imp[0].BidFloor = math.Inf(1) // cant be marshalled
		extra := &adapters.ExtraRequestInfo{}
		reqs, errs := kidoz.MakeRequests(request, extra)
		// cant assert message its different on different versions of go
		assert.Equal(t, 1, len(errs))
		assert.Contains(t, errs[0].Error(), "json")
		assert.Equal(t, 0, len(reqs))
	})
}

func TestMakeBids(t *testing.T) {
	kidoz := NewKidozBidder("http://example.com/prebid")

	t.Run("Handles response marshal failure", func(t *testing.T) {
		request := makeBidRequest()
		requestData := &adapters.RequestData{}
		responseData := &adapters.ResponseData{
			StatusCode: http.StatusOK,
		}

		resp, errs := kidoz.MakeBids(request, requestData, responseData)
		// cant assert message its different on different versions of go
		assert.Equal(t, 1, len(errs))
		assert.Contains(t, errs[0].Error(), "JSON")
		assert.Nil(t, resp)
	})
}

func TestGetMediaTypeForImp(t *testing.T) {
	imps := []openrtb.Imp{
		{
			ID:     "1",
			Banner: &openrtb.Banner{},
		},
		{
			ID:    "2",
			Video: &openrtb.Video{},
		},
		{
			ID:     "3",
			Native: &openrtb.Native{},
		},
		{
			ID:    "4",
			Audio: &openrtb.Audio{},
		},
	}

	t.Run("Bid not found is type empty string", func(t *testing.T) {
		actual := GetMediaTypeForImp("ARGLE_BARGLE", imps)
		assert.Equal(t, UndefinedMediaType, actual)
	})
	t.Run("Can find banner type", func(t *testing.T) {
		actual := GetMediaTypeForImp("1", imps)
		assert.Equal(t, openrtb_ext.BidTypeBanner, actual)
	})
	t.Run("Can find video type", func(t *testing.T) {
		actual := GetMediaTypeForImp("2", imps)
		assert.Equal(t, openrtb_ext.BidTypeVideo, actual)
	})
	t.Run("Can find native type", func(t *testing.T) {
		actual := GetMediaTypeForImp("3", imps)
		assert.Equal(t, openrtb_ext.BidTypeNative, actual)
	})
	t.Run("Can find audio type", func(t *testing.T) {
		actual := GetMediaTypeForImp("4", imps)
		assert.Equal(t, openrtb_ext.BidTypeAudio, actual)
	})
}
