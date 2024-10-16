package mixins

import (
	"context"
	"fmt"
	"time"

	gen "eidng8.cc/microservices/admin-area/ent"       // my whole folder with ent files
	"eidng8.cc/microservices/admin-area/ent/hook"      // auto generated
	"eidng8.cc/microservices/admin-area/ent/intercept" // auto generate after updating the ent/generate.go

	// all below are just imported ent files
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// SoftDeleteMixin implements the soft delete pattern for schemas.
type SoftDeleteMixin struct {
	mixin.Schema
}

// Fields of the SoftDeleteMixin.
// Once you declare "deleted_at" in here, you MUST DELETE IT from the entity that will use that Mixin
func (SoftDeleteMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("deleted_at").
			Optional(),
	}
}

type softDeleteKey struct{}

// SkipSoftDelete returns a new context that skips the soft-delete interceptor/mutators.
func SkipSoftDelete(parent context.Context) context.Context {
	return context.WithValue(parent, softDeleteKey{}, true)
}

// Interceptors of the SoftDeleteMixin.
func (d SoftDeleteMixin) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		intercept.TraverseFunc(
			func(ctx context.Context, q intercept.Query) error {
				// Skip soft-delete, means include soft-deleted entities.
				if skip, _ := ctx.Value(softDeleteKey{}).(bool); skip {
					return nil
				}
				d.P(q)
				return nil
			},
		),
	}
}

// Hooks of the SoftDeleteMixin.
func (d SoftDeleteMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return ent.MutateFunc(
					func(ctx context.Context, m ent.Mutation) (
						ent.Value, error,
					) {
						// Skip soft-delete, means delete the entity permanently.
						if skip, _ := ctx.Value(softDeleteKey{}).(bool); skip {
							return next.Mutate(ctx, m)
						}

						mx, ok := m.(interface {
							SetOp(ent.Op)
							Client() *gen.Client
							SetDeletedAt(time.Time) // That is the line that needs to be updated if you change column name to be deleted_at
							WhereP(...func(*sql.Selector))
						})
						if !ok {
							return nil, fmt.Errorf(
								"unexpected mutation type %T %+v", m, m,
							)
						}
						d.P(mx)
						mx.SetOp(ent.OpUpdate)
						mx.SetDeletedAt(time.Now())
						return mx.Client().Mutate(ctx, m)
					},
				)
			},
			ent.OpDeleteOne|ent.OpDelete,
		),
	}
}

// P adds a storage-level predicate to the queries and mutations.
func (d SoftDeleteMixin) P(w interface{ WhereP(...func(*sql.Selector)) }) {
	w.WhereP(
		sql.FieldIsNull(d.Fields()[0].Descriptor().Name),
	)
}
