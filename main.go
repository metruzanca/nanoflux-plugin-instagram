// Command nanoflux-plugin-instagram serves the Instagram plugin over gRPC so
// nanoflux can load it as an external plugin. Instagram has no public feed; the
// plugin reads the recent-post grid from a profile page and is best-effort by
// nature.
//
// Build it and drop the binary into the host's plugins directory
// (NF_PLUGINS_DIR, ./plugins by default):
//
//	GOOS=linux CGO_ENABLED=0 go build -o ../nanoflux-plugin-instagram .
//
// See instagram/readme.md for the quirks.
package main

import (
	goplugin "github.com/hashicorp/go-plugin"

	instagram "nanoflux-plugin-instagram/instagram"

	"github.com/metruzanca/nanoflux/pluginapi"
)

func main() {
	goplugin.Serve(&goplugin.ServeConfig{
		HandshakeConfig: pluginapi.Handshake,
		Plugins:         pluginapi.PluginSet(&instagram.Plugin{}),
		GRPCServer:      goplugin.DefaultGRPCServer,
	})
}
