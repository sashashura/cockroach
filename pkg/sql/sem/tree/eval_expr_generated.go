// Copyright 2022 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

// Code generated by eval_gen.go. DO NOT EDIT.
// Regenerate this file with either of the following commands:
//
//   ./dev generate go
//   go generate ./pkg/sql/sem/tree
//
package tree

// ExprEvaluator is used to evaluate TypedExpr expressions.
type ExprEvaluator interface {
	EvalAllColumnsSelector(*AllColumnsSelector) (Datum, error)
	EvalAndExpr(*AndExpr) (Datum, error)
	EvalArray(*Array) (Datum, error)
	EvalArrayFlatten(*ArrayFlatten) (Datum, error)
	EvalBinaryExpr(*BinaryExpr) (Datum, error)
	EvalCaseExpr(*CaseExpr) (Datum, error)
	EvalCastExpr(*CastExpr) (Datum, error)
	EvalCoalesceExpr(*CoalesceExpr) (Datum, error)
	EvalCollateExpr(*CollateExpr) (Datum, error)
	EvalColumnAccessExpr(*ColumnAccessExpr) (Datum, error)
	EvalColumnItem(*ColumnItem) (Datum, error)
	EvalComparisonExpr(*ComparisonExpr) (Datum, error)
	EvalDefaultVal(*DefaultVal) (Datum, error)
	EvalFuncExpr(*FuncExpr) (Datum, error)
	EvalIfErrExpr(*IfErrExpr) (Datum, error)
	EvalIfExpr(*IfExpr) (Datum, error)
	EvalIndexedVar(*IndexedVar) (Datum, error)
	EvalIndirectionExpr(*IndirectionExpr) (Datum, error)
	EvalIsNotNullExpr(*IsNotNullExpr) (Datum, error)
	EvalIsNullExpr(*IsNullExpr) (Datum, error)
	EvalIsOfTypeExpr(*IsOfTypeExpr) (Datum, error)
	EvalNotExpr(*NotExpr) (Datum, error)
	EvalNullIfExpr(*NullIfExpr) (Datum, error)
	EvalOrExpr(*OrExpr) (Datum, error)
	EvalParenExpr(*ParenExpr) (Datum, error)
	EvalPlaceholder(*Placeholder) (Datum, error)
	EvalRangeCond(*RangeCond) (Datum, error)
	EvalSubquery(*Subquery) (Datum, error)
	EvalTuple(*Tuple) (Datum, error)
	EvalTupleStar(*TupleStar) (Datum, error)
	EvalTypedDummy(*TypedDummy) (Datum, error)
	EvalUnaryExpr(*UnaryExpr) (Datum, error)
	EvalUnqualifiedStar(UnqualifiedStar) (Datum, error)
	EvalUnresolvedName(*UnresolvedName) (Datum, error)
}


// Eval is part of the TypedExpr interface.
func (node *AllColumnsSelector) Eval(v ExprEvaluator) (Datum, error) {
	return v.EvalAllColumnsSelector(node)
}

// Eval is part of the TypedExpr interface.
func (node *AndExpr) Eval(v ExprEvaluator) (Datum, error) {
	return v.EvalAndExpr(node)
}

// Eval is part of the TypedExpr interface.
func (node *Array) Eval(v ExprEvaluator) (Datum, error) {
	return v.EvalArray(node)
}

// Eval is part of the TypedExpr interface.
func (node *ArrayFlatten) Eval(v ExprEvaluator) (Datum, error) {
	return v.EvalArrayFlatten(node)
}

// Eval is part of the TypedExpr interface.
func (node *BinaryExpr) Eval(v ExprEvaluator) (Datum, error) {
	return v.EvalBinaryExpr(node)
}

// Eval is part of the TypedExpr interface.
func (node *CaseExpr) Eval(v ExprEvaluator) (Datum, error) {
	return v.EvalCaseExpr(node)
}

// Eval is part of the TypedExpr interface.
func (node *CastExpr) Eval(v ExprEvaluator) (Datum, error) {
	return v.EvalCastExpr(node)
}

// Eval is part of the TypedExpr interface.
func (node *CoalesceExpr) Eval(v ExprEvaluator) (Datum, error) {
	return v.EvalCoalesceExpr(node)
}

// Eval is part of the TypedExpr interface.
func (node *CollateExpr) Eval(v ExprEvaluator) (Datum, error) {
	return v.EvalCollateExpr(node)
}

