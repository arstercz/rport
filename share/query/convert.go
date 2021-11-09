package query

import (
	"fmt"
	"strings"
)

func ConvertListOptionsToQuery(lo *ListOptions, q string) (qOut string, params []interface{}) {
	return AppendOptionsToQuery(lo, q, nil)
}

func ConvertRetrieveOptionsToQuery(ro *RetrieveOptions, q string) string {
	qOut := ReplaceStarSelect(ro.Fields, q)

	return qOut
}

func AppendOptionsToQuery(o *ListOptions, q string, params []interface{}) (string, []interface{}) {
	q, params = addWhere(o.Filters, q, params)
	q = addOrderBy(o.Sorts, q)
	q = ReplaceStarSelect(o.Fields, q)
	q, params = addLimitOffset(o.Pagination, q, params)

	return q, params
}

func addWhere(filterOptions []FilterOption, q string, params []interface{}) (string, []interface{}) {
	if len(filterOptions) == 0 {
		return q, params
	}

	whereParts := make([]string, 0, len(filterOptions))
	for i, fo := range filterOptions {
		if IsLimitFilter(fo) {
			continue
		}
		if len(filterOptions[i].Values) == 1 {
			whereParts = append(whereParts, fmt.Sprintf("%s %s ?", filterOptions[i].Column, filterOptions[i].Operator.Code()))
			params = append(params, filterOptions[i].Values[0])
		} else {
			orParts := make([]string, 0, len(filterOptions[i].Values))
			for y := range filterOptions[i].Values {
				orParts = append(orParts, fmt.Sprintf("%s %s ?", filterOptions[i].Column, filterOptions[i].Operator.Code()))
				params = append(params, filterOptions[i].Values[y])
			}

			whereParts = append(whereParts, fmt.Sprintf("(%s)", strings.Join(orParts, " OR ")))
		}
	}

	concat := " WHERE "
	qUpper := strings.ToUpper(q)
	if strings.Contains(qUpper, " WHERE ") {
		concat = " AND "
	}
	q += concat + strings.Join(whereParts, " AND ")

	return q, params
}

func addLimit(filterOptions []FilterOption, q string) string {
	limit := ""
	for _, fo := range filterOptions {
		if !IsLimitFilter(fo) {
			continue
		}
		limit = limit + fmt.Sprintf(" LIMIT %s", fo.Values[0])
	}

	return q + limit
}

func addOrderBy(sortOptions []SortOption, q string) string {
	if len(sortOptions) == 0 {
		return q
	}
	orderByValues := make([]string, 0, len(sortOptions))
	for i := range sortOptions {
		direction := "ASC"
		if !sortOptions[i].IsASC {
			direction = "DESC"
		}
		orderByValues = append(orderByValues, fmt.Sprintf("%s %s", sortOptions[i].Column, direction))
	}
	if len(orderByValues) > 0 {
		q += " ORDER BY " + strings.Join(orderByValues, ", ")
	}

	return q
}

func ReplaceStarSelect(fieldOptions []FieldsOption, q string) string {
	if !strings.HasPrefix(strings.ToUpper(q), "SELECT * ") {
		return q
	}
	if len(fieldOptions) == 0 {
		return q
	}

	fields := []string{}
	for _, fo := range fieldOptions {
		for _, field := range fo.Fields {
			fields = append(fields, fmt.Sprintf("%s.%s", fo.Resource, field))
		}
	}

	return strings.Replace(q, "*", strings.Join(fields, ", "), 1)
}

func addLimitOffset(pagination *Pagination, q string, params []interface{}) (string, []interface{}) {
	if pagination == nil {
		return q, params
	}

	q += " LIMIT ? OFFSET ?"
	params = append(params, pagination.Limit, pagination.Offset)

	return q, params
}
