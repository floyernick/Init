package database

import (
	"fmt"
)

func (db DB) GetBuilder() queryBuilder {
	builder := queryBuilder{
		db: db,
	}
	return builder
}

type queryBuilder struct {
	db             DB
	clause         string
	params         []interface{}
	offset         *int
	limit          *int
	orderField     *string
	orderDirection *string
}

func (builder queryBuilder) Equals(field string, value interface{}) queryBuilder {
	builder.params = append(builder.params, value)
	builder.clause += fmt.Sprintf("%s = $%d ", field, len(builder.params))
	return builder
}

func (builder queryBuilder) NotEquals(field string, value interface{}) queryBuilder {
	builder.params = append(builder.params, value)
	builder.clause += fmt.Sprintf("%s != $%d ", field, len(builder.params))
	return builder
}

func (builder queryBuilder) Greater(field string, value interface{}) queryBuilder {
	builder.params = append(builder.params, value)
	builder.clause += fmt.Sprintf("%s > $%d ", field, len(builder.params))
	return builder
}

func (builder queryBuilder) GreaterOrEquals(field string, value interface{}) queryBuilder {
	builder.params = append(builder.params, value)
	builder.clause += fmt.Sprintf("%s >= $%d ", field, len(builder.params))
	return builder
}

func (builder queryBuilder) Less(field string, value interface{}) queryBuilder {
	builder.params = append(builder.params, value)
	builder.clause += fmt.Sprintf("%s < $%d ", field, len(builder.params))
	return builder
}

func (builder queryBuilder) LessOrEquals(field string, value interface{}) queryBuilder {
	builder.params = append(builder.params, value)
	builder.clause += fmt.Sprintf("%s <= $%d ", field, len(builder.params))
	return builder
}

func (builder queryBuilder) Contains(field string, value interface{}) queryBuilder {
	builder.params = append(builder.params, value)
	builder.clause += fmt.Sprintf("$%d = ANY(%s) ", len(builder.params), field)
	return builder
}

func (builder queryBuilder) Like(field string, value interface{}) queryBuilder {
	builder.params = append(builder.params, value)
	builder.clause += fmt.Sprintf("%s ILIKE '%%' || $%d || '%%' ", field, len(builder.params))
	return builder
}

func (builder queryBuilder) In(field string, value interface{}) queryBuilder {
	builder.params = append(builder.params, value)
	builder.clause += fmt.Sprintf("%s = ANY(%d) ", field, len(builder.params))
	return builder
}

func (builder queryBuilder) NotIn(field string, value interface{}) queryBuilder {
	builder.params = append(builder.params, value)
	builder.clause += fmt.Sprintf("%s = ANY(%d) ", field, len(builder.params))
	return builder
}

func (builder queryBuilder) And() queryBuilder {
	builder.clause += "AND "
	return builder
}

func (builder queryBuilder) Or() queryBuilder {
	builder.clause += "OR "
	return builder
}

func (builder queryBuilder) Group() queryBuilder {
	builder.clause += "("
	return builder
}

func (builder queryBuilder) EndGroup() queryBuilder {
	builder.clause += ") "
	return builder
}

func (builder queryBuilder) Offset(offset int) queryBuilder {
	builder.offset = &offset
	return builder
}

func (builder queryBuilder) Limit(limit int) queryBuilder {
	builder.limit = &limit
	return builder
}

func (builder queryBuilder) OrderBy(field string, direction string) queryBuilder {
	builder.orderField = &field
	builder.orderDirection = &direction
	return builder
}

func (builder queryBuilder) formatQuery(query string) string {
	if len(builder.params) != 0 {
		query += fmt.Sprintf(" WHERE %s", builder.clause)
	}
	if builder.orderField != nil && builder.orderDirection != nil {
		query += fmt.Sprintf(" ORDER BY %s %s", *builder.orderField, *builder.orderDirection)
	}
	if builder.offset != nil {
		query += fmt.Sprintf(" OFFSET %d", *builder.offset)
	}
	if builder.limit != nil {
		query += fmt.Sprintf(" LIMIT %d", *builder.limit)
	}
	return query
}
