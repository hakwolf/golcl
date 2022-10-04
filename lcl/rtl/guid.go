package rtl

import (
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/types"
)

// 将TGUID转为字符形式
//
// Convert TGUID to character form.
func GUIDToString(guid types.TGUID) string {
	return guid.ToString()
}

// 将字符形式的GUID转为TGUID结构
//
// Convert GUID in character form to TGUID structure.
func StringToGUID(str string) types.TGUID {
	return api.DStringToGUID(str)
}

// 创建一个新的GUID
//
// Create a new GUID.
func CreateGUID() types.TGUID {
	return api.DCreateGUID()
}