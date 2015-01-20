// It is autogenerated bindings for Android sdk class.
//
// See documentation about methods on: https://developer.android.com//reference/android/widget/TextView.html
package sdk

import "github.com/seletskiy/go-android-rpc/android"

type TextView struct {
	View
}

func NewTextView(id string) interface{} {
	obj := TextView{NewView(id).(View)}

	return obj
}

func init() {
	android.ViewTypeConstructors["TextView"] = NewTextView
}

func (obj TextView) Append(text_ string) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"append",
		text_,
	)
}

func (obj TextView) Append3sii(text_ string, start_ int, end_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"append",
		text_, start_, end_,
	)
}

func (obj TextView) BeginBatchEdit() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"beginBatchEdit",
	)
}

func (obj TextView) BringPointIntoView(offset_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"bringPointIntoView",
		offset_,
	)
}

func (obj TextView) CancelLongPress() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"cancelLongPress",
	)
}

func (obj TextView) ClearComposingText() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"clearComposingText",
	)
}

func (obj TextView) ComputeScroll() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"computeScroll",
	)
}

func (obj TextView) Debug(depth_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"debug",
		depth_,
	)
}

func (obj TextView) DidTouchFocusSelect() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"didTouchFocusSelect",
	)
}

func (obj TextView) DrawableHotspotChanged(x_ float64, y_ float64) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"drawableHotspotChanged",
		android.Float(x_), android.Float(y_),
	)
}

func (obj TextView) EndBatchEdit() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"endBatchEdit",
	)
}

func (obj TextView) GetAutoLinkMask() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getAutoLinkMask",
	)
}

func (obj TextView) GetBaseline() map[string]interface{} {
	return android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getBaseline",
	)
}

func (obj TextView) GetCompoundDrawablePadding() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getCompoundDrawablePadding",
	)
}

func (obj TextView) GetCompoundPaddingBottom() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getCompoundPaddingBottom",
	)
}

func (obj TextView) GetCompoundPaddingEnd() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getCompoundPaddingEnd",
	)
}

func (obj TextView) GetCompoundPaddingLeft() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getCompoundPaddingLeft",
	)
}

func (obj TextView) GetCompoundPaddingRight() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getCompoundPaddingRight",
	)
}

func (obj TextView) GetCompoundPaddingStart() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getCompoundPaddingStart",
	)
}

func (obj TextView) GetCompoundPaddingTop() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getCompoundPaddingTop",
	)
}

func (obj TextView) GetCurrentHintTextColor() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getCurrentHintTextColor",
	)
}

func (obj TextView) GetCurrentTextColor() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getCurrentTextColor",
	)
}

func (obj TextView) GetError() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getError",
	)
}

func (obj TextView) GetExtendedPaddingBottom() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getExtendedPaddingBottom",
	)
}

func (obj TextView) GetExtendedPaddingTop() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getExtendedPaddingTop",
	)
}

func (obj TextView) GetFontFeatureSettings() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getFontFeatureSettings",
	)
}

func (obj TextView) GetFreezesText() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getFreezesText",
	)
}

func (obj TextView) GetGravity() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getGravity",
	)
}

func (obj TextView) GetHighlightColor() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getHighlightColor",
	)
}

func (obj TextView) GetHint() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getHint",
	)
}

func (obj TextView) GetImeActionId() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getImeActionId",
	)
}

func (obj TextView) GetImeActionLabel() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getImeActionLabel",
	)
}

func (obj TextView) GetImeOptions() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getImeOptions",
	)
}

func (obj TextView) GetIncludeFontPadding() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getIncludeFontPadding",
	)
}

func (obj TextView) GetInputType() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getInputType",
	)
}

func (obj TextView) GetLetterSpacing() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getLetterSpacing",
	)
}

func (obj TextView) GetLineCount() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getLineCount",
	)
}

func (obj TextView) GetLineHeight() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getLineHeight",
	)
}

func (obj TextView) GetLineSpacingExtra() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getLineSpacingExtra",
	)
}

