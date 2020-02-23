// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/kcarretto/paragon/ent/job"
	"github.com/kcarretto/paragon/ent/user"
)

// Job is the model entity for the Job schema.
type Job struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "Name" field.
	Name string `json:"Name,omitempty"`
	// CreationTime holds the value of the "CreationTime" field.
	CreationTime time.Time `json:"CreationTime,omitempty"`
	// Content holds the value of the "Content" field.
	Content string `json:"Content,omitempty"`
	// Staged holds the value of the "Staged" field.
	Staged bool `json:"Staged,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the JobQuery when eager-loading is set.
	Edges     JobEdges `json:"edges"`
	job_next  *int
	user_jobs *int
}

// JobEdges holds the relations/edges for other nodes in the graph.
type JobEdges struct {
	// Tasks holds the value of the tasks edge.
	Tasks []*Task
	// Tags holds the value of the tags edge.
	Tags []*Tag
	// Prev holds the value of the prev edge.
	Prev *Job
	// Next holds the value of the next edge.
	Next *Job
	// Owner holds the value of the owner edge.
	Owner *User
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// TasksOrErr returns the Tasks value or an error if the edge
// was not loaded in eager-loading.
func (e JobEdges) TasksOrErr() ([]*Task, error) {
	if e.loadedTypes[0] {
		return e.Tasks, nil
	}
	return nil, &NotLoadedError{edge: "tasks"}
}

// TagsOrErr returns the Tags value or an error if the edge
// was not loaded in eager-loading.
func (e JobEdges) TagsOrErr() ([]*Tag, error) {
	if e.loadedTypes[1] {
		return e.Tags, nil
	}
	return nil, &NotLoadedError{edge: "tags"}
}

// PrevOrErr returns the Prev value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e JobEdges) PrevOrErr() (*Job, error) {
	if e.loadedTypes[2] {
		if e.Prev == nil {
			// The edge prev was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: job.Label}
		}
		return e.Prev, nil
	}
	return nil, &NotLoadedError{edge: "prev"}
}

// NextOrErr returns the Next value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e JobEdges) NextOrErr() (*Job, error) {
	if e.loadedTypes[3] {
		if e.Next == nil {
			// The edge next was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: job.Label}
		}
		return e.Next, nil
	}
	return nil, &NotLoadedError{edge: "next"}
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e JobEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[4] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Job) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // Name
		&sql.NullTime{},   // CreationTime
		&sql.NullString{}, // Content
		&sql.NullBool{},   // Staged
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Job) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // job_next
		&sql.NullInt64{}, // user_jobs
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Job fields.
func (j *Job) assignValues(values ...interface{}) error {
	if m, n := len(values), len(job.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	j.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Name", values[0])
	} else if value.Valid {
		j.Name = value.String
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field CreationTime", values[1])
	} else if value.Valid {
		j.CreationTime = value.Time
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Content", values[2])
	} else if value.Valid {
		j.Content = value.String
	}
	if value, ok := values[3].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field Staged", values[3])
	} else if value.Valid {
		j.Staged = value.Bool
	}
	values = values[4:]
	if len(values) == len(job.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field job_next", value)
		} else if value.Valid {
			j.job_next = new(int)
			*j.job_next = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field user_jobs", value)
		} else if value.Valid {
			j.user_jobs = new(int)
			*j.user_jobs = int(value.Int64)
		}
	}
	return nil
}

// QueryTasks queries the tasks edge of the Job.
func (j *Job) QueryTasks() *TaskQuery {
	return (&JobClient{j.config}).QueryTasks(j)
}

// QueryTags queries the tags edge of the Job.
func (j *Job) QueryTags() *TagQuery {
	return (&JobClient{j.config}).QueryTags(j)
}

// QueryPrev queries the prev edge of the Job.
func (j *Job) QueryPrev() *JobQuery {
	return (&JobClient{j.config}).QueryPrev(j)
}

// QueryNext queries the next edge of the Job.
func (j *Job) QueryNext() *JobQuery {
	return (&JobClient{j.config}).QueryNext(j)
}

// QueryOwner queries the owner edge of the Job.
func (j *Job) QueryOwner() *UserQuery {
	return (&JobClient{j.config}).QueryOwner(j)
}

// Update returns a builder for updating this Job.
// Note that, you need to call Job.Unwrap() before calling this method, if this Job
// was returned from a transaction, and the transaction was committed or rolled back.
func (j *Job) Update() *JobUpdateOne {
	return (&JobClient{j.config}).UpdateOne(j)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (j *Job) Unwrap() *Job {
	tx, ok := j.config.driver.(*txDriver)
	if !ok {
		panic("ent: Job is not a transactional entity")
	}
	j.config.driver = tx.drv
	return j
}

// String implements the fmt.Stringer.
func (j *Job) String() string {
	var builder strings.Builder
	builder.WriteString("Job(")
	builder.WriteString(fmt.Sprintf("id=%v", j.ID))
	builder.WriteString(", Name=")
	builder.WriteString(j.Name)
	builder.WriteString(", CreationTime=")
	builder.WriteString(j.CreationTime.Format(time.ANSIC))
	builder.WriteString(", Content=")
	builder.WriteString(j.Content)
	builder.WriteString(", Staged=")
	builder.WriteString(fmt.Sprintf("%v", j.Staged))
	builder.WriteByte(')')
	return builder.String()
}

// Jobs is a parsable slice of Job.
type Jobs []*Job

func (j Jobs) config(cfg config) {
	for _i := range j {
		j[_i].config = cfg
	}
}
