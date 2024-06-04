package models

import (
	"github.com/stdyum/api-common/models"
)

type Option struct {
	Type    string
	Subject models.Subject
	Group   models.Group
	Teacher models.Teacher
}
