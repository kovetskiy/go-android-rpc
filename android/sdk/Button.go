// It is autogenerated bindings for Android sdk class.
//
// See documentation about methods on: https://developer.android.com//reference/android/widget/Button.html
package sdk

import "github.com/seletskiy/go-android-rpc/android"

type Button struct {
	TextView
}

func NewButton(id string) interface{} {
	obj := Button{NewTextView(id).(TextView)}

	return obj
}


func init() {
	android.ViewTypeConstructors["Button"] = NewButton
}

