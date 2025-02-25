// Copyright (C) 2016  The GoHBase Authors.  All rights reserved.
// This file is part of GoHBase.
// Use of this source code is governed by the Apache License 2.0
// that can be found in the COPYING file.

package mock_test

import (
	"net"

	"github.com/LoneWolf38/gohbase/hrpc"
	"github.com/LoneWolf38/gohbase/test/mock"
	regionMock "github.com/LoneWolf38/gohbase/test/mock/region"
	zkMock "github.com/LoneWolf38/gohbase/test/mock/zk"
	"github.com/LoneWolf38/gohbase/zk"
	"github.com/tsuna/gohbase"
)

var (
	_ gohbase.Client      = (*mock.MockClient)(nil)
	_ gohbase.RPCClient   = (*mock.MockRPCClient)(nil)
	_ gohbase.AdminClient = (*mock.MockAdminClient)(nil)
	_ hrpc.Call           = (*mock.MockCall)(nil)
	_ net.Conn            = (*mock.MockConn)(nil)
	_ zk.Client           = (*zkMock.MockClient)(nil)
	_ hrpc.RegionClient   = (*regionMock.MockRegionClient)(nil)
)
