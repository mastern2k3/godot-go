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

//func NewCSGCylinderFromPointer(ptr gdnative.Pointer) CSGCylinder {
func newCSGCylinderFromPointer(ptr gdnative.Pointer) CSGCylinder {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := CSGCylinder{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type CSGCylinder struct {
	CSGPrimitive
	owner gdnative.Object
}

func (o *CSGCylinder) BaseClass() string {
	return "CSGCylinder"
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *CSGCylinder) GetHeight() gdnative.Real {
	//log.Println("Calling CSGCylinder.GetHeight()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGCylinder", "get_height")

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
	Args: [], Returns: Material
*/
func (o *CSGCylinder) GetMaterial() MaterialImplementer {
	//log.Println("Calling CSGCylinder.GetMaterial()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGCylinder", "get_material")

	// Call the parent method.
	// Material
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newMaterialFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(MaterialImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Material" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(MaterialImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *CSGCylinder) GetRadius() gdnative.Real {
	//log.Println("Calling CSGCylinder.GetRadius()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGCylinder", "get_radius")

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
	Args: [], Returns: int
*/
func (o *CSGCylinder) GetSides() gdnative.Int {
	//log.Println("Calling CSGCylinder.GetSides()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGCylinder", "get_sides")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *CSGCylinder) GetSmoothFaces() gdnative.Bool {
	//log.Println("Calling CSGCylinder.GetSmoothFaces()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGCylinder", "get_smooth_faces")

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
	Args: [], Returns: bool
*/
func (o *CSGCylinder) IsCone() gdnative.Bool {
	//log.Println("Calling CSGCylinder.IsCone()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGCylinder", "is_cone")

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
	Args: [{ false cone bool}], Returns: void
*/
func (o *CSGCylinder) SetCone(cone gdnative.Bool) {
	//log.Println("Calling CSGCylinder.SetCone()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(cone)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGCylinder", "set_cone")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false height float}], Returns: void
*/
func (o *CSGCylinder) SetHeight(height gdnative.Real) {
	//log.Println("Calling CSGCylinder.SetHeight()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(height)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGCylinder", "set_height")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false material Material}], Returns: void
*/
func (o *CSGCylinder) SetMaterial(material MaterialImplementer) {
	//log.Println("Calling CSGCylinder.SetMaterial()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(material.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGCylinder", "set_material")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false radius float}], Returns: void
*/
func (o *CSGCylinder) SetRadius(radius gdnative.Real) {
	//log.Println("Calling CSGCylinder.SetRadius()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(radius)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGCylinder", "set_radius")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false sides int}], Returns: void
*/
func (o *CSGCylinder) SetSides(sides gdnative.Int) {
	//log.Println("Calling CSGCylinder.SetSides()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(sides)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGCylinder", "set_sides")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false smooth_faces bool}], Returns: void
*/
func (o *CSGCylinder) SetSmoothFaces(smoothFaces gdnative.Bool) {
	//log.Println("Calling CSGCylinder.SetSmoothFaces()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(smoothFaces)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGCylinder", "set_smooth_faces")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// CSGCylinderImplementer is an interface that implements the methods
// of the CSGCylinder class.
type CSGCylinderImplementer interface {
	CSGPrimitiveImplementer
	GetHeight() gdnative.Real
	GetMaterial() MaterialImplementer
	GetRadius() gdnative.Real
	GetSides() gdnative.Int
	GetSmoothFaces() gdnative.Bool
	IsCone() gdnative.Bool
	SetCone(cone gdnative.Bool)
	SetHeight(height gdnative.Real)
	SetMaterial(material MaterialImplementer)
	SetRadius(radius gdnative.Real)
	SetSides(sides gdnative.Int)
	SetSmoothFaces(smoothFaces gdnative.Bool)
}