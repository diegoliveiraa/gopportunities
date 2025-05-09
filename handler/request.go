package handler

import (
	"fmt"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// Create opening

type CreateOpeningRequest struct {
	Role     string `json:"role" binding:"required"`
	Company  string `json:"company" binding:"required"`
	Location string `json:"location" binding:"required"`
	Remote   *bool  `json:"remote" binding:"required"`
	Link     string `json:"link" binding:"required"`
	Salary   int64  `json:"salary" binding:"required"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Remote == nil && r.Link == "" && r.Salary <= 0 {
		return fmt.Errorf("at least one of the params is required")
	}
	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if r.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if r.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if r.Remote == nil {
		return errParamIsRequired("remote", "boolean")
	}
	if r.Link == "" {
		return errParamIsRequired("link", "string")
	}
	if r.Salary <= 0 {
		return errParamIsRequired("salary", "int64")
	}
	return nil
}

type UpdateOpeningRequest struct {
	Role     string `json:"role" binding:"required"`
	Company  string `json:"company" binding:"required"`
	Location string `json:"location" binding:"required"`
	Remote   *bool  `json:"remote" binding:"required"`
	Link     string `json:"link" binding:"required"`
	Salary   int64  `json:"salary" binding:"required"`
}

func (r *UpdateOpeningRequest) Validate() error {
	// if any field is provided, validation is true
	if r.Role != "" && r.Company != "" && r.Location != "" && r.Remote != nil && r.Link != "" && r.Salary > 0 {
		return nil
	}
	//if none of the fields are provided, validation is false
	return fmt.Errorf("at least one of the params is required")
}
