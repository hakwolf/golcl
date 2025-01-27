// Automatically generated by the res2go IDE plug-in.
package main

import (
	"github.com/energye/golcl/lcl"
	_ "github.com/energye/golcl/pkgs/winappres"
)

func main() {

	lcl.Application.SetScaled(true)
	lcl.Application.SetTitle("videosrtgui")
	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)
	lcl.Application.CreateForm(&MainForm)
	lcl.Application.CreateForm(&AppSettings)
	lcl.Application.CreateForm(&NewBaiduTranslateEngine)
	lcl.Application.CreateForm(&NewTencentTranslateEngine)
	lcl.Application.CreateForm(&OSSSaveSettings)
	lcl.Application.CreateForm(&NewAliyunAudioEngine)
	lcl.Application.Run()
}
