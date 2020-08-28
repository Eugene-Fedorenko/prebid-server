# StroeerCore Adapter

# Setup

This document serves as a reminder of what to do when installing this adapter.


### Copy files

Copy the following to your copy/fork of prebid server:

* `adapters/stroeerCore/` folder and contents 
* `openrtb_ext/imp_stroeercore.go` file
* `static/bidder-info/stroeerCore.yaml` file
* `static/bidder-params/stroeerCore.json` file


### Bidder setup

Add BidderName constant in `bidders.go` for stroeerCore:
```go
const (
//...
BidderStroeerCore      BidderName = "stroeerCore"
//...
);
```

Add stroeerCore to the BidderMap in `bidders.go`:

```go
var BidderMap = map[string]BidderName{
// ...
"stroeerCore":       BidderStroeerCore,
// ...
}
```

Update the `newAdapterMap` function in `adapter_map.go`:

```go
func newAdapterMap(client *http.Client, cfg *config.Configuration, infos adapters.BidderInfos) map[openrtb_ext.BidderName]adaptedBidder {
	ortbBidders := map[openrtb_ext.BidderName]adapters.Bidder{
        // ...
        openrtb_ext.BidderStroeerCore:      stroeerCore.NewStroeerCoreBidder(cfg.Adapters[strings.ToLower(string(openrtb_ext.BidderStroeerCore))].Endpoint),
        // ...
    }
    // ...
}
```

Be sure this adapter is imported at the top of the `adapter_map.go` file:

```go
import ( 
    // ...
    "github.com/eugene-fedorenko/prebid-server/adapters/stroeerCore"
    // ...
)
```

And set the default value for `adapters.stroeercore.endpoint` in `config.go`:

```go
  v.SetDefault("adapters.stroeercore.endpoint", "http://mhb.adscale.de/s2sdsh")
```

### User synching setup

Be sure your build of universal prebid creative knows about your prebid server (https://github.com/prebid/prebid-universal-creative/).
This is achieved by modifying the `VALID_ENDPOINTS` object in `cookieSync.js`

```javascript 1.8
const VALID_ENDPOINTS = {
  rubicon: 'https://prebid-server.rubiconproject.com/cookie_sync',
  appnexus: 'https://prebid.adnxs.com/pbs/v1/cookie_sync',
  stroeerCore: "https://your.prebid.server.host/cookie_sync"
};
```

Now move your attention back to the prebid server source code. Register stroeerCore syncher in `config.go`:

```go
func (cfg *Configuration) setDerivedDefaults() {
   ...
   setDefaultUsersync(cfg.Adapters, openrtb_ext.BidderStroeerCore, "https://js.adscale.de/pbsync.html?gdpr={{.GDPR}}&gdpr_consent={{.GDPRConsent}}&redirect="+url.QueryEscape(externalURL)+"%2Fsetuid%3Fbidder%3DstroeerCore%26gdpr%3D{{.GDPR}}%26gdpr_consent%3D{{.GDPRConsent}}%26uid%3D")
   ...
}
```

The default user sync TTL for a bidder is 14 days. For us, this is too long. 

We do more than share our user id with prebid server. We trigger our own user syncing with our DSPs. And we need this to happen regularly. Thus the reason for a smaller TTL.

To do this, edit the `cookie.go` file and add `stroeerCore` to `customBidderTTLs`:

```go
// customBidderTTLs stores rules about how long a particular UID sync is valid for each bidder.
// If a bidder does a cookie sync *without* listing a rule here, then the DEFAULT_TTL will be used.
var customBidderTTLs = map[string]time.Duration{
    "stroeerCore": time.Second * 30,
}
```

In the AMP environment, user syncing on the page will look like this:

```html
<html amp lang="en">
    <head>
        <meta charset="utf-8">
        <script async src="https://cdn.ampproject.org/v0.js"></script>
        <script async custom-element="amp-iframe" src="https://cdn.ampproject.org/v0/amp-iframe-0.1.js"></script>
        <meta name="viewport" content="width=device-width,minimum-scale=1,initial-scale=1">
        <style amp-boilerplate>body{-webkit-animation:-amp-start 8s steps(1,end) 0s 1 normal both;-moz-animation:-amp-start 8s steps(1,end) 0s 1 normal both;-ms-animation:-amp-start 8s steps(1,end) 0s 1 normal both;animation:-amp-start 8s steps(1,end) 0s 1 normal both}@-webkit-keyframes -amp-start{from{visibility:hidden}to{visibility:visible}}@-moz-keyframes -amp-start{from{visibility:hidden}to{visibility:visible}}@-ms-keyframes -amp-start{from{visibility:hidden}to{visibility:visible}}@-o-keyframes -amp-start{from{visibility:hidden}to{visibility:visible}}@keyframes -amp-start{from{visibility:hidden}to{visibility:visible}}</style><noscript><style amp-boilerplate>body{-webkit-animation:none;-moz-animation:none;-ms-animation:none;animation:none}</style></noscript>
        <title>AMP Cookiesync Test Page</title>
    </head>
    <body>
        <h1>This is an AMP page</h1>
        <amp-iframe width="1" title="User Sync"
            height="1"
            sandbox="allow-script allow-same-origin"
            frameborder="0"
            src="https://[your-universal-prebid-host]/load-cookie.html?max_sync_count=1&endpoint=stroeerCore">
            <amp-img layout="fill" src="data:image/gif;base64,R0lGODlhAQABAIAAAP///wAAACH5BAEAAAAALAAAAAABAAEAAAICRAEAOw==" placeholder></amp-img>
        </amp-iframe>
    </body>
</html>
```

### Currency setup

Stroeer prebid server uses EUR as the currency.

Be sure the resolved request will have EUR as the only currency on the list.
To make this happen, be sure `cur` is on your stored request. Otherwise USD will be used as the default.

Example setup:
```
{
    "id": "auction-id",
    "cur": ["EUR"]
    ...
}
```

The conversion file used by prebid server is https://cdn.jsdelivr.net/gh/prebid/currency-file@1/latest.json

Bidders are not required to bid in EUR despite the resolved request having "EUR" as the currency. All Non-EUR bids will
be converted using the rates (or derived rates) from the conversion file.

### Prebid config setup

Remember your pbs.json or pbs.yaml file. This specifies how stored requests are retrieved, where the creative cache is, how long to cache creatives, etc.
Example pbs.json (this may be incorrect or incomplete for your case):
```
{
  "stored_requests": {
    "http": {
      "endpoint": "http://[your-stored-request-endpoint]/fetch-stored-requests",
      "amp_endpoint": "http://[your-amp-stored-request-endpoint]/fetch-stored-requests"
    }
  },
  "cache.host": "[your-prebid-cache-host]",
  "cache.scheme": "http",
  "cache.query": "uuid=%PBS_CACHE_UUID%",
  "port": 8300,
  "external_url": "https://[your-prebid-server-host]",
  "gdpr.usersync_if_ambiguous": true,
  "cache.default_ttl_seconds.banner" : 3000,
  "audiencenetwork": {
    "disabled": true
  },
  ... 
}

```
