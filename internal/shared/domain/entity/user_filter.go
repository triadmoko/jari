package entity

type FilterUser struct {
	ID       string `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
	Name     string `json:"name,omitempty"`
	IsActive *bool  `json:"is_active,omitempty"` // Pointer to distinguish between false and not set

	Role string `json:"role,omitempty"`

	Roles []string `json:"roles,omitempty"`

	Offset  int `json:"offset"`
	PerPage int `json:"per_page"`
}
