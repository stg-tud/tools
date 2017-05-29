package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	flags "github.com/jessevdk/go-flags"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	parser := flags.NewNamedParser("bblfsh-tools", flags.Default)
	parser.AddCommand("dummy", "", "Run dummy tool", &Dummy{})

	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok {
			if flagsErr.Type == flags.ErrHelp {
				os.Exit(0)
			} else {
				parser.WriteHelp(os.Stderr)
				os.Exit(1)
			}
		}

		logrus.Errorf("exiting with error: %s", err)
		os.Exit(1)
	}
	logrus.Debug("exiting without error")
}
