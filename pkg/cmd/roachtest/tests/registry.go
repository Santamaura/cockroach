// Copyright 2018 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package tests

import "github.com/cockroachdb/cockroach/pkg/cmd/roachtest/registry"

// RegisterTests registers all tests to the Registry. This powers `roachtest run`.
func RegisterTests(r registry.Registry) {
	registerAcceptance(r)
	registerActiveRecord(r)
	registerAllocator(r)
	registerAlterPK(r)
	registerAutoUpgrade(r)
	registerBackup(r)
	registerBackupMixedVersion(r)
	registerBackupNodeShutdown(r)
	registerCancel(r)
	registerCDC(r)
	registerClearRange(r)
	registerClockJumpTests(r)
	registerClockMonotonicTests(r)
	registerConnectionLatencyTest(r)
	registerCopy(r)
	registerDecommission(r)
	registerDiskFull(r)
	RegisterDiskStalledDetection(r)
	registerDjango(r)
	registerDrop(r)
	registerEncryption(r)
	registerEngineSwitch(r)
	registerFlowable(r)
	registerFollowerReads(r)
	registerGopg(r)
	registerGossip(r)
	registerGORM(r)
	registerHibernate(r, hibernateOpts)
	registerHibernate(r, hibernateSpatialOpts)
	registerHotSpotSplits(r)
	registerImportDecommissioned(r)
	registerImportMixedVersion(r)
	registerImportTPCC(r)
	registerImportTPCH(r)
	registerImportNodeShutdown(r)
	registerInconsistency(r)
	registerIndexes(r)
	RegisterJepsen(r)
	registerJobsMixedVersions(r)
	registerKnex(r)
	registerKV(r)
	registerKVContention(r)
	registerKVQuiescenceDead(r)
	registerKVGracefulDraining(r)
	registerKVScalability(r)
	registerKVSplits(r)
	registerKVRangeLookups(r)
	registerKVMultiStoreWithOverload(r)
	registerLargeRange(r)
	registerLedger(r)
	registerLibPQ(r)
	registerLiquibase(r)
	registerNetwork(r)
	registerPebble(r)
	registerPgjdbc(r)
	registerPgx(r)
	registerNodeJSPostgres(r)
	registerPop(r)
	registerPrivilegeVersionUpgrade(r)
	registerPsycopg(r)
	registerQueue(r)
	registerQuitAllNodes(r)
	registerQuitTransfersLeases(r)
	registerRebalanceLoad(r)
	registerReplicaGC(r)
	registerRestart(r)
	registerRestoreNodeShutdown(r)
	registerRestore(r)
	registerRoachmart(r)
	registerRoachtest(r)
	registerRubyPG(r)
	registerSchemaChangeBulkIngest(r)
	registerSchemaChangeDatabaseVersionUpgrade(r)
	registerSchemaChangeDuringKV(r)
	registerSchemaChangeIndexTPCC100(r)
	registerSchemaChangeIndexTPCC1000(r)
	registerSchemaChangeDuringTPCC1000(r)
	registerSchemaChangeInvertedIndex(r)
	registerSchemaChangeMixedVersions(r)
	registerSchemaChangeRandomLoad(r)
	registerScrubAllChecksTPCC(r)
	registerScrubIndexOnlyTPCC(r)
	registerSecondaryIndexesMultiVersionCluster(r)
	registerSecure(r)
	registerSequelize(r)
	registerSQLAlchemy(r)
	registerSQLSmith(r)
	registerSyncTest(r)
	registerSysbench(r)
	registerTLP(r)
	registerTPCC(r)
	registerTPCDSVec(r)
	registerTPCE(r)
	registerTPCHVec(r)
	registerKVBench(r)
	registerTypeORM(r)
	registerLoadSplits(r)
	registerVersion(r)
	registerYCSB(r)
	registerTPCHBench(r)
	registerOverload(r)
	registerMultiTenantUpgrade(r)
}

// RegisterBenchmarks registers all benchmarks to the registry. This powers `roachtest bench`.
//
// TODO(tbg): it's unclear that `roachtest bench` is that useful, perhaps we make everything
// a roachtest but use a `bench` tag to determine what tests to understand as benchmarks.
func RegisterBenchmarks(r registry.Registry) {
	registerIndexesBench(r)
	registerTPCCBench(r)
	registerKVBench(r)
	registerTPCHBench(r)
}
