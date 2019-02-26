/*
 * Copyright (C) Lenovo Corp. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package x509

import "sync"

var (
	once           sync.Once
	systemRoots    *CertPool
	systemRootsErr error
)

func systemRootsPool() *CertPool {
	once.Do(initSystemRoots)
	return systemRoots
}

func initSystemRoots() {
	systemRoots, systemRootsErr = loadSystemRoots()
}
