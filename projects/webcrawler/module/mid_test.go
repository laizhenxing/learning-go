package module

import (
	"fmt"
	"net"
	"strconv"
	"testing"
)

var legalMIDs = []MID{}

var legalTypes = []Type{
	TYPE_DOWNLOADER,
	TYPE_ANALYZER,
	TYPE_PIPELINE,
}

var illegalTypes = []Type{
	Type("OTHER_MODULE_TYPE"),
}

var illegalMIDs = []MID{
	MID("D"),
	MID("DZ"),
	MID("D1|"),
	MID("D1|127.0.0.1:-1"),
	MID("D1|127.0.0.1:"),
	MID("D1|127.0.0.1"),
	MID("D1|127.0.0."),
	MID("D1|127"),
	MID("D1|127.0.0.0.1:8080"),
	MID("DZ|127.0.0.1:8080"),
	MID("A"),
	MID("AZ"),
	MID("A1|"),
	MID("A1|127.0.0.1:-1"),
	MID("A1|127.0.0.1:"),
	MID("A1|127.0.0.1"),
	MID("A1|127.0.0."),
	MID("A1|127"),
	MID("A1|127.0.0.0.1:8080"),
	MID("AZ|127.0.0.1:8080"),
	MID("P"),
	MID("PZ"),
	MID("P1|"),
	MID("P1|127.0.0.1:-1"),
	MID("P1|127.0.0.1:"),
	MID("P1|127.0.0.1"),
	MID("P1|127.0.0."),
	MID("P1|127"),
	MID("P1|127.0.0.0.1:8080"),
	MID("PZ|127.0.0.1:8080"),
	MID("M1|127.0.0.1:8080"),
}

func init()  {
	for _, mt := range legalTypes {
		for mip := range legalIPMap {
			addr, _ := NewAddr("http", mip, 8080)
			mid, _ := GenMID(mt, DefaultSNGen.Get(), addr)
			legalMIDs = append(legalMIDs, mid)
		}
	}
	fmt.Println(legalMIDs)
}

func TestGenAdnSplitMID(t *testing.T) {
	addr,_ := NewAddr("http", "127.0.0.1", 8080)
	addrs := []net.Addr{nil, addr}
	for _, addr := range addrs {
		for _, mt := range legalTypes {
			expectedSN := DefaultSNGen.Get()
			mid, err := GenMID(mt, expectedSN, addr)
			if err != nil {
				t.Fatalf("An error occurs when generatin module ID: %s (type: %s, sn: %d, addr: %s)",
					err, mt, expectedSN, addr)
			}
			expectedLetter := legalTypeLetterMap[mt]
			var expectedAddr string
			if addr != nil {
				expectedAddr = addr.String()
			}
			parts, err := SplitMID(mid)
			if err != nil {
				t.Fatalf("An error occurs when split mid: %s (mid: %s)", err, mid)
			}
			letter, snStr, addrStr := parts[0], parts[1], parts[2]
			if letter != expectedLetter {
				t.Fatalf("Inconsistent letter in MID: expected: %s, actual: %s", expectedLetter, letter)
			}
			sn, err := strconv.ParseUint(snStr, 10, 64)
			if err != nil {
				t.Fatalf("Inconsistent sn in SN: %s, sn: %s", err,snStr)
			}
			if sn != expectedSN {
				t.Fatalf("Inconsistent sn in MID: expected: %d, actual: %d", expectedSN, sn)
			}
			if addrStr != expectedAddr {
				t.Fatalf("Inconsistent addr in MID: expected: %s, actual: %s", expectedAddr, addrStr)
			}
		}
	}

	// 使用错误类型
	for _, addr := range addrs {
		for _, mt := range illegalTypes {
			mid, err := GenMID(mt, DefaultSNGen.Get(), addr)
			if err == nil {
				t.Fatalf("It still can generate module mid with illegal type: %q", mt)
			}
			if string(mid) != "" {
				t.Fatalf("It still can generate module mid %q with illegal type: %q", mid, mt)
			}
		}
	}

	// 分割错误mid
	for _, mid := range illegalMIDs {
		if _, err := SplitMID(mid); err == nil {
			t.Fatalf("It still can split illegal module ID %q", mid)
		}
	}
}

func TestLegalMID(t *testing.T) {
	var (
		addr net.Addr
		mid MID
		err error
	)
	for _, mt := range legalTypes {
		for mip := range legalIPMap {
			sn := DefaultSNGen.Get()
			addr, err = NewAddr("http", mip, 8080)
			if err == nil {
				mid, err = GenMID(mt, sn, addr)
			}
			if err != nil {
				t.Fatalf("An error occurs when judging legality for MID: %s (type: %s, sn: %d, addr: %s)",
					err, mt, sn, addr)
			}
			if !LegalMID(mid) {
				t.Fatalf("The generated MID %q is legal, but don't be detected!", mid)
			}
		}
	}
	for _, illegalMID := range illegalMIDs {
		if LegalMID(illegalMID) {
			t.Fatalf("The MID %q is illegal, but don't be detected!", illegalMID)
		}
	}
}

