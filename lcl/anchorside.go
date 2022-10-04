//----------------------------------------
// The code is automatically generated by the GenlibLcl tool.
// Copyright © sxm. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	. "github.com/energye/golcl/lcl/api"
	. "github.com/energye/golcl/lcl/types"
	"unsafe"
)

type TAnchorSide struct {
	IObject
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsAnchorSide(obj interface{}) *TAnchorSide {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TAnchorSide{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
// Deprecated: use AsAnchorSide.
func AnchorSideFromInst(inst uintptr) *TAnchorSide {
	return AsAnchorSide(inst)
}

// 新建一个对象来自已经存在的对象实例。
//
// Create a new object from an existing object instance.
// Deprecated: use AsAnchorSide.
func AnchorSideFromObj(obj IObject) *TAnchorSide {
	return AsAnchorSide(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
//
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsAnchorSide.
func AnchorSideFromUnsafePointer(ptr unsafe.Pointer) *TAnchorSide {
	return AsAnchorSide(ptr)
}

// -------------------------- Deprecated end --------------------------
// 返回对象实例指针。
//
// Return object instance pointer.
func (a *TAnchorSide) Instance() uintptr {
	return a.instance
}

// 获取一个不安全的地址。
//
// Get an unsafe address.
func (a *TAnchorSide) UnsafeAddr() unsafe.Pointer {
	return a.ptr
}

// 检测地址是否为空。
//
// Check if the address is empty.
func (a *TAnchorSide) IsValid() bool {
	return a.instance != 0
}

// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (a *TAnchorSide) Is() TIs {
	return TIs(a.instance)
}

// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (a *TAnchorSide) As() TAs {
//    return TAs(a.instance)
//}

// 获取类信息指针。
//
// Get class information pointer.
func TAnchorSideClass() TClass {
	return AnchorSide_StaticClassType()
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (a *TAnchorSide) Assign(Source IObject) {
	AnchorSide_Assign(a.instance, CheckPtr(Source))
}

// 获取类名路径。
//
// Get the class name path.
func (a *TAnchorSide) GetNamePath() string {
	return AnchorSide_GetNamePath(a.instance)
}

// 获取类的类型信息。
//
// Get class type information.
func (a *TAnchorSide) ClassType() TClass {
	return AnchorSide_ClassType(a.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (a *TAnchorSide) ClassName() string {
	return AnchorSide_ClassName(a.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (a *TAnchorSide) InstanceSize() int32 {
	return AnchorSide_InstanceSize(a.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (a *TAnchorSide) InheritsFrom(AClass TClass) bool {
	return AnchorSide_InheritsFrom(a.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (a *TAnchorSide) Equals(Obj IObject) bool {
	return AnchorSide_Equals(a.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (a *TAnchorSide) GetHashCode() int32 {
	return AnchorSide_GetHashCode(a.instance)
}

// 文本类信息。
//
// Text information.
func (a *TAnchorSide) ToString() string {
	return AnchorSide_ToString(a.instance)
}

// 获取组件所有者。
//
// Get component owner.
func (a *TAnchorSide) Owner() *TControl {
	return AsControl(AnchorSide_GetOwner(a.instance))
}

func (a *TAnchorSide) Kind() TAnchorKind {
	return AnchorSide_GetKind(a.instance)
}

func (a *TAnchorSide) Control() *TControl {
	return AsControl(AnchorSide_GetControl(a.instance))
}

func (a *TAnchorSide) SetControl(value IControl) {
	AnchorSide_SetControl(a.instance, CheckPtr(value))
}

func (a *TAnchorSide) Side() TAnchorSideReference {
	return AnchorSide_GetSide(a.instance)
}

func (a *TAnchorSide) SetSide(value TAnchorSideReference) {
	AnchorSide_SetSide(a.instance, value)
}