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

// TreeDropModeFlags is an enum for DropModeFlags values.
type TreeDropModeFlags int

const (
	TreeDropModeDisabled  TreeDropModeFlags = 0
	TreeDropModeInbetween TreeDropModeFlags = 2
	TreeDropModeOnItem    TreeDropModeFlags = 1
)

// TreeSelectMode is an enum for SelectMode values.
type TreeSelectMode int

const (
	TreeSelectMulti  TreeSelectMode = 2
	TreeSelectRow    TreeSelectMode = 1
	TreeSelectSingle TreeSelectMode = 0
)

//func NewTreeFromPointer(ptr gdnative.Pointer) Tree {
func newTreeFromPointer(ptr gdnative.Pointer) Tree {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Tree{}
	obj.SetBaseObject(owner)

	return obj
}

/*
This shows a tree of items that can be selected, expanded and collapsed. The tree can have multiple columns with custom controls like text editing, buttons and popups. It can be useful for structured displays and interactions. Trees are built via code, using [TreeItem] objects to create the structure. They have a single root but multiple roots can be simulated if a dummy hidden root is added. [codeblock] func _ready(): var tree = Tree.new() var root = tree.create_item() tree.set_hide_root(true) var child1 = tree.create_item(root) var child2 = tree.create_item(root) var subchild1 = tree.create_item(child1) subchild1.set_text(0, "Subchild1") [/codeblock]
*/
type Tree struct {
	Control
	owner gdnative.Object
}

func (o *Tree) BaseClass() string {
	return "Tree"
}

