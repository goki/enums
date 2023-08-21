// Code generated by "enumgen "; DO NOT EDIT.

package testdata

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _MyEnumName = "SundayMondayTuesdayWednesdayThursdayFridaySaturday"

var _MyEnumIndex = [...]uint8{0, 6, 12, 19, 28, 36, 42, 50}

const _MyEnumLowerName = "sundaymondaytuesdaywednesdaythursdayfridaysaturday"

func (i MyEnum) String() string {
	if i < 0 || i >= MyEnum(len(_MyEnumIndex)-1) {
		return "MyEnum(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MyEnumName[_MyEnumIndex[i]:_MyEnumIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _MyEnumNoOp() {
	var x [1]struct{}
	_ = x[Sunday-(0)]
	_ = x[Monday-(1)]
	_ = x[Tuesday-(2)]
	_ = x[Wednesday-(3)]
	_ = x[Thursday-(4)]
	_ = x[Friday-(5)]
	_ = x[Saturday-(6)]
}

var _MyEnumValues = []MyEnum{Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday}

var _MyEnumNameToValueMap = map[string]MyEnum{
	_MyEnumName[0:6]:        Sunday,
	_MyEnumLowerName[0:6]:   Sunday,
	_MyEnumName[6:12]:       Monday,
	_MyEnumLowerName[6:12]:  Monday,
	_MyEnumName[12:19]:      Tuesday,
	_MyEnumLowerName[12:19]: Tuesday,
	_MyEnumName[19:28]:      Wednesday,
	_MyEnumLowerName[19:28]: Wednesday,
	_MyEnumName[28:36]:      Thursday,
	_MyEnumLowerName[28:36]: Thursday,
	_MyEnumName[36:42]:      Friday,
	_MyEnumLowerName[36:42]: Friday,
	_MyEnumName[42:50]:      Saturday,
	_MyEnumLowerName[42:50]: Saturday,
}

var _MyEnumNames = []string{
	_MyEnumName[0:6],
	_MyEnumName[6:12],
	_MyEnumName[12:19],
	_MyEnumName[19:28],
	_MyEnumName[28:36],
	_MyEnumName[36:42],
	_MyEnumName[42:50],
}

// SetString sets the enum value from its
// string representation, and returns an
// error if the string is invalid.
func (i *MyEnum) SetString(s string) error {
	if val, ok := _MyEnumNameToValueMap[s]; ok {
		*i = val
		return nil
	}

	if val, ok := _MyEnumNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return 0, fmt.Errorf("%s does not belong to MyEnum values", s)
}

// Values returns all possible values this
// enum type has. This slice will be in the
// same order as those returned by Strings and Descs.
func (i MyEnum) Values() []MyEnum {
	return _MyEnumValues
}

// Strings returns the string encodings of
// all possible values this enum type has.
// This slice will be in the same order as
// those returned by Values and Descs.
func (i MyEnum) Strings() []string {
	strs := make([]string, len(_MyEnumNames))
	copy(strs, _MyEnumNames)
	return strs
}

// IsValid returns whether the value is a
// valid option for its enum type.
func (i MyEnum) IsValid() bool {
	for _, v := range _MyEnumValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for MyEnum
func (i MyEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for MyEnum
func (i *MyEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("MyEnum should be a string, got %s", data)
	}

	var err error
	*i, err = MyEnumString(s)
	return err
}

// MarshalText implements the encoding.TextMarshaler interface for MyEnum
func (i MyEnum) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for MyEnum
func (i *MyEnum) UnmarshalText(text []byte) error {
	var err error
	*i, err = MyEnumString(string(text))
	return err
}

const _MyBitEnumName = "AppleOrangePeachBlueberryGrapefruitStrawberry"

var _MyBitEnumIndex = [...]uint8{0, 5, 11, 16, 25, 35, 45}

const _MyBitEnumLowerName = "appleorangepeachblueberrygrapefruitstrawberry"

func (i MyBitEnum) String() string {
	if i < 0 || i >= MyBitEnum(len(_MyBitEnumIndex)-1) {
		return "MyBitEnum(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MyBitEnumName[_MyBitEnumIndex[i]:_MyBitEnumIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _MyBitEnumNoOp() {
	var x [1]struct{}
	_ = x[Apple-(0)]
	_ = x[Orange-(1)]
	_ = x[Peach-(2)]
	_ = x[Blueberry-(3)]
	_ = x[Grapefruit-(4)]
	_ = x[Strawberry-(5)]
}

var _MyBitEnumValues = []MyBitEnum{Apple, Orange, Peach, Blueberry, Grapefruit, Strawberry}

var _MyBitEnumNameToValueMap = map[string]MyBitEnum{
	_MyBitEnumName[0:5]:        Apple,
	_MyBitEnumLowerName[0:5]:   Apple,
	_MyBitEnumName[5:11]:       Orange,
	_MyBitEnumLowerName[5:11]:  Orange,
	_MyBitEnumName[11:16]:      Peach,
	_MyBitEnumLowerName[11:16]: Peach,
	_MyBitEnumName[16:25]:      Blueberry,
	_MyBitEnumLowerName[16:25]: Blueberry,
	_MyBitEnumName[25:35]:      Grapefruit,
	_MyBitEnumLowerName[25:35]: Grapefruit,
	_MyBitEnumName[35:45]:      Strawberry,
	_MyBitEnumLowerName[35:45]: Strawberry,
}

var _MyBitEnumNames = []string{
	_MyBitEnumName[0:5],
	_MyBitEnumName[5:11],
	_MyBitEnumName[11:16],
	_MyBitEnumName[16:25],
	_MyBitEnumName[25:35],
	_MyBitEnumName[35:45],
}

// SetString sets the enum value from its
// string representation, and returns an
// error if the string is invalid.
func (i *MyBitEnum) SetString(s string) error {
	if val, ok := _MyBitEnumNameToValueMap[s]; ok {
		*i = val
		return nil
	}

	if val, ok := _MyBitEnumNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return 0, fmt.Errorf("%s does not belong to MyBitEnum values", s)
}

// Values returns all possible values this
// enum type has. This slice will be in the
// same order as those returned by Strings and Descs.
func (i MyBitEnum) Values() []MyBitEnum {
	return _MyBitEnumValues
}

// Strings returns the string encodings of
// all possible values this enum type has.
// This slice will be in the same order as
// those returned by Values and Descs.
func (i MyBitEnum) Strings() []string {
	strs := make([]string, len(_MyBitEnumNames))
	copy(strs, _MyBitEnumNames)
	return strs
}

// IsValid returns whether the value is a
// valid option for its enum type.
func (i MyBitEnum) IsValid() bool {
	for _, v := range _MyBitEnumValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for MyBitEnum
func (i MyBitEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for MyBitEnum
func (i *MyBitEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("MyBitEnum should be a string, got %s", data)
	}

	var err error
	*i, err = MyBitEnumString(s)
	return err
}

// MarshalText implements the encoding.TextMarshaler interface for MyBitEnum
func (i MyBitEnum) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for MyBitEnum
func (i *MyBitEnum) UnmarshalText(text []byte) error {
	var err error
	*i, err = MyBitEnumString(string(text))
	return err
}