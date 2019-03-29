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

//func NewOpenSimplexNoiseFromPointer(ptr gdnative.Pointer) OpenSimplexNoise {
func newOpenSimplexNoiseFromPointer(ptr gdnative.Pointer) OpenSimplexNoise {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := OpenSimplexNoise{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type OpenSimplexNoise struct {
	Resource
	owner gdnative.Object
}

func (o *OpenSimplexNoise) BaseClass() string {
	return "OpenSimplexNoise"
}

/*
        Undocumented
	Args: [{ false width int} { false height int}], Returns: Image
*/
func (o *OpenSimplexNoise) GetImage(width gdnative.Int, height gdnative.Int) ImageImplementer {
	//log.Println("Calling OpenSimplexNoise.GetImage()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(width)
	ptrArguments[1] = gdnative.NewPointerFromInt(height)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OpenSimplexNoise", "get_image")

	// Call the parent method.
	// Image
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newImageFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(ImageImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Image" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(ImageImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *OpenSimplexNoise) GetLacunarity() gdnative.Real {
	//log.Println("Calling OpenSimplexNoise.GetLacunarity()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OpenSimplexNoise", "get_lacunarity")

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
	Args: [{ false x float} { false y float}], Returns: float
*/
func (o *OpenSimplexNoise) GetNoise2D(x gdnative.Real, y gdnative.Real) gdnative.Real {
	//log.Println("Calling OpenSimplexNoise.GetNoise2D()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromReal(x)
	ptrArguments[1] = gdnative.NewPointerFromReal(y)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OpenSimplexNoise", "get_noise_2d")

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
	Args: [{ false pos Vector2}], Returns: float
*/
func (o *OpenSimplexNoise) GetNoise2Dv(pos gdnative.Vector2) gdnative.Real {
	//log.Println("Calling OpenSimplexNoise.GetNoise2Dv()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(pos)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OpenSimplexNoise", "get_noise_2dv")

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
	Args: [{ false x float} { false y float} { false z float}], Returns: float
*/
func (o *OpenSimplexNoise) GetNoise3D(x gdnative.Real, y gdnative.Real, z gdnative.Real) gdnative.Real {
	//log.Println("Calling OpenSimplexNoise.GetNoise3D()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromReal(x)
	ptrArguments[1] = gdnative.NewPointerFromReal(y)
	ptrArguments[2] = gdnative.NewPointerFromReal(z)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OpenSimplexNoise", "get_noise_3d")

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
	Args: [{ false pos Vector3}], Returns: float
*/
func (o *OpenSimplexNoise) GetNoise3Dv(pos gdnative.Vector3) gdnative.Real {
	//log.Println("Calling OpenSimplexNoise.GetNoise3Dv()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(pos)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OpenSimplexNoise", "get_noise_3dv")

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
	Args: [{ false x float} { false y float} { false z float} { false w float}], Returns: float
*/
func (o *OpenSimplexNoise) GetNoise4D(x gdnative.Real, y gdnative.Real, z gdnative.Real, w gdnative.Real) gdnative.Real {
	//log.Println("Calling OpenSimplexNoise.GetNoise4D()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromReal(x)
	ptrArguments[1] = gdnative.NewPointerFromReal(y)
	ptrArguments[2] = gdnative.NewPointerFromReal(z)
	ptrArguments[3] = gdnative.NewPointerFromReal(w)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OpenSimplexNoise", "get_noise_4d")

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
func (o *OpenSimplexNoise) GetOctaves() gdnative.Int {
	//log.Println("Calling OpenSimplexNoise.GetOctaves()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OpenSimplexNoise", "get_octaves")

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
func (o *OpenSimplexNoise) GetPeriod() gdnative.Real {
	//log.Println("Calling OpenSimplexNoise.GetPeriod()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OpenSimplexNoise", "get_period")

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
	Args: [], Returns: float
*/
func (o *OpenSimplexNoise) GetPersistence() gdnative.Real {
	//log.Println("Calling OpenSimplexNoise.GetPersistence()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OpenSimplexNoise", "get_persistence")

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
	Args: [{ false size int}], Returns: Image
*/
func (o *OpenSimplexNoise) GetSeamlessImage(size gdnative.Int) ImageImplementer {
	//log.Println("Calling OpenSimplexNoise.GetSeamlessImage()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(size)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OpenSimplexNoise", "get_seamless_image")

	// Call the parent method.
	// Image
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newImageFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(ImageImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Image" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(ImageImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *OpenSimplexNoise) GetSeed() gdnative.Int {
	//log.Println("Calling OpenSimplexNoise.GetSeed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OpenSimplexNoise", "get_seed")

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
	Args: [{ false lacunarity float}], Returns: void
*/
func (o *OpenSimplexNoise) SetLacunarity(lacunarity gdnative.Real) {
	//log.Println("Calling OpenSimplexNoise.SetLacunarity()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(lacunarity)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OpenSimplexNoise", "set_lacunarity")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false octave_count int}], Returns: void
*/
func (o *OpenSimplexNoise) SetOctaves(octaveCount gdnative.Int) {
	//log.Println("Calling OpenSimplexNoise.SetOctaves()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(octaveCount)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OpenSimplexNoise", "set_octaves")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false period float}], Returns: void
*/
func (o *OpenSimplexNoise) SetPeriod(period gdnative.Real) {
	//log.Println("Calling OpenSimplexNoise.SetPeriod()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(period)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OpenSimplexNoise", "set_period")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false persistence float}], Returns: void
*/
func (o *OpenSimplexNoise) SetPersistence(persistence gdnative.Real) {
	//log.Println("Calling OpenSimplexNoise.SetPersistence()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(persistence)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OpenSimplexNoise", "set_persistence")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false seed int}], Returns: void
*/
func (o *OpenSimplexNoise) SetSeed(seed gdnative.Int) {
	//log.Println("Calling OpenSimplexNoise.SetSeed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(seed)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("OpenSimplexNoise", "set_seed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// OpenSimplexNoiseImplementer is an interface that implements the methods
// of the OpenSimplexNoise class.
type OpenSimplexNoiseImplementer interface {
	ResourceImplementer
	GetImage(width gdnative.Int, height gdnative.Int) ImageImplementer
	GetLacunarity() gdnative.Real
	GetNoise2D(x gdnative.Real, y gdnative.Real) gdnative.Real
	GetNoise2Dv(pos gdnative.Vector2) gdnative.Real
	GetNoise3D(x gdnative.Real, y gdnative.Real, z gdnative.Real) gdnative.Real
	GetNoise3Dv(pos gdnative.Vector3) gdnative.Real
	GetNoise4D(x gdnative.Real, y gdnative.Real, z gdnative.Real, w gdnative.Real) gdnative.Real
	GetOctaves() gdnative.Int
	GetPeriod() gdnative.Real
	GetPersistence() gdnative.Real
	GetSeamlessImage(size gdnative.Int) ImageImplementer
	GetSeed() gdnative.Int
	SetLacunarity(lacunarity gdnative.Real)
	SetOctaves(octaveCount gdnative.Int)
	SetPeriod(period gdnative.Real)
	SetPersistence(persistence gdnative.Real)
	SetSeed(seed gdnative.Int)
}
