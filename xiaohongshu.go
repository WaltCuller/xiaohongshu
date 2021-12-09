package xiaohongshu

import (
	"github.com/WaltCuller/xiaohongshu/packages"
	"github.com/WaltCuller/xiaohongshu/refund"
	"strings"
)

func (a *BaseApp) PackagesList(arg packages.List) (rsp packages.ListRsp, err error) {
	err = a.Get(packages.URL_PACKAGES_LIST, arg, &rsp)
	return
}

func (a *BaseApp) PackagesDetail(packageID string) (rsp packages.DetailRsp, err error) {
	urlPackagesDetail := strings.Replace(packages.URL_PACKAGES_DETAIL, "{package_id}", packageID, 1)
	err = a.Get(urlPackagesDetail, nil, &rsp)
	return
}

func (a *BaseApp) RefundAudit(arg refund.Audit) {
	err := a.Post(refund.URL_REFUND_AUDIT, arg, nil)
	if err != nil {
		return
	}
}

func (a *BaseApp) RefundList(arg refund.List) (rsp refund.ListRsp, err error) {
	err = a.Get(refund.URL_REFUND_LIST, arg, &rsp)
	return
}

func (a *BaseApp) RefundDetail(returnsID string) (rsp refund.DetailRsp, err error) {
	urlRefundDetail := strings.Replace(refund.URL_REFUND_DETAIL, "{returns_id", returnsID, 1)
	err = a.Get(urlRefundDetail, nil, &rsp)
	return
}
