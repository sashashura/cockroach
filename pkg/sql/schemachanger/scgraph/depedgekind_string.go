// Code generated by "stringer"; DO NOT EDIT.

package scgraph

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[HappensAfter-1]
	_ = x[SameStage-2]
}

const _DepEdgeKind_name = "HappensAfterSameStage"

var _DepEdgeKind_index = [...]uint8{0, 12, 21}

func (i DepEdgeKind) String() string {
	i -= 1
	if i < 0 || i >= DepEdgeKind(len(_DepEdgeKind_index)-1) {
		return "DepEdgeKind(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _DepEdgeKind_name[_DepEdgeKind_index[i]:_DepEdgeKind_index[i+1]]
}