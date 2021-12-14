package xiaohongshu

import (
	"github.com/WaltCuller/xiaohongshu/packages"
	"github.com/WaltCuller/xiaohongshu/refund"
)

type App struct {
	base                  *BaseApp
	AccessToken           string `mapstructure:"access_token"`
	ExpiresAt             uint32 `mapstructure:"expires_at"`
	RefreshToken          string `mapstructure:"refresh_token"`
	RefreshTokenExpiresAt uint32 `mapstructure:"refresh_token_expires_at"`
	Code                  string `mapstructure:"code"`
	SellerID              uint64 `mapstructure:"seller_id"`
	SellerName            string `mapstructure:"seller_name"`
	Error                 error  `mapstructure:"-"`
}

func (a *App) PackagesList(arg packages.ReqList) (rsp packages.RspList, err error) {
	err = a.base.NewRequest(packages.METHOD_PACKAGES_LIST, arg, &rsp)
	return
}

func (a *App) PackagesDetail(arg packages.ReqDetail) (rsp packages.RspDetail, err error) {
	err = a.base.NewRequest(packages.METHOD_PACKAGES_DETAIL, arg, &rsp)
	return
}

func (a *App) RefundList(arg refund.ReqList) (rsp refund.RspList, err error) {
	err = a.base.NewRequest(refund.METHOD_REFUND_LIST, arg, &rsp)
	return
}

func (a *App) RefundDetail(arg refund.ReqDetail) (rsp refund.RspDetail, err error) {
	err = a.base.NewRequest(refund.METHOD_REFUND_DETAIL, arg, &rsp)
	return
}
