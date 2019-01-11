//Sample to add a hostroute policy to an overlay network using hcsshim.
//E.g. .\HostRoute -add -network

package main

import (
	"os"

	"github.com/Microsoft/hcsshim/hcn"
	"github.com/jessevdk/go-flags"
	"github.com/sirupsen/logrus"
)

type Options struct {
	// Example of verbosity with level
	NetworkName string `short:"n" long:"NetworkName" required:"true" description:"Network Name"`
	Remove      bool   `short:"r" long:"Remove" description:"Remove Host Route policy"`
}

var options Options
var parser = flags.NewParser(&options, flags.Default)

func updateHostRoute(networkName string, remove bool) error {

	network, err := hcn.GetNetworkByName(networkName)
	if err != nil {
		return err
	}

	logrus.Infof("Network = %v", network)
	hostRoutePolicy := hcn.NetworkPolicy{
		Type:     hcn.HostRoute,
		Settings: []byte("{}"),
	}

	networkRequest := hcn.PolicyNetworkRequest{
		Policies: []hcn.NetworkPolicy{hostRoutePolicy},
	}

	if remove {
		network.RemovePolicy(networkRequest)
	} else {
		network.AddPolicy(networkRequest)
	}

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
	logrus.Infof("%+v", options)

	err := updateHostRoute(options.NetworkName, options.Remove)
	if err != nil {
		logrus.Infof("error = %v", err)
	}
}
