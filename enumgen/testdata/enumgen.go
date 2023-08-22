// Code generated by "enumgen -test.paniconexit0 -test.timeout=10m0s"; DO NOT EDIT.

package testdata

import (
	"errors"
	"strconv"
	"strings"
	"sync/atomic"

	"github.com/goki/enums/enums"
)

const _DaysName = "SundayMondayTuesdayWednesdayThursdayFridaySaturday"

var _DaysIndex = [...]uint8{0, 6, 12, 19, 28, 36, 42, 50}

const _DaysLowerName = "sundaymondaytuesdaywednesdaythursdayfridaysaturday"

// String returns the string representation
// of this Days value.
func (i Days) String() string {
	if i < 0 || i >= Days(len(_DaysIndex)-1) {
		return "Days(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _DaysName[_DaysIndex[i]:_DaysIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _DaysNoOp() {
	var x [1]struct{}
	_ = x[Sunday-(0)]
	_ = x[Monday-(1)]
	_ = x[Tuesday-(2)]
	_ = x[Wednesday-(3)]
	_ = x[Thursday-(4)]
	_ = x[Friday-(5)]
	_ = x[Saturday-(6)]
}

var _DaysValues = []Days{Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday}

var _DaysNameToValueMap = map[string]Days{
	_DaysName[0:6]:        Sunday,
	_DaysLowerName[0:6]:   Sunday,
	_DaysName[6:12]:       Monday,
	_DaysLowerName[6:12]:  Monday,
	_DaysName[12:19]:      Tuesday,
	_DaysLowerName[12:19]: Tuesday,
	_DaysName[19:28]:      Wednesday,
	_DaysLowerName[19:28]: Wednesday,
	_DaysName[28:36]:      Thursday,
	_DaysLowerName[28:36]: Thursday,
	_DaysName[36:42]:      Friday,
	_DaysLowerName[36:42]: Friday,
	_DaysName[42:50]:      Saturday,
	_DaysLowerName[42:50]: Saturday,
}

var _DaysNames = []string{
	_DaysName[0:6],
	_DaysName[6:12],
	_DaysName[12:19],
	_DaysName[19:28],
	_DaysName[28:36],
	_DaysName[36:42],
	_DaysName[42:50],
}

var _DaysDescMap = map[Days]string{
	0: _DaysDescs[0],
	1: _DaysDescs[1],
	2: _DaysDescs[2],
	3: _DaysDescs[3],
	4: _DaysDescs[4],
	5: _DaysDescs[5],
	6: _DaysDescs[6],
}

var _DaysDescs = []string{
	`Sunday is the first day of the week`,
	`Monday is the second day of the week`,
	`Tuesday is the third day of the week`,
	`Wednesday is the fourth day of the week`,
	`Thursday is the fifth day of the week`,
	`Friday is the sixth day of the week`,
	`Saturday is the seventh day of the week`,
}

// SetString sets the Days value from its
// string representation, and returns an
// error if the string is invalid.
func (i *Days) SetString(s string) error {
	if val, ok := _DaysNameToValueMap[s]; ok {
		*i = val
		return nil
	}

	if val, ok := _DaysNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " does not belong to Days values")
}

// Int64 returns the Days value as an int64.
func (i Days) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the Days value from an int64.
func (i *Days) SetInt64(in int64) {
	*i = Days(in)
}

// Desc returns the description of the Days value.
func (i Days) Desc() string {
	if str, ok := _DaysDescMap[i]; ok {
		return str
	}
	return i.String()
}

// DaysValues returns all possible values of
// the type Days. This slice will be in the
// same order as those returned by the Values,
// Strings, and Descs methods on Days.
func DaysValues() []Days {
	return _DaysValues
}

// Values returns all possible values of
// type Days. This slice will be in the
// same order as those returned by Strings and Descs.
func (i Days) Values() []enums.Enum {
	res := make([]enums.Enum, len(_DaysValues))
	for i, d := range _DaysValues {
		res[i] = &d
	}
	return res
}

// Strings returns the string representations of
// all possible values of type Days.
// This slice will be in the same order as
// those returned by Values and Descs.
func (i Days) Strings() []string {
	return _DaysNames
}

// Descs returns the descriptions of all
// possible values of type Days.
// This slice will be in the same order as
// those returned by Values and Strings.
func (i Days) Descs() []string {
	return _DaysDescs
}

// IsValid returns whether the value is a
// valid option for type Days.
func (i Days) IsValid() bool {
	for _, v := range _DaysValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Days) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Days) UnmarshalText(text []byte) error {
	return i.SetString(string(text))
}

const _StatesName = "EnabledDisabledFocusedHoveredActiveSelected"

