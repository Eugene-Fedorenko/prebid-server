package emx_digital

import (
	"testing"

	"github.com/eugene-fedorenko/prebid-server/adapters/adapterstest"
)

func TestJsonSamples(t *testing.T) {
	emxDigitalAdapter := NewEmxDigitalBidder("https://hb.emxdgt.com")
	emxDigitalAdapter.testing = true
	adapterstest.RunJSONBidderTest(t, "emx_digitaltest", emxDigitalAdapter)
}
