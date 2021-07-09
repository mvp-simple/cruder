package cruder

import "reflect"

// oneModelCreator меняет создание сущности которая используется в crud
type oneModelCreator struct {
	creator reflect.Type
}

// Option имплементация интерфеса types.ICruderOption
func (o *oneModelCreator) Option(optionSetIn iCruderOptionHelper) {
	optionSetIn.SetModel(o.creator)
}

// OneModelCreator меняет создание сущности которая используется в crud
func OneModelCreator(model interface{}) (out ICruderOption) {
	return &oneModelCreator{creator: reflect.TypeOf(model)}
}
