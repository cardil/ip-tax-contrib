package cmd

import (
	"os"

	"github.com/cardil/ip-tax-contributions/internal"
	"github.com/cardil/ip-tax-contributions/pkg"
	"github.com/cardil/ip-tax-contributions/pkg/io"
	"github.com/spf13/cobra"
	"github.com/wavesoftware/go-commandline"
)

// Options to override the commandline for testing purposes.
var Options []commandline.Option //nolint:gochecknoglobals

// App is a main Ght application.
type App struct {
	internal.Args
}

func (a *App) Command() *cobra.Command {
	c := &cobra.Command{
		Use:          pkg.AppName,
		Short:        "Generate a list of contributions for 🇵🇱 IP Tax form",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return internal.Execute(
				io.WithOutput(cmd.Context(), cmd),
				a.Args,
			)
		},
	}
	c.SetOut(os.Stdout)
	a.SetFlags(c)
	return c
}

func (a *App) SetFlags(c *cobra.Command) {
	defs := a.Defaults()
	fl := c.Flags()
	fl.StringVarP(&a.ConfigPath, "config", "c",
		defs.ConfigPath, "path to configuration file")
	fl.StringVarP(&a.For, "for", "f", defs.For, "for date")
	fl.IntVarP(&a.MonthEnd, "month-end", "m", defs.MonthEnd, "month tax ip day")
}
