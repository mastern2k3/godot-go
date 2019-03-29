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

//func NewShapeFromPointer(ptr gdnative.Pointer) Shape {
func newShapeFromPointer(ptr gdnative.Pointer) Shape {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Shape{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Base class for all 3D shape resources. Nodes that inherit from this can be used as shapes for a [PhysicsBody] or [Area] objects.
*/
type Shape struct {
	Resource
	owner gdnative.Object
}

func (o *Shape) BaseClass() string {
	return "Shape"
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *Shape) GetMargin() gdnative.Real {
	//log.Println("Calling Shape.GetMargin()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Shape", "get_margin")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false margin float}], Returns: void
*/
func (o *Shape) SetMargin(margin gdnative.Real) {
	//log.Println("Calling Shape.SetMargin()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(margin)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Shape", "set_margin")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// ShapeImplementer is an interface that implements the methods
// of the Shape class.
type ShapeImplementer interface {
	ResourceImplementer
	GetMargin() gdnative.Real
	SetMargin(margin gdnative.Real)
}
