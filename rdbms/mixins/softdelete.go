package mixins

import (
	"context"
	"fmt"
	"time"

	// all below are just imported ent files
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"

	gen "eidng8.cc/microservices/admin-area/ent"
)

// SoftDeleteMixin implements the soft delete pattern for schemas.
type SoftDeleteMixin struct {
	mixin.Schema
}

// Fields of the SoftDeleteMixin.
// Once you declare "deleted_at" in here, you MUST DELETE IT from the entity that will use that Mixin
func (SoftDeleteMixin) Fields() []ent.Field {
	return []ent.Field{field.Time("deleted_at").Optional().Nillable()}
}

type softDeleteKey struct{}

// IncludeSoftDeleted returns a new context that skips the soft-delete interceptor/mutators.
func IncludeSoftDeleted(parent context.Context) context.Context {
	return context.WithValue(parent, softDeleteKey{}, true)
}

// Interceptors of the SoftDeleteMixin.
func (d SoftDeleteMixin) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		ent.InterceptFunc(
			func(next ent.Querier) ent.Querier {
				return ent.QuerierFunc(
					func(ctx context.Context, query ent.Query) (
						ent.Value, error,
					) {
						// Skip soft-delete, means include soft-deleted entities.
						if skip, _ := ctx.Value(softDeleteKey{}).(bool); skip {
							return next.Query(ctx, query)
						}
						return next.Query(ctx, query)
					},
				)
			},
		),
	}
}

// Hooks of the SoftDeleteMixin.
func (d SoftDeleteMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		func(next ent.Mutator) ent.Mutator {
			return ent.MutateFunc(
				func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
					op := m.Op()
					if op != ent.OpDelete && op != ent.OpDeleteOne {
						return next.Mutate(ctx, m)
					}
					// Skip soft-delete, means delete the entity permanently.
					if skip, _ := ctx.Value(softDeleteKey{}).(bool); skip {
						return next.Mutate(ctx, m)
					}
					mx, ok := m.(interface {
						SetOp(ent.Op)
						Client() *gen.Client
						SetDeletedAt(time.Time)
					})
					if !ok {
						return nil, fmt.Errorf(
							"unexpected mutation type %T %+v", m, m,
						)
					}
					mx.SetOp(ent.OpUpdate)
					mx.SetDeletedAt(time.Now())
					return mx.Client().Mutate(ctx, m)
				},
			)
		},
	}
}
