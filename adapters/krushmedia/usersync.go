package krushmedia

import (
	"text/template"

	"github.com/eugene-fedorenko/prebid-server/adapters"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

func NewKrushmediaSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("krushmedia", 0, temp, adapters.SyncTypeRedirect)
}
