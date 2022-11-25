package internal

import (
	"context"

	"github.com/cardil/ip-tax-contributions/pkg/config"
	"github.com/cardil/ip-tax-contributions/pkg/io"
	"github.com/cardil/ip-tax-contributions/pkg/iptax"
	"github.com/cardil/ip-tax-contributions/pkg/markdown"
)

func Execute(ctx context.Context, args Args) error {
	var (
		cfg           config.Config
		period        iptax.Period
		contributions []iptax.Contribution
		err           error
	)

	cfg, err = config.Load(args.ConfigPath)
	if err != nil {
		return err
	}
	period, err = iptax.NewPeriod(args.For, args.MonthEnd)
	if err != nil {
		return err
	}
	contributions, err = iptax.CollectContributions(period, cfg)
	if err != nil {
		return err
	}
	present(ctx, period, contributions)
	return nil
}

func present(ctx context.Context, period iptax.Period, contributions iptax.Contributions) {
	md := markdown.NewBuilder()
	md.Title("IP Tax Report for " + period.String())
	md.Text("Number of PRs:", len(contributions))
	md.Header("Pull Requests")
	for _, c := range contributions {
		u := c.URL()
		md.Link(c.Ref(), u.String())
	}
	md.Header("Repositories")
	for _, repo := range contributions.ToRepos() {
		u := repo.URL()
		md.Link(repo.Name, u.String())
	}
	md.WriteTo(io.OutputFrom(ctx))
}
