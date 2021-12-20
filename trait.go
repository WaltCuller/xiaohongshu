package xiaohongshu

import (
	"github.com/WaltCuller/xiaohongshu/packages"
	"github.com/WaltCuller/xiaohongshu/refund"
)

type Packages interface {
	PackagesList(arg packages.ReqList) (packages.RspList, error)
	PackagesDetail(arg packages.ReqDetail) (packages.RspDetail, error)
	PackagesCancelList(arg packages.ReqCancelList) (packages.RspCancelList, error)
}

type Refund interface {
	RefundList(arg refund.ReqList) (refund.RspList, error)
	RefundDetail(arg refund.ReqDetail) (refund.RspDetail, error)
}

type Oauth interface {
	GetAccessToken()
	RefreshToken()
}
