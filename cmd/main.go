package main

import (
	"fmt"
	"io"
	"os"

	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
	"github.com/thxcode/winnet/cmd/nets"
	"github.com/thxcode/winnet/cmd/routes"
	"github.com/urfave/cli"
)

var (
	Version = "dev"
	Commit  = "dev"
)

func main() {
	app := cli.NewApp()
	app.Version = Version
	app.Description = fmt.Sprintf("winnet %s", Commit)
	app.Writer = colorable.NewColorableStdout()
	app.ErrWriter = colorable.NewColorableStderr()
	app.CommandNotFound = func(cliCtx *cli.Context, s string) {
		fmt.Fprintf(cliCtx.App.Writer, "Invalid Command: %s \n\n", s)
		if pcliCtx := cliCtx.Parent(); pcliCtx == nil {
			cli.ShowAppHelpAndExit(cliCtx, 1)
		} else {
			cli.ShowCommandHelpAndExit(cliCtx, pcliCtx.Command.Name, 1)
		}
	}
	app.OnUsageError = func(cliCtx *cli.Context, err error, isSubcommand bool) error {
		fmt.Fprintf(cliCtx.App.Writer, "Incorrect Usage: %s \n\n", err.Error())
		if isSubcommand {
			cli.ShowSubcommandHelp(cliCtx)
		} else {
			cli.ShowAppHelp(cliCtx)
		}
		return nil
	}
	app.Before = func(cliCtx *cli.Context) error {
		logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true, FullTimestamp: true})
		logrus.SetOutput(cliCtx.App.Writer)
		return nil
	}

	app.Commands = []cli.Command{
		nets.NewCommand(),
		routes.NewCommand(),
	}

	if err := app.Run(os.Args); err != nil && err != io.EOF {
		logrus.Fatal(err)
	}

	logrus.Debug("Finished")
}
