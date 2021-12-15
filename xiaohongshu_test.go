package xiaohongshu

import (
	"testing"
)

var app *BaseApp

func TestMain(m *testing.M) {
	app = NewBaseApp("e2a4574fcb", "3c586efc292a6bc49b14f30f7ba8dec5")
	m.Run()
}
