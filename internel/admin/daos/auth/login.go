package auth

import (
	"github.com/jeffcail/ginframe/internel/admin/models"
)

func FindAdminByUsername(username string) (*models.Admin, error) {
	m := new(models.Admin)
	admin, err := m.FindAdminByUsername(username)
	if err != nil {
		return nil, err
	}
	return admin, nil
}