// Eval is part of the TypedExpr interface.
func (node *ColumnAccessExpr) Eval(v ExprEvaluator) (Datum, error) {
	return v.EvalColumnAccessExpr(node)
}

// Eval is part of the TypedExpr interface.
func (node *ColumnItem) Eval(v ExprEvaluator) (Datum, error) {
	return v.EvalColumnItem(node)
}

// Eval is part of the TypedExpr interface.
func (node *ComparisonExpr) Eval(v ExprEvaluator) (Datum, error) {
	return v.EvalComparisonExpr(node)
}

// Eval is part of the TypedExpr interface.
func (node *DArray) Eval(v ExprEvaluator) (Datum, error) {
	return node, nil
}

// Eval is part of the TypedExpr interface.
func (node *DBitArray) Eval(v ExprEvaluator) (Datum, error) {
	return node, nil
}

// Eval is part of the TypedExpr interface.
func (node *DBool) Eval(v ExprEvaluator) (Datum, error) {
	return node, nil
}

// Eval is part of the TypedExpr interface.
func (node *DBox2D) Eval(v ExprEvaluator) (Datum, error) {
	return node, nil
}

// Eval is part of the TypedExpr interface.
func (node *DBytes) Eval(v ExprEvaluator) (Datum, error) {
	return node, nil
}

// Eval is part of the TypedExpr interface.
func (node *DCollatedString) Eval(v ExprEvaluator) (Datum, error) {
	return node, nil
}

// Eval is part of the TypedExpr interface.
func (node *DDate) Eval(v ExprEvaluator) (Datum, error) {
	return node, nil
}

// Eval is part of the TypedExpr interface.
func (node *DDecimal) Eval(v ExprEvaluator) (Datum, error) {
	return node, nil
}

// Eval is part of the TypedExpr interface.
func (node *DEncodedKey) Eval(v ExprEvaluator) (Datum, error) {
	return node, nil
}

// Eval is part of the TypedExpr interface.
func (node *DEnum) Eval(v ExprEvaluator) (Datum, error) {
	return node, nil
}

// Eval is part of the TypedExpr interface.
func (node *DFloat) Eval(v ExprEvaluator) (Datum, error) {
	return node, nil
}

// Eval is part of the TypedExpr interface.
func (node *DGeography) Eval(v ExprEvaluator) (Datum, error) {
	return node, nil
}

// Eval is part of the TypedExpr interface.
func (node *DGeometry) Eval(v ExprEvaluator) (Datum, error) {
	return node, nil
}

// Eval is part of the TypedExpr interface.
func (node *DIPAddr) Eval(v ExprEvaluator) (Datum, error) {
	return node, nil
}

// Eval is part of the TypedExpr interface.
func (node *DInt) Eval(v ExprEvaluator) (Datum, error) {
	return node, nil
}

// Eval is part of the TypedExpr interface.
func (node *DInterval) Eval(v ExprEvaluator) (Datum, error) {
	return node, nil
}

// Eval is part of the TypedExpr interface.
func (node *DJSON) Eval(v ExprEvaluator) (Datum, error) {
	return node, nil
}

// Eval is part of the TypedExpr interface.
func (node *DOid) Eval(v ExprEvaluator) (Datum, error) {
	return node, nil
}

// Eval is part of the TypedExpr interface.
func (node *DOidWrapper) Eval(v ExprEvaluator) (Datum, error) {
	return node, nil
}

// Eval is part of the TypedExpr interface.
func (node *DString) Eval(v ExprEvaluator) (Datum, error) {
	return node, nil
}

// Eval is part of the TypedExpr interface.
func (node *DTime) Eval(v ExprEvaluator) (Datum, error) {
	return node, nil
}

// Eval is part of the TypedExpr interface.
func (node *DTimeTZ) Eval(v ExprEvaluator) (Datum, error) {
	return node, nil
}

// Eval is part of the TypedExpr interface.
func (node *DTimestamp) Eval(v ExprEvaluator) (Datum, error) {
	return node, nil
}

// Eval is part of the TypedExpr interface.
func (node *DTimestampTZ) Eval(v ExprEvaluator) (Datum, error) {
	return node, nil
}

// Eval is part of the TypedExpr interface.
func (node *DTuple) Eval(v ExprEvaluator) (Datum, error) {
	return node, nil
}

