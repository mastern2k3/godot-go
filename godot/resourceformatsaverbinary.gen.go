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

//func NewResourceFormatSaverBinaryFromPointer(ptr gdnative.Pointer) ResourceFormatSaverBinary {
func newResourceFormatSaverBinaryFromPointer(ptr gdnative.Pointer) ResourceFormatSaverBinary {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := ResourceFormatSaverBinary{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type ResourceFormatSaverBinary struct {
	ResourceFormatSaver
	owner gdnative.Object
}

func (o *ResourceFormatSaverBinary) BaseClass() string {
	return "ResourceFormatSaverBinary"
}

// ResourceFormatSaverBinaryImplementer is an interface that implements the methods
// of the ResourceFormatSaverBinary class.
type ResourceFormatSaverBinaryImplementer interface {
	ResourceFormatSaverImplementer
}
