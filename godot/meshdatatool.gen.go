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

//func NewMeshDataToolFromPointer(ptr gdnative.Pointer) MeshDataTool {
func newMeshDataToolFromPointer(ptr gdnative.Pointer) MeshDataTool {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := MeshDataTool{}
	obj.SetBaseObject(owner)

	return obj
}

/*
The MeshDataTool provides access to individual vertices in a [Mesh]. It allows users to read and edit vertex data of meshes. It also creates an array of faces and edges. To use the MeshDataTool, load a mesh with [method create_from_surface]. When you are finished editing the data commit the data to a mesh with [method commit_to_surface]. Below is an example of how the MeshDataTool may be used. [codeblock] var mdt = MeshDataTool.new() mdt.create_from_surface(mesh, 0) for i in range(mdt.get_vertex_count()): var vertex = mdt.get_vertex(i) ... mdt.set_vertex(i, vertex) mesh.surface_remove(0) mdt.commit_to_surface(mesh) [/codeblock]
*/
type MeshDataTool struct {
	Reference
	owner gdnative.Object
}

func (o *MeshDataTool) BaseClass() string {
	return "MeshDataTool"
}

/*
        Clears all data currently in MeshDataTool.
	Args: [], Returns: void
*/
func (o *MeshDataTool) Clear() {
	//log.Println("Calling MeshDataTool.Clear()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshDataTool", "clear")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds a new surface to specified [Mesh] with edited data.
	Args: [{ false mesh ArrayMesh}], Returns: enum.Error
*/
func (o *MeshDataTool) CommitToSurface(mesh ArrayMeshImplementer) gdnative.Error {
	//log.Println("Calling MeshDataTool.CommitToSurface()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(mesh.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshDataTool", "commit_to_surface")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Uses specified surface of given [Mesh] to populate data for MeshDataTool. Requires [Mesh] with primitive type [code]PRIMITIVE_TRIANGLES[/code].
	Args: [{ false mesh ArrayMesh} { false surface int}], Returns: enum.Error
*/
func (o *MeshDataTool) CreateFromSurface(mesh ArrayMeshImplementer, surface gdnative.Int) gdnative.Error {
	//log.Println("Calling MeshDataTool.CreateFromSurface()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromObject(mesh.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromInt(surface)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshDataTool", "create_from_surface")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Returns the number of edges in this [Mesh].
	Args: [], Returns: int
*/
func (o *MeshDataTool) GetEdgeCount() gdnative.Int {
	//log.Println("Calling MeshDataTool.GetEdgeCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshDataTool", "get_edge_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns array of faces that touch given edge.
	Args: [{ false idx int}], Returns: PoolIntArray
*/
func (o *MeshDataTool) GetEdgeFaces(idx gdnative.Int) gdnative.PoolIntArray {
	//log.Println("Calling MeshDataTool.GetEdgeFaces()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshDataTool", "get_edge_faces")

	// Call the parent method.
	// PoolIntArray
	retPtr := gdnative.NewEmptyPoolIntArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolIntArrayFromPointer(retPtr)
	return ret
}

/*
        Returns meta information assigned to given edge.
	Args: [{ false idx int}], Returns: Variant
*/
func (o *MeshDataTool) GetEdgeMeta(idx gdnative.Int) gdnative.Variant {
	//log.Println("Calling MeshDataTool.GetEdgeMeta()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshDataTool", "get_edge_meta")

	// Call the parent method.
	// Variant
	retPtr := gdnative.NewEmptyVariant()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVariantFromPointer(retPtr)
	return ret
}

/*
        Returns index of specified vertex connected to given edge. Vertex argument can only be 0 or 1 because edges are comprised of two vertices.
	Args: [{ false idx int} { false vertex int}], Returns: int
*/
func (o *MeshDataTool) GetEdgeVertex(idx gdnative.Int, vertex gdnative.Int) gdnative.Int {
	//log.Println("Calling MeshDataTool.GetEdgeVertex()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromInt(vertex)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshDataTool", "get_edge_vertex")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the number of faces in this [Mesh].
	Args: [], Returns: int
*/
func (o *MeshDataTool) GetFaceCount() gdnative.Int {
	//log.Println("Calling MeshDataTool.GetFaceCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshDataTool", "get_face_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns specified edge associated with given face. Edge argument must 2 or less because a face only has three edges.
	Args: [{ false idx int} { false edge int}], Returns: int
*/
func (o *MeshDataTool) GetFaceEdge(idx gdnative.Int, edge gdnative.Int) gdnative.Int {
	//log.Println("Calling MeshDataTool.GetFaceEdge()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromInt(edge)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshDataTool", "get_face_edge")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns meta data associated with given face.
	Args: [{ false idx int}], Returns: Variant
*/
func (o *MeshDataTool) GetFaceMeta(idx gdnative.Int) gdnative.Variant {
	//log.Println("Calling MeshDataTool.GetFaceMeta()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshDataTool", "get_face_meta")

	// Call the parent method.
	// Variant
	retPtr := gdnative.NewEmptyVariant()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVariantFromPointer(retPtr)
	return ret
}

/*
        Calculates and returns face normal of given face.
	Args: [{ false idx int}], Returns: Vector3
*/
func (o *MeshDataTool) GetFaceNormal(idx gdnative.Int) gdnative.Vector3 {
	//log.Println("Calling MeshDataTool.GetFaceNormal()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshDataTool", "get_face_normal")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Returns specified vertex of given face. Vertex argument must be 2 or less because faces contain three vertices.
	Args: [{ false idx int} { false vertex int}], Returns: int
*/
func (o *MeshDataTool) GetFaceVertex(idx gdnative.Int, vertex gdnative.Int) gdnative.Int {
	//log.Println("Calling MeshDataTool.GetFaceVertex()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromInt(vertex)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshDataTool", "get_face_vertex")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns format of [Mesh]. Format is an integer made up of [Mesh] format flags combined together. For example, a mesh containing both vertices and normals would return a format of [code]3[/code] because [code]ARRAY_FORMAT_VERTEX[/code] is [code]1[/code] and [code]ARRAY_FORMAT_NORMAL[/code] is [code]2[/code]. For list of format flags see [ArrayMesh].
	Args: [], Returns: int
*/
func (o *MeshDataTool) GetFormat() gdnative.Int {
	//log.Println("Calling MeshDataTool.GetFormat()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshDataTool", "get_format")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns material assigned to the [Mesh].
	Args: [], Returns: Material
*/
func (o *MeshDataTool) GetMaterial() MaterialImplementer {
	//log.Println("Calling MeshDataTool.GetMaterial()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshDataTool", "get_material")

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
        Returns the vertex at given index.
	Args: [{ false idx int}], Returns: Vector3
*/
func (o *MeshDataTool) GetVertex(idx gdnative.Int) gdnative.Vector3 {
	//log.Println("Calling MeshDataTool.GetVertex()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshDataTool", "get_vertex")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Returns the bones of the given vertex.
	Args: [{ false idx int}], Returns: PoolIntArray
*/
func (o *MeshDataTool) GetVertexBones(idx gdnative.Int) gdnative.PoolIntArray {
	//log.Println("Calling MeshDataTool.GetVertexBones()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshDataTool", "get_vertex_bones")

	// Call the parent method.
	// PoolIntArray
	retPtr := gdnative.NewEmptyPoolIntArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolIntArrayFromPointer(retPtr)
	return ret
}

/*
        Returns the color of the given vertex.
	Args: [{ false idx int}], Returns: Color
*/
func (o *MeshDataTool) GetVertexColor(idx gdnative.Int) gdnative.Color {
	//log.Println("Calling MeshDataTool.GetVertexColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshDataTool", "get_vertex_color")

	// Call the parent method.
	// Color
	retPtr := gdnative.NewEmptyColor()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewColorFromPointer(retPtr)
	return ret
}

/*
        Returns the total number of vertices in [Mesh].
	Args: [], Returns: int
*/
func (o *MeshDataTool) GetVertexCount() gdnative.Int {
	//log.Println("Calling MeshDataTool.GetVertexCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshDataTool", "get_vertex_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns array of edges that share given vertex.
	Args: [{ false idx int}], Returns: PoolIntArray
*/
func (o *MeshDataTool) GetVertexEdges(idx gdnative.Int) gdnative.PoolIntArray {
	//log.Println("Calling MeshDataTool.GetVertexEdges()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshDataTool", "get_vertex_edges")

	// Call the parent method.
	// PoolIntArray
	retPtr := gdnative.NewEmptyPoolIntArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolIntArrayFromPointer(retPtr)
	return ret
}

/*
        Returns array of faces that share given vertex.
	Args: [{ false idx int}], Returns: PoolIntArray
*/
func (o *MeshDataTool) GetVertexFaces(idx gdnative.Int) gdnative.PoolIntArray {
	//log.Println("Calling MeshDataTool.GetVertexFaces()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshDataTool", "get_vertex_faces")

	// Call the parent method.
	// PoolIntArray
	retPtr := gdnative.NewEmptyPoolIntArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolIntArrayFromPointer(retPtr)
	return ret
}

/*
        Returns meta data associated with given vertex.
	Args: [{ false idx int}], Returns: Variant
*/
func (o *MeshDataTool) GetVertexMeta(idx gdnative.Int) gdnative.Variant {
	//log.Println("Calling MeshDataTool.GetVertexMeta()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshDataTool", "get_vertex_meta")

	// Call the parent method.
	// Variant
	retPtr := gdnative.NewEmptyVariant()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVariantFromPointer(retPtr)
	return ret
}

/*
        Returns normal of given vertex.
	Args: [{ false idx int}], Returns: Vector3
*/
func (o *MeshDataTool) GetVertexNormal(idx gdnative.Int) gdnative.Vector3 {
	//log.Println("Calling MeshDataTool.GetVertexNormal()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshDataTool", "get_vertex_normal")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Returns tangent of given vertex.
	Args: [{ false idx int}], Returns: Plane
*/
func (o *MeshDataTool) GetVertexTangent(idx gdnative.Int) gdnative.Plane {
	//log.Println("Calling MeshDataTool.GetVertexTangent()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshDataTool", "get_vertex_tangent")

	// Call the parent method.
	// Plane
	retPtr := gdnative.NewEmptyPlane()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPlaneFromPointer(retPtr)
	return ret
}

/*
        Returns UV of given vertex.
	Args: [{ false idx int}], Returns: Vector2
*/
func (o *MeshDataTool) GetVertexUv(idx gdnative.Int) gdnative.Vector2 {
	//log.Println("Calling MeshDataTool.GetVertexUv()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshDataTool", "get_vertex_uv")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*
        Returns UV2 of given vertex.
	Args: [{ false idx int}], Returns: Vector2
*/
func (o *MeshDataTool) GetVertexUv2(idx gdnative.Int) gdnative.Vector2 {
	//log.Println("Calling MeshDataTool.GetVertexUv2()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshDataTool", "get_vertex_uv2")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*
        Returns bone weights of given vertex.
	Args: [{ false idx int}], Returns: PoolRealArray
*/
func (o *MeshDataTool) GetVertexWeights(idx gdnative.Int) gdnative.PoolRealArray {
	//log.Println("Calling MeshDataTool.GetVertexWeights()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshDataTool", "get_vertex_weights")

	// Call the parent method.
	// PoolRealArray
	retPtr := gdnative.NewEmptyPoolRealArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolRealArrayFromPointer(retPtr)
	return ret
}

/*
        Sets the meta data of given edge.
	Args: [{ false idx int} { false meta Variant}], Returns: void
*/
func (o *MeshDataTool) SetEdgeMeta(idx gdnative.Int, meta gdnative.Variant) {
	//log.Println("Calling MeshDataTool.SetEdgeMeta()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromVariant(meta)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshDataTool", "set_edge_meta")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the meta data of given face.
	Args: [{ false idx int} { false meta Variant}], Returns: void
*/
func (o *MeshDataTool) SetFaceMeta(idx gdnative.Int, meta gdnative.Variant) {
	//log.Println("Calling MeshDataTool.SetFaceMeta()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromVariant(meta)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshDataTool", "set_face_meta")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the material to be used by newly constructed [Mesh].
	Args: [{ false material Material}], Returns: void
*/
func (o *MeshDataTool) SetMaterial(material MaterialImplementer) {
	//log.Println("Calling MeshDataTool.SetMaterial()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(material.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshDataTool", "set_material")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the position of given vertex.
	Args: [{ false idx int} { false vertex Vector3}], Returns: void
*/
func (o *MeshDataTool) SetVertex(idx gdnative.Int, vertex gdnative.Vector3) {
	//log.Println("Calling MeshDataTool.SetVertex()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromVector3(vertex)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshDataTool", "set_vertex")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the bones of given vertex.
	Args: [{ false idx int} { false bones PoolIntArray}], Returns: void
*/
func (o *MeshDataTool) SetVertexBones(idx gdnative.Int, bones gdnative.PoolIntArray) {
	//log.Println("Calling MeshDataTool.SetVertexBones()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromPoolIntArray(bones)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshDataTool", "set_vertex_bones")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the color of given vertex.
	Args: [{ false idx int} { false color Color}], Returns: void
*/
func (o *MeshDataTool) SetVertexColor(idx gdnative.Int, color gdnative.Color) {
	//log.Println("Calling MeshDataTool.SetVertexColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromColor(color)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshDataTool", "set_vertex_color")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the meta data associated with given vertex.
	Args: [{ false idx int} { false meta Variant}], Returns: void
*/
func (o *MeshDataTool) SetVertexMeta(idx gdnative.Int, meta gdnative.Variant) {
	//log.Println("Calling MeshDataTool.SetVertexMeta()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromVariant(meta)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshDataTool", "set_vertex_meta")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the normal of given vertex.
	Args: [{ false idx int} { false normal Vector3}], Returns: void
*/
func (o *MeshDataTool) SetVertexNormal(idx gdnative.Int, normal gdnative.Vector3) {
	//log.Println("Calling MeshDataTool.SetVertexNormal()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromVector3(normal)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshDataTool", "set_vertex_normal")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the tangent of given vertex.
	Args: [{ false idx int} { false tangent Plane}], Returns: void
*/
func (o *MeshDataTool) SetVertexTangent(idx gdnative.Int, tangent gdnative.Plane) {
	//log.Println("Calling MeshDataTool.SetVertexTangent()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromPlane(tangent)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshDataTool", "set_vertex_tangent")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the UV of given vertex.
	Args: [{ false idx int} { false uv Vector2}], Returns: void
*/
func (o *MeshDataTool) SetVertexUv(idx gdnative.Int, uv gdnative.Vector2) {
	//log.Println("Calling MeshDataTool.SetVertexUv()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromVector2(uv)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshDataTool", "set_vertex_uv")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the UV2 of given vertex.
	Args: [{ false idx int} { false uv2 Vector2}], Returns: void
*/
func (o *MeshDataTool) SetVertexUv2(idx gdnative.Int, uv2 gdnative.Vector2) {
	//log.Println("Calling MeshDataTool.SetVertexUv2()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromVector2(uv2)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshDataTool", "set_vertex_uv2")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the bone weights of given vertex.
	Args: [{ false idx int} { false weights PoolRealArray}], Returns: void
*/
func (o *MeshDataTool) SetVertexWeights(idx gdnative.Int, weights gdnative.PoolRealArray) {
	//log.Println("Calling MeshDataTool.SetVertexWeights()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromPoolRealArray(weights)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MeshDataTool", "set_vertex_weights")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// MeshDataToolImplementer is an interface that implements the methods
// of the MeshDataTool class.
type MeshDataToolImplementer interface {
	ReferenceImplementer
	Clear()
	GetEdgeCount() gdnative.Int
	GetEdgeFaces(idx gdnative.Int) gdnative.PoolIntArray
	GetEdgeMeta(idx gdnative.Int) gdnative.Variant
	GetEdgeVertex(idx gdnative.Int, vertex gdnative.Int) gdnative.Int
	GetFaceCount() gdnative.Int
	GetFaceEdge(idx gdnative.Int, edge gdnative.Int) gdnative.Int
	GetFaceMeta(idx gdnative.Int) gdnative.Variant
	GetFaceNormal(idx gdnative.Int) gdnative.Vector3
	GetFaceVertex(idx gdnative.Int, vertex gdnative.Int) gdnative.Int
	GetFormat() gdnative.Int
	GetMaterial() MaterialImplementer
	GetVertex(idx gdnative.Int) gdnative.Vector3
	GetVertexBones(idx gdnative.Int) gdnative.PoolIntArray
	GetVertexColor(idx gdnative.Int) gdnative.Color
	GetVertexCount() gdnative.Int
	GetVertexEdges(idx gdnative.Int) gdnative.PoolIntArray
	GetVertexFaces(idx gdnative.Int) gdnative.PoolIntArray
	GetVertexMeta(idx gdnative.Int) gdnative.Variant
	GetVertexNormal(idx gdnative.Int) gdnative.Vector3
	GetVertexTangent(idx gdnative.Int) gdnative.Plane
	GetVertexUv(idx gdnative.Int) gdnative.Vector2
	GetVertexUv2(idx gdnative.Int) gdnative.Vector2
	GetVertexWeights(idx gdnative.Int) gdnative.PoolRealArray
	SetEdgeMeta(idx gdnative.Int, meta gdnative.Variant)
	SetFaceMeta(idx gdnative.Int, meta gdnative.Variant)
	SetMaterial(material MaterialImplementer)
	SetVertex(idx gdnative.Int, vertex gdnative.Vector3)
	SetVertexBones(idx gdnative.Int, bones gdnative.PoolIntArray)
	SetVertexColor(idx gdnative.Int, color gdnative.Color)
	SetVertexMeta(idx gdnative.Int, meta gdnative.Variant)
	SetVertexNormal(idx gdnative.Int, normal gdnative.Vector3)
	SetVertexTangent(idx gdnative.Int, tangent gdnative.Plane)
	SetVertexUv(idx gdnative.Int, uv gdnative.Vector2)
	SetVertexUv2(idx gdnative.Int, uv2 gdnative.Vector2)
	SetVertexWeights(idx gdnative.Int, weights gdnative.PoolRealArray)
}
