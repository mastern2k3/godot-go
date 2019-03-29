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

//func NewAnimationNodeStateMachineFromPointer(ptr gdnative.Pointer) AnimationNodeStateMachine {
func newAnimationNodeStateMachineFromPointer(ptr gdnative.Pointer) AnimationNodeStateMachine {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AnimationNodeStateMachine{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type AnimationNodeStateMachine struct {
	AnimationRootNode
	owner gdnative.Object
}

func (o *AnimationNodeStateMachine) BaseClass() string {
	return "AnimationNodeStateMachine"
}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *AnimationNodeStateMachine) X_TreeChanged() {
	//log.Println("Calling AnimationNodeStateMachine.X_TreeChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeStateMachine", "_tree_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false name String} { false node AnimationNode} {(0, 0) true position Vector2}], Returns: void
*/
func (o *AnimationNodeStateMachine) AddNode(name gdnative.String, node AnimationNodeImplementer, position gdnative.Vector2) {
	//log.Println("Calling AnimationNodeStateMachine.AddNode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromString(name)
	ptrArguments[1] = gdnative.NewPointerFromObject(node.GetBaseObject())
	ptrArguments[2] = gdnative.NewPointerFromVector2(position)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeStateMachine", "add_node")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false from String} { false to String} { false transition AnimationNodeStateMachineTransition}], Returns: void
*/
func (o *AnimationNodeStateMachine) AddTransition(from gdnative.String, to gdnative.String, transition AnimationNodeStateMachineTransitionImplementer) {
	//log.Println("Calling AnimationNodeStateMachine.AddTransition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromString(from)
	ptrArguments[1] = gdnative.NewPointerFromString(to)
	ptrArguments[2] = gdnative.NewPointerFromObject(transition.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeStateMachine", "add_transition")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [], Returns: String
*/
func (o *AnimationNodeStateMachine) GetEndNode() gdnative.String {
	//log.Println("Calling AnimationNodeStateMachine.GetEndNode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeStateMachine", "get_end_node")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*

	Args: [], Returns: Vector2
*/
func (o *AnimationNodeStateMachine) GetGraphOffset() gdnative.Vector2 {
	//log.Println("Calling AnimationNodeStateMachine.GetGraphOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeStateMachine", "get_graph_offset")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*

	Args: [{ false name String}], Returns: AnimationNode
*/
func (o *AnimationNodeStateMachine) GetNode(name gdnative.String) AnimationNodeImplementer {
	//log.Println("Calling AnimationNodeStateMachine.GetNode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeStateMachine", "get_node")

	// Call the parent method.
	// AnimationNode
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newAnimationNodeFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(AnimationNodeImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "AnimationNode" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(AnimationNodeImplementer)
	}

	return &ret
}

/*

	Args: [{ false node AnimationNode}], Returns: String
*/
func (o *AnimationNodeStateMachine) GetNodeName(node AnimationNodeImplementer) gdnative.String {
	//log.Println("Calling AnimationNodeStateMachine.GetNodeName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(node.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeStateMachine", "get_node_name")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false name String}], Returns: Vector2
*/
func (o *AnimationNodeStateMachine) GetNodePosition(name gdnative.String) gdnative.Vector2 {
	//log.Println("Calling AnimationNodeStateMachine.GetNodePosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeStateMachine", "get_node_position")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*

	Args: [], Returns: String
*/
func (o *AnimationNodeStateMachine) GetStartNode() gdnative.String {
	//log.Println("Calling AnimationNodeStateMachine.GetStartNode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeStateMachine", "get_start_node")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false idx int}], Returns: AnimationNodeStateMachineTransition
*/
func (o *AnimationNodeStateMachine) GetTransition(idx gdnative.Int) AnimationNodeStateMachineTransitionImplementer {
	//log.Println("Calling AnimationNodeStateMachine.GetTransition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeStateMachine", "get_transition")

	// Call the parent method.
	// AnimationNodeStateMachineTransition
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newAnimationNodeStateMachineTransitionFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(AnimationNodeStateMachineTransitionImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "AnimationNodeStateMachineTransition" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(AnimationNodeStateMachineTransitionImplementer)
	}

	return &ret
}

/*

	Args: [], Returns: int
*/
func (o *AnimationNodeStateMachine) GetTransitionCount() gdnative.Int {
	//log.Println("Calling AnimationNodeStateMachine.GetTransitionCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeStateMachine", "get_transition_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false idx int}], Returns: String
*/
func (o *AnimationNodeStateMachine) GetTransitionFrom(idx gdnative.Int) gdnative.String {
	//log.Println("Calling AnimationNodeStateMachine.GetTransitionFrom()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeStateMachine", "get_transition_from")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false idx int}], Returns: String
*/
func (o *AnimationNodeStateMachine) GetTransitionTo(idx gdnative.Int) gdnative.String {
	//log.Println("Calling AnimationNodeStateMachine.GetTransitionTo()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeStateMachine", "get_transition_to")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false name String}], Returns: bool
*/
func (o *AnimationNodeStateMachine) HasNode(name gdnative.String) gdnative.Bool {
	//log.Println("Calling AnimationNodeStateMachine.HasNode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeStateMachine", "has_node")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false from String} { false to String}], Returns: bool
*/
func (o *AnimationNodeStateMachine) HasTransition(from gdnative.String, to gdnative.String) gdnative.Bool {
	//log.Println("Calling AnimationNodeStateMachine.HasTransition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(from)
	ptrArguments[1] = gdnative.NewPointerFromString(to)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeStateMachine", "has_transition")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false name String}], Returns: void
*/
func (o *AnimationNodeStateMachine) RemoveNode(name gdnative.String) {
	//log.Println("Calling AnimationNodeStateMachine.RemoveNode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeStateMachine", "remove_node")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false from String} { false to String}], Returns: void
*/
func (o *AnimationNodeStateMachine) RemoveTransition(from gdnative.String, to gdnative.String) {
	//log.Println("Calling AnimationNodeStateMachine.RemoveTransition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(from)
	ptrArguments[1] = gdnative.NewPointerFromString(to)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeStateMachine", "remove_transition")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false idx int}], Returns: void
*/
func (o *AnimationNodeStateMachine) RemoveTransitionByIndex(idx gdnative.Int) {
	//log.Println("Calling AnimationNodeStateMachine.RemoveTransitionByIndex()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeStateMachine", "remove_transition_by_index")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false name String} { false new_name String}], Returns: void
*/
func (o *AnimationNodeStateMachine) RenameNode(name gdnative.String, newName gdnative.String) {
	//log.Println("Calling AnimationNodeStateMachine.RenameNode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(name)
	ptrArguments[1] = gdnative.NewPointerFromString(newName)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeStateMachine", "rename_node")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false name String}], Returns: void
*/
func (o *AnimationNodeStateMachine) SetEndNode(name gdnative.String) {
	//log.Println("Calling AnimationNodeStateMachine.SetEndNode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeStateMachine", "set_end_node")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false name Vector2}], Returns: void
*/
func (o *AnimationNodeStateMachine) SetGraphOffset(name gdnative.Vector2) {
	//log.Println("Calling AnimationNodeStateMachine.SetGraphOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeStateMachine", "set_graph_offset")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false name String} { false position Vector2}], Returns: void
*/
func (o *AnimationNodeStateMachine) SetNodePosition(name gdnative.String, position gdnative.Vector2) {
	//log.Println("Calling AnimationNodeStateMachine.SetNodePosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(name)
	ptrArguments[1] = gdnative.NewPointerFromVector2(position)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeStateMachine", "set_node_position")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false name String}], Returns: void
*/
func (o *AnimationNodeStateMachine) SetStartNode(name gdnative.String) {
	//log.Println("Calling AnimationNodeStateMachine.SetStartNode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeStateMachine", "set_start_node")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// AnimationNodeStateMachineImplementer is an interface that implements the methods
// of the AnimationNodeStateMachine class.
type AnimationNodeStateMachineImplementer interface {
	AnimationRootNodeImplementer
	X_TreeChanged()
	AddNode(name gdnative.String, node AnimationNodeImplementer, position gdnative.Vector2)
	AddTransition(from gdnative.String, to gdnative.String, transition AnimationNodeStateMachineTransitionImplementer)
	GetEndNode() gdnative.String
	GetGraphOffset() gdnative.Vector2
	GetNode(name gdnative.String) AnimationNodeImplementer
	GetNodeName(node AnimationNodeImplementer) gdnative.String
	GetNodePosition(name gdnative.String) gdnative.Vector2
	GetStartNode() gdnative.String
	GetTransition(idx gdnative.Int) AnimationNodeStateMachineTransitionImplementer
	GetTransitionCount() gdnative.Int
	GetTransitionFrom(idx gdnative.Int) gdnative.String
	GetTransitionTo(idx gdnative.Int) gdnative.String
	HasNode(name gdnative.String) gdnative.Bool
	HasTransition(from gdnative.String, to gdnative.String) gdnative.Bool
	RemoveNode(name gdnative.String)
	RemoveTransition(from gdnative.String, to gdnative.String)
	RemoveTransitionByIndex(idx gdnative.Int)
	RenameNode(name gdnative.String, newName gdnative.String)
	SetEndNode(name gdnative.String)
	SetGraphOffset(name gdnative.Vector2)
	SetNodePosition(name gdnative.String, position gdnative.Vector2)
	SetStartNode(name gdnative.String)
}