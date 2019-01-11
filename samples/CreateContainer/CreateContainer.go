//Sample to create an overlay network using hcsshim.

package main

import (
	"os"

	flags "github.com/jessevdk/go-flags"

	"github.com/sirupsen/logrus"
)

type options struct {
	// Example of verbosity with level
	NetworkID string `short:"n" long:"NetworkName" required:"true" description:"Network Name"`
}

var opt options
var parser = flags.NewParser(&opt, flags.Default)

func createContainer(networkID string) error {

	//craete container, endpoint, and attach.

	return nil
}

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
	logrus.Infof("%+v", opt)

	err := createContainer(opt.NetworkID)
	if err != nil {
		logrus.Infof("error = %v", err)
	}
}