func (obj TextView) GetLineSpacingMultiplier() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getLineSpacingMultiplier",
	)
}

func (obj TextView) GetLinksClickable() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getLinksClickable",
	)
}

func (obj TextView) GetMarqueeRepeatLimit() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getMarqueeRepeatLimit",
	)
}

func (obj TextView) GetMaxEms() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getMaxEms",
	)
}

func (obj TextView) GetMaxHeight() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getMaxHeight",
	)
}

func (obj TextView) GetMaxLines() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getMaxLines",
	)
}

func (obj TextView) GetMaxWidth() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getMaxWidth",
	)
}

func (obj TextView) GetMinEms() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getMinEms",
	)
}

func (obj TextView) GetMinHeight() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getMinHeight",
	)
}

func (obj TextView) GetMinLines() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getMinLines",
	)
}

func (obj TextView) GetMinWidth() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getMinWidth",
	)
}

func (obj TextView) GetOffsetForPosition(x_ float64, y_ float64) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getOffsetForPosition",
		android.Float(x_), android.Float(y_),
	)
}

func (obj TextView) GetPaintFlags() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getPaintFlags",
	)
}

func (obj TextView) GetPrivateImeOptions() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getPrivateImeOptions",
	)
}

func (obj TextView) GetSelectionEnd() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getSelectionEnd",
	)
}

func (obj TextView) GetSelectionStart() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getSelectionStart",
	)
}

func (obj TextView) GetShadowColor() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getShadowColor",
	)
}

func (obj TextView) GetShadowDx() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getShadowDx",
	)
}

func (obj TextView) GetShadowDy() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getShadowDy",
	)
}

func (obj TextView) GetShadowRadius() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getShadowRadius",
	)
}

func (obj TextView) GetShowSoftInputOnFocus() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getShowSoftInputOnFocus",
	)
}

func (obj TextView) GetText() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getText",
	)
}

func (obj TextView) GetTextScaleX() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getTextScaleX",
	)
}

func (obj TextView) GetTextSize() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getTextSize",
	)
}

func (obj TextView) GetTotalPaddingBottom() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getTotalPaddingBottom",
	)
}

func (obj TextView) GetTotalPaddingEnd() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getTotalPaddingEnd",
	)
}

func (obj TextView) GetTotalPaddingLeft() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getTotalPaddingLeft",
	)
}

func (obj TextView) GetTotalPaddingRight() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getTotalPaddingRight",
	)
}

func (obj TextView) GetTotalPaddingStart() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getTotalPaddingStart",
	)
}

func (obj TextView) GetTotalPaddingTop() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"getTotalPaddingTop",
	)
}

func (obj TextView) HasOverlappingRendering() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"hasOverlappingRendering",
	)
}

func (obj TextView) HasSelection() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"hasSelection",
	)
}

func (obj TextView) IsCursorVisible() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"isCursorVisible",
	)
}

func (obj TextView) IsInputMethodTarget() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"isInputMethodTarget",
	)
}

func (obj TextView) IsSuggestionsEnabled() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"isSuggestionsEnabled",
	)
}

func (obj TextView) IsTextSelectable() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"isTextSelectable",
	)
}

func (obj TextView) JumpDrawablesToCurrentState() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"jumpDrawablesToCurrentState",
	)
}

func (obj TextView) Length() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"length",
	)
}

func (obj TextView) MoveCursorToVisibleOffset() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"moveCursorToVisibleOffset",
	)
}

func (obj TextView) OnBeginBatchEdit() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"onBeginBatchEdit",
	)
}

func (obj TextView) OnCheckIsTextEditor() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"onCheckIsTextEditor",
	)
}

func (obj TextView) OnEditorAction(actionCode_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"onEditorAction",
		actionCode_,
	)
}

func (obj TextView) OnEndBatchEdit() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"onEndBatchEdit",
	)
}

func (obj TextView) OnFinishTemporaryDetach() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"onFinishTemporaryDetach",
	)
}

func (obj TextView) OnPreDraw() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"onPreDraw",
	)
}

