package databricks

import (
	"get.porter.sh/porter/pkg/mixin"
	"get.porter.sh/porter/pkg/pkgmgmt"
	"get.porter.sh/porter/pkg/porter/version"
	"github.com/squillace/porter-databricks/pkg"
)

func (m *Mixin) PrintVersion(opts version.Options) error {
	metadata := mixin.Metadata{
		Name: "databricks",
		VersionInfo: pkgmgmt.VersionInfo{
			Version: pkg.Version,
			Commit:  pkg.Commit,
			Author:  "ralph squillace",
		},
	}
	return version.PrintVersion(m.Context, opts, metadata)
}
