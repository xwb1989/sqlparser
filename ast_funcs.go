package sqlparser

//NewSelect is used to create a select statement
func NewSelect(comments Comments, exprs SelectExprs, selectOptions []string, from TableExprs, where *Where, groupBy GroupBy, having *Where) *Select {
	////var cache *bool
	////var distinct, straightJoinHint, sqlFoundRows bool
	//
	//for _, option := range selectOptions {
	//	switch strings.ToLower(option) {
	//	case DistinctStr:
	//		distinct = true
	//	case SQLCacheStr:
	//		truth := true
	//		cache = &truth
	//	case SQLNoCacheStr:
	//		truth := false
	//		cache = &truth
	//	case StraightJoinHint:
	//		straightJoinHint = true
	//		//case SQLCalcFoundRowsStr:
	//		//	sqlFoundRows = true
	//	}
	//}
	return &Select{
		//Cache:    cache,
		Comments: comments,
		//Distinct: distinct,
		//StraightJoinHint: straightJoinHint,
		//SQLCalcFoundRows: sqlFoundRows,
		SelectExprs: exprs,
		From:        from,
		Where:       where,
		GroupBy:     groupBy,
		Having:      having,
	}
}

//Unionize returns a UNION, either creating one or adding SELECT to an existing one
func Unionize(lhs, rhs SelectStatement, unionType string, by OrderBy, limit *Limit, lock string) *Union {

	union, isUnion := lhs.(*Union)
	if isUnion {
		union.UnionSelects = append(union.UnionSelects, &UnionSelect{UnionType: unionType, Statement: rhs})
		union.OrderBy = by
		union.Limit = limit
		union.Lock = lock
		return union
	}

	return &Union{FirstStatement: lhs, UnionSelects: []*UnionSelect{{UnionType: unionType, Statement: rhs}}, OrderBy: by, Limit: limit, Lock: lock}
}
