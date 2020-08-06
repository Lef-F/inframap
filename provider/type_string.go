// Code generated by "enumer -type=Type -transform=lower -output=type_string.go"; DO NOT EDIT.

package provider

import (
	"fmt"
)

const _TypeName = "rawawsflexibleengineopenstackgoogleazurerm"

var _TypeIndex = [...]uint8{0, 3, 6, 20, 29, 35, 42}

const _TypeLowerName = "rawawsflexibleengineopenstackgoogleazurerm"

func (i Type) String() string {
	if i < 0 || i >= Type(len(_TypeIndex)-1) {
		return fmt.Sprintf("Type(%d)", i)
	}
	return _TypeName[_TypeIndex[i]:_TypeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _TypeNoOp() {
	var x [1]struct{}
	_ = x[Raw-(0)]
	_ = x[AWS-(1)]
	_ = x[FlexibleEngine-(2)]
	_ = x[OpenStack-(3)]
	_ = x[Google-(4)]
	_ = x[Azurerm-(5)]
}

var _TypeValues = []Type{Raw, AWS, FlexibleEngine, OpenStack, Google, Azurerm}

var _TypeNameToValueMap = map[string]Type{
	_TypeName[0:3]:        Raw,
	_TypeLowerName[0:3]:   Raw,
	_TypeName[3:6]:        AWS,
	_TypeLowerName[3:6]:   AWS,
	_TypeName[6:20]:       FlexibleEngine,
	_TypeLowerName[6:20]:  FlexibleEngine,
	_TypeName[20:29]:      OpenStack,
	_TypeLowerName[20:29]: OpenStack,
	_TypeName[29:35]:      Google,
	_TypeLowerName[29:35]: Google,
	_TypeName[35:42]:      Azurerm,
	_TypeLowerName[35:42]: Azurerm,
}

var _TypeNames = []string{
	_TypeName[0:3],
	_TypeName[3:6],
	_TypeName[6:20],
	_TypeName[20:29],
	_TypeName[29:35],
	_TypeName[35:42],
}

// TypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func TypeString(s string) (Type, error) {
	if val, ok := _TypeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Type values", s)
}

// TypeValues returns all values of the enum
func TypeValues() []Type {
	return _TypeValues
}

// TypeStrings returns a slice of all String values of the enum
func TypeStrings() []string {
	strs := make([]string, len(_TypeNames))
	copy(strs, _TypeNames)
	return strs
}

// IsAType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Type) IsAType() bool {
	for _, v := range _TypeValues {
		if i == v {
			return true
		}
	}
	return false
}
