package xiaohongshu

import (
	packages2 "github.com/WaltCuller/xiaohongshu/packages"
	refund2 "github.com/WaltCuller/xiaohongshu/refund"
	"testing"
)

var app *BaseApp

func TestMain(m *testing.M) {
	app = NewBaseApp("6888235097809946125", "2c4a1fa3-d4f0-4f47-ab57-1e44d4237e07")
	m.Run()
}

func TestSign_Get(t *testing.T) {
	header := app.headerMap()
	header["timestamp"] = "1639046926"
	header["start_time"] = "1"
	header["end_time"] = "2"
	SortKeyList = append(SortKeyList, "start_time", "end_time")
	sign := Sign(header, app.Secret, "/ark/open_api/v0/packages")
	t.Log(sign)
}

func TestSign_Post(t *testing.T) {
	header := app.headerMap()
	header["timestamp"] = "1639046794"
	sign := Sign(header, app.Secret, "/ark/open_api/v0/refund/audit")
	t.Log(sign)
}

func TestApp_PackagesList(t *testing.T) {
	packages := Packages(app)

	arg := packages2.List{
		StartTime: 1,
		EndTime:   2,
	}

	list, err := packages.PackagesList(arg)
	if err != nil {
		return
	}
	t.Log(list.CurrentPage)
}

func TestBaseApp_RefundAudit(t *testing.T) {
	refund := Refund(app)

	arg := refund2.Audit{
		ReturnsId: "R8413773191496918",
	}

	refund.RefundAudit(arg)
}
