package engagebdr

import (
	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
	"text/template"
)

func NewEngageBDRSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("engagebdr", 62, temp, adapters.SyncTypeIframe)
}
