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

//func NewEditorImportPluginFromPointer(ptr gdnative.Pointer) EditorImportPlugin {
func newEditorImportPluginFromPointer(ptr gdnative.Pointer) EditorImportPlugin {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := EditorImportPlugin{}
	obj.SetBaseObject(owner)

	return obj
}

/*
EditorImportPlugins provide a way to extend the editor's resource import functionality. Use them to import resources from custom files or to provide alternatives to the editor's existing importers. Register your [EditorPlugin] with [method EditorPlugin.add_import_plugin]. EditorImportPlugins work by associating with specific file extensions and a resource type. See [method get_recognized_extensions] and [method get_resource_type]). They may optionally specify some import presets that affect the import process. EditorImportPlugins are responsible for creating the resources and saving them in the [code].import[/code] directory. Below is an example EditorImportPlugin that imports a [Mesh] from a file with the extension ".special" or ".spec": [codeblock] tool extends EditorImportPlugin func get_importer_name(): return "my.special.plugin" func get_visible_name(): return "Special Mesh Importer" func get_recognized_extensions(): return ["special", "spec"] func get_save_extension(): return "mesh" func get_resource_type(): return "Mesh" func get_preset_count(): return 1 func get_preset_name(i): return "Default" func get_import_options(i): return [{"name": "my_option", "default_value": false}] func import(source_file, save_path, options, platform_variants, gen_files): var file = File.new() if file.open(source_file, File.READ) != OK: return FAILED var mesh = Mesh.new() # Fill the Mesh with data read in 'file', left as exercise to the reader var filename = save_path + "." + get_save_extension() ResourceSaver.save(filename, mesh) return OK [/codeblock]
*/
type EditorImportPlugin struct {
	Reference
	owner gdnative.Object
}

func (o *EditorImportPlugin) BaseClass() string {
	return "EditorImportPlugin"
}

/*
        Get the options and default values for the preset at this index. Returns an Array of Dictionaries with the following keys: [code]name[/code], [code]default_value[/code], [code]property_hint[/code] (optional), [code]hint_string[/code] (optional), [code]usage[/code] (optional).
	Args: [{ false preset int}], Returns: Array
*/
func (o *EditorImportPlugin) GetImportOptions(preset gdnative.Int) gdnative.Array {
	//log.Println("Calling EditorImportPlugin.GetImportOptions()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(preset)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorImportPlugin", "get_import_options")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*
        Get the order of this importer to be run when importing resources. Higher values will be called later. Use this to ensure the importer runs after the dependencies are already imported.
	Args: [], Returns: int
*/
func (o *EditorImportPlugin) GetImportOrder() gdnative.Int {
	//log.Println("Calling EditorImportPlugin.GetImportOrder()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorImportPlugin", "get_import_order")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Get the unique name of the importer.
	Args: [], Returns: String
*/
func (o *EditorImportPlugin) GetImporterName() gdnative.String {
	//log.Println("Calling EditorImportPlugin.GetImporterName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorImportPlugin", "get_importer_name")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false option String} { false options Dictionary}], Returns: bool
*/
func (o *EditorImportPlugin) GetOptionVisibility(option gdnative.String, options gdnative.Dictionary) gdnative.Bool {
	//log.Println("Calling EditorImportPlugin.GetOptionVisibility()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(option)
	ptrArguments[1] = gdnative.NewPointerFromDictionary(options)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorImportPlugin", "get_option_visibility")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Get the number of initial presets defined by the plugin. Use [method get_import_options] to get the default options for the preset and [method get_preset_name] to get the name of the preset.
	Args: [], Returns: int
*/
func (o *EditorImportPlugin) GetPresetCount() gdnative.Int {
	//log.Println("Calling EditorImportPlugin.GetPresetCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorImportPlugin", "get_preset_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Get the name of the options preset at this index.
	Args: [{ false preset int}], Returns: String
*/
func (o *EditorImportPlugin) GetPresetName(preset gdnative.Int) gdnative.String {
	//log.Println("Calling EditorImportPlugin.GetPresetName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(preset)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorImportPlugin", "get_preset_name")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Get the priority of this plugin for the recognized extension. Higher priority plugins will be preferred. Default value is 1.0.
	Args: [], Returns: float
*/
func (o *EditorImportPlugin) GetPriority() gdnative.Real {
	//log.Println("Calling EditorImportPlugin.GetPriority()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorImportPlugin", "get_priority")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Get the list of file extensions to associate with this loader (case insensitive). e.g. [code]["obj"][/code].
	Args: [], Returns: Array
*/
func (o *EditorImportPlugin) GetRecognizedExtensions() gdnative.Array {
	//log.Println("Calling EditorImportPlugin.GetRecognizedExtensions()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorImportPlugin", "get_recognized_extensions")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*
        Get the Godot resource type associated with this loader. e.g. [code]"Mesh"[/code] or [code]"Animation"[/code].
	Args: [], Returns: String
*/
func (o *EditorImportPlugin) GetResourceType() gdnative.String {
	//log.Println("Calling EditorImportPlugin.GetResourceType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorImportPlugin", "get_resource_type")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Get the extension used to save this resource in the [code].import[/code] directory.
	Args: [], Returns: String
*/
func (o *EditorImportPlugin) GetSaveExtension() gdnative.String {
	//log.Println("Calling EditorImportPlugin.GetSaveExtension()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorImportPlugin", "get_save_extension")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Get the name to display in the import window.
	Args: [], Returns: String
*/
func (o *EditorImportPlugin) GetVisibleName() gdnative.String {
	//log.Println("Calling EditorImportPlugin.GetVisibleName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorImportPlugin", "get_visible_name")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false source_file String} { false save_path String} { false options Dictionary} { false platform_variants Array} { false gen_files Array}], Returns: int
*/
func (o *EditorImportPlugin) Import(sourceFile gdnative.String, savePath gdnative.String, options gdnative.Dictionary, platformVariants gdnative.Array, genFiles gdnative.Array) gdnative.Int {
	//log.Println("Calling EditorImportPlugin.Import()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 5, 5)
	ptrArguments[0] = gdnative.NewPointerFromString(sourceFile)
	ptrArguments[1] = gdnative.NewPointerFromString(savePath)
	ptrArguments[2] = gdnative.NewPointerFromDictionary(options)
	ptrArguments[3] = gdnative.NewPointerFromArray(platformVariants)
	ptrArguments[4] = gdnative.NewPointerFromArray(genFiles)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorImportPlugin", "import")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

// EditorImportPluginImplementer is an interface that implements the methods
// of the EditorImportPlugin class.
type EditorImportPluginImplementer interface {
	ReferenceImplementer
	GetImportOptions(preset gdnative.Int) gdnative.Array
	GetImportOrder() gdnative.Int
	GetImporterName() gdnative.String
	GetOptionVisibility(option gdnative.String, options gdnative.Dictionary) gdnative.Bool
	GetPresetCount() gdnative.Int
	GetPresetName(preset gdnative.Int) gdnative.String
	GetPriority() gdnative.Real
	GetRecognizedExtensions() gdnative.Array
	GetResourceType() gdnative.String
	GetSaveExtension() gdnative.String
	GetVisibleName() gdnative.String
	Import(sourceFile gdnative.String, savePath gdnative.String, options gdnative.Dictionary, platformVariants gdnative.Array, genFiles gdnative.Array) gdnative.Int
}
