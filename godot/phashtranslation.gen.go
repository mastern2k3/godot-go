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

//func NewPHashTranslationFromPointer(ptr gdnative.Pointer) PHashTranslation {
func newPHashTranslationFromPointer(ptr gdnative.Pointer) PHashTranslation {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := PHashTranslation{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Optimized translation. Uses real-time compressed translations, which results in very small dictionaries.
*/
type PHashTranslation struct {
	Translation
	owner gdnative.Object
}

func (o *PHashTranslation) BaseClass() string {
	return "PHashTranslation"
}

/*
        Generates and sets an optimized translation from the given [Translation] resource.
	Args: [{ false from Translation}], Returns: void
*/
func (o *PHashTranslation) Generate(from TranslationImplementer) {
	//log.Println("Calling PHashTranslation.Generate()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(from.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PHashTranslation", "generate")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// PHashTranslationImplementer is an interface that implements the methods
// of the PHashTranslation class.
type PHashTranslationImplementer interface {
	TranslationImplementer
	Generate(from TranslationImplementer)
}
