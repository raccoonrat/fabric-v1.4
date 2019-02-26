/*
 * Copyright (C) Lenovo Corp. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package sm2

import (
	"bufio"
        "compress/bzip2"
	"crypto/rand"
	"crypto/sha1"
        "crypto/sha256"
        "crypto/sha512"
	"encoding/hex"
	"math/big"
	"hash"
        "io"
        "os"
        "strings"

	"fmt"
	"testing"
	"github.com/lenovo/crypto/sm/sm3"
)

func TestSignVerify(t *testing.T) {
	msg := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	priv, err := GenerateKey(rand.Reader)
	if err != nil {
		panic("GenerateKey failed")
	}

	hfunc := sha256.New()
	hfunc.Write(msg)
	hash := hfunc.Sum(nil)

	r, s, err := Sign(rand.Reader, priv, hash)
	if err != nil {
		panic(err)
	}

	ret := Verify(&priv.PublicKey, hash, r, s)
	fmt.Println(ret)
}

func TestBase(t *testing.T) {
	msg := []byte{1,2,3,4}
	priv, err := GenerateKey(rand.Reader)
	if err != nil {
		panic("GenerateKey failed")
	}
	fmt.Printf("D:%s\n" , priv.D.Text(16))
	fmt.Printf("X:%s\n" , priv.X.Text(16))
	fmt.Printf("Y:%s\n" , priv.Y.Text(16))

	hfunc := sm3.New()
	hfunc.Write(msg)
	hash := hfunc.Sum(nil)
	fmt.Printf("hash:%02X\n", hash)

	r, s, err := Sign(rand.Reader, priv, hash)
	if err != nil {
		panic(err)
	}

	fmt.Printf("R:%s\n" , r.Text(16))
	fmt.Printf("S:%s\n" , s.Text(16))


	ret := Verify(&priv.PublicKey, hash, r, s)
	fmt.Println(ret)
}

func testKeyGeneration(t *testing.T, tag string) {
        priv, err := GenerateKey(rand.Reader)
        if err != nil {
                t.Errorf("%s: error: %s", tag, err)
                return
        }
        if !priv.PublicKey.Curve.IsOnCurve(priv.PublicKey.X, priv.PublicKey.Y) {
                t.Errorf("%s: public key invalid: %s", tag, err)
        }
}

func TestKeyGeneration(t *testing.T) {
	testKeyGeneration(t, "sm2p256")
        if testing.Short() {
                return
        }
}

func BenchmarkSign(b *testing.B) {
	b.ResetTimer()
	origin := []byte("testing")
	hashed  := sm3.SumSM3(origin)
	priv, _ := GenerateKey(rand.Reader)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = Sign(rand.Reader, priv, hashed[:])
	}
}

func BenchmarkVerifyP256(b *testing.B) {
        b.ResetTimer()
	origin := []byte("testing")
	hash := sm3.New()
	hash.Write(origin)
	hashed := hash.Sum(nil)
        priv, _ := GenerateKey(rand.Reader)
        r, s, _ := Sign(rand.Reader, priv, hashed)

	b.ReportAllocs()
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                Verify(&priv.PublicKey, hashed, r, s)
        }
}

func BenchmarkSignAndVerify(b *testing.B) {
	b.ResetTimer()
	origin := []byte("testing")
	hash := sm3.New()
	hash.Write(origin)
	hashed := hash.Sum(nil)
	priv, _ := GenerateKey(rand.Reader)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r, s, _ := Sign(rand.Reader, priv, hashed[:])
		Verify(&priv.PublicKey, hashed, r, s)
	}
}

func TestSignAndVerify(t *testing.T) {
	priv, _ := GenerateKey(rand.Reader)

	origin := []byte("testintestintestintestintestintestinggggggtesting")
	hash := sm3.New()
	hash.Write(origin)
	hashed := hash.Sum(nil)
	r, s, err := Sign(rand.Reader, priv, hashed)
	if err != nil {
		t.Errorf(" error signing: %s", err)
		return
	}

	if !Verify(&priv.PublicKey, hashed, r, s) {
		t.Errorf(" Verify failed")
	}

	//hashed[0] ^= 0xff
	hashed[0] = 0x53
	for i := 0; i < len(hashed); i++ {
		hashed[i] = byte(i)
	}
	if Verify(&priv.PublicKey, hashed, r, s) {
		t.Errorf("Verify always works!")
	}
}

func testNonceSafety(t *testing.T, tag string) {
        priv, _ := GenerateKey(rand.Reader)

        origin := []byte("testing")
	hash := sm3.New()
	hash.Write(origin)
	hashed := hash.Sum(nil)
        r0, s0, err := Sign(zeroReader, priv, hashed)
        if err != nil {
                t.Errorf("%s: error signing: %s", tag, err)
                return
        }

        origin = []byte("testing...")
	hash = sm3.New()
	hash.Write(origin)
	hashed = hash.Sum(nil)
        r1, s1, err := Sign(zeroReader, priv, hashed)
        if err != nil {
                t.Errorf("%s: error signing: %s", tag, err)
                return
        }

        if s0.Cmp(s1) == 0 {
                // This should never happen.
                t.Errorf("%s: the signatures on two different messages were the same", tag)
        }

        if r0.Cmp(r1) == 0 {
                t.Errorf("%s: the nonce used for two diferent messages was the same", tag)
        }
}

func TestNonceSafety(t *testing.T) {
        testNonceSafety(t, "sm2p256")
}

func testINDCCA(t *testing.T, tag string) {
        priv, _ := GenerateKey(rand.Reader)

        origin := []byte("testing")
	hash := sm3.New()
	hash.Write(origin)
	hashed := hash.Sum(nil)
        r0, s0, err := Sign(rand.Reader, priv, hashed)
        if err != nil {
                t.Errorf("%s: error signing: %s", tag, err)
                return
        }

        r1, s1, err := Sign(rand.Reader, priv, hashed)
        if err != nil {
                t.Errorf("%s: error signing: %s", tag, err)
                return
        }

        if s0.Cmp(s1) == 0 {
                t.Errorf("%s: two signatures of the same message produced the same result", tag)
        }

        if r0.Cmp(r1) == 0 {
                t.Errorf("%s: two signatures of the same message produced the same nonce", tag)
        }
}

func TestBaseZA(t *testing.T) {

        msg := []byte{1, 2, 3, 4}
        uid := []byte{1, 1, 1, 1}

        priv, err := GenerateKey(rand.Reader)
        if err != nil {
                panic("GenerateKey failed")
        }
        fmt.Printf("D:%s\n", priv.D.Text(16))
        fmt.Printf("X:%s\n", priv.X.Text(16))
        fmt.Printf("Y:%s\n", priv.Y.Text(16))

        za, err := ZA(&priv.PublicKey, uid)
        if err != nil {
                panic("gen ZA failed")
        }

        hfunc := sm3.New()
        hfunc.Write(za)
        hfunc.Write(msg)
        hash := hfunc.Sum(nil)
        fmt.Printf("hash:%02X\n", hash)

        r, s, err := Sign(rand.Reader, priv, hash)
        if err != nil {
                panic(err)
        }

        fmt.Printf("R:%s\n", r.Text(16))
        fmt.Printf("S:%s\n", s.Text(16))

        ret := Verify(&priv.PublicKey, hash, r, s)
        fmt.Println(ret)
}

func BenchmarkSignZA(b *testing.B) {
        b.ResetTimer()
        origin := []byte("testing")
        uid := []byte("Alice")
        priv, _ := GenerateKey(rand.Reader)
        b.ReportAllocs()
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                za, _ := ZA(&priv.PublicKey, uid)
                hfunc := sm3.New()
                hfunc.Write(za)
                hfunc.Write(origin)
                hashed := hfunc.Sum(nil)
                _, _, _ = Sign(rand.Reader, priv, hashed[:])
        }
}

func BenchmarkVerifyP256ZA(b *testing.B) {
        b.ResetTimer()
        origin := []byte("testing")
        uid := []byte("Alice")
        priv, _ := GenerateKey(rand.Reader)
        zai, _ := ZA(&priv.PublicKey, uid)
        hash := sm3.New()
        hash.Write(zai)
        hash.Write(origin)
        hashedi := hash.Sum(nil)
        r, s, _ := Sign(rand.Reader, priv, hashedi)

        b.ReportAllocs()
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
                za, _ := ZA(&priv.PublicKey, uid)
                hash := sm3.New()
                hash.Write(za)
                hash.Write(origin)
                hashed := hash.Sum(nil)
                Verify(&priv.PublicKey, hashed, r, s)
        }
}

func TestINDCCA(t *testing.T) {
        testINDCCA(t, "sm2p256")
}

func fromHex(s string) *big.Int {
        r, ok := new(big.Int).SetString(s, 16)
        if !ok {
                panic("bad hex")
        }
        return r
}

func TestVectors(t *testing.T) {
	// This test runs the full set of NIST test vectors from
	// http://csrc.nist.gov/groups/STM/cavp/documents/dss/186-3ecdsatestvectors.zip
	//
	// The SigVer.rsp file has been edited to remove test vectors for
	// unsupported algorithms and has been compressed.

	if testing.Short() {
		return
	}

	f, err := os.Open("testdata/SigVer.rsp.bz2")
	if err != nil {
		t.Fatal(err)
	}

	buf := bufio.NewReader(bzip2.NewReader(f))

	lineNo := 1
	var h hash.Hash
	var msg []byte
	var hashed []byte
	var r, s *big.Int
	pub := new(PublicKey)

	for {
		line, err := buf.ReadString('\n')
		if len(line) == 0 {
			if err == io.EOF {
				break
			}
			t.Fatalf("error reading from input: %s", err)
		}
		lineNo++
		// Need to remove \r\n from the end of the line.
		if !strings.HasSuffix(line, "\r\n") {
			t.Fatalf("bad line ending (expected \\r\\n) on line %d", lineNo)
		}
		line = line[:len(line)-2]

		if len(line) == 0 || line[0] == '#' {
			continue
		}

		if line[0] == '[' {
			line = line[1 : len(line)-1]
			parts := strings.SplitN(line, ",", 2)

			switch parts[0] {
			case "P-256":
				pub.Curve = SM2P256V1()
			default:
				pub.Curve = nil
			}

			switch parts[1] {
			case "SHA-1":
				h = sha1.New()
			case "SHA-224":
				h = sha256.New224()
			case "SHA-256":
				h = sha256.New()
			case "SHA-384":
				h = sha512.New384()
			case "SHA-512":
				h = sha512.New()
			default:
				h = nil
			}

			continue
		}

		if h == nil || pub.Curve == nil {
			continue
		}

		switch {
		case strings.HasPrefix(line, "Msg = "):
			if msg, err = hex.DecodeString(line[6:]); err != nil {
				t.Fatalf("failed to decode message on line %d: %s", lineNo, err)
			}
		case strings.HasPrefix(line, "Qx = "):
			pub.X = fromHex(line[5:])
		case strings.HasPrefix(line, "Qy = "):
			pub.Y = fromHex(line[5:])
		case strings.HasPrefix(line, "R = "):
			r = fromHex(line[4:])
		case strings.HasPrefix(line, "S = "):
			s = fromHex(line[4:])
		case strings.HasPrefix(line, "Result = "):
			expected := line[9] == 'P'
			h.Reset()
			h.Write(msg)
			hashed := h.Sum(hashed[:0])
			if Verify(pub, hashed, r, s) != expected {
				t.Fatalf("incorrect result on line %d", lineNo)
			}
		default:
			t.Fatalf("unknown variable on line %d: %s", lineNo, line)
		}
	}
}

