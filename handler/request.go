package handler

import "fmt"

type CreateOpeningRequest struct {
	Role     string `json:"role" binding:"required"`
	Company  string `json:"company" binding:"required"`
	Location string `json:"location" binding:"required"`
	Remote   *bool  `json:"remote" binding:"required"`
	Link     string `json:"link" binding:"required"`
	Salary   int64  `json:"salary" binding:"required"`
}

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
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
		return errParamIsRequired("remote", "bool")
	}

	if r.Link == "" {
		return errParamIsRequired("link", "string")
	}

	if r.Salary == 0 {
		return errParamIsRequired("salary", "int64")
	}

	return nil
}

func (r *UpdateOpeningRequest) Validate() error {
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Remote != nil || r.Link != "" || r.Salary != 0 {
		return nil
	}

	return fmt.Errorf("at least one field must be provided")
}

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s)", name, typ)
}
