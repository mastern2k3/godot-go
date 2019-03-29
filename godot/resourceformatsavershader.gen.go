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

//func NewResourceFormatSaverShaderFromPointer(ptr gdnative.Pointer) ResourceFormatSaverShader {
func newResourceFormatSaverShaderFromPointer(ptr gdnative.Pointer) ResourceFormatSaverShader {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := ResourceFormatSaverShader{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type ResourceFormatSaverShader struct {
	ResourceFormatSaver
	owner gdnative.Object
}

func (o *ResourceFormatSaverShader) BaseClass() string {
	return "ResourceFormatSaverShader"
}

// ResourceFormatSaverShaderImplementer is an interface that implements the methods
// of the ResourceFormatSaverShader class.
type ResourceFormatSaverShaderImplementer interface {
	ResourceFormatSaverImplementer
}
