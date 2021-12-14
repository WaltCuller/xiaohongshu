package xiaohongshu

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

func Sign(pm ParamMap, secret string) string {
	paramJSON := pm["param_json"].(ParamMap)
	if len(paramJSON) == 0 {
		pm["param_json"] = "{}"
	} else {
		var ks []string
		for k := range paramJSON {
			ks = append(ks, k)
		}
		sort.Strings(ks)
		for i, k := range ks {
			ks[i] = fmt.Sprintf(`"%v":"%v"`, k, paramJSON[k])
		}
		pm["param_json"] = "{" + strings.Join(ks, ",") + "}"
	}
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
