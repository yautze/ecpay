package service

import (
	"ECPay/config"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"reflect"
	"sort"
	"strings"
)

// GenerateCheckMacValue -
func (s *service) GenerateCheckMacValue(ctx context.Context, params map[string]interface{}) string {
	delete(params, "CheckMacValue")
	delete(params, "HashKey")
	delete(params, "HashIV")

	encodedParams := sortMap(params)
	encodedParams = fmt.Sprintf("HashKey=%s&%s&HashIV=%s", config.HashKey, encodedParams, config.HashIV)
	encodedParams = formURLEncode(encodedParams)
	encodedParams = strings.ToLower(encodedParams)
	sum := sha256.Sum256([]byte(encodedParams))
	checkMac := strings.ToUpper(hex.EncodeToString(sum[:]))

	return checkMac
}

// sortMap -
func sortMap(m map[string]interface{}) string {
	// sort map by key
	keys := make([]string, 0, len(m))
	for k := range m {
		if IsNil(m[k]) {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// build str
	str := ""
	for i, k := range keys {
		if i+1 == len(keys) {
			str += fmt.Sprintf("%s=%v", k, m[k])
			continue
		}
		str += fmt.Sprintf("%s=%v&", k, m[k])
	}
	return str
}

// formURLEncode - replace char
func formURLEncode(s string) string {
	s = url.QueryEscape(s)
	s = strings.ReplaceAll(s, "%2d", "-")
	s = strings.ReplaceAll(s, "%5f", "_")
	s = strings.ReplaceAll(s, "%2e", ".")
	s = strings.ReplaceAll(s, "%21", "!")
	s = strings.ReplaceAll(s, "%2a", "*")
	s = strings.ReplaceAll(s, "%28", "(")
	s = strings.ReplaceAll(s, "%29", ")")
	return s
}

// IsNil -
func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}
