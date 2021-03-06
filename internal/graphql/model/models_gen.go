// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql/introspection"
)

type Integration struct {
	ID                   string `json:"id" `
	Name                 string `json:"name" `
	Specification        string `json:"specification" `
	SpecificationFileURL string `json:"specificationFileURL" `
	CreatorID            string `json:"creatorID" `
}
type IntegrationInput struct {
	Specification string `json:"specification" `
}
type Preferences struct {
	ID      string        `json:"id" `
	UserID  string        `json:"userID" `
	Details *PrefsDetails `json:"details" `
}
type PrefsAutomation struct {
	ID                 string           `json:"id" `
	FocusWorkflowIndex *int             `json:"focusWorkflowIndex" `
	Workflows          []*PrefsWorkflow `json:"workflows" `
}
type PrefsBase struct {
	ID              string        `json:"id" `
	FocusTableIndex *int          `json:"focusTableIndex" `
	Tables          []*PrefsTable `json:"tables" `
}
type PrefsBoard struct {
	ID string `json:"id" `
}
type PrefsDeck struct {
	ID              string       `json:"id" `
	FocusBoardIndex *int         `json:"focusBoardIndex" `
	Decks           []*PrefsDeck `json:"decks" `
}
type PrefsDetails struct {
	FocusWorkspaceIndex *int              `json:"focusWorkspaceIndex" `
	Workspaces          []*PrefsWorkspace `json:"workspaces" `
}
type PrefsLayout struct {
	MainShortcuts  []*Shortcut `json:"mainShortcuts" `
	QuickShortcuts []*Shortcut `json:"quickShortcuts" `
	OtherShortcuts []*Shortcut `json:"otherShortcuts" `
}
type PrefsSpace struct {
	ID                   string             `json:"id" `
	FocusDeckIndex       *int               `json:"focusDeckIndex" `
	FocusAppIndex        *int               `json:"focusAppIndex" `
	FocusAutomationIndex *int               `json:"focusAutomationIndex" `
	FocusBaseIndex       *int               `json:"focusBaseIndex" `
	Decks                []*PrefsDeck       `json:"decks" `
	Automations          []*PrefsAutomation `json:"automations" `
	Bases                []*PrefsBase       `json:"bases" `
}
type PrefsTable struct {
	ID string `json:"id" `
}
type PrefsWorkflow struct {
	ID string `json:"id" `
}
type PrefsWorkspace struct {
	ID              string        `json:"id" `
	FocusSpaceIndex *int          `json:"focusSpaceIndex" `
	Spaces          []*PrefsSpace `json:"spaces" `
	Layout          *PrefsLayout  `json:"layout" `
}
type Profile struct {
	ID        string  `json:"id" `
	Username  *string `json:"username" `
	FirstName *string `json:"firstName" `
	LastName  *string `json:"lastName" `
	Email     *string `json:"email" `
	AvatarURL *string `json:"avatarURL" `
	UserID    string  `json:"userID" `
}
type ProfileInput struct {
	Username  *string `json:"username" `
	FirstName *string `json:"firstName" `
	LastName  *string `json:"lastName" `
	Email     *string `json:"email" `
	AvatarURL *string `json:"avatarURL" `
}
type SessionUser struct {
	ID          string       `json:"id" `
	Profile     *Profile     `json:"profile" `
	Preferences *Preferences `json:"preferences" `
}
type Shortcut struct {
	IconName   string `json:"iconName" `
	EntityName string `json:"entityName" `
	Route      string `json:"route" `
}
type User struct {
	ID string `json:"id" `
}
type Directive struct {
	Name        string                      `json:"name" `
	Description *string                     `json:"description" `
	Locations   []string                    `json:"locations" `
	Args        []*introspection.InputValue `json:"args" `
}
type EnumValue struct {
	Name              string  `json:"name" `
	Description       *string `json:"description" `
	IsDeprecated      bool    `json:"isDeprecated" `
	DeprecationReason *string `json:"deprecationReason" `
}
type Field struct {
	Name              string                      `json:"name" `
	Description       *string                     `json:"description" `
	Args              []*introspection.InputValue `json:"args" `
	Type              *introspection.Type         `json:"type" `
	IsDeprecated      bool                        `json:"isDeprecated" `
	DeprecationReason *string                     `json:"deprecationReason" `
}
type InputValue struct {
	Name         string              `json:"name" `
	Description  *string             `json:"description" `
	Type         *introspection.Type `json:"type" `
	DefaultValue *string             `json:"defaultValue" `
}
type Schema struct {
	Types            []*introspection.Type      `json:"types" `
	QueryType        *introspection.Type        `json:"queryType" `
	MutationType     *introspection.Type        `json:"mutationType" `
	SubscriptionType *introspection.Type        `json:"subscriptionType" `
	Directives       []*introspection.Directive `json:"directives" `
}
type Type struct {
	Kind          string                      `json:"kind" `
	Name          *string                     `json:"name" `
	Description   *string                     `json:"description" `
	Fields        []*introspection.Field      `json:"fields" `
	Interfaces    []*introspection.Type       `json:"interfaces" `
	PossibleTypes []*introspection.Type       `json:"possibleTypes" `
	EnumValues    []*introspection.EnumValue  `json:"enumValues" `
	InputFields   []*introspection.InputValue `json:"inputFields" `
	OfType        *introspection.Type         `json:"ofType" `
}

