package constants

import "regexp"

type ValidationCode int

const (
	Required ValidationCode = iota
	InvalidFormat
	TooShort
	TooLong
	InvalidEmail

	PasswordTooShort
	UsernameTooShort
	UsernameTooLong
)

var validationMessages = map[ValidationCode]map[Lang]string{
	Required: {
		LangEN: "%s is required",
		LangID: "%s wajib diisi",
	},
	InvalidFormat: {
		LangEN: "%s format is invalid",
		LangID: "format %s tidak valid",
	},
	TooShort: {
		LangEN: "%s is too short (minimum %d characters)",
		LangID: "%s terlalu pendek (minimal %d karakter)",
	},
	TooLong: {
		LangEN: "%s is too long (maximum %d characters)",
		LangID: "%s terlalu panjang (maksimal %d karakter)",
	},
	InvalidEmail: {
		LangEN: "invalid email format",
		LangID: "format email tidak valid",
	},
	PasswordTooShort: {
		LangEN: "password must be at least %d characters",
		LangID: "password minimal %d karakter",
	},
	UsernameTooShort: {
		LangEN: "username must be at least %d characters",
		LangID: "username minimal %d karakter",
	},
	UsernameTooLong: {
		LangEN: "username must be at most %d characters",
		LangID: "username maksimal %d karakter",
	},
}

func GetValidationMessage(code ValidationCode, lang Lang) string {
	if msgs, ok := validationMessages[code]; ok {
		if msg, ok := msgs[lang]; ok {
			return msg
		}
		if msg, ok := msgs[LangEN]; ok {
			return msg
		}
	}
	return "validation error"
}

type Validator interface {
	Validate(lang Lang) map[string][]string
}

func IsValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func IsNotEmpty(s string) bool {
	return s != ""
}

func MinLength(s string, min int) bool {
	return len(s) >= min
}

func MaxLength(s string, max int) bool {
	return len(s) <= max
}
