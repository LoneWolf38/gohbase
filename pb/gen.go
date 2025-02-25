// Copyright (C) 2020  The GoHBase Authors.  All rights reserved.
// This file is part of GoHBase.
// Use of this source code is governed by the Apache License 2.0
// that can be found in the COPYING file.

// To run this command you need protoc.
//go:generate go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
//go:generate protoc --proto_path=. --go_out=. Admin.proto Aggregate.proto Authentication.proto Backup.proto BootstrapNode.proto BucketCacheEntry.proto Cell.proto Client.proto ClusterId.proto ClusterStatus.proto ColumnAggregationNullResponseProtocol.proto ColumnAggregationProtocol.proto ColumnAggregationWithErrorsProtocol.proto Comparator.proto DummyRegionServerEndpoint.proto Encryption.proto ErrorHandling.proto Export.proto FS.proto Filter.proto HBase.proto HFile.proto IncrementCounterProcessor.proto LoadBalancer.proto LockService.proto MapReduce.proto Master.proto PingProtocol.proto Procedure.proto Quota.proto RPC.proto RSGroup.proto RSGroupAdmin.proto RecentLogs.proto RegionNormalizer.proto RegionServerStatus.proto Registry.proto Replication.proto ShellExecEndpoint.proto Snapshot.proto SnapshotCleanup.proto StoreFileTracker.proto TestProcedure.proto TooSlowLog.proto Tracing.proto VisibilityLabels.proto WAL.proto ZooKeeper.proto any.proto test.proto test_rpc_service.proto

package pb
