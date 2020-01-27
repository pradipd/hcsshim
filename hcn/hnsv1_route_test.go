// +build integration

package hcn

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Microsoft/hcsshim/internal/hns"
)

func TestCreateDeleteRouteUsingV1(t *testing.T) {
	network, err := CreateTestOverlayNetwork()
	if err != nil {
		t.Fatal(err)
	}
	endpoint, err := HcnCreateTestEndpoint(network)
	if err != nil {
		t.Fatal(err)
	}

	//The v2 route API is missing rs5, so we must use v1 route API.
	//Therefore, we need to convert the v2 endpoint to v1 endpoint in order to call the v1 route API.
	endpoints := []hns.HNSEndpoint{
		{
			Id: endpoint.Id,
		},
	}
	//Note: This is using the v1 route API
	pl_route, err := hns.AddRoute(endpoints, "10.0.0.0/24", "10.0.0.1", true)
	if err != nil {
		t.Fatal(err)
	}

	jsonString, err := json.Marshal(pl_route)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("SDN Route JSON:\n%s \n", jsonString)

	_, err = pl_route.Delete()
	if err != nil {
		t.Fatal(err)
	}

	err = endpoint.Delete()
	if err != nil {
		t.Fatal(err)
	}
	err = network.Delete()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetRouteByIdV1(t *testing.T) {
	network, err := CreateTestOverlayNetwork()
	if err != nil {
		t.Fatal(err)
	}
	endpoint, err := HcnCreateTestEndpoint(network)
	if err != nil {
		t.Fatal(err)
	}

	//The v2 route API is missing rs5, so we must use v1 route API.
	//Therefore, we need to convert the v2 endpoint to v1 endpoint in order to call the v1 route API.
	endpoints := []hns.HNSEndpoint{
		{
			Id: endpoint.Id,
		},
	}
	//Note: This is using the v1 route API
	pl_route, err := hns.AddRoute(endpoints, "10.0.0.0/24", "10.0.0.1", true)
	if err != nil {
		t.Fatal(err)
	}

	jsonString, err := json.Marshal(pl_route)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("SDN Route JSON:\n%s \n", jsonString)

	foundRoute, err := hns.GetPolicyListByID(pl_route.ID)
	if err != nil {
		t.Fatal(err)
	}
	if foundRoute == nil {
		t.Fatalf("No SDN route found")
	}
	jsonString, err = json.Marshal(foundRoute)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Found SDN Route JSON:\n%s \n", jsonString)

	_, err = pl_route.Delete()
	if err != nil {
		t.Fatal(err)
	}

	err = endpoint.Delete()
	if err != nil {
		t.Fatal(err)
	}
	err = network.Delete()
	if err != nil {
		t.Fatal(err)
	}
}

func TestListRoutesV1(t *testing.T) {
	network, err := CreateTestOverlayNetwork()
	if err != nil {
		t.Fatal(err)
	}
	endpoint, err := HcnCreateTestEndpoint(network)
	if err != nil {
		t.Fatal(err)
	}

	//The v2 route API is missing rs5, so we must use v1 route API.
	//Therefore, we need to convert the v2 endpoint to v1 endpoint in order to call the v1 route API.
	endpoints := []hns.HNSEndpoint{
		{
			Id: endpoint.Id,
		},
	}
	//Note: This is using the v1 route API
	pl_route, err := hns.AddRoute(endpoints, "10.0.0.0/24", "10.0.0.1", true)
	if err != nil {
		t.Fatal(err)
	}

	jsonString, err := json.Marshal(pl_route)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("SDN Route JSON:\n%s \n", jsonString)

	all_policies, err := hns.HNSListPolicyListRequest()
	if err != nil {
		t.Fatal(err)
	}

	//Note: This returns all policies.  Not just route policies.
	//We need to iterate over the collection to get route policies.
	jsonString, err = json.Marshal(all_policies)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("all policies\n%s \n", jsonString)

	_, err = pl_route.Delete()
	if err != nil {
		t.Fatal(err)
	}

	err = endpoint.Delete()
	if err != nil {
		t.Fatal(err)
	}
	err = network.Delete()
	if err != nil {
		t.Fatal(err)
	}
}

func TestRouteAddRemoveEndpointV1(t *testing.T) {
	network, err := CreateTestOverlayNetwork()
	if err != nil {
		t.Fatal(err)
	}
	endpoint, err := HcnCreateTestEndpoint(network)
	if err != nil {
		t.Fatal(err)
	}

	//The v2 route API is missing rs5, so we must use v1 route API.
	//Therefore, we need to convert the v2 endpoint to v1 endpoint in order to call the v1 route API.
	endpoints := []hns.HNSEndpoint{
		{
			Id: endpoint.Id,
		},
	}
	//Note: This is using the v1 route API
	pl_route, err := hns.AddRoute(endpoints, "10.0.0.0/24", "10.0.0.1", true)
	if err != nil {
		t.Fatal(err)
	}

	jsonString, err := json.Marshal(pl_route)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("SDN Route JSON:\n%s \n", jsonString)

	secondEndpoint, err := HcnCreateTestEndpoint(network)
	if err != nil {
		t.Fatal(err)
	}

	//The v2 route API is missing rs5, so we must use v1 route API.
	//Therefore, we need to convert the v2 endpoint to v1 endpoint in order to call the v1 route API.
	secondHNSEndpoint := hns.HNSEndpoint{
		Id: secondEndpoint.Id,
	}

	new_pl_route, err := pl_route.AddEndpoint(&secondHNSEndpoint)
	if err != nil {
		t.Fatal(err)
	}

	if len(new_pl_route.EndpointReferences) != 2 {
		t.Fatalf("Endpoint not added to SDN Route")
	}

	new_pl_route1, err := new_pl_route.RemoveEndpoint(&secondHNSEndpoint)
	if err != nil {
		t.Fatal(err)
	}
	if len(new_pl_route1.EndpointReferences) != 1 {
		t.Fatalf("Endpoint not removed from SDN Route")
	}

	_, err = pl_route.Delete()
	if err != nil {
		t.Fatal(err)
	}
	err = endpoint.Delete()
	if err != nil {
		t.Fatal(err)
	}
	err = network.Delete()
	if err != nil {
		t.Fatal(err)
	}
}

/*
Note: We already used hns.AddRoute above.
func TestAddRoute(t *testing.T) {
	network, err := CreateTestOverlayNetwork()
	if err != nil {
		t.Fatal(err)
	}
	endpoint, err := HcnCreateTestEndpoint(network)
	if err != nil {
		t.Fatal(err)
	}
	route, err := AddRoute([]HostComputeEndpoint{*endpoint}, "169.254.169.254/24", "127.10.0.33", false)
	if err != nil {
		t.Fatal(err)
	}
	jsonString, err := json.Marshal(route)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("SDN Route JSON:\n%s \n", jsonString)
	foundRoute, err := GetRouteByID(route.ID)
	if err != nil {
		t.Fatal(err)
	}
	if foundRoute == nil {
		t.Fatalf("No SDN route found")
	}
	err = route.Delete()
	if err != nil {
		t.Fatal(err)
	}
	err = endpoint.Delete()
	if err != nil {
		t.Fatal(err)
	}
	err = network.Delete()
	if err != nil {
		t.Fatal(err)
	}
}
*/
