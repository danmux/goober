// Code generated by "stringer -type=typeID"; DO NOT EDIT.

package goblin

import "strconv"

const (
	_typeID_name_0 = "tBooltInttUinttFloattBytestStringtComplextInterfacetSlicetMaptStruct"
	_typeID_name_1 = "tWireTypetArrayTypetCommonTypetSliceTypetStructTypetFieldType"
	_typeID_name_2 = "tMapType"
	_typeID_name_3 = "minUserType"
)

var (
	_typeID_index_0 = [...]uint8{0, 5, 9, 14, 20, 26, 33, 41, 51, 57, 61, 68}
	_typeID_index_1 = [...]uint8{0, 9, 19, 30, 40, 51, 61}
)

func (i typeID) String() string {
	switch {
	case 1 <= i && i <= 11:
		i -= 1
		return _typeID_name_0[_typeID_index_0[i]:_typeID_index_0[i+1]]
	case 16 <= i && i <= 21:
		i -= 16
		return _typeID_name_1[_typeID_index_1[i]:_typeID_index_1[i+1]]
	case i == 23:
		return _typeID_name_2
	case i == 30:
		return _typeID_name_3
	default:
		return "typeID(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}