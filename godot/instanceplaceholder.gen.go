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

//func NewInstancePlaceholderFromPointer(ptr gdnative.Pointer) InstancePlaceholder {
func newInstancePlaceholderFromPointer(ptr gdnative.Pointer) InstancePlaceholder {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := InstancePlaceholder{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Turning on the option [b]Load As Placeholder[/b] for an instanced scene in the editor causes it to be replaced by an InstancePlaceholder when running the game. This makes it possible to delay actually loading the scene until calling [method replace_by_instance]. This is useful to avoid loading large scenes all at once by loading parts of it selectively. The InstancePlaceholder does not have a transform. This causes any child nodes to be positioned relatively to the Viewport from point (0,0), rather than their parent as displayed in the editor. Replacing the placeholder with a scene with a transform will transform children relatively to their parent again.
*/
type InstancePlaceholder struct {
	Node
	owner gdnative.Object
}

func (o *InstancePlaceholder) BaseClass() string {
	return "InstancePlaceholder"
}

/*

	Args: [{False true replace bool} {Null true custom_scene PackedScene}], Returns: Node
*/
func (o *InstancePlaceholder) CreateInstance(replace gdnative.Bool, customScene PackedSceneImplementer) NodeImplementer {
	//log.Println("Calling InstancePlaceholder.CreateInstance()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromBool(replace)
	ptrArguments[1] = gdnative.NewPointerFromObject(customScene.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InstancePlaceholder", "create_instance")

	// Call the parent method.
	// Node
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newNodeFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(NodeImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Node" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(NodeImplementer)
	}

	return &ret
}

/*
        Retrieve the path to the [PackedScene] resource file that is loaded by default when calling [method replace_by_instance].
	Args: [], Returns: String
*/
func (o *InstancePlaceholder) GetInstancePath() gdnative.String {
	//log.Println("Calling InstancePlaceholder.GetInstancePath()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InstancePlaceholder", "get_instance_path")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*

	Args: [{False true with_order bool}], Returns: Dictionary
*/
func (o *InstancePlaceholder) GetStoredValues(withOrder gdnative.Bool) gdnative.Dictionary {
	//log.Println("Calling InstancePlaceholder.GetStoredValues()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(withOrder)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InstancePlaceholder", "get_stored_values")

	// Call the parent method.
	// Dictionary
	retPtr := gdnative.NewEmptyDictionary()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewDictionaryFromPointer(retPtr)
	return ret
}

/*
        Replace this placeholder by the scene handed as an argument, or the original scene if no argument is given. As for all resources, the scene is loaded only if it's not loaded already. By manually loading the scene beforehand, delays caused by this function can be avoided.
	Args: [{Null true custom_scene PackedScene}], Returns: void
*/
func (o *InstancePlaceholder) ReplaceByInstance(customScene PackedSceneImplementer) {
	//log.Println("Calling InstancePlaceholder.ReplaceByInstance()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(customScene.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InstancePlaceholder", "replace_by_instance")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// InstancePlaceholderImplementer is an interface that implements the methods
// of the InstancePlaceholder class.
type InstancePlaceholderImplementer interface {
	NodeImplementer
	CreateInstance(replace gdnative.Bool, customScene PackedSceneImplementer) NodeImplementer
	GetInstancePath() gdnative.String
	GetStoredValues(withOrder gdnative.Bool) gdnative.Dictionary
	ReplaceByInstance(customScene PackedSceneImplementer)
}
