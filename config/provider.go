/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"saagie.io/provider-ovh/config/database"
	"saagie.io/provider-ovh/config/kube"
	"saagie.io/provider-ovh/config/user"
)

const (
	resourcePrefix = "ovh"
	modulePath     = "saagie.io/provider-ovh"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("ovh.saagie.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		database.Configure,
		kube.Configure,
		user.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
