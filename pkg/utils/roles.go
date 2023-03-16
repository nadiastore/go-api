package utils

import (
	"fmt"

	// TODO: Buat reponya biar bisa diinstall
	"github.com/nadiastore/go-api/pkg/repository"
)

func VerifyRole(role string) (string, error) {
	switch role {
	case repository.AdminRoleName:
		// Langsung terverifikasi.

	case repository.ModeratorRoleName:
		// Langsung terverifikasi.

	case repository.UserRoleName:
		// Langsung terverifikasi.

	default:
		return "", fmt.Errorf("role '%v' does not exist", role)
	}

	return role, nil
}
