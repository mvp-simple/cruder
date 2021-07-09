package cruder

import "reflect"

// manyModelCreator меняет создание слайса сущности которая используется в crud
type manyModelCreator struct {
	modelSlice reflect.Type
}

// Option имплементация интерфеса types.ICruderOption
func (o *manyModelCreator) Option(optionSetIn iCruderOptionHelper) {
	optionSetIn.SetModelSlice(o.modelSlice)
}

// ManyModelCreator меняет создание слайса сущности которая используется в crud
func ManyModelCreator(modelSlice interface{}) (out ICruderOption) {
	return &manyModelCreator{modelSlice: reflect.TypeOf(modelSlice)}
}
