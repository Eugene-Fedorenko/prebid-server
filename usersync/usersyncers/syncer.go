package usersyncers

import (
	"strings"
	"text/template"

	"github.com/golang/glog"
	ttx "github.com/eugene-fedorenko/prebid-server/adapters/33across"
	"github.com/eugene-fedorenko/prebid-server/adapters/adform"
	"github.com/eugene-fedorenko/prebid-server/adapters/adkernel"
	"github.com/eugene-fedorenko/prebid-server/adapters/adkernelAdn"
	"github.com/eugene-fedorenko/prebid-server/adapters/adman"
	"github.com/eugene-fedorenko/prebid-server/adapters/admixer"
	"github.com/eugene-fedorenko/prebid-server/adapters/adocean"
	"github.com/eugene-fedorenko/prebid-server/adapters/adpone"
	"github.com/eugene-fedorenko/prebid-server/adapters/adtarget"
	"github.com/eugene-fedorenko/prebid-server/adapters/adtelligent"
	"github.com/eugene-fedorenko/prebid-server/adapters/advangelists"
	"github.com/eugene-fedorenko/prebid-server/adapters/aja"
	"github.com/eugene-fedorenko/prebid-server/adapters/appnexus"
	"github.com/eugene-fedorenko/prebid-server/adapters/audienceNetwork"
	"github.com/eugene-fedorenko/prebid-server/adapters/avocet"
	"github.com/eugene-fedorenko/prebid-server/adapters/beachfront"
	"github.com/eugene-fedorenko/prebid-server/adapters/beintoo"
	"github.com/eugene-fedorenko/prebid-server/adapters/brightroll"
	"github.com/eugene-fedorenko/prebid-server/adapters/consumable"
	"github.com/eugene-fedorenko/prebid-server/adapters/conversant"
	"github.com/eugene-fedorenko/prebid-server/adapters/cpmstar"
	"github.com/eugene-fedorenko/prebid-server/adapters/datablocks"
	"github.com/eugene-fedorenko/prebid-server/adapters/dmx"
	"github.com/eugene-fedorenko/prebid-server/adapters/emx_digital"
	"github.com/eugene-fedorenko/prebid-server/adapters/engagebdr"
	"github.com/eugene-fedorenko/prebid-server/adapters/eplanning"
	"github.com/eugene-fedorenko/prebid-server/adapters/gamma"
	"github.com/eugene-fedorenko/prebid-server/adapters/gamoshi"
	"github.com/eugene-fedorenko/prebid-server/adapters/grid"
	"github.com/eugene-fedorenko/prebid-server/adapters/gumgum"
	"github.com/eugene-fedorenko/prebid-server/adapters/improvedigital"
	"github.com/eugene-fedorenko/prebid-server/adapters/ix"
	"github.com/eugene-fedorenko/prebid-server/adapters/lifestreet"
	"github.com/eugene-fedorenko/prebid-server/adapters/lockerdome"
	"github.com/eugene-fedorenko/prebid-server/adapters/logicad"
	"github.com/eugene-fedorenko/prebid-server/adapters/lunamedia"
	"github.com/eugene-fedorenko/prebid-server/adapters/marsmedia"
	"github.com/eugene-fedorenko/prebid-server/adapters/mgid"
	"github.com/eugene-fedorenko/prebid-server/adapters/nanointeractive"
	"github.com/eugene-fedorenko/prebid-server/adapters/ninthdecimal"
	"github.com/eugene-fedorenko/prebid-server/adapters/openx"
	"github.com/eugene-fedorenko/prebid-server/adapters/pubmatic"
	"github.com/eugene-fedorenko/prebid-server/adapters/pulsepoint"
	"github.com/eugene-fedorenko/prebid-server/adapters/rhythmone"
	"github.com/eugene-fedorenko/prebid-server/adapters/rtbhouse"
	"github.com/eugene-fedorenko/prebid-server/adapters/rubicon"
	"github.com/eugene-fedorenko/prebid-server/adapters/sharethrough"
	"github.com/eugene-fedorenko/prebid-server/adapters/smartadserver"
	"github.com/eugene-fedorenko/prebid-server/adapters/smartrtb"
	"github.com/eugene-fedorenko/prebid-server/adapters/somoaudience"
	"github.com/eugene-fedorenko/prebid-server/adapters/sonobi"
	"github.com/eugene-fedorenko/prebid-server/adapters/sovrn"
	"github.com/eugene-fedorenko/prebid-server/adapters/synacormedia"
	"github.com/eugene-fedorenko/prebid-server/adapters/telaria"
	"github.com/eugene-fedorenko/prebid-server/adapters/triplelift"
	"github.com/eugene-fedorenko/prebid-server/adapters/triplelift_native"
	"github.com/eugene-fedorenko/prebid-server/adapters/ucfunnel"
	"github.com/eugene-fedorenko/prebid-server/adapters/unruly"
	"github.com/eugene-fedorenko/prebid-server/adapters/valueimpression"
	"github.com/eugene-fedorenko/prebid-server/adapters/verizonmedia"
	"github.com/eugene-fedorenko/prebid-server/adapters/visx"
	"github.com/eugene-fedorenko/prebid-server/adapters/vrtcal"
	"github.com/eugene-fedorenko/prebid-server/adapters/yieldlab"
	"github.com/eugene-fedorenko/prebid-server/adapters/yieldmo"
	"github.com/eugene-fedorenko/prebid-server/adapters/yieldone"
	"github.com/eugene-fedorenko/prebid-server/adapters/zeroclickfraud"
	"github.com/eugene-fedorenko/prebid-server/adapters/stroeerCore"
	"github.com/eugene-fedorenko/prebid-server/config"
	"github.com/eugene-fedorenko/prebid-server/openrtb_ext"
	"github.com/eugene-fedorenko/prebid-server/usersync"
)

