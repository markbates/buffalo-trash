package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gobuffalo/buffalo/meta"
	"github.com/markbates/inflect"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var name inflect.Name
var app meta.App

// trashCmd represents the trash command
var trashCmd = &cobra.Command{
	Use:                "trash",
	Short:              "destroys and recreates a buffalo app",
	DisableFlagParsing: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("you must pass the name of the app to recreate")
		}

		name = inflect.Name(args[0])

		nargs := []string{"new"}

		if len(args) > 1 {
			nargs = append(nargs, args[:len(args)]...)
		}

		nargs = append(nargs, "-f", string(name))
		if err := run("buffalo", nargs...); err != nil {
			return errors.WithStack(err)
		}

		pwd, err := os.Getwd()
		if err != nil {
			return errors.WithStack(err)
		}

		app := meta.New(filepath.Join(pwd, string(name)))

		// it's not using pop, so that's it. nothing to see here. move along.
		if !app.WithPop {
			return nil
		}

		err = os.Chdir(app.Root)
		if err != nil {
			return errors.WithStack(err)
		}

		// purposefully ignoring the drop error since
		// its possible the database didn't exist in the first place.
		run("buffalo", "db", "drop", "-a", "-d")
		if err := run("buffalo", "db", "create", "-a", "-d"); err != nil {
			return errors.WithStack(err)
		}

		return nil
	},
}

// b new coke -f; cd coke; b db drop -d; b db create -d;
func run(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	fmt.Println(strings.Join(cmd.Args, " "))
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func init() {
	rootCmd.AddCommand(trashCmd)
}
