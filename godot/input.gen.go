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

// InputCursorShape is an enum for CursorShape values.
type InputCursorShape int

const (
	InputCursorArrow        InputCursorShape = 0
	InputCursorBdiagsize    InputCursorShape = 11
	InputCursorBusy         InputCursorShape = 5
	InputCursorCanDrop      InputCursorShape = 7
	InputCursorCross        InputCursorShape = 3
	InputCursorDrag         InputCursorShape = 6
	InputCursorFdiagsize    InputCursorShape = 12
	InputCursorForbidden    InputCursorShape = 8
	InputCursorHelp         InputCursorShape = 16
	InputCursorHsize        InputCursorShape = 10
	InputCursorHsplit       InputCursorShape = 15
	InputCursorIbeam        InputCursorShape = 1
	InputCursorMove         InputCursorShape = 13
	InputCursorPointingHand InputCursorShape = 2
	InputCursorVsize        InputCursorShape = 9
	InputCursorVsplit       InputCursorShape = 14
	InputCursorWait         InputCursorShape = 4
)

// InputMouseMode is an enum for MouseMode values.
type InputMouseMode int

const (
	InputMouseModeCaptured InputMouseMode = 2
	InputMouseModeConfined InputMouseMode = 3
	InputMouseModeHidden   InputMouseMode = 1
	InputMouseModeVisible  InputMouseMode = 0
)

//func NewinputFromPointer(ptr gdnative.Pointer) input {
func newInputFromPointer(ptr gdnative.Pointer) input {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := input{}
	obj.SetBaseObject(owner)

	return obj
}

func newSingletonInput() *input {
	return &input{}
}

/*
   A Singleton that deals with inputs. This includes key presses, mouse buttons and movement, joypads, and input actions. Actions and their events can be set in the Project Settings / Input Map tab. Or be set with [InputMap].
*/
var Input = newSingletonInput()

/*
A Singleton that deals with inputs. This includes key presses, mouse buttons and movement, joypads, and input actions. Actions and their events can be set in the Project Settings / Input Map tab. Or be set with [InputMap].
*/
type input struct {
	Object
	owner       gdnative.Object
	initialized bool
}

// EnsureSingleton will check to see if we have an object for it. If not, it will fetch its
// GDNative object and set it.
func (o *input) ensureSingleton() {
	if o.initialized == true {
		return
	}
	//log.Println("Singleton not found. Fetching from GDNative...")
	base := gdnative.GetSingleton("Input")
	o.SetBaseObject(base)
	o.initialized = true
}

func (o *input) BaseClass() string {
	return "Input"
}

