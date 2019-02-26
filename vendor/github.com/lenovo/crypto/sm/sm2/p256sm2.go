/*
 * Copyright (C) Lenovo Corp. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package sm2

import (
	"math/big"
	"crypto/elliptic"
	"sync"
)


type p256Sm2Curve struct {
		*elliptic.CurveParams
}

var p256Sm2Params *elliptic.CurveParams
var p256sm2Curve p256Sm2Curve
var initoncesm2 sync.Once

// 取自elliptic的p256.go文件，修改曲线参数为sm2
// See FIPS 186-3, section D.2.3
func initP256Sm2() {
	p256Sm2Params = &elliptic.CurveParams{Name: "SM2-P-256"} // 注明为SM2
	//SM2椭	椭 圆 曲 线 公 钥 密 码 算 法 推 荐 曲 线 参 数
	p256Sm2Params.P, _ = new(big.Int).SetString("FFFFFFFEFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF00000000FFFFFFFFFFFFFFFF", 16)
	p256Sm2Params.N, _ = new(big.Int).SetString("FFFFFFFEFFFFFFFFFFFFFFFFFFFFFFFF7203DF6B21C6052B53BBF40939D54123", 16)
	p256Sm2Params.B, _ = new(big.Int).SetString("28E9FA9E9D9F5E344D5A9E4BCF6509A7F39789F515AB8F92DDBCBD414D940E93", 16)
	p256Sm2Params.Gx, _ = new(big.Int).SetString("32C4AE2C1F1981195F9904466A39C9948FE30BBFF2660BE1715A4589334C74C7", 16)
	p256Sm2Params.Gy, _ = new(big.Int).SetString("BC3736A2F4F6779C59BDCEE36B692153D0A9877CC62A474002DF32E52139F0A0", 16)
	p256Sm2Params.BitSize = 256

	p256sm2Curve = p256Sm2Curve{p256Sm2Params}
}

func P256Sm2() elliptic.Curve {
	initoncesm2.Do(initP256Sm2)
	return p256sm2Curve
}
