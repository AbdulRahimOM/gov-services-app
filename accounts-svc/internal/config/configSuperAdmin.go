package config

import (
	"os"
)

// GetSuperAdminCredentials
func GetSuperAdminCredentials() (string, string) {
	return os.Getenv("SUPER_ADMIN_USERNAME"), os.Getenv("SUPER_ADMIN_PASSWORD")
}