/*
        This will simulate pressing the specified action. The strength can be used for non-boolean actions, it's ranged between 0 and 1 representing the intensity of the given action.
	Args: [{ false action String} {1 true strength float}], Returns: void
*/
func (o *input) ActionPress(action gdnative.String, strength gdnative.Real) {
	o.ensureSingleton()
	//log.Println("Calling Input.ActionPress()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(action)
	ptrArguments[1] = gdnative.NewPointerFromReal(strength)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Input", "action_press")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        If the specified action is already pressed, this will release it.
	Args: [{ false action String}], Returns: void
*/
func (o *input) ActionRelease(action gdnative.String) {
	o.ensureSingleton()
	//log.Println("Calling Input.ActionRelease()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(action)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Input", "action_release")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Add a new mapping entry (in SDL2 format) to the mapping database. Optionally update already connected devices.
	Args: [{ false mapping String} {False true update_existing bool}], Returns: void
*/
func (o *input) AddJoyMapping(mapping gdnative.String, updateExisting gdnative.Bool) {
	o.ensureSingleton()
	//log.Println("Calling Input.AddJoyMapping()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(mapping)
	ptrArguments[1] = gdnative.NewPointerFromBool(updateExisting)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Input", "add_joy_mapping")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        If the device has an accelerometer, this will return the acceleration. Otherwise, it returns an empty [Vector3]. Note this method returns an empty [Vector3] when running from the editor even when your device has an accelerometer. You must export your project to a supported device to read values from the accelerometer.
	Args: [], Returns: Vector3
*/
func (o *input) GetAccelerometer() gdnative.Vector3 {
	o.ensureSingleton()
	//log.Println("Calling Input.GetAccelerometer()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Input", "get_accelerometer")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Returns a value between 0 and 1 representing the intensity of the given action. In a joypad, for example, the further away the axis (analog sticks or L2, R2 triggers) is from the dead zone, the closer the value will be to 1. If the action is mapped to a control that has no axis as the keyboard, the value returned will be 0 or 1.
	Args: [{ false action String}], Returns: float
*/
func (o *input) GetActionStrength(action gdnative.String) gdnative.Real {
	o.ensureSingleton()
	//log.Println("Calling Input.GetActionStrength()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(action)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Input", "get_action_strength")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Returns an [Array] containing the device IDs of all currently connected joypads.
	Args: [], Returns: Array
*/
func (o *input) GetConnectedJoypads() gdnative.Array {
	o.ensureSingleton()
	//log.Println("Calling Input.GetConnectedJoypads()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Input", "get_connected_joypads")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*
        If the device has an accelerometer, this will return the gravity. Otherwise, it returns an empty [Vector3].
	Args: [], Returns: Vector3
*/
func (o *input) GetGravity() gdnative.Vector3 {
	o.ensureSingleton()
	//log.Println("Calling Input.GetGravity()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Input", "get_gravity")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        If the device has a gyroscope, this will return the rate of rotation in rad/s around a device's x, y, and z axis. Otherwise, it returns an empty [Vector3].
	Args: [], Returns: Vector3
*/
func (o *input) GetGyroscope() gdnative.Vector3 {
	o.ensureSingleton()
	//log.Println("Calling Input.GetGyroscope()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Input", "get_gyroscope")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Returns the current value of the joypad axis at given index (see [code]JOY_*[/code] constants in [@GlobalScope])
	Args: [{ false device int} { false axis int}], Returns: float
*/
func (o *input) GetJoyAxis(device gdnative.Int, axis gdnative.Int) gdnative.Real {
	o.ensureSingleton()
	//log.Println("Calling Input.GetJoyAxis()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(device)
	ptrArguments[1] = gdnative.NewPointerFromInt(axis)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Input", "get_joy_axis")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Returns the index of the provided axis name.
	Args: [{ false axis String}], Returns: int
*/
func (o *input) GetJoyAxisIndexFromString(axis gdnative.String) gdnative.Int {
	o.ensureSingleton()
	//log.Println("Calling Input.GetJoyAxisIndexFromString()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(axis)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Input", "get_joy_axis_index_from_string")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Receives a [code]JOY_AXIS_*[/code] Enum and returns its equivalent name as a string.
	Args: [{ false axis_index int}], Returns: String
*/
func (o *input) GetJoyAxisString(axisIndex gdnative.Int) gdnative.String {
	o.ensureSingleton()
	//log.Println("Calling Input.GetJoyAxisString()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(axisIndex)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Input", "get_joy_axis_string")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Returns the index of the provided button name.
	Args: [{ false button String}], Returns: int
*/
func (o *input) GetJoyButtonIndexFromString(button gdnative.String) gdnative.Int {
	o.ensureSingleton()
	//log.Println("Calling Input.GetJoyButtonIndexFromString()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(button)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Input", "get_joy_button_index_from_string")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Receives a [code]JOY_BUTTON_*[/code] Enum and returns its equivalent name as a string.
	Args: [{ false button_index int}], Returns: String
*/
func (o *input) GetJoyButtonString(buttonIndex gdnative.Int) gdnative.String {
	o.ensureSingleton()
	//log.Println("Calling Input.GetJoyButtonString()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(buttonIndex)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Input", "get_joy_button_string")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Returns a SDL2 compatible device guid on platforms that use gamepad remapping. Returns "Default Gamepad" otherwise.
	Args: [{ false device int}], Returns: String
*/
func (o *input) GetJoyGuid(device gdnative.Int) gdnative.String {
	o.ensureSingleton()
	//log.Println("Calling Input.GetJoyGuid()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(device)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Input", "get_joy_guid")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Returns the name of the joypad at the specified device index
	Args: [{ false device int}], Returns: String
*/
func (o *input) GetJoyName(device gdnative.Int) gdnative.String {
	o.ensureSingleton()
	//log.Println("Calling Input.GetJoyName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(device)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Input", "get_joy_name")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Returns the duration of the current vibration effect in seconds.
	Args: [{ false device int}], Returns: float
*/
func (o *input) GetJoyVibrationDuration(device gdnative.Int) gdnative.Real {
	o.ensureSingleton()
	//log.Println("Calling Input.GetJoyVibrationDuration()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(device)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Input", "get_joy_vibration_duration")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Returns the strength of the joypad vibration: x is the strength of the weak motor, and y is the strength of the strong motor.
	Args: [{ false device int}], Returns: Vector2
*/
func (o *input) GetJoyVibrationStrength(device gdnative.Int) gdnative.Vector2 {
	o.ensureSingleton()
	//log.Println("Calling Input.GetJoyVibrationStrength()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(device)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Input", "get_joy_vibration_strength")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*
        Returns the mouse speed for the last time the cursor was moved, and this until the next frame where the mouse moves. This means that even if the mouse is not moving, this function will still return the value of the last motion.
	Args: [], Returns: Vector2
*/
func (o *input) GetLastMouseSpeed() gdnative.Vector2 {
	o.ensureSingleton()
	//log.Println("Calling Input.GetLastMouseSpeed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Input", "get_last_mouse_speed")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*
        If the device has a magnetometer, this will return the magnetic field strength in micro-Tesla for all axes.
	Args: [], Returns: Vector3
*/
func (o *input) GetMagnetometer() gdnative.Vector3 {
	o.ensureSingleton()
	//log.Println("Calling Input.GetMagnetometer()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Input", "get_magnetometer")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Returns mouse buttons as a bitmask. If multiple mouse buttons are pressed at the same time the bits are added together.
	Args: [], Returns: int
*/
func (o *input) GetMouseButtonMask() gdnative.Int {
	o.ensureSingleton()
	//log.Println("Calling Input.GetMouseButtonMask()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Input", "get_mouse_button_mask")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Return the mouse mode. See the constants for more information.
	Args: [], Returns: enum.Input::MouseMode
*/
func (o *input) GetMouseMode() InputMouseMode {
	o.ensureSingleton()
	//log.Println("Calling Input.GetMouseMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Input", "get_mouse_mode")

	// Call the parent method.
	// enum.Input::MouseMode
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return InputMouseMode(ret)
}

/*
        Returns [code]true[/code] when the user starts pressing the action event, meaning it's true only on the frame that the user pressed down the button. This is useful for code that needs to run only once when an action is pressed, instead of every frame while it's pressed.
	Args: [{ false action String}], Returns: bool
*/
func (o *input) IsActionJustPressed(action gdnative.String) gdnative.Bool {
	o.ensureSingleton()
	//log.Println("Calling Input.IsActionJustPressed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(action)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Input", "is_action_just_pressed")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] when the user stops pressing the action event, meaning it's true only on the frame that the user released the button.
	Args: [{ false action String}], Returns: bool
*/
func (o *input) IsActionJustReleased(action gdnative.String) gdnative.Bool {
	o.ensureSingleton()
	//log.Println("Calling Input.IsActionJustReleased()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(action)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Input", "is_action_just_released")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] if you are pressing the action event.
	Args: [{ false action String}], Returns: bool
*/
func (o *input) IsActionPressed(action gdnative.String) gdnative.Bool {
	o.ensureSingleton()
	//log.Println("Calling Input.IsActionPressed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(action)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Input", "is_action_pressed")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] if you are pressing the joypad button. (see [code]JOY_*[/code] constants in [@GlobalScope])
	Args: [{ false device int} { false button int}], Returns: bool
*/
func (o *input) IsJoyButtonPressed(device gdnative.Int, button gdnative.Int) gdnative.Bool {
	o.ensureSingleton()
	//log.Println("Calling Input.IsJoyButtonPressed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(device)
	ptrArguments[1] = gdnative.NewPointerFromInt(button)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Input", "is_joy_button_pressed")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] if the system knows the specified device. This means that it sets all button and axis indices exactly as defined in the [code]JOY_*[/code] constants (see [@GlobalScope]). Unknown joypads are not expected to match these constants, but you can still retrieve events from them.
	Args: [{ false device int}], Returns: bool
*/
func (o *input) IsJoyKnown(device gdnative.Int) gdnative.Bool {
	o.ensureSingleton()
	//log.Println("Calling Input.IsJoyKnown()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(device)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Input", "is_joy_known")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] if you are pressing the key. You can pass [code]KEY_*[/code], which are pre-defined constants listed in [@GlobalScope].
	Args: [{ false scancode int}], Returns: bool
*/
func (o *input) IsKeyPressed(scancode gdnative.Int) gdnative.Bool {
	o.ensureSingleton()
	//log.Println("Calling Input.IsKeyPressed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(scancode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Input", "is_key_pressed")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] if you are pressing the mouse button. You can pass [code]BUTTON_*[/code], which are pre-defined constants listed in [@GlobalScope].
	Args: [{ false button int}], Returns: bool
*/
func (o *input) IsMouseButtonPressed(button gdnative.Int) gdnative.Bool {
	o.ensureSingleton()
	//log.Println("Calling Input.IsMouseButtonPressed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(button)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Input", "is_mouse_button_pressed")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false device int} { false connected bool} { false name String} { false guid String}], Returns: void
*/
func (o *input) JoyConnectionChanged(device gdnative.Int, connected gdnative.Bool, name gdnative.String, guid gdnative.String) {
	o.ensureSingleton()
	//log.Println("Calling Input.JoyConnectionChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromInt(device)
	ptrArguments[1] = gdnative.NewPointerFromBool(connected)
	ptrArguments[2] = gdnative.NewPointerFromString(name)
	ptrArguments[3] = gdnative.NewPointerFromString(guid)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Input", "joy_connection_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Feeds an [InputEvent] to the game. Can be used to artificially trigger input events from code.
	Args: [{ false event InputEvent}], Returns: void
*/
func (o *input) ParseInputEvent(event InputEventImplementer) {
	o.ensureSingleton()
	//log.Println("Calling Input.ParseInputEvent()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(event.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Input", "parse_input_event")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Removes all mappings from the internal db that match the given uid.
	Args: [{ false guid String}], Returns: void
*/
func (o *input) RemoveJoyMapping(guid gdnative.String) {
	o.ensureSingleton()
	//log.Println("Calling Input.RemoveJoyMapping()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(guid)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Input", "remove_joy_mapping")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets a custom mouse cursor image, which is only visible inside the game window. The hotspot can also be specified. Passing [code]null[/code] to the image parameter resets to the system cursor. See enum [code]CURSOR_*[/code] for the list of shapes. [code]image[/code]'s size must be lower than 256x256. [code]hotspot[/code] must be within [code]image[/code]'s size.
	Args: [{ false image Resource} {0 true shape int} {(0, 0) true hotspot Vector2}], Returns: void
*/
func (o *input) SetCustomMouseCursor(image ResourceImplementer, shape gdnative.Int, hotspot gdnative.Vector2) {
	o.ensureSingleton()
	//log.Println("Calling Input.SetCustomMouseCursor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromObject(image.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromInt(shape)
	ptrArguments[2] = gdnative.NewPointerFromVector2(hotspot)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Input", "set_custom_mouse_cursor")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the default cursor shape to be used in the viewport instead of [code]CURSOR_ARROW[/code]. Note that if you want to change the default cursor shape for [Control]'s nodes, use [member Control.mouse_default_cursor_shape] instead.
	Args: [{0 true shape int}], Returns: void
*/
func (o *input) SetDefaultCursorShape(shape gdnative.Int) {
	o.ensureSingleton()
	//log.Println("Calling Input.SetDefaultCursorShape()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(shape)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Input", "set_default_cursor_shape")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Set the mouse mode. See the constants for more information.
	Args: [{ false mode int}], Returns: void
*/
func (o *input) SetMouseMode(mode gdnative.Int) {
	o.ensureSingleton()
	//log.Println("Calling Input.SetMouseMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Input", "set_mouse_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Whether to accumulate similar input events sent by the operating system. Defaults to [code]true[/code].
	Args: [{ false enable bool}], Returns: void
*/
func (o *input) SetUseAccumulatedInput(enable gdnative.Bool) {
	o.ensureSingleton()
	//log.Println("Calling Input.SetUseAccumulatedInput()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Input", "set_use_accumulated_input")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Starts to vibrate the joypad. Joypads usually come with two rumble motors, a strong and a weak one. weak_magnitude is the strength of the weak motor (between 0 and 1) and strong_magnitude is the strength of the strong motor (between 0 and 1). duration is the duration of the effect in seconds (a duration of 0 will try to play the vibration indefinitely). Note that not every hardware is compatible with long effect durations, it is recommended to restart an effect if in need to play it for more than a few seconds.
	Args: [{ false device int} { false weak_magnitude float} { false strong_magnitude float} {0 true duration float}], Returns: void
*/
func (o *input) StartJoyVibration(device gdnative.Int, weakMagnitude gdnative.Real, strongMagnitude gdnative.Real, duration gdnative.Real) {
	o.ensureSingleton()
	//log.Println("Calling Input.StartJoyVibration()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromInt(device)
	ptrArguments[1] = gdnative.NewPointerFromReal(weakMagnitude)
	ptrArguments[2] = gdnative.NewPointerFromReal(strongMagnitude)
	ptrArguments[3] = gdnative.NewPointerFromReal(duration)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Input", "start_joy_vibration")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Stops the vibration of the joypad.
	Args: [{ false device int}], Returns: void
*/
func (o *input) StopJoyVibration(device gdnative.Int) {
	o.ensureSingleton()
	//log.Println("Calling Input.StopJoyVibration()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(device)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Input", "stop_joy_vibration")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the mouse position to the specified vector.
	Args: [{ false to Vector2}], Returns: void
*/
func (o *input) WarpMousePosition(to gdnative.Vector2) {
	o.ensureSingleton()
	//log.Println("Calling Input.WarpMousePosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(to)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Input", "warp_mouse_position")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// InputImplementer is an interface that implements the methods
// of the Input class.
type InputImplementer interface {
	ObjectImplementer
	ActionPress(action gdnative.String, strength gdnative.Real)
	ActionRelease(action gdnative.String)
	AddJoyMapping(mapping gdnative.String, updateExisting gdnative.Bool)
	GetAccelerometer() gdnative.Vector3
	GetActionStrength(action gdnative.String) gdnative.Real
	GetConnectedJoypads() gdnative.Array
	GetGravity() gdnative.Vector3
	GetGyroscope() gdnative.Vector3
	GetJoyAxis(device gdnative.Int, axis gdnative.Int) gdnative.Real
	GetJoyAxisIndexFromString(axis gdnative.String) gdnative.Int
	GetJoyAxisString(axisIndex gdnative.Int) gdnative.String
	GetJoyButtonIndexFromString(button gdnative.String) gdnative.Int
	GetJoyButtonString(buttonIndex gdnative.Int) gdnative.String
	GetJoyGuid(device gdnative.Int) gdnative.String
	GetJoyName(device gdnative.Int) gdnative.String
	GetJoyVibrationDuration(device gdnative.Int) gdnative.Real
	GetJoyVibrationStrength(device gdnative.Int) gdnative.Vector2
	GetLastMouseSpeed() gdnative.Vector2
	GetMagnetometer() gdnative.Vector3
	GetMouseButtonMask() gdnative.Int
	IsActionJustPressed(action gdnative.String) gdnative.Bool
	IsActionJustReleased(action gdnative.String) gdnative.Bool
	IsActionPressed(action gdnative.String) gdnative.Bool
	IsJoyButtonPressed(device gdnative.Int, button gdnative.Int) gdnative.Bool
	IsJoyKnown(device gdnative.Int) gdnative.Bool
	IsKeyPressed(scancode gdnative.Int) gdnative.Bool
	IsMouseButtonPressed(button gdnative.Int) gdnative.Bool
	JoyConnectionChanged(device gdnative.Int, connected gdnative.Bool, name gdnative.String, guid gdnative.String)
	ParseInputEvent(event InputEventImplementer)
	RemoveJoyMapping(guid gdnative.String)
	SetCustomMouseCursor(image ResourceImplementer, shape gdnative.Int, hotspot gdnative.Vector2)
	SetDefaultCursorShape(shape gdnative.Int)
	SetMouseMode(mode gdnative.Int)
	SetUseAccumulatedInput(enable gdnative.Bool)
	StartJoyVibration(device gdnative.Int, weakMagnitude gdnative.Real, strongMagnitude gdnative.Real, duration gdnative.Real)
	StopJoyVibration(device gdnative.Int)
	WarpMousePosition(to gdnative.Vector2)
}
