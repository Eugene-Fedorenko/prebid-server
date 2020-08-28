package smartadserver

import (
	"testing"

	"github.com/eugene-fedorenko/prebid-server/adapters/adapterstest"
)

func TestJsonSamples(t *testing.T) {
	adapterstest.RunJSONBidderTest(t, "smartadservertest", NewSmartadserverBidder("https://ssb.smartadserver.com"))
}
