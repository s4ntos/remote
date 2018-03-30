package models

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-gorp/gorp"
	"github.com/revel/revel"
)

type Role struct {
	RoleId      int
	Role        string
	Description string
	Privileges  string
}

type Roles struct {
	RoleId    int
	ProfileId int
}

var RoleNameRegex = regexp.MustCompile("^[^#@]+$")

var RoleRegex = regexp.MustCompile("^[a-zA-Z0-9_.-]+$")

func (r *Role) String() string {
	return fmt.Sprintf("Role(%s)", r.Role)
}

func (role *Role) RoleValidate(v *revel.Validation) {
	ValidateRoleName(v, role.Role)
	ValidateRoleDescription(v, role.Description)
	ValidateRolePrivileges(v, role.Privileges)
}

func ValidateRoleName(v *revel.Validation, role string) *revel.ValidationResult {
	result := v.Required(role).Message("Role name required")
	if !result.Ok {
		return result
	}

	result = v.MaxSize(role, 64).Message("Role name can not exceed 64 characters")
	if !result.Ok {
		return result
	}

	result = v.Match(role, RoleNameRegex).Message("Invalid User name. Alphanumerics or '-' '.' '_' allowed only")
	if !result.Ok {
		return result
	}

	return result
}

func ValidateRoleDescription(v *revel.Validation, description string) *revel.ValidationResult {
	result := v.MaxSize(description, 400).Message("Profile description cannot exceed 400 characters")

	return result
}

func ValidateRolePrivileges(v *revel.Validation, privileges string) *revel.ValidationResult {
	result := v.MaxSize(privileges, 200).Message("Privileges cannot exceed 200 characters")
	// needs to be extend to validate other things TO-DO
	return result
}

func (r *Role) PreInsert(_ gorp.SqlExecutor) error {
	r.Role = strings.ToLower(r.Role)
	return nil
}

func (r *Role) PreUpdate(_ gorp.SqlExecutor) error {
	r.Role = strings.ToLower(r.Role)
	return nil
}

/*
func (r *Role) PostGet(exe gorp.SqlExecutor) error {
	var (
		obj interface{}
		err error
	)

	obj, err = exe.Get(Role{}, r.RoleId)
	if err != nil {
		return fmt.Errorf("Error loading a profile's user (%d): %s", r.RoleId, err)
	}
	r.Role = obj.(*Role)

	return nil
} */
