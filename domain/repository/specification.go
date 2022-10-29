package repository

import (
	"fmt"
	"reflect"
	"strings"
)

type Relation string

const (
	Equals      Relation = "="
	GreaterThan Relation = ">"
	GreaterOrEq Relation = ">="
	LessThan    Relation = "<"
	LessOrEq    Relation = "<="
)

type Specification interface {
	Query() string
	Value() []any
}

func param(r any) any {
	t := reflect.TypeOf(r)
	valType := t.String()
	val := r

	switch valType {
	case "string":
		if val != "?" {
			val = "'" + val.(string) + "'"
		}
	default:

	}
	return val
}

type EqualsSpecification struct {
	field string
	value any
}
type FunctionSpecification func() string

func (fs FunctionSpecification) Query() string {
	return fs()
}

func (fs FunctionSpecification) Value() []interface{} {
	return nil
}

type AndSpecification struct {
	specifications []Specification
}

func (a AndSpecification) Query() string {
	var queries []string
	for _, specification := range a.specifications {
		queries = append(queries, specification.Query())
	}
	query := strings.Join(queries, " AND ")
	return fmt.Sprintf("(%s)", query)
}
func (a AndSpecification) Value() []interface{} {
	var values []interface{}
	for _, specification := range a.specifications {
		values = append(values, specification.Value()...)
	}
	return values
}

type OrSpecification struct {
	specifications []Specification
}

func (o OrSpecification) Query() string {
	var queries []string
	for _, specification := range o.specifications {
		queries = append(queries, specification.Query())
	}
	query := strings.Join(queries, " OR ")
	return fmt.Sprintf("(%s)", query)
}
func (o OrSpecification) Value() []interface{} {
	var values []interface{}
	for _, specification := range o.specifications {
		values = append(values, specification.Value()...)
	}
	return values
}

type RelationSpecification struct {
	field    string
	operator Relation
	value    any
}

func (r RelationSpecification) Query() string {

	return fmt.Sprintf(" %s %s %v ", r.field, r.operator, param(r.value))
}

func (r RelationSpecification) Value() []any {
	return []any{
		r.field,
		r.operator,
		r.value,
	}
}

func NewEqualsSpecification(field string, value any) EqualsSpecification {
	return EqualsSpecification{
		field: field,
		value: value,
	}
}
func (e EqualsSpecification) Query() string {
	return fmt.Sprintf("%s = %v", e.field, param(e.value))
}
func (e EqualsSpecification) Value() []any {
	return []any{e.value}
}

func NewRelationSpecification(field string, op Relation, value any) RelationSpecification {
	return RelationSpecification{
		field:    field,
		operator: op,
		value:    value,
	}
}
func NewAndSpecification(parts ...Specification) AndSpecification {
	return AndSpecification{specifications: parts}
}
func NewOrSpecification(parts ...Specification) OrSpecification {
	return OrSpecification{
		specifications: parts,
	}
}
