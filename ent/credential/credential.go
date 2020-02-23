// Code generated by entc, DO NOT EDIT.

package credential

import (
	"fmt"

	"github.com/kcarretto/paragon/ent/schema"
)

const (
	// Label holds the string label denoting the credential type in the database.
	Label = "credential"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPrincipal holds the string denoting the principal vertex property in the database.
	FieldPrincipal = "principal"
	// FieldSecret holds the string denoting the secret vertex property in the database.
	FieldSecret = "secret"
	// FieldKind holds the string denoting the kind vertex property in the database.
	FieldKind = "kind"
	// FieldFails holds the string denoting the fails vertex property in the database.
	FieldFails = "fails"

	// Table holds the table name of the credential in the database.
	Table = "credentials"
	// TargetTable is the table the holds the target relation/edge.
	TargetTable = "credentials"
	// TargetInverseTable is the table name for the Target entity.
	// It exists in this package in order to avoid circular dependency with the "target" package.
	TargetInverseTable = "targets"
	// TargetColumn is the table column denoting the target relation/edge.
	TargetColumn = "target_credentials"
)

// Columns holds all SQL columns for credential fields.
var Columns = []string{
	FieldID,
	FieldPrincipal,
	FieldSecret,
	FieldKind,
	FieldFails,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Credential type.
var ForeignKeys = []string{
	"target_credentials",
}

var (
	fields = schema.Credential{}.Fields()

	// descSecret is the schema descriptor for secret field.
	descSecret = fields[1].Descriptor()
	// SecretValidator is a validator for the "secret" field. It is called by the builders before save.
	SecretValidator = descSecret.Validators[0].(func(string) error)

	// descFails is the schema descriptor for fails field.
	descFails = fields[3].Descriptor()
	// DefaultFails holds the default value on creation for the fails field.
	DefaultFails = descFails.Default.(int)
	// FailsValidator is a validator for the "fails" field. It is called by the builders before save.
	FailsValidator = descFails.Validators[0].(func(int) error)
)

// Kind defines the type for the kind enum field.
type Kind string

// Kind values.
const (
	KindPassword    Kind = "password"
	KindKey         Kind = "key"
	KindCertificate Kind = "certificate"
)

func (s Kind) String() string {
	return string(s)
}

// KindValidator is a validator for the "k" field enum values. It is called by the builders before save.
func KindValidator(k Kind) error {
	switch k {
	case KindPassword, KindKey, KindCertificate:
		return nil
	default:
		return fmt.Errorf("credential: invalid enum value for kind field: %q", k)
	}
}
