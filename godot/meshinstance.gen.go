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

//func NewMeshInstanceFromPointer(ptr gdnative.Pointer) MeshInstance {
func newMeshInstanceFromPointer(ptr gdnative.Pointer) MeshInstance {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := MeshInstance{}
	obj.SetBaseObject(owner)

	return obj
}

/*
MeshInstance is a node that takes a [Mesh] resource and adds it to the current scenario by creating an instance of it. This is the class most often used to get 3D geometry rendered and can be used to instance a single [Mesh] in many places. This allows to reuse geometry and save on resources. When a [Mesh] has to be instanced more than thousands of times at close proximity, consider using a [MultiMesh] in a [MultiMeshInstance] instead.
*/
type MeshInstance struct {
	GeometryInstance
	owner gdnative.Object
}

func (o *MeshInstance) BaseClass() string {
	return "MeshInstance"
}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *MeshInstance) X_MeshChanged() {
	//log.Println("Calling MeshInstance.X_MeshChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshInstance", "_mesh_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        This helper creates a [StaticBody] child node with a [ConvexPolygonShape] collision shape calculated from the mesh geometry. It's mainly used for testing.
	Args: [], Returns: void
*/
func (o *MeshInstance) CreateConvexCollision() {
	//log.Println("Calling MeshInstance.CreateConvexCollision()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshInstance", "create_convex_collision")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        This helper creates a [MeshInstance] child node with gizmos at every vertex calculated from the mesh geometry. It's mainly used for testing.
	Args: [], Returns: void
*/
func (o *MeshInstance) CreateDebugTangents() {
	//log.Println("Calling MeshInstance.CreateDebugTangents()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshInstance", "create_debug_tangents")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        This helper creates a [StaticBody] child node with a [ConcavePolygonShape] collision shape calculated from the mesh geometry. It's mainly used for testing.
	Args: [], Returns: void
*/
func (o *MeshInstance) CreateTrimeshCollision() {
	//log.Println("Calling MeshInstance.CreateTrimeshCollision()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshInstance", "create_trimesh_collision")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: Mesh
*/
func (o *MeshInstance) GetMesh() MeshImplementer {
	//log.Println("Calling MeshInstance.GetMesh()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshInstance", "get_mesh")

	// Call the parent method.
	// Mesh
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newMeshFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(MeshImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Mesh" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(MeshImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: NodePath
*/
func (o *MeshInstance) GetSkeletonPath() gdnative.NodePath {
	//log.Println("Calling MeshInstance.GetSkeletonPath()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshInstance", "get_skeleton_path")

	// Call the parent method.
	// NodePath
	retPtr := gdnative.NewEmptyNodePath()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewNodePathFromPointer(retPtr)
	return ret
}

/*
        Returns the [Material] for a surface of the [Mesh] resource.
	Args: [{ false surface int}], Returns: Material
*/
func (o *MeshInstance) GetSurfaceMaterial(surface gdnative.Int) MaterialImplementer {
	//log.Println("Calling MeshInstance.GetSurfaceMaterial()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(surface)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshInstance", "get_surface_material")

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

	Args: [], Returns: int
*/
func (o *MeshInstance) GetSurfaceMaterialCount() gdnative.Int {
	//log.Println("Calling MeshInstance.GetSurfaceMaterialCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshInstance", "get_surface_material_count")

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
	Args: [{ false mesh Mesh}], Returns: void
*/
func (o *MeshInstance) SetMesh(mesh MeshImplementer) {
	//log.Println("Calling MeshInstance.SetMesh()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(mesh.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshInstance", "set_mesh")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false skeleton_path NodePath}], Returns: void
*/
func (o *MeshInstance) SetSkeletonPath(skeletonPath gdnative.NodePath) {
	//log.Println("Calling MeshInstance.SetSkeletonPath()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromNodePath(skeletonPath)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshInstance", "set_skeleton_path")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the [Material] for a surface of the [Mesh] resource.
	Args: [{ false surface int} { false material Material}], Returns: void
*/
func (o *MeshInstance) SetSurfaceMaterial(surface gdnative.Int, material MaterialImplementer) {
	//log.Println("Calling MeshInstance.SetSurfaceMaterial()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(surface)
	ptrArguments[1] = gdnative.NewPointerFromObject(material.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshInstance", "set_surface_material")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// MeshInstanceImplementer is an interface that implements the methods
// of the MeshInstance class.
type MeshInstanceImplementer interface {
	GeometryInstanceImplementer
	X_MeshChanged()
	CreateConvexCollision()
	CreateDebugTangents()
	CreateTrimeshCollision()
	GetMesh() MeshImplementer
	GetSkeletonPath() gdnative.NodePath
	GetSurfaceMaterial(surface gdnative.Int) MaterialImplementer
	GetSurfaceMaterialCount() gdnative.Int
	SetMesh(mesh MeshImplementer)
	SetSkeletonPath(skeletonPath gdnative.NodePath)
	SetSurfaceMaterial(surface gdnative.Int, material MaterialImplementer)
}
