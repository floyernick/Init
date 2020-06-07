package storage

import (
	"fmt"
)

type SelectQuery struct {
	storage Service
	clause  string
	params  []interface{}
	offset  *int
	limit   *int
}

func (builder *SelectQuery) Equals(field string, value interface{}) SelectQuerier {
	builder.params = append(builder.params, value)
	builder.clause += fmt.Sprintf("%s = $%d ", field, len(builder.params))
	return builder
}

func (builder *SelectQuery) NotEquals(field string, value interface{}) SelectQuerier {
	builder.params = append(builder.params, value)
	builder.clause += fmt.Sprintf("%s != $%d ", field, len(builder.params))
	return builder
}

func (builder *SelectQuery) Greater(field string, value interface{}) SelectQuerier {
	builder.params = append(builder.params, value)
	builder.clause += fmt.Sprintf("%s > $%d ", field, len(builder.params))
	return builder
}

func (builder *SelectQuery) GreaterOrEquals(field string, value interface{}) SelectQuerier {
	builder.params = append(builder.params, value)
	builder.clause += fmt.Sprintf("%s >= $%d ", field, len(builder.params))
	return builder
}

func (builder *SelectQuery) Less(field string, value interface{}) SelectQuerier {
	builder.params = append(builder.params, value)
	builder.clause += fmt.Sprintf("%s < $%d ", field, len(builder.params))
	return builder
}

func (builder *SelectQuery) LessOrEquals(field string, value interface{}) SelectQuerier {
	builder.params = append(builder.params, value)
	builder.clause += fmt.Sprintf("%s <= $%d ", field, len(builder.params))
	return builder
}

func (builder *SelectQuery) Contains(field string, value interface{}) SelectQuerier {
	builder.params = append(builder.params, value)
	builder.clause += fmt.Sprintf("$%d = ANY(%s) ", len(builder.params), field)
	return builder
}

func (builder *SelectQuery) Like(field string, value interface{}) SelectQuerier {
	builder.params = append(builder.params, value)
	builder.clause += fmt.Sprintf("%s ILIKE '%%' || $%d || '%%' ", field, len(builder.params))
	return builder
}

func (builder *SelectQuery) In(field string, value interface{}) SelectQuerier {
	builder.params = append(builder.params, value)
	builder.clause += fmt.Sprintf("%s = ANY(%d) ", field, len(builder.params))
	return builder
}

func (builder *SelectQuery) NotIn(field string, value interface{}) SelectQuerier {
	builder.params = append(builder.params, value)
	builder.clause += fmt.Sprintf("%s = ANY(%d) ", field, len(builder.params))
	return builder
}

func (builder *SelectQuery) And() SelectQuerier {
	builder.clause += "AND "
	return builder
}

func (builder *SelectQuery) Or() SelectQuerier {
	builder.clause += "OR "
	return builder
}

func (builder *SelectQuery) Group() SelectQuerier {
	builder.clause += "("
	return builder
}

func (builder *SelectQuery) EndGroup() SelectQuerier {
	builder.clause += ") "
	return builder
}

func (builder *SelectQuery) Paginate(offset int, limit int) SelectQuerier {
	return builder
}

func (builder *SelectQuery) Order(field string, value string) SelectQuerier {
	return builder
}

func (builder *SelectQuery) formatQuery(query string) string {
	return fmt.Sprintf("%s WHERE %s", query, builder.clause)
}
