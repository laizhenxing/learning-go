package module

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"webcrawler/errors"
)

// MID代表组件ID
type MID string

// midTemplate 代表组件ID的模板
var midTemplate = "%s%d|%s"

// DefaultSnGen 代表默认的组件序列号生成器
var DefaultSNGen = NewGenerator(1, 0)

// GenMID 会根据给定的参数生成组件ID
func GenMID(mtype Type, sn uint64, maddr net.Addr) (MID, error) {
	if !LegalType(mtype) {
		errMsg := fmt.Sprintf("illegal module type: %s", mtype)
		return "", errors.NewIllegalParameterError(errMsg)
	}
	letter := legalTypeLetterMap[mtype]
	var midStr string
	if maddr == nil {
		midStr = fmt.Sprintf(midTemplate, letter, sn, "")
		midStr = midStr[:len(midStr)-1]
	} else {
		midStr = fmt.Sprintf(midTemplate, letter, sn, maddr)
	}

	return MID(midStr), nil
}

// 判断给定的组件ID是否合法
func LegalMID(mid MID) bool {
	if _, err := SplitMID(mid); err == nil {
		return true
	}

	return false
}

// SplitMID用于分解组件ID
// 依次包含组件类型字母，序列号和组件网络地址（如果有的话）
func SplitMID(mid MID) ([]string, error) {
	var (
		ok                  bool
		letter, snStr, addr string
	)
	midStr := string(mid)
	if len(midStr) <= 1 {
		return nil, errors.NewIllegalParameterError("insufficient MID")
	}
	letter = midStr[:1]
	// 错误类型判断
	if _, ok = legalLetterTypeMap[letter]; !ok {
		return nil, errors.NewIllegalParameterError(
			fmt.Sprintf("illegal module type letter: %s", letter))
	}

	snAndAddr := midStr[1:]
	index := strings.LastIndex(snAndAddr, "|")
	// 只有序列号，没有网络地址
	if index < 0 {
		snStr = snAndAddr
		if !legalSN(snStr) {
			return nil, errors.NewIllegalParameterError(
				fmt.Sprintf("illegal module SN: %s", snStr))
		}
	} else {// 包含网络地址
		snStr = snAndAddr[:index]
		// 判断序列号是否正确
		if !legalSN(snStr) {
			return nil, errors.NewIllegalParameterError(
				fmt.Sprintf("illegal module SN: %s", snStr))
		}
		addr = snAndAddr[index+1:]
		// 分解ip:port => ip port
		index = strings.LastIndex(addr, ":")
		if index < 0 {
			return nil, errors.NewIllegalParameterError(
				fmt.Sprintf("illegal module address: %s", addr))
		}

		ipStr := addr[:index]
		// 判断IP是否正确
		if ip := net.ParseIP(ipStr); ip == nil {
			return nil, errors.NewIllegalParameterError(
				fmt.Sprintf("illegal module IP: %s", ip))
		}
		portStr := addr[index+1:]
		// 判断端口是否正确
		if _, err := strconv.ParseUint(portStr, 10, 64); err != nil {
			return nil, errors.NewIllegalParameterError(
				fmt.Sprintf("illegal module port: %s", portStr))
		}
	}

	return []string{letter, snStr, addr}, nil
}

// 判断序列号的合法性
func legalSN(snStr string) bool {
	// string转uint64
	_, err := strconv.ParseUint(snStr, 10, 64)
	if err != nil {
		return false
	}
	return true
}
