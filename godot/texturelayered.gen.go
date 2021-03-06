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

// TextureLayeredFlags is an enum for Flags values.
type TextureLayeredFlags int

const (
	TextureLayeredFlagsDefault TextureLayeredFlags = 4
	TextureLayeredFlagFilter   TextureLayeredFlags = 4
	TextureLayeredFlagMipmaps  TextureLayeredFlags = 1
	TextureLayeredFlagRepeat   TextureLayeredFlags = 2
)

//func NewTextureLayeredFromPointer(ptr gdnative.Pointer) TextureLayered {
func newTextureLayeredFromPointer(ptr gdnative.Pointer) TextureLayered {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := TextureLayered{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type TextureLayered struct {
	Resource
	owner gdnative.Object
}

func (o *TextureLayered) BaseClass() string {
	return "TextureLayered"
}

/*
        Undocumented
	Args: [], Returns: Dictionary
*/
func (o *TextureLayered) X_GetData() gdnative.Dictionary {
	//log.Println("Calling TextureLayered.X_GetData()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureLayered", "_get_data")

	// Call the parent method.
	// Dictionary
	retPtr := gdnative.NewEmptyDictionary()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewDictionaryFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false data Dictionary}], Returns: void
*/
func (o *TextureLayered) X_SetData(data gdnative.Dictionary) {
	//log.Println("Calling TextureLayered.X_SetData()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromDictionary(data)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureLayered", "_set_data")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false width int} { false height int} { false depth int} { false format int} {4 true flags int}], Returns: void
*/
func (o *TextureLayered) Create(width gdnative.Int, height gdnative.Int, depth gdnative.Int, format gdnative.Int, flags gdnative.Int) {
	//log.Println("Calling TextureLayered.Create()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 5, 5)
	ptrArguments[0] = gdnative.NewPointerFromInt(width)
	ptrArguments[1] = gdnative.NewPointerFromInt(height)
	ptrArguments[2] = gdnative.NewPointerFromInt(depth)
	ptrArguments[3] = gdnative.NewPointerFromInt(format)
	ptrArguments[4] = gdnative.NewPointerFromInt(flags)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureLayered", "create")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [], Returns: int
*/
func (o *TextureLayered) GetDepth() gdnative.Int {
	//log.Println("Calling TextureLayered.GetDepth()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureLayered", "get_depth")

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
	Args: [], Returns: int
*/
func (o *TextureLayered) GetFlags() gdnative.Int {
	//log.Println("Calling TextureLayered.GetFlags()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureLayered", "get_flags")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*

	Args: [], Returns: enum.Image::Format
*/
func (o *TextureLayered) GetFormat() ImageFormat {
	//log.Println("Calling TextureLayered.GetFormat()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureLayered", "get_format")

	// Call the parent method.
	// enum.Image::Format
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ImageFormat(ret)
}

/*

	Args: [], Returns: int
*/
func (o *TextureLayered) GetHeight() gdnative.Int {
	//log.Println("Calling TextureLayered.GetHeight()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureLayered", "get_height")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false layer int}], Returns: Image
*/
func (o *TextureLayered) GetLayerData(layer gdnative.Int) ImageImplementer {
	//log.Println("Calling TextureLayered.GetLayerData()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(layer)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureLayered", "get_layer_data")

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

	Args: [], Returns: int
*/
func (o *TextureLayered) GetWidth() gdnative.Int {
	//log.Println("Calling TextureLayered.GetWidth()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureLayered", "get_width")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false image Image} { false x_offset int} { false y_offset int} { false layer int} {0 true mipmap int}], Returns: void
*/
func (o *TextureLayered) SetDataPartial(image ImageImplementer, xOffset gdnative.Int, yOffset gdnative.Int, layer gdnative.Int, mipmap gdnative.Int) {
	//log.Println("Calling TextureLayered.SetDataPartial()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 5, 5)
	ptrArguments[0] = gdnative.NewPointerFromObject(image.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromInt(xOffset)
	ptrArguments[2] = gdnative.NewPointerFromInt(yOffset)
	ptrArguments[3] = gdnative.NewPointerFromInt(layer)
	ptrArguments[4] = gdnative.NewPointerFromInt(mipmap)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureLayered", "set_data_partial")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false flags int}], Returns: void
*/
func (o *TextureLayered) SetFlags(flags gdnative.Int) {
	//log.Println("Calling TextureLayered.SetFlags()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(flags)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureLayered", "set_flags")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false image Image} { false layer int}], Returns: void
*/
func (o *TextureLayered) SetLayerData(image ImageImplementer, layer gdnative.Int) {
	//log.Println("Calling TextureLayered.SetLayerData()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromObject(image.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromInt(layer)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureLayered", "set_layer_data")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// TextureLayeredImplementer is an interface that implements the methods
// of the TextureLayered class.
type TextureLayeredImplementer interface {
	ResourceImplementer
	X_GetData() gdnative.Dictionary
	X_SetData(data gdnative.Dictionary)
	Create(width gdnative.Int, height gdnative.Int, depth gdnative.Int, format gdnative.Int, flags gdnative.Int)
	GetDepth() gdnative.Int
	GetFlags() gdnative.Int
	GetHeight() gdnative.Int
	GetLayerData(layer gdnative.Int) ImageImplementer
	GetWidth() gdnative.Int
	SetDataPartial(image ImageImplementer, xOffset gdnative.Int, yOffset gdnative.Int, layer gdnative.Int, mipmap gdnative.Int)
	SetFlags(flags gdnative.Int)
	SetLayerData(image ImageImplementer, layer gdnative.Int)
}
