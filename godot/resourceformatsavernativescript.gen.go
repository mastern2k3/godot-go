package godot

import (
	"github.com/shadowapex/godot-go/gdnative"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

//func NewResourceFormatSaverNativeScriptFromPointer(ptr gdnative.Pointer) ResourceFormatSaverNativeScript {
func newResourceFormatSaverNativeScriptFromPointer(ptr gdnative.Pointer) ResourceFormatSaverNativeScript {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := ResourceFormatSaverNativeScript{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type ResourceFormatSaverNativeScript struct {
	ResourceFormatSaver
	owner gdnative.Object
}

func (o *ResourceFormatSaverNativeScript) BaseClass() string {
	return "ResourceFormatSaverNativeScript"
}

// ResourceFormatSaverNativeScriptImplementer is an interface that implements the methods
// of the ResourceFormatSaverNativeScript class.
type ResourceFormatSaverNativeScriptImplementer interface {
	ResourceFormatSaverImplementer
}
