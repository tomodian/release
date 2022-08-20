package cmd

import (
	"fmt"
	"log"
	"os"

	"release/cmd/task"

	"github.com/urfave/cli/v2"
)

// Run it.
func Run(args []string) error {
	wd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	app := &cli.App{
		Name:  "release",
		Usage: "Manage changelog for your release process 🚀",
		Commands: []*cli.Command{
			task.Targets(wd),
			task.Latest(wd),
			task.Unreleased(wd),
			task.Show(wd),
			task.To(wd),
			task.Next(wd),
		},
	}

	if err := app.Run(args); err != nil {
		fmt.Println("Error!")
		fmt.Println(err)
		return err
	}

	return nil
}
