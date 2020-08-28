package zeroclickfraud

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewZeroClickFraudSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("zeroclickfraud", 0, temp, adapters.SyncTypeIframe)
}
