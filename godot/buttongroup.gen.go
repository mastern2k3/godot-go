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

//func NewButtonGroupFromPointer(ptr gdnative.Pointer) ButtonGroup {
func newButtonGroupFromPointer(ptr gdnative.Pointer) ButtonGroup {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := ButtonGroup{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Group of [Button]. All direct and indirect children buttons become radios. Only one allows being pressed. [member BaseButton.toggle_mode] should be [code]true[/code].
*/
type ButtonGroup struct {
	Resource
	owner gdnative.Object
}

func (o *ButtonGroup) BaseClass() string {
	return "ButtonGroup"
}

/*
        Returns an [Array] of [Button]s who have this as their [code]ButtonGroup[/code] (see [member BaseButton.group]).
	Args: [], Returns: Array
*/
func (o *ButtonGroup) GetButtons() gdnative.Array {
	//log.Println("Calling ButtonGroup.GetButtons()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ButtonGroup", "get_buttons")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*
        Returns the current pressed button.
	Args: [], Returns: BaseButton
*/
func (o *ButtonGroup) GetPressedButton() BaseButtonImplementer {
	//log.Println("Calling ButtonGroup.GetPressedButton()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ButtonGroup", "get_pressed_button")

	// Call the parent method.
	// BaseButton
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newBaseButtonFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(BaseButtonImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "BaseButton" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(BaseButtonImplementer)
	}

	return &ret
}

// ButtonGroupImplementer is an interface that implements the methods
// of the ButtonGroup class.
type ButtonGroupImplementer interface {
	ResourceImplementer
	GetButtons() gdnative.Array
	GetPressedButton() BaseButtonImplementer
}
