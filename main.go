package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/chyroc/emommit/internal"
	"github.com/urfave/cli/v2"
)

var (
	appName  = "emommit"
	appFlags = []cli.Flag{}
)

func main() {
	app := &cli.App{
		Name:   appName,
		Usage:  "git commit with emoji",
		Flags:  appFlags,
		Action: runApp,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func runApp(c *cli.Context) error {
	emoji, err := internal.SelectEmoji()
	if err != nil {
		return err
	}

	fmt.Println(c.Args())

	msg, err := internal.Input("Input Message")
	if err != nil {
		return err
	}

	msg = fmt.Sprintf("%s %s: %s", emoji.Text, emoji.CommitType, msg)

	// fmt.Printf(msg)

	// git commit [-- args] -m "<generate message>"
	args := c.Args().Slice()
	args = append([]string{"commit"}, args...)
	args = append(args, "-m", msg)

	cmd := exec.Command("git", args...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Env = os.Environ()
	return cmd.Run()
}
