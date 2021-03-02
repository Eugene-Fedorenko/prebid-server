package audienceNetwork

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewFacebookSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("audienceNetwork", 0, temp, adapters.SyncTypeRedirect)
}
