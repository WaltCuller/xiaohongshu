package xiaohongshu

import (
	packages2 "github.com/WaltCuller/xiaohongshu/packages"
	"testing"
)

var app *BaseApp

func TestMain(m *testing.M) {
	app = NewBaseApp("6888235097809946125", "2c4a1fa3-d4f0-4f47-ab57-1e44d4237e07")
	m.Run()
}

func TestApp_PackagesList(t *testing.T) {
	packages := Packages(app)

	arg := packages2.List{
		StartTime: 1,
		EndTime:   2,
	}

	packages.PackagesList(arg)
}
