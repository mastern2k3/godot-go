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

//func NewAnimationNodeStateMachinePlaybackFromPointer(ptr gdnative.Pointer) AnimationNodeStateMachinePlayback {
func newAnimationNodeStateMachinePlaybackFromPointer(ptr gdnative.Pointer) AnimationNodeStateMachinePlayback {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AnimationNodeStateMachinePlayback{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type AnimationNodeStateMachinePlayback struct {
	Resource
	owner gdnative.Object
}

func (o *AnimationNodeStateMachinePlayback) BaseClass() string {
	return "AnimationNodeStateMachinePlayback"
}

/*

	Args: [], Returns: String
*/
func (o *AnimationNodeStateMachinePlayback) GetCurrentNode() gdnative.String {
	//log.Println("Calling AnimationNodeStateMachinePlayback.GetCurrentNode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeStateMachinePlayback", "get_current_node")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*

	Args: [], Returns: PoolStringArray
*/
func (o *AnimationNodeStateMachinePlayback) GetTravelPath() gdnative.PoolStringArray {
	//log.Println("Calling AnimationNodeStateMachinePlayback.GetTravelPath()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeStateMachinePlayback", "get_travel_path")

	// Call the parent method.
	// PoolStringArray
	retPtr := gdnative.NewEmptyPoolStringArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolStringArrayFromPointer(retPtr)
	return ret
}

/*

	Args: [], Returns: bool
*/
func (o *AnimationNodeStateMachinePlayback) IsPlaying() gdnative.Bool {
	//log.Println("Calling AnimationNodeStateMachinePlayback.IsPlaying()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeStateMachinePlayback", "is_playing")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false node String}], Returns: void
*/
func (o *AnimationNodeStateMachinePlayback) Start(node gdnative.String) {
	//log.Println("Calling AnimationNodeStateMachinePlayback.Start()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(node)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeStateMachinePlayback", "start")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [], Returns: void
*/
func (o *AnimationNodeStateMachinePlayback) Stop() {
	//log.Println("Calling AnimationNodeStateMachinePlayback.Stop()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeStateMachinePlayback", "stop")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false to_node String}], Returns: void
*/
func (o *AnimationNodeStateMachinePlayback) Travel(toNode gdnative.String) {
	//log.Println("Calling AnimationNodeStateMachinePlayback.Travel()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(toNode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeStateMachinePlayback", "travel")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// AnimationNodeStateMachinePlaybackImplementer is an interface that implements the methods
// of the AnimationNodeStateMachinePlayback class.
type AnimationNodeStateMachinePlaybackImplementer interface {
	ResourceImplementer
	GetCurrentNode() gdnative.String
	GetTravelPath() gdnative.PoolStringArray
	IsPlaying() gdnative.Bool
	Start(node gdnative.String)
	Stop()
	Travel(toNode gdnative.String)
}
