package stroeerCore

import (
	"github.com/prebid/prebid-server/privacy"
	"github.com/prebid/prebid-server/privacy/gdpr"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

func TestStroeerCoreSyncer(t *testing.T) {
	temp := template.Must(template.New("sync-template").Parse("http://js.lsd.test/pbsync.html?gdpr={{.GDPR}}&gdpr_consent={{.GDPRConsent}}&redirect=http%3A%2F%2Fprebidserver.com%2Fsetuid%3Fbidder%3DstroeerCore%26gdpr%3D{{.GDPR}}%26gdpr_consent%3D{{.GDPRConsent}}%26uid%3D"))
	syncer := NewStroeerCoreSyncer(temp)
	syncInfo, err := syncer.GetUsersyncInfo(privacy.Policies{
		GDPR: gdpr.Policy{
			Signal:  "1",
			Consent: "ABCDEF",
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, "http://js.lsd.test/pbsync.html?gdpr=1&gdpr_consent=ABCDEF&redirect=http%3A%2F%2Fprebidserver.com%2Fsetuid%3Fbidder%3DstroeerCore%26gdpr%3D1%26gdpr_consent%3DABCDEF%26uid%3D", syncInfo.URL)
	assert.Equal(t, "iframe", syncInfo.Type)
	assert.EqualValues(t, 136, syncer.GDPRVendorID())
	assert.Equal(t, false, syncInfo.SupportCORS)
	assert.Equal(t, "stroeerCore", syncer.FamilyName())
}
