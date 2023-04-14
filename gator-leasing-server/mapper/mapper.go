package mapper

import (
	"github.com/dranikpg/dto-mapper"

	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/errors"
)

type Mapper[T, U any] struct {
	mapper dto.Mapper
}

func NewMapper[T, U any](*T, *U) *Mapper[T, U] {
	mapper := &Mapper[T, U]{mapper: dto.Mapper{}}
	return mapper
}

func (m *Mapper[T, U]) Map(from *T) (*U, error) {
	to := new(U)
	err := m.mapper.Map(to, from)
	if err != nil {
		return nil, &errors.InternalServerError{Msg: err.Error()}
	}
	return to, nil
}

func (m *Mapper[T, U]) MapSlice(fromSlice []*T) ([]*U, error) {
	toSlice := []*U{}
	for _, from := range fromSlice {
		to, err := m.Map(from)
		if err != nil {
			return nil, err
		}
		toSlice = append(toSlice, to)
	}

	return toSlice, nil
}
