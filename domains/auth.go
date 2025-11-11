package domains

import "github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/models"

type AuthService interface {
	Authorize(tokenString string) (bool, error)
	CreateToken(models.User) string
}
