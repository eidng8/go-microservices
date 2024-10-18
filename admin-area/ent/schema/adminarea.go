package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"eidng8.cc/microservices/rdbms/mixins"
)

type AdminArea struct {
	ent.Schema
}

func (AdminArea) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		schema.Comment("行政区划表"),
	}
}

func (AdminArea) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique().Immutable(),
		field.String("name").NotEmpty().Comment("行政区划名称"),
		field.String("memo").Optional().Nillable().Comment("备注"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Optional().Nillable().UpdateDefault(time.Now),
		field.Int("lft").Default(0),
		field.Int("rgt").Default(0),
		field.Uint64("parent_id").Optional().Nillable(),
	}
}

func (AdminArea) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", AdminArea.Type).
			From("parent").Field("parent_id").Unique(),
	}
}

func (AdminArea) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.SoftDeleteMixin{},
	}
}
