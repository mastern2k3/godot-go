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

//func NewAnimatedSprite3DFromPointer(ptr gdnative.Pointer) AnimatedSprite3D {
func newAnimatedSprite3DFromPointer(ptr gdnative.Pointer) AnimatedSprite3D {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AnimatedSprite3D{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Animations are created using a [SpriteFrames] resource, which can be configured in the editor via the SpriteFrames panel.
*/
type AnimatedSprite3D struct {
	SpriteBase3D
	owner gdnative.Object
}

func (o *AnimatedSprite3D) BaseClass() string {
	return "AnimatedSprite3D"
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *AnimatedSprite3D) X_IsPlaying() gdnative.Bool {
	//log.Println("Calling AnimatedSprite3D.X_IsPlaying()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite3D", "_is_playing")

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
	Args: [], Returns: void
*/
func (o *AnimatedSprite3D) X_ResChanged() {
	//log.Println("Calling AnimatedSprite3D.X_ResChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite3D", "_res_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false playing bool}], Returns: void
*/
func (o *AnimatedSprite3D) X_SetPlaying(playing gdnative.Bool) {
	//log.Println("Calling AnimatedSprite3D.X_SetPlaying()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(playing)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite3D", "_set_playing")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *AnimatedSprite3D) GetAnimation() gdnative.String {
	//log.Println("Calling AnimatedSprite3D.GetAnimation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite3D", "get_animation")

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
	Args: [], Returns: int
*/
func (o *AnimatedSprite3D) GetFrame() gdnative.Int {
	//log.Println("Calling AnimatedSprite3D.GetFrame()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite3D", "get_frame")

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
	Args: [], Returns: SpriteFrames
*/
func (o *AnimatedSprite3D) GetSpriteFrames() SpriteFramesImplementer {
	//log.Println("Calling AnimatedSprite3D.GetSpriteFrames()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite3D", "get_sprite_frames")

	// Call the parent method.
	// SpriteFrames
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newSpriteFramesFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(SpriteFramesImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "SpriteFrames" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(SpriteFramesImplementer)
	}

	return &ret
}

/*
        Return true if an animation if currently being played.
	Args: [], Returns: bool
*/
func (o *AnimatedSprite3D) IsPlaying() gdnative.Bool {
	//log.Println("Calling AnimatedSprite3D.IsPlaying()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite3D", "is_playing")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Play the animation set in parameter. If no parameter is provided, the current animation is played.
	Args: [{ true anim String}], Returns: void
*/
func (o *AnimatedSprite3D) Play(anim gdnative.String) {
	//log.Println("Calling AnimatedSprite3D.Play()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(anim)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite3D", "play")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false animation String}], Returns: void
*/
func (o *AnimatedSprite3D) SetAnimation(animation gdnative.String) {
	//log.Println("Calling AnimatedSprite3D.SetAnimation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(animation)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite3D", "set_animation")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false frame int}], Returns: void
*/
func (o *AnimatedSprite3D) SetFrame(frame gdnative.Int) {
	//log.Println("Calling AnimatedSprite3D.SetFrame()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(frame)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite3D", "set_frame")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false sprite_frames SpriteFrames}], Returns: void
*/
func (o *AnimatedSprite3D) SetSpriteFrames(spriteFrames SpriteFramesImplementer) {
	//log.Println("Calling AnimatedSprite3D.SetSpriteFrames()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(spriteFrames.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite3D", "set_sprite_frames")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Stop the current animation (does not reset the frame counter).
	Args: [], Returns: void
*/
func (o *AnimatedSprite3D) Stop() {
	//log.Println("Calling AnimatedSprite3D.Stop()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite3D", "stop")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// AnimatedSprite3DImplementer is an interface that implements the methods
// of the AnimatedSprite3D class.
type AnimatedSprite3DImplementer interface {
	SpriteBase3DImplementer
	X_IsPlaying() gdnative.Bool
	X_ResChanged()
	X_SetPlaying(playing gdnative.Bool)
	GetAnimation() gdnative.String
	GetFrame() gdnative.Int
	GetSpriteFrames() SpriteFramesImplementer
	IsPlaying() gdnative.Bool
	Play(anim gdnative.String)
	SetAnimation(animation gdnative.String)
	SetFrame(frame gdnative.Int)
	SetSpriteFrames(spriteFrames SpriteFramesImplementer)
	Stop()
}