func (obj TextView) OnRtlPropertiesChanged(layoutDirection_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"onRtlPropertiesChanged",
		layoutDirection_,
	)
}

func (obj TextView) OnScreenStateChanged(screenState_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"onScreenStateChanged",
		screenState_,
	)
}

func (obj TextView) OnStartTemporaryDetach() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"onStartTemporaryDetach",
	)
}

func (obj TextView) OnTextContextMenuItem(id_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"onTextContextMenuItem",
		id_,
	)
}

func (obj TextView) OnWindowFocusChanged(hasWindowFocus_ bool) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"onWindowFocusChanged",
		hasWindowFocus_,
	)
}

func (obj TextView) PerformLongClick() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"performLongClick",
	)
}

func (obj TextView) SendAccessibilityEvent(eventType_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"sendAccessibilityEvent",
		eventType_,
	)
}

func (obj TextView) SetAllCaps(allCaps_ bool) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setAllCaps",
		allCaps_,
	)
}

func (obj TextView) SetAutoLinkMask(mask_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setAutoLinkMask",
		mask_,
	)
}

func (obj TextView) SetCompoundDrawablePadding(pad_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setCompoundDrawablePadding",
		pad_,
	)
}

func (obj TextView) SetCompoundDrawablesRelativeWithIntrinsicBounds(start_ int, top_ int, end_ int, bottom_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setCompoundDrawablesRelativeWithIntrinsicBounds",
		start_, top_, end_, bottom_,
	)
}

func (obj TextView) SetCompoundDrawablesWithIntrinsicBounds(left_ int, top_ int, right_ int, bottom_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setCompoundDrawablesWithIntrinsicBounds",
		left_, top_, right_, bottom_,
	)
}

func (obj TextView) SetCursorVisible(visible_ bool) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setCursorVisible",
		visible_,
	)
}

func (obj TextView) SetElegantTextHeight(elegant_ bool) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setElegantTextHeight",
		elegant_,
	)
}

func (obj TextView) SetEms(ems_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setEms",
		ems_,
	)
}

func (obj TextView) SetEnabled(enabled_ bool) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setEnabled",
		enabled_,
	)
}

func (obj TextView) SetError(error_ string) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setError",
		error_,
	)
}

func (obj TextView) SetFilters() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setFilters",
	)
}

func (obj TextView) SetFontFeatureSettings(fontFeatureSettings_ string) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setFontFeatureSettings",
		fontFeatureSettings_,
	)
}

func (obj TextView) SetFreezesText(freezesText_ bool) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setFreezesText",
		freezesText_,
	)
}

func (obj TextView) SetGravity(gravity_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setGravity",
		gravity_,
	)
}

func (obj TextView) SetHeight(pixels_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setHeight",
		pixels_,
	)
}

func (obj TextView) SetHighlightColor(color_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setHighlightColor",
		color_,
	)
}

func (obj TextView) SetHint(hint_ string) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setHint",
		hint_,
	)
}

func (obj TextView) SetHint1i(resid_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setHint",
		resid_,
	)
}

func (obj TextView) SetHintTextColor(color_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setHintTextColor",
		color_,
	)
}

func (obj TextView) SetHorizontallyScrolling(whether_ bool) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setHorizontallyScrolling",
		whether_,
	)
}

func (obj TextView) SetImeActionLabel(label_ string, actionId_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setImeActionLabel",
		label_, actionId_,
	)
}

func (obj TextView) SetImeOptions(imeOptions_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setImeOptions",
		imeOptions_,
	)
}

func (obj TextView) SetIncludeFontPadding(includepad_ bool) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setIncludeFontPadding",
		includepad_,
	)
}

func (obj TextView) SetInputExtras(xmlResId_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setInputExtras",
		xmlResId_,
	)
}

func (obj TextView) SetInputType(type_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setInputType",
		type_,
	)
}

func (obj TextView) SetLetterSpacing(letterSpacing_ float64) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setLetterSpacing",
		android.Float(letterSpacing_),
	)
}

