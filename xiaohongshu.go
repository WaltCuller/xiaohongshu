package xiaohongshu

import "github.com/WaltCuller/xiaohongshu/packages"

func (a *BaseApp) PackagesList(arg packages.List) {
	err := a.Get("/ark/open_api/v0/packages", arg, nil)
	if err != nil {
		return
	}
}
