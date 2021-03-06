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

//func NewCSGPrimitiveFromPointer(ptr gdnative.Pointer) CSGPrimitive {
func newCSGPrimitiveFromPointer(ptr gdnative.Pointer) CSGPrimitive {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := CSGPrimitive{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type CSGPrimitive struct {
	CSGShape
	owner gdnative.Object
}

func (o *CSGPrimitive) BaseClass() string {
	return "CSGPrimitive"
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *CSGPrimitive) IsInvertingFaces() gdnative.Bool {
	//log.Println("Calling CSGPrimitive.IsInvertingFaces()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGPrimitive", "is_inverting_faces")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false invert_faces bool}], Returns: void
*/
func (o *CSGPrimitive) SetInvertFaces(invertFaces gdnative.Bool) {
	//log.Println("Calling CSGPrimitive.SetInvertFaces()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(invertFaces)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGPrimitive", "set_invert_faces")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// CSGPrimitiveImplementer is an interface that implements the methods
// of the CSGPrimitive class.
type CSGPrimitiveImplementer interface {
	CSGShapeImplementer
	IsInvertingFaces() gdnative.Bool
	SetInvertFaces(invertFaces gdnative.Bool)
}
