// Code generated by go-enum
// DO NOT EDIT!

package document

import (
	"fmt"
)

const (
	// DirectiveLocationUNKNOWN is a DirectiveLocation of type UNKNOWN
	DirectiveLocationUNKNOWN DirectiveLocation = iota
	// DirectiveLocationQUERY is a DirectiveLocation of type QUERY
	DirectiveLocationQUERY
	// DirectiveLocationMUTATION is a DirectiveLocation of type MUTATION
	DirectiveLocationMUTATION
	// DirectiveLocationSUBSCRIPTION is a DirectiveLocation of type SUBSCRIPTION
	DirectiveLocationSUBSCRIPTION
	// DirectiveLocationFIELD is a DirectiveLocation of type FIELD
	DirectiveLocationFIELD
	// DirectiveLocationFRAGMENT_DEFINITION is a DirectiveLocation of type FRAGMENT_DEFINITION
	DirectiveLocationFRAGMENT_DEFINITION
	// DirectiveLocationFRAGMENT_SPREAD is a DirectiveLocation of type FRAGMENT_SPREAD
	DirectiveLocationFRAGMENT_SPREAD
	// DirectiveLocationINLINE_FRAGMENT is a DirectiveLocation of type INLINE_FRAGMENT
	DirectiveLocationINLINE_FRAGMENT
	// DirectiveLocationSCHEMA is a DirectiveLocation of type SCHEMA
	DirectiveLocationSCHEMA
	// DirectiveLocationSCALAR is a DirectiveLocation of type SCALAR
	DirectiveLocationSCALAR
	// DirectiveLocationOBJECT is a DirectiveLocation of type OBJECT
	DirectiveLocationOBJECT
	// DirectiveLocationFIELD_DEFINITION is a DirectiveLocation of type FIELD_DEFINITION
	DirectiveLocationFIELD_DEFINITION
	// DirectiveLocationARGUMENT_DEFINITION is a DirectiveLocation of type ARGUMENT_DEFINITION
	DirectiveLocationARGUMENT_DEFINITION
	// DirectiveLocationINTERFACE is a DirectiveLocation of type INTERFACE
	DirectiveLocationINTERFACE
	// DirectiveLocationUNION is a DirectiveLocation of type UNION
	DirectiveLocationUNION
	// DirectiveLocationENUM is a DirectiveLocation of type ENUM
	DirectiveLocationENUM
	// DirectiveLocationENUM_VALUE is a DirectiveLocation of type ENUM_VALUE
	DirectiveLocationENUM_VALUE
	// DirectiveLocationINPUT_OBJECT is a DirectiveLocation of type INPUT_OBJECT
	DirectiveLocationINPUT_OBJECT
	// DirectiveLocationINPUT_FIELD_DEFINITION is a DirectiveLocation of type INPUT_FIELD_DEFINITION
	DirectiveLocationINPUT_FIELD_DEFINITION
)

const _DirectiveLocationName = "UNKNOWNQUERYMUTATIONSUBSCRIPTIONFIELDFRAGMENT_DEFINITIONFRAGMENT_SPREADINLINE_FRAGMENTSCHEMASCALAROBJECTFIELD_DEFINITIONARGUMENT_DEFINITIONINTERFACEUNIONENUMENUM_VALUEINPUT_OBJECTINPUT_FIELD_DEFINITION"

var _DirectiveLocationMap = map[DirectiveLocation]string{
	0:  _DirectiveLocationName[0:7],
	1:  _DirectiveLocationName[7:12],
	2:  _DirectiveLocationName[12:20],
	3:  _DirectiveLocationName[20:32],
	4:  _DirectiveLocationName[32:37],
	5:  _DirectiveLocationName[37:56],
	6:  _DirectiveLocationName[56:71],
	7:  _DirectiveLocationName[71:86],
	8:  _DirectiveLocationName[86:92],
	9:  _DirectiveLocationName[92:98],
	10: _DirectiveLocationName[98:104],
	11: _DirectiveLocationName[104:120],
	12: _DirectiveLocationName[120:139],
	13: _DirectiveLocationName[139:148],
	14: _DirectiveLocationName[148:153],
	15: _DirectiveLocationName[153:157],
	16: _DirectiveLocationName[157:167],
	17: _DirectiveLocationName[167:179],
	18: _DirectiveLocationName[179:201],
}

// String implements the Stringer interface.
func (x DirectiveLocation) String() string {
	if str, ok := _DirectiveLocationMap[x]; ok {
		return str
	}
	return fmt.Sprintf("DirectiveLocation(%d)", x)
}

var _DirectiveLocationValue = map[string]DirectiveLocation{
	_DirectiveLocationName[0:7]:     0,
	_DirectiveLocationName[7:12]:    1,
	_DirectiveLocationName[12:20]:   2,
	_DirectiveLocationName[20:32]:   3,
	_DirectiveLocationName[32:37]:   4,
	_DirectiveLocationName[37:56]:   5,
	_DirectiveLocationName[56:71]:   6,
	_DirectiveLocationName[71:86]:   7,
	_DirectiveLocationName[86:92]:   8,
	_DirectiveLocationName[92:98]:   9,
	_DirectiveLocationName[98:104]:  10,
	_DirectiveLocationName[104:120]: 11,
	_DirectiveLocationName[120:139]: 12,
	_DirectiveLocationName[139:148]: 13,
	_DirectiveLocationName[148:153]: 14,
	_DirectiveLocationName[153:157]: 15,
	_DirectiveLocationName[157:167]: 16,
	_DirectiveLocationName[167:179]: 17,
	_DirectiveLocationName[179:201]: 18,
}

// ParseDirectiveLocation attempts to convert a string to a DirectiveLocation
func ParseDirectiveLocation(name []byte) (DirectiveLocation, error) {
	if x, ok := _DirectiveLocationValue[string(name)]; ok {
		return x, nil
	}
	return DirectiveLocation(0), fmt.Errorf("%s is not a valid DirectiveLocation", name)
}
