// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/testifysec/archivist/ent/signature"
	"github.com/testifysec/archivist/ent/timestamp"
)

// TimestampCreate is the builder for creating a Timestamp entity.
type TimestampCreate struct {
	config
	mutation *TimestampMutation
	hooks    []Hook
}

// SetType sets the "type" field.
func (tc *TimestampCreate) SetType(s string) *TimestampCreate {
	tc.mutation.SetType(s)
	return tc
}

// SetTimestamp sets the "timestamp" field.
func (tc *TimestampCreate) SetTimestamp(t time.Time) *TimestampCreate {
	tc.mutation.SetTimestamp(t)
	return tc
}

// SetSignatureID sets the "signature" edge to the Signature entity by ID.
func (tc *TimestampCreate) SetSignatureID(id int) *TimestampCreate {
	tc.mutation.SetSignatureID(id)
	return tc
}

// SetNillableSignatureID sets the "signature" edge to the Signature entity by ID if the given value is not nil.
func (tc *TimestampCreate) SetNillableSignatureID(id *int) *TimestampCreate {
	if id != nil {
		tc = tc.SetSignatureID(*id)
	}
	return tc
}

// SetSignature sets the "signature" edge to the Signature entity.
func (tc *TimestampCreate) SetSignature(s *Signature) *TimestampCreate {
	return tc.SetSignatureID(s.ID)
}

// Mutation returns the TimestampMutation object of the builder.
func (tc *TimestampCreate) Mutation() *TimestampMutation {
	return tc.mutation
}

// Save creates the Timestamp in the database.
func (tc *TimestampCreate) Save(ctx context.Context) (*Timestamp, error) {
	var (
		err  error
		node *Timestamp
	)
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TimestampMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Timestamp)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TimestampMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TimestampCreate) SaveX(ctx context.Context) *Timestamp {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TimestampCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TimestampCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TimestampCreate) check() error {
	if _, ok := tc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Timestamp.type"`)}
	}
	if _, ok := tc.mutation.Timestamp(); !ok {
		return &ValidationError{Name: "timestamp", err: errors.New(`ent: missing required field "Timestamp.timestamp"`)}
	}
	return nil
}

func (tc *TimestampCreate) sqlSave(ctx context.Context) (*Timestamp, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (tc *TimestampCreate) createSpec() (*Timestamp, *sqlgraph.CreateSpec) {
	var (
		_node = &Timestamp{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: timestamp.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: timestamp.FieldID,
			},
		}
	)
	if value, ok := tc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: timestamp.FieldType,
		})
		_node.Type = value
	}
	if value, ok := tc.mutation.Timestamp(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: timestamp.FieldTimestamp,
		})
		_node.Timestamp = value
	}
	if nodes := tc.mutation.SignatureIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   timestamp.SignatureTable,
			Columns: []string{timestamp.SignatureColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: signature.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.signature_timestamps = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TimestampCreateBulk is the builder for creating many Timestamp entities in bulk.
type TimestampCreateBulk struct {
	config
	builders []*TimestampCreate
}

// Save creates the Timestamp entities in the database.
func (tcb *TimestampCreateBulk) Save(ctx context.Context) ([]*Timestamp, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Timestamp, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TimestampMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TimestampCreateBulk) SaveX(ctx context.Context) []*Timestamp {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TimestampCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TimestampCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
