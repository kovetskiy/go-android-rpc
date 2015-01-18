// It is autogenerated bindings for Android sdk class.
//
// See documentation about methods on: https://developer.android.com//reference/android/widget/CheckedTextView.html
package sdk

import "github.com/seletskiy/go-android-rpc/android"

type CheckedTextView struct {
	View
}

func NewCheckedTextView(id string) interface{} {
	obj := CheckedTextView{NewView(id).(View)}

	return obj
}


func init() {
	android.ViewTypeConstructors["CheckedTextView"] = NewCheckedTextView
}

func (obj CheckedTextView) DrawableHotspotChanged(x_ float64, y_ float64) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"CheckedTextView",
		"drawableHotspotChanged",
		android.Float(x_), android.Float(y_),
	)
}

func (obj CheckedTextView) IsChecked() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"CheckedTextView",
		"isChecked",
	)
}

func (obj CheckedTextView) JumpDrawablesToCurrentState() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"CheckedTextView",
		"jumpDrawablesToCurrentState",
	)
}

func (obj CheckedTextView) OnRtlPropertiesChanged(layoutDirection_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"CheckedTextView",
		"onRtlPropertiesChanged",
		layoutDirection_,
	)
}

func (obj CheckedTextView) SetCheckMarkDrawable(resid_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"CheckedTextView",
		"setCheckMarkDrawable",
		resid_,
	)
}

func (obj CheckedTextView) SetChecked(checked_ bool) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"CheckedTextView",
		"setChecked",
		checked_,
	)
}

func (obj CheckedTextView) SetVisibility(visibility_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"CheckedTextView",
		"setVisibility",
		visibility_,
	)
}

func (obj CheckedTextView) Toggle() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"CheckedTextView",
		"toggle",
	)
}