type DirectiveLocation string

const (
	DirectiveLocationQuery                DirectiveLocation = "QUERY"
	DirectiveLocationMutation             DirectiveLocation = "MUTATION"
	DirectiveLocationSubscription         DirectiveLocation = "SUBSCRIPTION"
	DirectiveLocationField                DirectiveLocation = "FIELD"
	DirectiveLocationFragmentDefinition   DirectiveLocation = "FRAGMENT_DEFINITION"
	DirectiveLocationFragmentSpread       DirectiveLocation = "FRAGMENT_SPREAD"
	DirectiveLocationInlineFragment       DirectiveLocation = "INLINE_FRAGMENT"
	DirectiveLocationSchema               DirectiveLocation = "SCHEMA"
	DirectiveLocationScalar               DirectiveLocation = "SCALAR"
	DirectiveLocationObject               DirectiveLocation = "OBJECT"
	DirectiveLocationFieldDefinition      DirectiveLocation = "FIELD_DEFINITION"
	DirectiveLocationArgumentDefinition   DirectiveLocation = "ARGUMENT_DEFINITION"
	DirectiveLocationInterface            DirectiveLocation = "INTERFACE"
	DirectiveLocationUnion                DirectiveLocation = "UNION"
	DirectiveLocationEnum                 DirectiveLocation = "ENUM"
	DirectiveLocationEnumValue            DirectiveLocation = "ENUM_VALUE"
	DirectiveLocationInputObject          DirectiveLocation = "INPUT_OBJECT"
	DirectiveLocationInputFieldDefinition DirectiveLocation = "INPUT_FIELD_DEFINITION"
)

var AllDirectiveLocation = []DirectiveLocation{
	DirectiveLocationQuery,
	DirectiveLocationMutation,
	DirectiveLocationSubscription,
	DirectiveLocationField,
	DirectiveLocationFragmentDefinition,
	DirectiveLocationFragmentSpread,
	DirectiveLocationInlineFragment,
	DirectiveLocationSchema,
	DirectiveLocationScalar,
	DirectiveLocationObject,
	DirectiveLocationFieldDefinition,
	DirectiveLocationArgumentDefinition,
	DirectiveLocationInterface,
	DirectiveLocationUnion,
	DirectiveLocationEnum,
	DirectiveLocationEnumValue,
	DirectiveLocationInputObject,
	DirectiveLocationInputFieldDefinition,
}

func (e DirectiveLocation) IsValid() bool {
	switch e {
	case DirectiveLocationQuery, DirectiveLocationMutation, DirectiveLocationSubscription, DirectiveLocationField, DirectiveLocationFragmentDefinition, DirectiveLocationFragmentSpread, DirectiveLocationInlineFragment, DirectiveLocationSchema, DirectiveLocationScalar, DirectiveLocationObject, DirectiveLocationFieldDefinition, DirectiveLocationArgumentDefinition, DirectiveLocationInterface, DirectiveLocationUnion, DirectiveLocationEnum, DirectiveLocationEnumValue, DirectiveLocationInputObject, DirectiveLocationInputFieldDefinition:
		return true
	}
	return false
}

func (e DirectiveLocation) String() string {
	return string(e)
}

func (e *DirectiveLocation) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DirectiveLocation(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid __DirectiveLocation", str)
	}
	return nil
}

func (e DirectiveLocation) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TypeKind string

const (
	TypeKindScalar      TypeKind = "SCALAR"
	TypeKindObject      TypeKind = "OBJECT"
	TypeKindInterface   TypeKind = "INTERFACE"
	TypeKindUnion       TypeKind = "UNION"
	TypeKindEnum        TypeKind = "ENUM"
	TypeKindInputObject TypeKind = "INPUT_OBJECT"
	TypeKindList        TypeKind = "LIST"
	TypeKindNonNull     TypeKind = "NON_NULL"
)

var AllTypeKind = []TypeKind{
	TypeKindScalar,
	TypeKindObject,
	TypeKindInterface,
	TypeKindUnion,
	TypeKindEnum,
	TypeKindInputObject,
	TypeKindList,
	TypeKindNonNull,
}

func (e TypeKind) IsValid() bool {
	switch e {
	case TypeKindScalar, TypeKindObject, TypeKindInterface, TypeKindUnion, TypeKindEnum, TypeKindInputObject, TypeKindList, TypeKindNonNull:
		return true
	}
	return false
}

func (e TypeKind) String() string {
	return string(e)
}

func (e *TypeKind) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TypeKind(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid __TypeKind", str)
	}
	return nil
}

func (e TypeKind) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
