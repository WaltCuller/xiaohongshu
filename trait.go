package xiaohongshu

import (
	"github.com/WaltCuller/xiaohongshu/packages"
	"github.com/WaltCuller/xiaohongshu/refund"
)

type Packages interface {
	PackagesList(arg packages.List) (packages.ListRsp, error)
	PackagesDetail(packageID string) (packages.DetailRsp, error)
}

type Refund interface {
	RefundList(arg refund.List) (refund.ListRsp, error)
	RefundDetail(returnsID string) (refund.DetailRsp, error)
	RefundAudit(arg refund.Audit)
}