// NewSyncerMap returns a map of all the usersyncer objects.
// The same keys should exist in this map as in the exchanges map.
// Static syncer map will be removed when adapter isolation is complete.
func NewSyncerMap(cfg *config.Configuration) map[openrtb_ext.BidderName]usersync.Usersyncer {
	syncers := make(map[openrtb_ext.BidderName]usersync.Usersyncer, len(cfg.Adapters))

	insertIntoMap(cfg, syncers, openrtb_ext.Bidder33Across, ttx.New33AcrossSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderAdform, adform.NewAdformSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderAdkernel, adkernel.NewAdkernelSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderAdkernelAdn, adkernelAdn.NewAdkernelAdnSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderAdman, adman.NewAdmanSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderAdmixer, admixer.NewAdmixerSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderAdOcean, adocean.NewAdOceanSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderAdpone, adpone.NewadponeSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderAdtarget, adtarget.NewAdtargetSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderAdtelligent, adtelligent.NewAdtelligentSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderAdvangelists, advangelists.NewAdvangelistsSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderAJA, aja.NewAJASyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderAppnexus, appnexus.NewAppnexusSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderAvocet, avocet.NewAvocetSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderBeachfront, beachfront.NewBeachfrontSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderBeintoo, beintoo.NewBeintooSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderBrightroll, brightroll.NewBrightrollSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderConsumable, consumable.NewConsumableSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderConversant, conversant.NewConversantSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderCpmstar, cpmstar.NewCpmstarSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderDatablocks, datablocks.NewDatablocksSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderDmx, dmx.NewDmxSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderEmxDigital, emx_digital.NewEMXDigitalSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderEngageBDR, engagebdr.NewEngageBDRSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderEPlanning, eplanning.NewEPlanningSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderFacebook, audienceNetwork.NewFacebookSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderGamma, gamma.NewGammaSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderGamoshi, gamoshi.NewGamoshiSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderGrid, grid.NewGridSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderGumGum, gumgum.NewGumGumSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderImprovedigital, improvedigital.NewImprovedigitalSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderIx, ix.NewIxSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderLifestreet, lifestreet.NewLifestreetSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderLockerDome, lockerdome.NewLockerDomeSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderLogicad, logicad.NewLogicadSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderLunaMedia, lunamedia.NewLunaMediaSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderMarsmedia, marsmedia.NewMarsmediaSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderMgid, mgid.NewMgidSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderNanoInteractive, nanointeractive.NewNanoInteractiveSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderNinthDecimal, ninthdecimal.NewNinthDecimalSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderOpenx, openx.NewOpenxSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderPubmatic, pubmatic.NewPubmaticSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderPulsepoint, pulsepoint.NewPulsepointSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderRhythmone, rhythmone.NewRhythmoneSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderRTBHouse, rtbhouse.NewRTBHouseSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderRubicon, rubicon.NewRubiconSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderSharethrough, sharethrough.NewSharethroughSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderSomoaudience, somoaudience.NewSomoaudienceSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderSonobi, sonobi.NewSonobiSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderSovrn, sovrn.NewSovrnSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderSmartadserver, smartadserver.NewSmartadserverSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderSmartRTB, smartrtb.NewSmartRTBSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderSynacormedia, synacormedia.NewSynacorMediaSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderTelaria, telaria.NewTelariaSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderTriplelift, triplelift.NewTripleliftSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderTripleliftNative, triplelift_native.NewTripleliftSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderUcfunnel, ucfunnel.NewUcfunnelSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderUnruly, unruly.NewUnrulySyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderValueImpression, valueimpression.NewValueImpressionSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderVerizonMedia, verizonmedia.NewVerizonMediaSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderVisx, visx.NewVisxSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderVrtcal, vrtcal.NewVrtcalSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderYieldlab, yieldlab.NewYieldlabSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderYieldmo, yieldmo.NewYieldmoSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderYieldone, yieldone.NewYieldoneSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderZeroClickFraud, zeroclickfraud.NewZeroClickFraudSyncer)
	insertIntoMap(cfg, syncers, openrtb_ext.BidderStroeerCore, stroeerCore.NewStroeerCoreSyncer)

	return syncers
}

func insertIntoMap(cfg *config.Configuration, syncers map[openrtb_ext.BidderName]usersync.Usersyncer, bidder openrtb_ext.BidderName, syncerFactory func(*template.Template) usersync.Usersyncer) {
	lowercased := strings.ToLower(string(bidder))
	urlString := cfg.Adapters[lowercased].UserSyncURL
	if urlString == "" {
		glog.Warningf("adapters." + string(bidder) + ".usersync_url was not defined, and their usersync API isn't flexible enough for Prebid Server to choose a good default. No usersyncs will be performed with " + string(bidder))
		return
	}
	syncers[bidder] = syncerFactory(template.Must(template.New(lowercased + "_usersync_url").Parse(urlString)))
}
