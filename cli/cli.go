package cli

import (
	"backup-automation/backup"
	"backup-automation/logs"
	"github.com/urfave/cli/v2"
	"os"
	"time"
)

func CLI() {
	app := &cli.App{
		Name:     "CLI automation",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name: "José Castro",
			},
		},
		ArgsUsage: "[args]",
		Commands: []*cli.Command{
			{
				Name:    "backup",
				Aliases: []string{"back"},
				Usage:   "back up files",
				Action: func(context *cli.Context) error {
					backup.Backup(context.Args().Slice()[0], context.Args().Slice()[1])
					return nil
				},
			},
			{
				Name:    "logs",
				Aliases: []string{"l"},
				Usage:   "show logs",
				Action: func(context *cli.Context) error {
					logs.ShowLogs(context.Args().Slice())
					return nil
				},
			},
		},
		SkipFlagParsing: false,
		HideHelp:        false,
	}

	app.Setup()

	app.Run(os.Args)
}
