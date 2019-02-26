/*
 * Copyright (C) Lenovo Corp. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package x509

import "syscall"

func init() {
	v, err := syscall.GetVersion()
	if err != nil {
		return
	}
	if major := byte(v); major < 6 {
		// Windows XP SP2 and Windows 2003 do not support SHA2.
		// http://blogs.technet.com/b/pki/archive/2010/09/30/sha2-and-windows.aspx
		supportSHA2 = false
	}
}
