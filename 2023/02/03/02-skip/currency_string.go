// Code generated by "stringer -type=Currency"; DO NOT EDIT.

package skip

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[USD-0]
	_ = x[MXN-1]
	_ = x[PHP-2]
}

const _Currency_name = "USDMXNPHP"

var _Currency_index = [...]uint8{0, 3, 6, 9}

func (i Currency) String() string {
	if i < 0 || i >= Currency(len(_Currency_index)-1) {
		return "Currency(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Currency_name[_Currency_index[i]:_Currency_index[i+1]]
}
