// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/testifysec/archivist/ent/attestationcollection"
	"github.com/testifysec/archivist/ent/statement"
)

// Statement is the model entity for the Statement schema.
type Statement struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Predicate holds the value of the "predicate" field.
	Predicate string `json:"predicate,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StatementQuery when eager-loading is set.
	Edges StatementEdges `json:"edges"`
}

// StatementEdges holds the relations/edges for other nodes in the graph.
type StatementEdges struct {
	// Subjects holds the value of the subjects edge.
	Subjects []*Subject `json:"subjects,omitempty"`
	// AttestationCollections holds the value of the attestation_collections edge.
	AttestationCollections *AttestationCollection `json:"attestation_collections,omitempty"`
	// Dsse holds the value of the dsse edge.
	Dsse []*Dsse `json:"dsse,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [3]*int
}

// SubjectsOrErr returns the Subjects value or an error if the edge
// was not loaded in eager-loading.
func (e StatementEdges) SubjectsOrErr() ([]*Subject, error) {
	if e.loadedTypes[0] {
		return e.Subjects, nil
	}
	return nil, &NotLoadedError{edge: "subjects"}
}

// AttestationCollectionsOrErr returns the AttestationCollections value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StatementEdges) AttestationCollectionsOrErr() (*AttestationCollection, error) {
	if e.loadedTypes[1] {
		if e.AttestationCollections == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: attestationcollection.Label}
		}
		return e.AttestationCollections, nil
	}
	return nil, &NotLoadedError{edge: "attestation_collections"}
}

// DsseOrErr returns the Dsse value or an error if the edge
// was not loaded in eager-loading.
func (e StatementEdges) DsseOrErr() ([]*Dsse, error) {
	if e.loadedTypes[2] {
		return e.Dsse, nil
	}
	return nil, &NotLoadedError{edge: "dsse"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Statement) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case statement.FieldID:
			values[i] = new(sql.NullInt64)
		case statement.FieldPredicate:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Statement", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Statement fields.
func (s *Statement) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case statement.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case statement.FieldPredicate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field predicate", values[i])
			} else if value.Valid {
				s.Predicate = value.String
			}
		}
	}
	return nil
}

// QuerySubjects queries the "subjects" edge of the Statement entity.
func (s *Statement) QuerySubjects() *SubjectQuery {
	return (&StatementClient{config: s.config}).QuerySubjects(s)
}

// QueryAttestationCollections queries the "attestation_collections" edge of the Statement entity.
func (s *Statement) QueryAttestationCollections() *AttestationCollectionQuery {
	return (&StatementClient{config: s.config}).QueryAttestationCollections(s)
}

// QueryDsse queries the "dsse" edge of the Statement entity.
func (s *Statement) QueryDsse() *DsseQuery {
	return (&StatementClient{config: s.config}).QueryDsse(s)
}

// Update returns a builder for updating this Statement.
// Note that you need to call Statement.Unwrap() before calling this method if this Statement
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Statement) Update() *StatementUpdateOne {
	return (&StatementClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Statement entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Statement) Unwrap() *Statement {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Statement is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Statement) String() string {
	var builder strings.Builder
	builder.WriteString("Statement(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("predicate=")
	builder.WriteString(s.Predicate)
	builder.WriteByte(')')
	return builder.String()
}

// Statements is a parsable slice of Statement.
type Statements []*Statement

func (s Statements) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