var _StatesIndex = [...]uint8{0, 7, 15, 22, 29, 35, 43}

const _StatesLowerName = "enableddisabledfocusedhoveredactiveselected"

// String returns the string representation
// of this States value.
func (i States) String() string {
	if i < 0 || i >= States(len(_StatesIndex)-1) {
		return "States(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _StatesName[_StatesIndex[i]:_StatesIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _StatesNoOp() {
	var x [1]struct{}
	_ = x[Enabled-(0)]
	_ = x[Disabled-(1)]
	_ = x[Focused-(2)]
	_ = x[Hovered-(3)]
	_ = x[Active-(4)]
	_ = x[Selected-(5)]
}

var _StatesValues = []States{Enabled, Disabled, Focused, Hovered, Active, Selected}

var _StatesNameToValueMap = map[string]States{
	_StatesName[0:7]:        Enabled,
	_StatesLowerName[0:7]:   Enabled,
	_StatesName[7:15]:       Disabled,
	_StatesLowerName[7:15]:  Disabled,
	_StatesName[15:22]:      Focused,
	_StatesLowerName[15:22]: Focused,
	_StatesName[22:29]:      Hovered,
	_StatesLowerName[22:29]: Hovered,
	_StatesName[29:35]:      Active,
	_StatesLowerName[29:35]: Active,
	_StatesName[35:43]:      Selected,
	_StatesLowerName[35:43]: Selected,
}

var _StatesNames = []string{
	_StatesName[0:7],
	_StatesName[7:15],
	_StatesName[15:22],
	_StatesName[22:29],
	_StatesName[29:35],
	_StatesName[35:43],
}

var _StatesDescMap = map[States]string{
	0: _StatesDescs[0],
	1: _StatesDescs[1],
	2: _StatesDescs[2],
	3: _StatesDescs[3],
	4: _StatesDescs[4],
	5: _StatesDescs[5],
}

var _StatesDescs = []string{
	`Enabled indicates the widget is enabled`,
	`Disabled indicates the widget is disabled`,
	`Focused indicates the widget has keyboard focus`,
	`Hovered indicates the widget is being hovered over`,
	`Active indicates the widget is being interacted with`,
	`Selected indicates the widget is selected`,
}

// SetString sets the States value from its
// string representation, and returns an
// error if the string is invalid.
func (i *States) SetString(s string) error {
	if val, ok := _StatesNameToValueMap[s]; ok {
		*i = val
		return nil
	}

	if val, ok := _StatesNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " does not belong to States values")
}

// Int64 returns the States value as an int64.
func (i States) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the States value from an int64.
func (i *States) SetInt64(in int64) {
	*i = States(in)
}

// Desc returns the description of the States value.
func (i States) Desc() string {
	if str, ok := _StatesDescMap[i]; ok {
		return str
	}
	return i.String()
}

// StatesValues returns all possible values of
// the type States. This slice will be in the
// same order as those returned by the Values,
// Strings, and Descs methods on States.
func StatesValues() []States {
	return _StatesValues
}

// Values returns all possible values of
// type States. This slice will be in the
// same order as those returned by Strings and Descs.
func (i States) Values() []enums.Enum {
	res := make([]enums.Enum, len(_StatesValues))
	for i, d := range _StatesValues {
		res[i] = &d
	}
	return res
}

// Strings returns the string representations of
// all possible values of type States.
// This slice will be in the same order as
// those returned by Values and Descs.
func (i States) Strings() []string {
	return _StatesNames
}

// Descs returns the descriptions of all
// possible values of type States.
// This slice will be in the same order as
// those returned by Values and Strings.
func (i States) Descs() []string {
	return _StatesDescs
}

// IsValid returns whether the value is a
// valid option for type States.
func (i States) IsValid() bool {
	for _, v := range _StatesValues {
		if i == v {
			return true
		}
	}
	return false
}

// HasBitFlag returns whether these
// bit flags have the given bit flag set.
func (i *States) HasBitFlag(f enums.BitFlag) bool {
	return atomic.LoadInt64((*int64)(i))&(1<<uint32(f.Int64())) != 0
}

// HasBitFlag returns whether these
// bit flags have the given bit flag set.
func (i *States) SetBitFlag(on bool, f ...enums.BitFlag) {
	var mask int64
	for _, v := range f {
		mask |= 1 << v.Int64()
	}
	in := int64(*i)
	if on {
		in |= mask
		atomic.StoreInt64((*int64)(i), in)
	} else {
		in &^= mask
		atomic.StoreInt64((*int64)(i), in)
	}
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i States) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *States) UnmarshalText(text []byte) error {
	return i.SetString(string(text))
}