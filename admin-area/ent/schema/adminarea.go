package schema

import (
	"database/sql"
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
		field.Int("id").Unique().Immutable(),
		field.String("name").NotEmpty().Comment("行政区划名称"),
		field.String("memo").Optional().GoType(&sql.NullString{}).Nillable().Comment("备注"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Optional().UpdateDefault(time.Now),
		field.Int("lft").Default(0),
		field.Int("rgt").Default(0),
		field.Int("parent_id").Optional().Nillable(),
	}
}

func (AdminArea) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", AdminArea.Type).From("parent"),
	}
}

func (AdminArea) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.SoftDeleteMixin{},
	}
}