// Eval is part of the TypedExpr interface.
func (node *DUuid) Eval(v ExprEvaluator) (Datum, error) {
	return node, nil
}

// Eval is part of the TypedExpr interface.
func (node *DVoid) Eval(v ExprEvaluator) (Datum, error) {
	return node, nil
}

// Eval is part of the TypedExpr interface.
func (node *DefaultVal) Eval(v ExprEvaluator) (Datum, error) {
	return v.EvalDefaultVal(node)
}

// Eval is part of the TypedExpr interface.
func (node *FuncExpr) Eval(v ExprEvaluator) (Datum, error) {
	return v.EvalFuncExpr(node)
}

// Eval is part of the TypedExpr interface.
func (node *IfErrExpr) Eval(v ExprEvaluator) (Datum, error) {
	return v.EvalIfErrExpr(node)
}

// Eval is part of the TypedExpr interface.
func (node *IfExpr) Eval(v ExprEvaluator) (Datum, error) {
	return v.EvalIfExpr(node)
}

// Eval is part of the TypedExpr interface.
func (node *IndexedVar) Eval(v ExprEvaluator) (Datum, error) {
	return v.EvalIndexedVar(node)
}

// Eval is part of the TypedExpr interface.
func (node *IndirectionExpr) Eval(v ExprEvaluator) (Datum, error) {
	return v.EvalIndirectionExpr(node)
}

// Eval is part of the TypedExpr interface.
func (node *IsNotNullExpr) Eval(v ExprEvaluator) (Datum, error) {
	return v.EvalIsNotNullExpr(node)
}

// Eval is part of the TypedExpr interface.
func (node *IsNullExpr) Eval(v ExprEvaluator) (Datum, error) {
	return v.EvalIsNullExpr(node)
}

// Eval is part of the TypedExpr interface.
func (node *IsOfTypeExpr) Eval(v ExprEvaluator) (Datum, error) {
	return v.EvalIsOfTypeExpr(node)
}

// Eval is part of the TypedExpr interface.
func (node *NotExpr) Eval(v ExprEvaluator) (Datum, error) {
	return v.EvalNotExpr(node)
}

// Eval is part of the TypedExpr interface.
func (node *NullIfExpr) Eval(v ExprEvaluator) (Datum, error) {
	return v.EvalNullIfExpr(node)
}

// Eval is part of the TypedExpr interface.
func (node *OrExpr) Eval(v ExprEvaluator) (Datum, error) {
	return v.EvalOrExpr(node)
}

// Eval is part of the TypedExpr interface.
func (node *ParenExpr) Eval(v ExprEvaluator) (Datum, error) {
	return v.EvalParenExpr(node)
}

// Eval is part of the TypedExpr interface.
func (node *Placeholder) Eval(v ExprEvaluator) (Datum, error) {
	return v.EvalPlaceholder(node)
}

// Eval is part of the TypedExpr interface.
func (node *RangeCond) Eval(v ExprEvaluator) (Datum, error) {
	return v.EvalRangeCond(node)
}

// Eval is part of the TypedExpr interface.
func (node *Subquery) Eval(v ExprEvaluator) (Datum, error) {
	return v.EvalSubquery(node)
}

// Eval is part of the TypedExpr interface.
func (node *Tuple) Eval(v ExprEvaluator) (Datum, error) {
	return v.EvalTuple(node)
}

// Eval is part of the TypedExpr interface.
func (node *TupleStar) Eval(v ExprEvaluator) (Datum, error) {
	return v.EvalTupleStar(node)
}

// Eval is part of the TypedExpr interface.
func (node *TypedDummy) Eval(v ExprEvaluator) (Datum, error) {
	return v.EvalTypedDummy(node)
}

// Eval is part of the TypedExpr interface.
func (node *UnaryExpr) Eval(v ExprEvaluator) (Datum, error) {
	return v.EvalUnaryExpr(node)
}

// Eval is part of the TypedExpr interface.
func (node UnqualifiedStar) Eval(v ExprEvaluator) (Datum, error) {
	return v.EvalUnqualifiedStar(node)
}

// Eval is part of the TypedExpr interface.
func (node *UnresolvedName) Eval(v ExprEvaluator) (Datum, error) {
	return v.EvalUnresolvedName(node)
}

// Eval is part of the TypedExpr interface.
func (node dNull) Eval(v ExprEvaluator) (Datum, error) {
	return node, nil
}
