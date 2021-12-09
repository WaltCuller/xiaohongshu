package xiaohongshu

import (
	"crypto/md5"
	"encoding/hex"
	"net/url"
	"sort"
)

// Sign 参数签名
func Sign(param HeaderMap, secret, urlStr string) string {
	query := url.Values{}

	sort.Strings(SortKeyList)

	for _, k := range SortKeyList {
		query.Add(k, param[k])
	}
	queryStr := query.Encode()
	if len(SortKeyList) == 2 {
		queryStr = "&" + queryStr
	}
	signStr := urlStr + "?" + queryStr + secret
	h := md5.New()
	h.Write([]byte(signStr))
	return hex.EncodeToString(h.Sum(nil))
}
