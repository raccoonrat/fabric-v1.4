/*
 * Copyright (C) Lenovo Corp. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */
// +build !cgo

package x509

func loadSystemRoots() (*CertPool, error) {
	return execSecurityRoots()
}
