package commands

import (
	"github.com/cloudfoundry/cli/cf/terminal"
	"github.com/cloudfoundry/cli/plugin"
	"github.com/SAP/cf-mta-plugin/log"
	"github.com/SAP/cf-mta-plugin/ui"
)

type PurgeConfigCommand struct {
	BaseCommand
}

func (c *PurgeConfigCommand) GetPluginCommand() plugin.Command {
	return plugin.Command{
		Name:     "purge-mta-config",
		HelpText: "Purge no longer valid configuration entries",
		UsageDetails: plugin.Usage{
			Usage: "cf purge-mta-config [-u URL]",
			Options: map[string]string{
				deployServiceURLOpt: "Deploy service URL, by default 'deploy-service.<system-domain>'",
			},
		},
	}
}

func (c *PurgeConfigCommand) Execute(args []string) ExecutionStatus {
	log.Tracef("Executing command %q with args %v\n", c.name, args)

	var host string
	flags, err := c.CreateFlags(&host)
	if err != nil {
		ui.Failed(err.Error())
		return Failure
	}
	err = c.ParseFlags(args, nil, flags, nil)
	if err != nil {
		c.Usage(err.Error())
		return Failure
	}

	context, err := c.GetContext()
	if err != nil {
		ui.Failed(err.Error())
		return Failure
	}

	ui.Say("Purging configuration entries in org %s / space %s as %s",
		terminal.EntityNameColor(context.Org),
		terminal.EntityNameColor(context.Space),
		terminal.EntityNameColor(context.Username))

	rc, err := c.NewRestClient(host)
	if err != nil {
		ui.Failed(err.Error())
		return Failure
	}
	if _, err := rc.GetComponents(); err != nil {
		c.reportError(err)
		return Failure
	}
	// TODO: Why do the same thing twice?
	if _, err := rc.GetComponents(); err != nil {
		c.reportError(err)
		return Failure
	}

	// TODO(ivan): the basePath construction for the rest client should
	// be part of each API call, not once when the client is created.
	rc = c.clientFactory.NewRestClient(host, "", "", c.transport, c.jar, c.tokenFactory)
	if err := rc.PurgeConfiguration(context.Org, context.Space); err != nil {
		c.reportError(err)
		return Failure
	}
	ui.Ok()
	return Success
}

func (c *PurgeConfigCommand) reportError(err error) {
	ui.Failed("Could not purge configuration: %v\n", err)
}