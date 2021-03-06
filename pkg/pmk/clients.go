// Copyright © 2020 The Platform9 Systems Inc.
package pmk

import (
	"github.com/platform9/pf9ctl/pkg/qbert"
	"github.com/platform9/pf9ctl/pkg/keystone"
	"github.com/platform9/pf9ctl/pkg/resmgr"
	"github.com/platform9/pf9ctl/pkg/cmdexec"

)

const HTTPMaxRetry = 5

// Clients struct encapsulate the collection of
// external services
type Client struct {
	Resmgr   resmgr.Resmgr
	Keystone keystone.Keystone
	Qbert    qbert.Qbert
	Executor cmdexec.Executor
	Segment  Segment
}

// New creates the clients needed by the CLI
// to interact with the external services.
func NewClient(fqdn string, executor cmdexec.Executor) (Client, error) {
	return Client{
		Resmgr:   resmgr.NewResmgr(fqdn, HTTPMaxRetry),
		Keystone: keystone.NewKeystone(fqdn),
		Qbert:    qbert.NewQbert(fqdn),
		Executor: executor,
		Segment:  NewSegment(fqdn),
	}, nil
}
