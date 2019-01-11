package main

import (
	"flag"
	"fmt"

	"github.com/Microsoft/hcsshim/hcn"
)

func main() {
	networkName := flag.String("network", "", "Name of the Network")
	flag.Parse()

	//Query a network.
	hnsNetwork, err := hcn.GetNetworkByName(*networkName)
	if err != nil {
		fmt.Printf("Unable to find network %s: %v\n", *networkName, err)
		return
	}
	fmt.Printf("network (v)\n%v\n", hnsNetwork)
	fmt.Printf("network (+v)\n%+v\n", hnsNetwork)
	fmt.Printf("network (#v)\n%#v\n", hnsNetwork)
	fmt.Printf("network (T)\n%T\n", hnsNetwork)

	/*
		localNetworkName := flag.String("localNetworkName", "", "local network")
		remoteDestinationPrefix := flag.String("remoteDestinationPrefix", "", "remote destination prefix")
		remoteVXLANMAC := flag.String("remoteVXLANMAC", "", "remove VXLAN MAC address (linux)/DR MAC address (windows)")
		remotePA := flag.String("remotePA", "", "remote PA")
		remoteVSID := flag.Uint("remoteVSID", 0, "remote network VSID")

		flag.Parse()

		if *localNetworkName == "" ||
			*remoteDestinationPrefix == "" ||
			*remoteVXLANMAC == "" ||
			*remotePA == "" ||
			*remoteVSID == 0 {

			flag.PrintDefaults()
			return
		}

		fmt.Println("localNetworkName:", *localNetworkName)
		fmt.Println("remoteDestinationPrefix:", *remoteDestinationPrefix)
		fmt.Println("remoteVXLANMAC:", *remoteVXLANMAC)
		fmt.Println("remotePA:", *remotePA)
		fmt.Println("remoteVSID:", *remoteVSID)

		hnsnetwork, err := hcn.GetNetworkByName(*localNetworkName)
		if err != nil {
			logrus.Debugf("Unable to find network %v, error: %v", *localNetworkName, err)
			return
		}

		networkPolicySettings := hcn.RemoteSubnetRoutePolicySetting{
			IsolationId:                 uint16(*remoteVSID),
			DistributedRouterMacAddress: *remoteVXLANMAC,
			ProviderAddress:             *remotePA,
			DestinationPrefix:           *remoteDestinationPrefix,
		}

		rawJSON, err := json.Marshal(networkPolicySettings)
		networkPolicy := hcn.NetworkPolicy{
			Type:     hcn.RemoteSubnetRoute,
			Settings: rawJSON,
		}

		policyNetworkRequest := hcn.PolicyNetworkRequest{
			Policies: []hcn.NetworkPolicy{networkPolicy},
		}

		hnsnetwork.AddPolicy(policyNetworkRequest)
	*/
}
