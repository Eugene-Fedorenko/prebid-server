package grid

import (
	"testing"
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/privacy"
	"github.com/eugene-fedorenko/prebid-server/privacy/gdpr"
	"github.com/stretchr/testify/assert"
)

func TestGridSyncer(t *testing.T) {
	syncURL := "//not_localhost/synclocalhost%2Fsetuid%3Fbidder%3Dgrid%26gdpr%3D{{.GDPR}}%26gdpr_consent%3D{{.GDPRConsent}}%26uid%3D%24UID"
	syncURLTemplate := template.Must(
		template.New("sync-template").Parse(syncURL),
	)

	syncer := NewGridSyncer(syncURLTemplate)
	syncInfo, err := syncer.GetUsersyncInfo(privacy.Policies{
		GDPR: gdpr.Policy{
			Signal: "0",
		},
	})

	assert.NoError(t, err)
	assert.Equal(t, "//not_localhost/synclocalhost%2Fsetuid%3Fbidder%3Dgrid%26gdpr%3D0%26gdpr_consent%3D%26uid%3D%24UID", syncInfo.URL)
	assert.Equal(t, "redirect", syncInfo.Type)
	assert.EqualValues(t, 686, syncer.GDPRVendorID())
	assert.Equal(t, false, syncInfo.SupportCORS)
}
