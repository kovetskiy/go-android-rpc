// It is autogenerated bindings for Android sdk class.
//
// See documentation about methods on: https://developer.android.com//reference/android/widget/GridView.html
package sdk

import "github.com/seletskiy/go-android-rpc/android"

type GridView struct {
	View
}

func NewGridView(id string) interface{} {
	obj := GridView{NewView(id).(View)}

	return obj
}


func init() {
	android.ViewTypeConstructors["GridView"] = NewGridView
}

func (obj GridView) GetColumnWidth() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.GridView",
		"getColumnWidth",
	)
}

func (obj GridView) GetGravity() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.GridView",
		"getGravity",
	)
}

func (obj GridView) GetHorizontalSpacing() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.GridView",
		"getHorizontalSpacing",
	)
}

func (obj GridView) GetNumColumns() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.GridView",
		"getNumColumns",
	)
}

func (obj GridView) GetRequestedColumnWidth() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.GridView",
		"getRequestedColumnWidth",
	)
}

func (obj GridView) GetRequestedHorizontalSpacing() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.GridView",
		"getRequestedHorizontalSpacing",
	)
}

func (obj GridView) GetStretchMode() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.GridView",
		"getStretchMode",
	)
}

func (obj GridView) GetVerticalSpacing() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.GridView",
		"getVerticalSpacing",
	)
}

func (obj GridView) SetColumnWidth(columnWidth_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.GridView",
		"setColumnWidth",
		columnWidth_,
	)
}

func (obj GridView) SetGravity(gravity_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.GridView",
		"setGravity",
		gravity_,
	)
}

func (obj GridView) SetHorizontalSpacing(horizontalSpacing_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.GridView",
		"setHorizontalSpacing",
		horizontalSpacing_,
	)
}

func (obj GridView) SetNumColumns(numColumns_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.GridView",
		"setNumColumns",
		numColumns_,
	)
}

func (obj GridView) SetSelection(position_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.GridView",
		"setSelection",
		position_,
	)
}

func (obj GridView) SetStretchMode(stretchMode_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.GridView",
		"setStretchMode",
		stretchMode_,
	)
}

func (obj GridView) SetVerticalSpacing(verticalSpacing_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.GridView",
		"setVerticalSpacing",
		verticalSpacing_,
	)
}

func (obj GridView) SmoothScrollByOffset(offset_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.GridView",
		"smoothScrollByOffset",
		offset_,
	)
}

func (obj GridView) SmoothScrollToPosition(position_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.GridView",
		"smoothScrollToPosition",
		position_,
	)
}

