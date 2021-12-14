package xiaohongshu

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func Sign(pm ParamMap, secret string) string {
	signStr := ""
	for _, k := range SortKeyList {
		if len(pm[k].(string)) == 0 {
			continue
		}
		signStr += fmt.Sprintf("%v%v", k, pm[k])
	}
	signStr = ReplaceSpecial(secret + signStr + secret)
	h := md5.New()
	h.Write([]byte(signStr))
	return hex.EncodeToString(h.Sum(nil))
}
