//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// energy extension
//
//govcl/vcl/api -> golcl/lcl/api/energy_api.go

package api

import "github.com/energye/golcl/lcl/api/dllimports"

func APIInit() {
	uiLib = loadUILib()
}

// 调用Energy导入的API列表中的函数
func EnergyDefSyscallN(trap int, args ...uintptr) uintptr {
	r1, _, _ := dllimports.GetEnergyImportDefFunc(uiLib, trap).Call(args...)
	return r1
}

func GetEnergyImport(index int) *dllimports.ImportTable {
	return dllimports.GetEnergyImport(index)
}
