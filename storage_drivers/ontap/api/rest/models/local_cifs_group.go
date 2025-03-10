// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LocalCifsGroup local cifs group
//
// swagger:model local_cifs_group
type LocalCifsGroup struct {

	// links
	Links *LocalCifsGroupInlineLinks `json:"_links,omitempty"`

	// Description for the local group.
	//
	// Example: This is a local group
	// Max Length: 256
	Description *string `json:"description,omitempty"`

	// local cifs group inline members
	// Read Only: true
	LocalCifsGroupInlineMembers []*LocalCifsGroupInlineMembersInlineArrayItem `json:"members,omitempty"`

	// Local group name. The maximum supported length of a group name is 256 characters.
	//
	// Example: SMB_SERVER01\\group
	// Min Length: 1
	Name *string `json:"name,omitempty"`

	// The security ID of the local group which uniquely identifies the group. The group SID is automatically generated in POST and it is retrieved using the GET method.
	//
	// Example: S-1-5-21-256008430-3394229847-3930036330-1001
	// Read Only: true
	Sid *string `json:"sid,omitempty"`

	// svm
	Svm *LocalCifsGroupInlineSvm `json:"svm,omitempty"`
}

// Validate validates this local cifs group
func (m *LocalCifsGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalCifsGroupInlineMembers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LocalCifsGroup) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *LocalCifsGroup) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", *m.Description, 256); err != nil {
		return err
	}

	return nil
}

func (m *LocalCifsGroup) validateLocalCifsGroupInlineMembers(formats strfmt.Registry) error {
	if swag.IsZero(m.LocalCifsGroupInlineMembers) { // not required
		return nil
	}

	for i := 0; i < len(m.LocalCifsGroupInlineMembers); i++ {
		if swag.IsZero(m.LocalCifsGroupInlineMembers[i]) { // not required
			continue
		}

		if m.LocalCifsGroupInlineMembers[i] != nil {
			if err := m.LocalCifsGroupInlineMembers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("members" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LocalCifsGroup) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	return nil
}

func (m *LocalCifsGroup) validateSvm(formats strfmt.Registry) error {
	if swag.IsZero(m.Svm) { // not required
		return nil
	}

	if m.Svm != nil {
		if err := m.Svm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this local cifs group based on the context it is used
func (m *LocalCifsGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocalCifsGroupInlineMembers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSid(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LocalCifsGroup) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *LocalCifsGroup) contextValidateLocalCifsGroupInlineMembers(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "members", "body", []*LocalCifsGroupInlineMembersInlineArrayItem(m.LocalCifsGroupInlineMembers)); err != nil {
		return err
	}

	for i := 0; i < len(m.LocalCifsGroupInlineMembers); i++ {

		if m.LocalCifsGroupInlineMembers[i] != nil {
			if err := m.LocalCifsGroupInlineMembers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("members" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LocalCifsGroup) contextValidateSid(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "sid", "body", m.Sid); err != nil {
		return err
	}

	return nil
}

func (m *LocalCifsGroup) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

	if m.Svm != nil {
		if err := m.Svm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LocalCifsGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LocalCifsGroup) UnmarshalBinary(b []byte) error {
	var res LocalCifsGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LocalCifsGroupInlineLinks local cifs group inline links
//
// swagger:model local_cifs_group_inline__links
type LocalCifsGroupInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this local cifs group inline links
func (m *LocalCifsGroupInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LocalCifsGroupInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this local cifs group inline links based on the context it is used
func (m *LocalCifsGroupInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LocalCifsGroupInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LocalCifsGroupInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LocalCifsGroupInlineLinks) UnmarshalBinary(b []byte) error {
	var res LocalCifsGroupInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LocalCifsGroupInlineMembersInlineArrayItem local cifs group inline members inline array item
//
// swagger:model local_cifs_group_inline_members_inline_array_item
type LocalCifsGroupInlineMembersInlineArrayItem struct {

	// Local user, Active Directory user, or Active Directory group which is a member of the specified local group.
	//
	Name *string `json:"name,omitempty"`
}

// Validate validates this local cifs group inline members inline array item
func (m *LocalCifsGroupInlineMembersInlineArrayItem) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this local cifs group inline members inline array item based on context it is used
func (m *LocalCifsGroupInlineMembersInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LocalCifsGroupInlineMembersInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LocalCifsGroupInlineMembersInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res LocalCifsGroupInlineMembersInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LocalCifsGroupInlineSvm SVM, applies only to SVM-scoped objects.
//
// swagger:model local_cifs_group_inline_svm
type LocalCifsGroupInlineSvm struct {

	// links
	Links *LocalCifsGroupInlineSvmInlineLinks `json:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this local cifs group inline svm
func (m *LocalCifsGroupInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LocalCifsGroupInlineSvm) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this local cifs group inline svm based on the context it is used
func (m *LocalCifsGroupInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LocalCifsGroupInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LocalCifsGroupInlineSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LocalCifsGroupInlineSvm) UnmarshalBinary(b []byte) error {
	var res LocalCifsGroupInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LocalCifsGroupInlineSvmInlineLinks local cifs group inline svm inline links
//
// swagger:model local_cifs_group_inline_svm_inline__links
type LocalCifsGroupInlineSvmInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this local cifs group inline svm inline links
func (m *LocalCifsGroupInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LocalCifsGroupInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this local cifs group inline svm inline links based on the context it is used
func (m *LocalCifsGroupInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LocalCifsGroupInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LocalCifsGroupInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LocalCifsGroupInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res LocalCifsGroupInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
