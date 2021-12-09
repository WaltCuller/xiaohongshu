package xiaohongshu

import (
	"github.com/WaltCuller/xiaohongshu/packages"
	"github.com/WaltCuller/xiaohongshu/refund"
)

type Packages interface {
	PackagesList(arg packages.List)
}

type Refund interface {
	RefundAudit(arg refund.Audit)
}
