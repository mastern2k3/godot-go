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

//func NewVisualShaderNodeUniformFromPointer(ptr gdnative.Pointer) VisualShaderNodeUniform {
func newVisualShaderNodeUniformFromPointer(ptr gdnative.Pointer) VisualShaderNodeUniform {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualShaderNodeUniform{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type VisualShaderNodeUniform struct {
	VisualShaderNode
	owner gdnative.Object
}

func (o *VisualShaderNodeUniform) BaseClass() string {
	return "VisualShaderNodeUniform"
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *VisualShaderNodeUniform) GetUniformName() gdnative.String {
	//log.Println("Calling VisualShaderNodeUniform.GetUniformName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShaderNodeUniform", "get_uniform_name")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false name String}], Returns: void
*/
func (o *VisualShaderNodeUniform) SetUniformName(name gdnative.String) {
	//log.Println("Calling VisualShaderNodeUniform.SetUniformName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShaderNodeUniform", "set_uniform_name")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// VisualShaderNodeUniformImplementer is an interface that implements the methods
// of the VisualShaderNodeUniform class.
type VisualShaderNodeUniformImplementer interface {
	VisualShaderNodeImplementer
	GetUniformName() gdnative.String
	SetUniformName(name gdnative.String)
}
