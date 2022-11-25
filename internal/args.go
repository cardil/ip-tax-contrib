package internal

import (
	"path"
	"time"

	"github.com/cardil/ip-tax-contributions/pkg"
	"github.com/cardil/ip-tax-contributions/pkg/iptax"
	"github.com/kirsle/configdir"
)

type Args struct {
	ConfigPath string
	For        string
	MonthEnd   int
}

func (a Args) Defaults() Args {
	configPath := configdir.LocalConfig(pkg.AppName)
	err := configdir.MakePath(configPath) // Ensure it exists.
	if err != nil {
		panic(err)
	}
	settingsPth := path.Join(configPath, "settings.yaml")

	dt := time.Now()

	return Args{
		ConfigPath: settingsPth,
		For:        dt.Format(iptax.PeriodLayout),
		MonthEnd:   iptax.MonthEnd,
	}
}
