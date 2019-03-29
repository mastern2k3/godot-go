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

//func NewResourceFormatSaverGDScriptFromPointer(ptr gdnative.Pointer) ResourceFormatSaverGDScript {
func newResourceFormatSaverGDScriptFromPointer(ptr gdnative.Pointer) ResourceFormatSaverGDScript {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := ResourceFormatSaverGDScript{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type ResourceFormatSaverGDScript struct {
	ResourceFormatSaver
	owner gdnative.Object
}

func (o *ResourceFormatSaverGDScript) BaseClass() string {
	return "ResourceFormatSaverGDScript"
}

// ResourceFormatSaverGDScriptImplementer is an interface that implements the methods
// of the ResourceFormatSaverGDScript class.
type ResourceFormatSaverGDScriptImplementer interface {
	ResourceFormatSaverImplementer
}