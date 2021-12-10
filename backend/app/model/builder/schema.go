package builder

import "skema/app/model"

type SchemaBuilder struct {
	Id    uint
	Owner uint
	Name  string
}

func (b *SchemaBuilder) Build() *model.Schema {
	return &model.Schema{
		Id:    b.Id,
		Owner: b.Owner,
		Name:  b.Name,
	}
}

type SchemaMother struct{}

func (m *SchemaMother) Obj0() *model.Schema {
	builder := SchemaBuilder{
		Id:    1,
		Owner: 2,
		Name:  "Uml class schema",
	}

	return builder.Build()
}

func (m *SchemaMother) Obj1() *model.Schema {
	builder := SchemaBuilder{
		Id:    2,
		Owner: 1,
		Name:  "Er model template for database",
	}

	return builder.Build()
}

func (m *SchemaMother) Obj2() *model.Schema {
	builder := SchemaBuilder{
		Id:    3,
		Owner: 2,
		Name:  "Use case diagram",
	}

	return builder.Build()
}

func (m *SchemaMother) Obj3() *model.Schema {
	builder := SchemaBuilder{
		Id:    4,
		Owner: 3,
		Name:  "Use case diagram for sites with logic",
	}

	return builder.Build()
}
