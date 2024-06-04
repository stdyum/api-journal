package controllers

import (
	"github.com/google/uuid"
	"github.com/stdyum/api-journal/internal/app/controllers/models"
	"github.com/stdyum/api-journal/internal/modules/types_registry"
)

type OptionsBuilder struct {
	options  []models.Option
	typesIds types_registry.TypesIds
}

func NewOptionsBuilder() *OptionsBuilder {
	return &OptionsBuilder{}
}

func (b *OptionsBuilder) Build() ([]models.Option, types_registry.TypesIds) {
	return b.options, b.typesIds
}

func (b *OptionsBuilder) Append(options ...models.Option) *OptionsBuilder {
	b.options = append(b.options, options...)
	b.appendTypeIdsFromOptionsArray(options)
	return b
}

func (b *OptionsBuilder) appendTypeIdsFromOptionsArray(options []models.Option) {
	for _, option := range options {
		b.appendTypeIdsFromOption(option)
	}
}

func (b *OptionsBuilder) appendTypeIdsFromOption(option models.Option) {
	if option.Group.ID != uuid.Nil {
		b.typesIds.GroupsIds = append(b.typesIds.GroupsIds, option.Group.ID)
	}

	if option.Subject.ID != uuid.Nil {
		b.typesIds.SubjectsIds = append(b.typesIds.SubjectsIds, option.Subject.ID)
	}

	if option.Teacher.ID != uuid.Nil {
		b.typesIds.TeachersIds = append(b.typesIds.TeachersIds, option.Teacher.ID)
	}
}
