package admixer

import (
	"testing"
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/privacy"
	"github.com/eugene-fedorenko/prebid-server/privacy/ccpa"
	"github.com/eugene-fedorenko/prebid-server/privacy/gdpr"
	"github.com/stretchr/testify/assert"
)

func TestAdmixerSyncer(t *testing.T) {
	syncURL := "http://anyHost/anyPath"
	syncURLTemplate := template.Must(
		template.New("sync-template").Parse(syncURL),
	)

	syncer := NewAdmixerSyncer(syncURLTemplate)
	syncInfo, err := syncer.GetUsersyncInfo(privacy.Policies{
		GDPR: gdpr.Policy{
			Signal:  "A",
			Consent: "B",
		},
		CCPA: ccpa.Policy{
			Consent: "C",
		},
	})

	assert.NoError(t, err)
	assert.Equal(t, "http://anyHost/anyPath", syncInfo.URL)
	assert.Equal(t, "redirect", syncInfo.Type)
	assert.EqualValues(t, 511, syncer.GDPRVendorID())
	assert.Equal(t, false, syncInfo.SupportCORS)
}
