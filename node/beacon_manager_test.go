// Copyright (C) 2022-2023, Dijets Inc, All rights reserved.
// See the file LICENSE for licensing terms.

package node

import (
	"sync"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/require"

	"github.com/lasthyphen/dijetsnodego/ids"
	"github.com/lasthyphen/dijetsnodego/snow/networking/router"
	"github.com/lasthyphen/dijetsnodego/snow/validators"
	"github.com/lasthyphen/dijetsnodego/utils/constants"
	"github.com/lasthyphen/dijetsnodego/utils/timer"
	"github.com/lasthyphen/dijetsnodego/version"
)

const numValidators = 5_000

// Tests that reconnects that mutate the beacon manager's current total stake
// weight is consistent. Test is not deterministic.
func TestBeaconManager_DataRace(t *testing.T) {
	r := require.New(t)

	validatorIDs := make([]ids.NodeID, 0, numValidators)
	validatorSet := validators.NewSet()
	for i := 0; i < numValidators; i++ {
		nodeID := ids.GenerateTestNodeID()

		r.NoError(validatorSet.Add(nodeID, nil, ids.Empty, 1))
		validatorIDs = append(validatorIDs, nodeID)
	}

	wg := &sync.WaitGroup{}

	ctrl := gomock.NewController(t)
	mockRouter := router.NewMockRouter(ctrl)

	b := beaconManager{
		Router:        mockRouter,
		timer:         timer.NewTimer(nil),
		beacons:       validatorSet,
		requiredConns: numValidators,
	}

	// connect numValidators validators, each with a weight of 1
	wg.Add(2 * numValidators)
	mockRouter.EXPECT().
		Connected(gomock.Any(), gomock.Any(), gomock.Any()).
		Times(2 * numValidators).
		Do(func(ids.NodeID, *version.Application, ids.ID) {
			wg.Done()
		})

	for _, nodeID := range validatorIDs {
		nodeID := nodeID
		go func() {
			b.Connected(nodeID, version.CurrentApp, constants.PrimaryNetworkID)
			b.Connected(nodeID, version.CurrentApp, ids.GenerateTestID())
		}()
	}
	wg.Wait()

	// we should have a weight of numValidators now
	r.EqualValues(numValidators, b.numConns)

	// disconnect numValidators validators
	wg.Add(numValidators)
	mockRouter.EXPECT().
		Disconnected(gomock.Any()).
		Times(numValidators).
		Do(func(ids.NodeID) {
			wg.Done()
		})

	for _, nodeID := range validatorIDs {
		go b.Disconnected(nodeID)
	}
	wg.Wait()

	// we should a weight of zero now
	r.Zero(b.numConns)
}
