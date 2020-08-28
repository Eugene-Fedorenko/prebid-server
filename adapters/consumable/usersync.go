package consumable

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

var VENDOR_ID uint16 = 591

func NewConsumableSyncer(temp *template.Template) usersync.Usersyncer {

	return adapters.NewSyncer(
		"consumable",
		VENDOR_ID,
		temp,
		adapters.SyncTypeRedirect)
}
