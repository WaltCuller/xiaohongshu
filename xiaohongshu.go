package xiaohongshu

import (
	"github.com/WaltCuller/xiaohongshu/packages"
	"github.com/WaltCuller/xiaohongshu/refund"
)

func (a *BaseApp) PackagesList(arg packages.List) {
	err := a.Get("/ark/open_api/v0/packages", arg, nil)
	if err != nil {
		return
	}
}

func (a *BaseApp) RefundAudit(arg refund.Audit) {
	err := a.Post("/ark/open_api/v0/refund/audit", arg, nil)
	if err != nil {
		return
	}
}