func (obj TextView) SetLineSpacing(add_ float64, mult_ float64) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setLineSpacing",
		android.Float(add_), android.Float(mult_),
	)
}

func (obj TextView) SetLines(lines_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setLines",
		lines_,
	)
}

func (obj TextView) SetLinkTextColor(color_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setLinkTextColor",
		color_,
	)
}

func (obj TextView) SetLinksClickable(whether_ bool) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setLinksClickable",
		whether_,
	)
}

func (obj TextView) SetMarqueeRepeatLimit(marqueeLimit_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setMarqueeRepeatLimit",
		marqueeLimit_,
	)
}

func (obj TextView) SetMaxEms(maxems_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setMaxEms",
		maxems_,
	)
}

func (obj TextView) SetMaxHeight(maxHeight_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setMaxHeight",
		maxHeight_,
	)
}

func (obj TextView) SetMaxLines(maxlines_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setMaxLines",
		maxlines_,
	)
}

func (obj TextView) SetMaxWidth(maxpixels_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setMaxWidth",
		maxpixels_,
	)
}

func (obj TextView) SetMinEms(minems_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setMinEms",
		minems_,
	)
}

func (obj TextView) SetMinHeight(minHeight_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setMinHeight",
		minHeight_,
	)
}

func (obj TextView) SetMinLines(minlines_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setMinLines",
		minlines_,
	)
}

func (obj TextView) SetMinWidth(minpixels_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setMinWidth",
		minpixels_,
	)
}

func (obj TextView) SetPadding(left_ int, top_ int, right_ int, bottom_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setPadding",
		left_, top_, right_, bottom_,
	)
}

func (obj TextView) SetPaddingRelative(start_ int, top_ int, end_ int, bottom_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setPaddingRelative",
		start_, top_, end_, bottom_,
	)
}

func (obj TextView) SetPaintFlags(flags_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setPaintFlags",
		flags_,
	)
}

func (obj TextView) SetPrivateImeOptions(type_ string) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setPrivateImeOptions",
		type_,
	)
}

func (obj TextView) SetRawInputType(type_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setRawInputType",
		type_,
	)
}

func (obj TextView) SetSelectAllOnFocus(selectAllOnFocus_ bool) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setSelectAllOnFocus",
		selectAllOnFocus_,
	)
}

func (obj TextView) SetSelected(selected_ bool) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setSelected",
		selected_,
	)
}

func (obj TextView) SetShadowLayer(radius_ float64, dx_ float64, dy_ float64, color_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setShadowLayer",
		android.Float(radius_), android.Float(dx_), android.Float(dy_), color_,
	)
}

func (obj TextView) SetShowSoftInputOnFocus(show_ bool) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setShowSoftInputOnFocus",
		show_,
	)
}

func (obj TextView) SetSingleLine() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setSingleLine",
	)
}

func (obj TextView) SetSingleLine1b(singleLine_ bool) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setSingleLine",
		singleLine_,
	)
}

func (obj TextView) SetText(resid_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setText",
		resid_,
	)
}

func (obj TextView) SetText2ii(start_ int, len_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setText",
		start_, len_,
	)
}

func (obj TextView) SetText1s(text_ string) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setText",
		text_,
	)
}

func (obj TextView) SetTextColor(color_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setTextColor",
		color_,
	)
}

func (obj TextView) SetTextIsSelectable(selectable_ bool) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setTextIsSelectable",
		selectable_,
	)
}

func (obj TextView) SetTextKeepState(text_ string) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setTextKeepState",
		text_,
	)
}

func (obj TextView) SetTextScaleX(size_ float64) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setTextScaleX",
		android.Float(size_),
	)
}

func (obj TextView) SetTextSize(size_ float64) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setTextSize",
		android.Float(size_),
	)
}

func (obj TextView) SetTextSize2if(unit_ int, size_ float64) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setTextSize",
		unit_, android.Float(size_),
	)
}

func (obj TextView) SetWidth(pixels_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"TextView",
		"setWidth",
		pixels_,
	)
}
