// It is autogenerated bindings for Android sdk class.
//
// See documentation about methods on: https://developer.android.com//reference/android/widget/HorizontalScrollView.html
package sdk

import "github.com/seletskiy/go-android-rpc/android"

type HorizontalScrollView struct {
	View
}

func NewHorizontalScrollView(id string) interface{} {
	obj := HorizontalScrollView{NewView(id).(View)}

	return obj
}

func (obj HorizontalScrollView) GetClassName() string {
	return "android.widget.HorizontalScrollView"
}

func init() {
	android.ViewTypeConstructors["android.widget.HorizontalScrollView"] = NewHorizontalScrollView
}

func (obj HorizontalScrollView) ArrowScroll(direction_ int) (bool, error) {
	return return_bool(android.CallViewMethod(
		obj.GetId(),
		"android.widget.HorizontalScrollView",
		"arrowScroll",
		direction_,
	))
}

func (obj HorizontalScrollView) ComputeScroll() error {
	return return_error(android.CallViewMethod(
		obj.GetId(),
		"android.widget.HorizontalScrollView",
		"computeScroll",
	))
}

func (obj HorizontalScrollView) Fling(velocityX_ int) error {
	return return_error(android.CallViewMethod(
		obj.GetId(),
		"android.widget.HorizontalScrollView",
		"fling",
		velocityX_,
	))
}

func (obj HorizontalScrollView) FullScroll(direction_ int) (bool, error) {
	return return_bool(android.CallViewMethod(
		obj.GetId(),
		"android.widget.HorizontalScrollView",
		"fullScroll",
		direction_,
	))
}

func (obj HorizontalScrollView) GetMaxScrollAmount() (int, error) {
	return return_int(android.CallViewMethod(
		obj.GetId(),
		"android.widget.HorizontalScrollView",
		"getMaxScrollAmount",
	))
}

func (obj HorizontalScrollView) IsFillViewport() (bool, error) {
	return return_bool(android.CallViewMethod(
		obj.GetId(),
		"android.widget.HorizontalScrollView",
		"isFillViewport",
	))
}

func (obj HorizontalScrollView) IsSmoothScrollingEnabled() (bool, error) {
	return return_bool(android.CallViewMethod(
		obj.GetId(),
		"android.widget.HorizontalScrollView",
		"isSmoothScrollingEnabled",
	))
}

func (obj HorizontalScrollView) PageScroll(direction_ int) (bool, error) {
	return return_bool(android.CallViewMethod(
		obj.GetId(),
		"android.widget.HorizontalScrollView",
		"pageScroll",
		direction_,
	))
}

func (obj HorizontalScrollView) RequestDisallowInterceptTouchEvent(disallowIntercept_ bool) error {
	return return_error(android.CallViewMethod(
		obj.GetId(),
		"android.widget.HorizontalScrollView",
		"requestDisallowInterceptTouchEvent",
		disallowIntercept_,
	))
}

func (obj HorizontalScrollView) RequestLayout() error {
	return return_error(android.CallViewMethod(
		obj.GetId(),
		"android.widget.HorizontalScrollView",
		"requestLayout",
	))
}

func (obj HorizontalScrollView) ScrollTo(x_ int, y_ int) error {
	return return_error(android.CallViewMethod(
		obj.GetId(),
		"android.widget.HorizontalScrollView",
		"scrollTo",
		x_, y_,
	))
}

func (obj HorizontalScrollView) SetFillViewport(fillViewport_ bool) error {
	return return_error(android.CallViewMethod(
		obj.GetId(),
		"android.widget.HorizontalScrollView",
		"setFillViewport",
		fillViewport_,
	))
}

func (obj HorizontalScrollView) SetOverScrollMode(mode_ int) error {
	return return_error(android.CallViewMethod(
		obj.GetId(),
		"android.widget.HorizontalScrollView",
		"setOverScrollMode",
		mode_,
	))
}

func (obj HorizontalScrollView) SetSmoothScrollingEnabled(smoothScrollingEnabled_ bool) error {
	return return_error(android.CallViewMethod(
		obj.GetId(),
		"android.widget.HorizontalScrollView",
		"setSmoothScrollingEnabled",
		smoothScrollingEnabled_,
	))
}

func (obj HorizontalScrollView) ShouldDelayChildPressedState() (bool, error) {
	return return_bool(android.CallViewMethod(
		obj.GetId(),
		"android.widget.HorizontalScrollView",
		"shouldDelayChildPressedState",
	))
}

func (obj HorizontalScrollView) SmoothScrollBy(dx_ int, dy_ int) error {
	return return_error(android.CallViewMethod(
		obj.GetId(),
		"android.widget.HorizontalScrollView",
		"smoothScrollBy",
		dx_, dy_,
	))
}

func (obj HorizontalScrollView) SmoothScrollTo(x_ int, y_ int) error {
	return return_error(android.CallViewMethod(
		obj.GetId(),
		"android.widget.HorizontalScrollView",
		"smoothScrollTo",
		x_, y_,
	))
}

