package mapper

import (
	"github.com/dranikpg/dto-mapper"
)

type Mapper[T, U any] struct {
	Mapper dto.Mapper
}

func NewMapper[T, U any](*T, *U) *Mapper[T, U] {
	mapper := &Mapper[T, U]{Mapper: dto.Mapper{}}
	return mapper
}

func (m *Mapper[T, U]) Map(from *T) (*U, error) {
	to := new(U)
	err := m.Mapper.Map(to, from)
	return to, err
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
