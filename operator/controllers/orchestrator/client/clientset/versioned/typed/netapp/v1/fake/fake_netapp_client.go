// Copyright 2021 NetApp, Inc. All Rights Reserved.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/netapp/trident/operator/controllers/orchestrator/client/clientset/versioned/typed/netapp/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeTridentV1 struct {
	*testing.Fake
}

// TridentOrchestrators returns a TridentOrchestratorInterface.
// Returns:
//   * FakeTridentOrchestrators
// Example:
//   * tridentOrchestrators := fake.TridentOrchestrators()
//
// -- Doc autogenerated on 2022-05-12 18:43:38.177255 --
func (c *FakeTridentV1) TridentOrchestrators() v1.TridentOrchestratorInterface {
	return &FakeTridentOrchestrators{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeTridentV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
