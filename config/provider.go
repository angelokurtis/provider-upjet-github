/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/angelokurtis/provider-github/config/branch"
	"github.com/angelokurtis/provider-github/config/repository"
)

const (
	resourcePrefix = "github"
	modulePath     = "github.com/angelokurtis/provider-github"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("kurtis.dev.br"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		branch.Configure,
		repository.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()

	return pc
}
