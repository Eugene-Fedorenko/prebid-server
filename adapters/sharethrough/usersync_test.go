package sharethrough

import (
	"testing"
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/privacy"
	"github.com/eugene-fedorenko/prebid-server/privacy/gdpr"
	"github.com/stretchr/testify/assert"
)

func TestSharethroughSyncer(t *testing.T) {
	syncURL := "https://match.sharethrough.com?gdpr={{.GDPR}}&gdpr_consent={{.GDPRConsent}}"
	syncURLTemplate := template.Must(
		template.New("sync-template").Parse(syncURL),
	)

	syncer := NewSharethroughSyncer(syncURLTemplate)
	syncInfo, err := syncer.GetUsersyncInfo(privacy.Policies{
		GDPR: gdpr.Policy{
			Signal: "0",
		},
	})

	assert.NoError(t, err)
	assert.Equal(t, "https://match.sharethrough.com?gdpr=0&gdpr_consent=", syncInfo.URL)
	assert.Equal(t, "redirect", syncInfo.Type)
	assert.EqualValues(t, 80, syncer.GDPRVendorID())
	assert.Equal(t, false, syncInfo.SupportCORS)
	assert.Equal(t, "sharethrough", syncer.FamilyName())
}
