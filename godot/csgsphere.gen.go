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

//func NewCSGSphereFromPointer(ptr gdnative.Pointer) CSGSphere {
func newCSGSphereFromPointer(ptr gdnative.Pointer) CSGSphere {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := CSGSphere{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type CSGSphere struct {
	CSGPrimitive
	owner gdnative.Object
}

func (o *CSGSphere) BaseClass() string {
	return "CSGSphere"
}

/*
        Undocumented
	Args: [], Returns: Material
*/
func (o *CSGSphere) GetMaterial() MaterialImplementer {
	//log.Println("Calling CSGSphere.GetMaterial()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGSphere", "get_material")

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
	Args: [], Returns: int
*/
func (o *CSGSphere) GetRadialSegments() gdnative.Int {
	//log.Println("Calling CSGSphere.GetRadialSegments()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGSphere", "get_radial_segments")

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
	Args: [], Returns: float
*/
func (o *CSGSphere) GetRadius() gdnative.Real {
	//log.Println("Calling CSGSphere.GetRadius()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGSphere", "get_radius")

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
func (o *CSGSphere) GetRings() gdnative.Int {
	//log.Println("Calling CSGSphere.GetRings()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGSphere", "get_rings")

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
func (o *CSGSphere) GetSmoothFaces() gdnative.Bool {
	//log.Println("Calling CSGSphere.GetSmoothFaces()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGSphere", "get_smooth_faces")

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
	Args: [{ false material Material}], Returns: void
*/
func (o *CSGSphere) SetMaterial(material MaterialImplementer) {
	//log.Println("Calling CSGSphere.SetMaterial()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(material.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGSphere", "set_material")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false radial_segments int}], Returns: void
*/
func (o *CSGSphere) SetRadialSegments(radialSegments gdnative.Int) {
	//log.Println("Calling CSGSphere.SetRadialSegments()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(radialSegments)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGSphere", "set_radial_segments")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false radius float}], Returns: void
*/
func (o *CSGSphere) SetRadius(radius gdnative.Real) {
	//log.Println("Calling CSGSphere.SetRadius()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(radius)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGSphere", "set_radius")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false rings int}], Returns: void
*/
func (o *CSGSphere) SetRings(rings gdnative.Int) {
	//log.Println("Calling CSGSphere.SetRings()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(rings)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGSphere", "set_rings")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false smooth_faces bool}], Returns: void
*/
func (o *CSGSphere) SetSmoothFaces(smoothFaces gdnative.Bool) {
	//log.Println("Calling CSGSphere.SetSmoothFaces()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(smoothFaces)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CSGSphere", "set_smooth_faces")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// CSGSphereImplementer is an interface that implements the methods
// of the CSGSphere class.
type CSGSphereImplementer interface {
	CSGPrimitiveImplementer
	GetMaterial() MaterialImplementer
	GetRadialSegments() gdnative.Int
	GetRadius() gdnative.Real
	GetRings() gdnative.Int
	GetSmoothFaces() gdnative.Bool
	SetMaterial(material MaterialImplementer)
	SetRadialSegments(radialSegments gdnative.Int)
	SetRadius(radius gdnative.Real)
	SetRings(rings gdnative.Int)
	SetSmoothFaces(smoothFaces gdnative.Bool)
}
