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

//func NewResourceFormatLoaderTextureLayeredFromPointer(ptr gdnative.Pointer) ResourceFormatLoaderTextureLayered {
func newResourceFormatLoaderTextureLayeredFromPointer(ptr gdnative.Pointer) ResourceFormatLoaderTextureLayered {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := ResourceFormatLoaderTextureLayered{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type ResourceFormatLoaderTextureLayered struct {
	ResourceFormatLoader
	owner gdnative.Object
}

func (o *ResourceFormatLoaderTextureLayered) BaseClass() string {
	return "ResourceFormatLoaderTextureLayered"
}

// ResourceFormatLoaderTextureLayeredImplementer is an interface that implements the methods
// of the ResourceFormatLoaderTextureLayered class.
type ResourceFormatLoaderTextureLayeredImplementer interface {
	ResourceFormatLoaderImplementer
}
