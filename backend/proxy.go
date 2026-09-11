package backend

import (
	"net/http"
	"net/url"
)

func HTTPClientAppendProxySetting(c *http.Client) *http.Client {
	settings, err := LoadConfigSettings()
	if err != nil || settings == nil {
		return c
	}

	scheme, ok := settings["proxyScheme"].(string)
	addr, ok1 := settings["proxyAddr"].(string)
	port, ok2 := settings["proxyPort"].(string)
	var proxyUrl *url.URL

	if !ok || !ok1 || !ok2 {
		return c
	}

	if scheme == "no" {
		return c
	}

	proxyUrl = &url.URL{
		Scheme: scheme,
		Host:   addr + ":" + port,
	}

	c.Transport = &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}

	return c
}
