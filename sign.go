package xiaohongshu

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
)

// HeaderMap 用于包装请求数据
type HeaderMap map[string]string

// GatewayURL 地址
const GatewayURL = "http://flssandbox.xiaohongshu.com"

// SortKeyList 公共参数排序后的字段列表，签名时用到
// https://school.xiaohongshu.com/open/quick-start/sign.html
// 公共参数为app-key和timestamp
var SortKeyList = []string{
	"app-key",
	"timestamp",
}

// BaseApp 应用的基础配置
type BaseApp struct {
	Key        string
	Secret     string
	gatewayURL string
}

// NewBaseApp 实例化基础应用
func NewBaseApp(k, s string) *BaseApp {
	return &BaseApp{Key: k, Secret: s, gatewayURL: GatewayURL}
}

// SetGatewayURL 重置地址
func (b *BaseApp) SetGatewayURL(u string) *BaseApp {
	b.gatewayURL = u
	return b
}

type BaseResp struct {
	Data      interface{} `json:"data"`
	ErrorCode int         `json:"error_code"`
	ErrorMsg  string      `json:"error_msg"`
	Success   bool        `json:"success"`
}

func (b *BaseApp) request(heads HeaderMap, method, urlStr string, rsp interface{}, body io.Reader) error {
	req, err := http.NewRequest(method, urlStr, body)
	if err != nil {
		return err
	}
	// https://school.xiaohongshu.com/open/quick-start/system-parameter.html
	req.Header.Add("timestamp", heads["timestamp"])
	req.Header.Add("app-key", heads["app-key"])
	req.Header.Add("sign", heads["sign"])
	// content-type为application/json;charset=utf-8
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("charset", "utf-8")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	var ret BaseResp
	//r, _ := ioutil.ReadAll(resp.Body)
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return err
	}
	if ret.ErrorCode != 0 || ret.ErrorMsg != "success" {
		return fmt.Errorf("response error %d %s", ret.ErrorCode, ret.ErrorMsg)
	}
	if rsp == nil {
		return nil
	}
	if ret.Data == nil {
		return errors.New("response error data is nil")
	}
	if reflect.TypeOf(rsp).Elem().Kind() == reflect.Interface {
		rd := reflect.ValueOf(ret.Data)
		reflect.ValueOf(rsp).Elem().Set(rd)
		return nil
	}
	return mapstructure.Decode(ret.Data, rsp)
}

func (b *BaseApp) headerMap() HeaderMap {
	return HeaderMap{
		"app-key":   b.Key,
		"timestamp": strconv.FormatInt(time.Now().Unix(), 10),
		"sign":      "",
	}
}

func (b *BaseApp) Put(urlStr string, req, rsp interface{}) error {
	headers := b.headerMap()

	headers["sign"] = Sign(headers, b.Secret, urlStr)
	d, err := json.Marshal(req)
	if err != nil {
		return err
	}
	body := strings.NewReader(string(d))
	return b.request(headers, "PUT", b.gatewayURL+urlStr, rsp, body)
}

func (b *BaseApp) Post(urlStr string, req, rsp interface{}) error {
	headers := b.headerMap()

	headers["sign"] = Sign(headers, b.Secret, urlStr)
	d, err := json.Marshal(req)
	if err != nil {
		return err
	}
	body := strings.NewReader(string(d))
	return b.request(headers, "POST", b.gatewayURL+urlStr, rsp, body)
}

// Get 执行请求
func (b *BaseApp) Get(urlStr string, req, rsp interface{}) error {
	headers := b.headerMap()

	query := url.Values{}
	r := structs.Map(req)
	for k, v := range r {
		vv := fmt.Sprint(v)
		query.Add(k, vv)
		SortKeyList = append(SortKeyList, k)
		headers[k] = vv
	}
	headers["sign"] = Sign(headers, b.Secret, urlStr)

	return b.request(headers, "GET", b.gatewayURL+urlStr+"?"+query.Encode(), rsp, nil)
}

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
