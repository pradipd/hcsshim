//Sample to create an overlay network using hcsshim.
//E.g. .\createnetwork.exe -n foo -a "10.0.0.0/24" -g "10.0.0.1" -v 5000

package main

import (
	"encoding/json"
	"os"

	"github.com/Microsoft/hcsshim/hcn"
	"github.com/jessevdk/go-flags"
	"github.com/sirupsen/logrus"
)

type Options struct {
	// Example of verbosity with level
	NetworkName    string `short:"n" long:"NetworkName" required:"true" description:"Network Name"`
	AddressPrefix  string `short:"a" long:"AddressPrefix" required:"true" description:"Address Prefix"`
	GatewayAddress string `short:"g" long:"GatewayAddress" required:"true" description:"Gateway Address"`
	VSID           uint32 `short:"v" long:"VSID" required:"true" description:"VSID"`
}

var options Options
var parser = flags.NewParser(&options, flags.Default)

func createOverlayNetwork(name, addressPrefix, gatewayAddress string, vsid uint32) error {

	network := &hcn.HostComputeNetwork{
		Type: "Overlay",
		Name: name,
		Ipams: []hcn.Ipam{
			{
				Type: "Static",
				Subnets: []hcn.Subnet{
					{
						IpAddressPrefix: addressPrefix,
						Routes: []hcn.Route{
							{
								NextHop:           gatewayAddress,
								DestinationPrefix: "0.0.0.0/0",
							},
						},
					},
				},
			},
		},
		SchemaVersion: hcn.SchemaVersion{
			Major: 2,
			Minor: 0,
		},
	}

	vsidPolicy := &hcn.VsidPolicySetting{
		IsolationId: vsid,
	}
	vsidJSON, err := json.Marshal(vsidPolicy)
	if err != nil {
		return err
	}

	sp := &hcn.SubnetPolicy{
		Type: hcn.VSID,
	}
	sp.Settings = vsidJSON

	spJSON, err := json.Marshal(sp)
	if err != nil {
		return err
	}

	network.Ipams[0].Subnets[0].Policies = append(network.Ipams[0].Subnets[0].Policies, spJSON)

	network, err = network.Create()
	if err != nil {
		return err
	}

	logrus.Infof("Network ID = %s", network.Id)

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

	err := createOverlayNetwork(options.NetworkName, options.AddressPrefix, options.GatewayAddress, options.VSID)
	if err != nil {
		logrus.Infof("error = %v", err)
	}
}
