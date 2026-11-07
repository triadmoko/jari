package entity

// FilterUser represents the filtering options for user queries
type FilterUser struct {
	// Basic filters
	ID       string `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
	Name     string `json:"name,omitempty"`
	IsActive *bool  `json:"is_active,omitempty"` // Pointer to distinguish between false and not set

	Role string `json:"role,omitempty"`

	// Array filters for IN queries
	Roles []string `json:"roles,omitempty"`

	// Pagination
	Offset  int `json:"offset"`
	PerPage int `json:"per_page"`
}
