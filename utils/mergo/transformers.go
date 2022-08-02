package mergo

import (
	"reflect"

	"github.com/imdario/mergo"
)

var _ mergo.Transformers = BoolPtrTransformer{}
var _ mergo.Transformers = CombinedTransformer{}

type CombinedTransformer struct {
	Transformers []mergo.Transformers
}

func (ct CombinedTransformer) Transformer(typ reflect.Type) func(dst, src reflect.Value) error {
	for _, t := range ct.Transformers {
		if f := t.Transformer(typ); f != nil {
			return f
		}
	}
	return nil
}

type BoolPtrTransformer struct{}

func (t BoolPtrTransformer) Transformer(typ reflect.Type) func(dst, src reflect.Value) error {
	if typ == reflect.TypeOf(new(bool)) {
		return func(dst, src reflect.Value) error {
			if !src.IsNil() && dst.CanSet() {
				dst.Set(src)
			}
			return nil
		}
	}
	return nil
}
