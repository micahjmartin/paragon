// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/kcarretto/paragon/ent/target"
)

// Target is the model entity for the Target schema.
type Target struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "Name" field.
	Name string `json:"Name,omitempty"`
	// PrimaryIP holds the value of the "PrimaryIP" field.
	PrimaryIP string `json:"PrimaryIP,omitempty"`
	// MachineUUID holds the value of the "MachineUUID" field.
	MachineUUID string `json:"MachineUUID,omitempty"`
	// PublicIP holds the value of the "PublicIP" field.
	PublicIP string `json:"PublicIP,omitempty"`
	// PrimaryMAC holds the value of the "PrimaryMAC" field.
	PrimaryMAC string `json:"PrimaryMAC,omitempty"`
	// Hostname holds the value of the "Hostname" field.
	Hostname string `json:"Hostname,omitempty"`
	// LastSeen holds the value of the "LastSeen" field.
	LastSeen time.Time `json:"LastSeen,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TargetQuery when eager-loading is set.
	Edges TargetEdges `json:"edges"`
}

// TargetEdges holds the relations/edges for other nodes in the graph.
type TargetEdges struct {
	// Tasks holds the value of the tasks edge.
	Tasks []*Task
	// Tags holds the value of the tags edge.
	Tags []*Tag
	// Credentials holds the value of the credentials edge.
	Credentials []*Credential
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// TasksOrErr returns the Tasks value or an error if the edge
// was not loaded in eager-loading.
func (e TargetEdges) TasksOrErr() ([]*Task, error) {
	if e.loadedTypes[0] {
		return e.Tasks, nil
	}
	return nil, &NotLoadedError{edge: "tasks"}
}

// TagsOrErr returns the Tags value or an error if the edge
// was not loaded in eager-loading.
func (e TargetEdges) TagsOrErr() ([]*Tag, error) {
	if e.loadedTypes[1] {
		return e.Tags, nil
	}
	return nil, &NotLoadedError{edge: "tags"}
}

// CredentialsOrErr returns the Credentials value or an error if the edge
// was not loaded in eager-loading.
func (e TargetEdges) CredentialsOrErr() ([]*Credential, error) {
	if e.loadedTypes[2] {
		return e.Credentials, nil
	}
	return nil, &NotLoadedError{edge: "credentials"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Target) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // Name
		&sql.NullString{}, // PrimaryIP
		&sql.NullString{}, // MachineUUID
		&sql.NullString{}, // PublicIP
		&sql.NullString{}, // PrimaryMAC
		&sql.NullString{}, // Hostname
		&sql.NullTime{},   // LastSeen
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Target fields.
func (t *Target) assignValues(values ...interface{}) error {
	if m, n := len(values), len(target.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	t.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Name", values[0])
	} else if value.Valid {
		t.Name = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field PrimaryIP", values[1])
	} else if value.Valid {
		t.PrimaryIP = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field MachineUUID", values[2])
	} else if value.Valid {
		t.MachineUUID = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field PublicIP", values[3])
	} else if value.Valid {
		t.PublicIP = value.String
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field PrimaryMAC", values[4])
	} else if value.Valid {
		t.PrimaryMAC = value.String
	}
	if value, ok := values[5].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Hostname", values[5])
	} else if value.Valid {
		t.Hostname = value.String
	}
	if value, ok := values[6].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field LastSeen", values[6])
	} else if value.Valid {
		t.LastSeen = value.Time
	}
	return nil
}

// QueryTasks queries the tasks edge of the Target.
func (t *Target) QueryTasks() *TaskQuery {
	return (&TargetClient{t.config}).QueryTasks(t)
}

// QueryTags queries the tags edge of the Target.
func (t *Target) QueryTags() *TagQuery {
	return (&TargetClient{t.config}).QueryTags(t)
}

// QueryCredentials queries the credentials edge of the Target.
func (t *Target) QueryCredentials() *CredentialQuery {
	return (&TargetClient{t.config}).QueryCredentials(t)
}

// Update returns a builder for updating this Target.
// Note that, you need to call Target.Unwrap() before calling this method, if this Target
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Target) Update() *TargetUpdateOne {
	return (&TargetClient{t.config}).UpdateOne(t)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (t *Target) Unwrap() *Target {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Target is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Target) String() string {
	var builder strings.Builder
	builder.WriteString("Target(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", Name=")
	builder.WriteString(t.Name)
	builder.WriteString(", PrimaryIP=")
	builder.WriteString(t.PrimaryIP)
	builder.WriteString(", MachineUUID=")
	builder.WriteString(t.MachineUUID)
	builder.WriteString(", PublicIP=")
	builder.WriteString(t.PublicIP)
	builder.WriteString(", PrimaryMAC=")
	builder.WriteString(t.PrimaryMAC)
	builder.WriteString(", Hostname=")
	builder.WriteString(t.Hostname)
	builder.WriteString(", LastSeen=")
	builder.WriteString(t.LastSeen.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Targets is a parsable slice of Target.
type Targets []*Target

func (t Targets) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
