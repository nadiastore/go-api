package utils

import (
	"fmt"

	// TODO: Buat reponya biar bisa diinstall
	"github.com/nadiastore/go-api/pkg/repository"
)

func GetCredentialsByRole(role string) ([]string, error) {
	var credentials []string

	switch role {
	// admin (full access)
	case repository.AdminRoleName:
		credentials = []string{
			repository.BookCreateCredential,
			repository.BookUpdateCredential,
			repository.BookDeleteCredential,
		}

	// moderator (hanya create & update)
	case repository.ModeratorRoleName:
		credentials = []string{
			repository.BookCreateCredential,
			repository.BookUpdateCredential,
		}

	// user (hanya create)
	case repository.UserRoleName:
		credentials = []string{
			repository.BookCreateCredential,
		}

	default:
		return nil, fmt.Errorf("role '%v' does not exist", role)
	}

	return credentials, nil
}
