// It is autogenerated bindings for Android sdk class.
//
// See documentation about methods on: https://developer.android.com//reference/android/widget/ExpandableListView.html
package sdk

import "github.com/seletskiy/go-android-rpc/android"

type ExpandableListView struct {
	View
}

func NewExpandableListView(id string) interface{} {
	obj := ExpandableListView{NewView(id).(View)}

	return obj
}


func init() {
	android.ViewTypeConstructors["ExpandableListView"] = NewExpandableListView
}

func (obj ExpandableListView) CollapseGroup(groupPos_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ExpandableListView",
		"collapseGroup",
		groupPos_,
	)
}

func (obj ExpandableListView) ExpandGroup(groupPos_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ExpandableListView",
		"expandGroup",
		groupPos_,
	)
}

func (obj ExpandableListView) ExpandGroup2ib(groupPos_ int, animate_ bool) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ExpandableListView",
		"expandGroup",
		groupPos_, animate_,
	)
}

func (obj ExpandableListView) IsGroupExpanded(groupPosition_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ExpandableListView",
		"isGroupExpanded",
		groupPosition_,
	)
}

func (obj ExpandableListView) OnRtlPropertiesChanged(layoutDirection_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ExpandableListView",
		"onRtlPropertiesChanged",
		layoutDirection_,
	)
}

func (obj ExpandableListView) SetChildIndicatorBounds(left_ int, right_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ExpandableListView",
		"setChildIndicatorBounds",
		left_, right_,
	)
}

func (obj ExpandableListView) SetChildIndicatorBoundsRelative(start_ int, end_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ExpandableListView",
		"setChildIndicatorBoundsRelative",
		start_, end_,
	)
}

func (obj ExpandableListView) SetIndicatorBounds(left_ int, right_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ExpandableListView",
		"setIndicatorBounds",
		left_, right_,
	)
}

func (obj ExpandableListView) SetIndicatorBoundsRelative(start_ int, end_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ExpandableListView",
		"setIndicatorBoundsRelative",
		start_, end_,
	)
}

func (obj ExpandableListView) SetSelectedChild(groupPosition_ int, childPosition_ int, shouldExpandGroup_ bool) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ExpandableListView",
		"setSelectedChild",
		groupPosition_, childPosition_, shouldExpandGroup_,
	)
}

func (obj ExpandableListView) SetSelectedGroup(groupPosition_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ExpandableListView",
		"setSelectedGroup",
		groupPosition_,
	)
}
