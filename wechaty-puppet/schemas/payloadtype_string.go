// Code generated by "stringer -type=PayloadType"; DO NOT EDIT.

package schemas

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[PayloadTypeUnknown-0]
}

const _PayloadType_name = "PayloadTypeUnknown"

var _PayloadType_index = [...]uint8{0, 18}

func (i PayloadType) String() string {
	if i < 0 || i >= PayloadType(len(_PayloadType_index)-1) {
		return "PayloadType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _PayloadType_name[_PayloadType_index[i]:_PayloadType_index[i+1]]
}