/*
        Undocumented
	Args: [{ false arg0 InputEvent}], Returns: void
*/
func (o *Tree) X_GuiInput(arg0 InputEventImplementer) {
	//log.Println("Calling Tree.X_GuiInput()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(arg0.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "_gui_input")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 int}], Returns: void
*/
func (o *Tree) X_PopupSelect(arg0 gdnative.Int) {
	//log.Println("Calling Tree.X_PopupSelect()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "_popup_select")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *Tree) X_RangeClickTimeout() {
	//log.Println("Calling Tree.X_RangeClickTimeout()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "_range_click_timeout")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 float}], Returns: void
*/
func (o *Tree) X_ScrollMoved(arg0 gdnative.Real) {
	//log.Println("Calling Tree.X_ScrollMoved()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "_scroll_moved")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 String}], Returns: void
*/
func (o *Tree) X_TextEditorEnter(arg0 gdnative.String) {
	//log.Println("Calling Tree.X_TextEditorEnter()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "_text_editor_enter")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *Tree) X_TextEditorModalClose() {
	//log.Println("Calling Tree.X_TextEditorModalClose()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "_text_editor_modal_close")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 float}], Returns: void
*/
func (o *Tree) X_ValueEditorChanged(arg0 gdnative.Real) {
	//log.Println("Calling Tree.X_ValueEditorChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "_value_editor_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns [code]true[/code] if the column titles are being shown.
	Args: [], Returns: bool
*/
func (o *Tree) AreColumnTitlesVisible() gdnative.Bool {
	//log.Println("Calling Tree.AreColumnTitlesVisible()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "are_column_titles_visible")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Clears the tree. This removes all items.
	Args: [], Returns: void
*/
func (o *Tree) Clear() {
	//log.Println("Calling Tree.Clear()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "clear")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Create an item in the tree and add it as the last child of [code]parent[/code]. If parent is not given, it will be added as the root's last child, or it'll the be the root itself if the tree is empty.
	Args: [{Null true parent Object} {-1 true idx int}], Returns: TreeItem
*/
func (o *Tree) CreateItem(parent ObjectImplementer, idx gdnative.Int) TreeItemImplementer {
	//log.Println("Calling Tree.CreateItem()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromObject(parent.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "create_item")

	// Call the parent method.
	// TreeItem
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newTreeItemFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(TreeItemImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "TreeItem" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(TreeItemImplementer)
	}

	return &ret
}

/*
        Makes the currently selected item visible. This will scroll the tree to make sure the selected item is visible.
	Args: [], Returns: void
*/
func (o *Tree) EnsureCursorIsVisible() {
	//log.Println("Calling Tree.EnsureCursorIsVisible()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "ensure_cursor_is_visible")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *Tree) GetAllowReselect() gdnative.Bool {
	//log.Println("Calling Tree.GetAllowReselect()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "get_allow_reselect")

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
	Args: [], Returns: bool
*/
func (o *Tree) GetAllowRmbSelect() gdnative.Bool {
	//log.Println("Calling Tree.GetAllowRmbSelect()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "get_allow_rmb_select")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns the column index under the given point.
	Args: [{ false position Vector2}], Returns: int
*/
func (o *Tree) GetColumnAtPosition(position gdnative.Vector2) gdnative.Int {
	//log.Println("Calling Tree.GetColumnAtPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(position)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "get_column_at_position")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the column's title.
	Args: [{ false column int}], Returns: String
*/
func (o *Tree) GetColumnTitle(column gdnative.Int) gdnative.String {
	//log.Println("Calling Tree.GetColumnTitle()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(column)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "get_column_title")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Returns the column's width in pixels.
	Args: [{ false column int}], Returns: int
*/
func (o *Tree) GetColumnWidth(column gdnative.Int) gdnative.Int {
	//log.Println("Calling Tree.GetColumnWidth()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(column)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "get_column_width")

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
func (o *Tree) GetColumns() gdnative.Int {
	//log.Println("Calling Tree.GetColumns()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "get_columns")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the rectangle for custom popups. Helper to create custom cell controls that display a popup. See [method TreeItem.set_cell_mode].
	Args: [], Returns: Rect2
*/
func (o *Tree) GetCustomPopupRect() gdnative.Rect2 {
	//log.Println("Calling Tree.GetCustomPopupRect()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "get_custom_popup_rect")

	// Call the parent method.
	// Rect2
	retPtr := gdnative.NewEmptyRect2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRect2FromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *Tree) GetDropModeFlags() gdnative.Int {
	//log.Println("Calling Tree.GetDropModeFlags()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "get_drop_mode_flags")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        If [member drop_mode_flags] includes [code]DROP_MODE_INBETWEEN[/code], returns -1 if [code]position[/code] is the upper part of a tree item at that position, 1 for the lower part, and additionally 0 for the middle part if [member drop_mode_flags] includes [code]DROP_MODE_ON_ITEM[/code]. Otherwise, returns 0. If there are no tree item at [code]position[/code], returns -100.
	Args: [{ false position Vector2}], Returns: int
*/
func (o *Tree) GetDropSectionAtPosition(position gdnative.Vector2) gdnative.Int {
	//log.Println("Calling Tree.GetDropSectionAtPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(position)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "get_drop_section_at_position")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the currently edited item. This is only available for custom cell mode.
	Args: [], Returns: TreeItem
*/
func (o *Tree) GetEdited() TreeItemImplementer {
	//log.Println("Calling Tree.GetEdited()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "get_edited")

	// Call the parent method.
	// TreeItem
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newTreeItemFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(TreeItemImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "TreeItem" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(TreeItemImplementer)
	}

	return &ret
}

/*
        Returns the column for the currently edited item. This is only available for custom cell mode.
	Args: [], Returns: int
*/
func (o *Tree) GetEditedColumn() gdnative.Int {
	//log.Println("Calling Tree.GetEditedColumn()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "get_edited_column")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the rectangle area for the specified item. If column is specified, only get the position and size of that column, otherwise get the rectangle containing all columns.
	Args: [{ false item Object} {-1 true column int}], Returns: Rect2
*/
func (o *Tree) GetItemAreaRect(item ObjectImplementer, column gdnative.Int) gdnative.Rect2 {
	//log.Println("Calling Tree.GetItemAreaRect()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromObject(item.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromInt(column)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "get_item_area_rect")

	// Call the parent method.
	// Rect2
	retPtr := gdnative.NewEmptyRect2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRect2FromPointer(retPtr)
	return ret
}

/*
        Returns the tree item at the specified position (relative to the tree origin position).
	Args: [{ false position Vector2}], Returns: TreeItem
*/
func (o *Tree) GetItemAtPosition(position gdnative.Vector2) TreeItemImplementer {
	//log.Println("Calling Tree.GetItemAtPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(position)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "get_item_at_position")

	// Call the parent method.
	// TreeItem
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newTreeItemFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(TreeItemImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "TreeItem" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(TreeItemImplementer)
	}

	return &ret
}

/*
        Returns the next selected item after the given one.
	Args: [{ false from Object}], Returns: TreeItem
*/
func (o *Tree) GetNextSelected(from ObjectImplementer) TreeItemImplementer {
	//log.Println("Calling Tree.GetNextSelected()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(from.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "get_next_selected")

	// Call the parent method.
	// TreeItem
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newTreeItemFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(TreeItemImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "TreeItem" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(TreeItemImplementer)
	}

	return &ret
}

/*
        Returns the last pressed button's index.
	Args: [], Returns: int
*/
func (o *Tree) GetPressedButton() gdnative.Int {
	//log.Println("Calling Tree.GetPressedButton()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "get_pressed_button")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the tree's root item.
	Args: [], Returns: TreeItem
*/
func (o *Tree) GetRoot() TreeItemImplementer {
	//log.Println("Calling Tree.GetRoot()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "get_root")

	// Call the parent method.
	// TreeItem
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newTreeItemFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(TreeItemImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "TreeItem" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(TreeItemImplementer)
	}

	return &ret
}

/*
        Returns the current scrolling position.
	Args: [], Returns: Vector2
*/
func (o *Tree) GetScroll() gdnative.Vector2 {
	//log.Println("Calling Tree.GetScroll()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "get_scroll")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: enum.Tree::SelectMode
*/
func (o *Tree) GetSelectMode() TreeSelectMode {
	//log.Println("Calling Tree.GetSelectMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "get_select_mode")

	// Call the parent method.
	// enum.Tree::SelectMode
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return TreeSelectMode(ret)
}

/*
        Returns the currently selected item.
	Args: [], Returns: TreeItem
*/
func (o *Tree) GetSelected() TreeItemImplementer {
	//log.Println("Calling Tree.GetSelected()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "get_selected")

	// Call the parent method.
	// TreeItem
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newTreeItemFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(TreeItemImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "TreeItem" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(TreeItemImplementer)
	}

	return &ret
}

/*
        Returns the current selection's column.
	Args: [], Returns: int
*/
func (o *Tree) GetSelectedColumn() gdnative.Int {
	//log.Println("Calling Tree.GetSelectedColumn()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "get_selected_column")

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
	Args: [], Returns: bool
*/
func (o *Tree) IsFoldingHidden() gdnative.Bool {
	//log.Println("Calling Tree.IsFoldingHidden()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "is_folding_hidden")

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
	Args: [], Returns: bool
*/
func (o *Tree) IsRootHidden() gdnative.Bool {
	//log.Println("Calling Tree.IsRootHidden()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "is_root_hidden")

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
	Args: [{ false allow bool}], Returns: void
*/
func (o *Tree) SetAllowReselect(allow gdnative.Bool) {
	//log.Println("Calling Tree.SetAllowReselect()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(allow)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "set_allow_reselect")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false allow bool}], Returns: void
*/
func (o *Tree) SetAllowRmbSelect(allow gdnative.Bool) {
	//log.Println("Calling Tree.SetAllowRmbSelect()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(allow)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "set_allow_rmb_select")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        If [code]true[/code], the column will have the "Expand" flag of [Control].
	Args: [{ false column int} { false expand bool}], Returns: void
*/
func (o *Tree) SetColumnExpand(column gdnative.Int, expand gdnative.Bool) {
	//log.Println("Calling Tree.SetColumnExpand()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(column)
	ptrArguments[1] = gdnative.NewPointerFromBool(expand)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "set_column_expand")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Set the minimum width of a column.
	Args: [{ false column int} { false min_width int}], Returns: void
*/
func (o *Tree) SetColumnMinWidth(column gdnative.Int, minWidth gdnative.Int) {
	//log.Println("Calling Tree.SetColumnMinWidth()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(column)
	ptrArguments[1] = gdnative.NewPointerFromInt(minWidth)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "set_column_min_width")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Set the title of a column.
	Args: [{ false column int} { false title String}], Returns: void
*/
func (o *Tree) SetColumnTitle(column gdnative.Int, title gdnative.String) {
	//log.Println("Calling Tree.SetColumnTitle()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(column)
	ptrArguments[1] = gdnative.NewPointerFromString(title)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "set_column_title")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        If [code]true[/code], column titles are visible.
	Args: [{ false visible bool}], Returns: void
*/
func (o *Tree) SetColumnTitlesVisible(visible gdnative.Bool) {
	//log.Println("Calling Tree.SetColumnTitlesVisible()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(visible)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "set_column_titles_visible")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false amount int}], Returns: void
*/
func (o *Tree) SetColumns(amount gdnative.Int) {
	//log.Println("Calling Tree.SetColumns()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(amount)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "set_columns")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false flags int}], Returns: void
*/
func (o *Tree) SetDropModeFlags(flags gdnative.Int) {
	//log.Println("Calling Tree.SetDropModeFlags()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(flags)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "set_drop_mode_flags")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false hide bool}], Returns: void
*/
func (o *Tree) SetHideFolding(hide gdnative.Bool) {
	//log.Println("Calling Tree.SetHideFolding()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(hide)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "set_hide_folding")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *Tree) SetHideRoot(enable gdnative.Bool) {
	//log.Println("Calling Tree.SetHideRoot()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "set_hide_root")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/
func (o *Tree) SetSelectMode(mode gdnative.Int) {
	//log.Println("Calling Tree.SetSelectMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Tree", "set_select_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// TreeImplementer is an interface that implements the methods
// of the Tree class.
type TreeImplementer interface {
	ControlImplementer
	X_PopupSelect(arg0 gdnative.Int)
	X_RangeClickTimeout()
	X_ScrollMoved(arg0 gdnative.Real)
	X_TextEditorEnter(arg0 gdnative.String)
	X_TextEditorModalClose()
	X_ValueEditorChanged(arg0 gdnative.Real)
	AreColumnTitlesVisible() gdnative.Bool
	Clear()
	CreateItem(parent ObjectImplementer, idx gdnative.Int) TreeItemImplementer
	EnsureCursorIsVisible()
	GetAllowReselect() gdnative.Bool
	GetAllowRmbSelect() gdnative.Bool
	GetColumnAtPosition(position gdnative.Vector2) gdnative.Int
	GetColumnTitle(column gdnative.Int) gdnative.String
	GetColumnWidth(column gdnative.Int) gdnative.Int
	GetColumns() gdnative.Int
	GetCustomPopupRect() gdnative.Rect2
	GetDropModeFlags() gdnative.Int
	GetDropSectionAtPosition(position gdnative.Vector2) gdnative.Int
	GetEdited() TreeItemImplementer
	GetEditedColumn() gdnative.Int
	GetItemAreaRect(item ObjectImplementer, column gdnative.Int) gdnative.Rect2
	GetItemAtPosition(position gdnative.Vector2) TreeItemImplementer
	GetNextSelected(from ObjectImplementer) TreeItemImplementer
	GetPressedButton() gdnative.Int
	GetRoot() TreeItemImplementer
	GetScroll() gdnative.Vector2
	GetSelected() TreeItemImplementer
	GetSelectedColumn() gdnative.Int
	IsFoldingHidden() gdnative.Bool
	IsRootHidden() gdnative.Bool
	SetAllowReselect(allow gdnative.Bool)
	SetAllowRmbSelect(allow gdnative.Bool)
	SetColumnExpand(column gdnative.Int, expand gdnative.Bool)
	SetColumnMinWidth(column gdnative.Int, minWidth gdnative.Int)
	SetColumnTitle(column gdnative.Int, title gdnative.String)
	SetColumnTitlesVisible(visible gdnative.Bool)
	SetColumns(amount gdnative.Int)
	SetDropModeFlags(flags gdnative.Int)
	SetHideFolding(hide gdnative.Bool)
	SetHideRoot(enable gdnative.Bool)
	SetSelectMode(mode gdnative.Int)
}
