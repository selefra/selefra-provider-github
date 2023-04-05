package faker

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"

	"github.com/selefra/selefra-provider-github/constants"
)

type faker struct {
	maxDepth int
}

type Option func(*faker)

func (f faker) getFakedValue(a interface{}) (reflect.Value, error) {
	t := reflect.TypeOf(a)
	if t == nil {
		return reflect.Value{}, fmt.Errorf(constants.Interfacenotallowed)
	}
	f.maxDepth--
	if f.maxDepth < 0 {
		return reflect.Value{}, fmt.Errorf(constants.Maxdepthreached)
	}
	k := t.Kind()

	switch k {
	case reflect.Ptr:
		v := reflect.New(t.Elem())
		var val reflect.Value
		var err error
		if a != reflect.Zero(reflect.TypeOf(a)).Interface() {
			val, err = f.getFakedValue(reflect.ValueOf(a).Elem().Interface())
		} else {
			val, err = f.getFakedValue(v.Elem().Interface())
		}
		if err != nil {
			return reflect.Value{}, err
		}
		v.Elem().Set(val.Convert(t.Elem()))
		return v, nil
	case reflect.Struct:
		switch t.String() {
		case constants.TimeTime:
			ft := time.Now().Add(time.Duration(rand.Int63()))
			return reflect.ValueOf(ft), nil
		default:
			v := reflect.New(t).Elem()
			for i := 0; i < v.NumField(); i++ {
				if !v.Field(i).CanSet() {
					continue
				}
				val, err := f.getFakedValue(v.Field(i).Interface())
				if err != nil {
					fmt.Println(err)
					continue

				}
				val = val.Convert(v.Field(i).Type())
				v.Field(i).Set(val)
			}
			return v, nil
		}
	case reflect.String:
		return reflect.ValueOf(constants.Teststring), nil
	case reflect.Slice:
		sliceLen := 1
		v := reflect.MakeSlice(t, sliceLen, sliceLen)
		for i := 0; i < v.Len(); i++ {
			val, err := f.getFakedValue(v.Index(i).Interface())
			if err != nil {
				return reflect.Value{}, err
			}
			val = val.Convert(v.Index(i).Type())
			v.Index(i).Set(val)
		}
		return v, nil
	case reflect.Array:
		v := reflect.New(t).Elem()
		for i := 0; i < v.Len(); i++ {
			val, err := f.getFakedValue(v.Index(i).Interface())
			if err != nil {
				return reflect.Value{}, err
			}
			val = val.Convert(v.Index(i).Type())
			v.Index(i).Set(val)
		}
		return v, nil
	case reflect.Int:
		return reflect.ValueOf(int(123)), nil
	case reflect.Int8:
		return reflect.ValueOf(int8(123)), nil
	case reflect.Int16:
		return reflect.ValueOf(int16(123)), nil
	case reflect.Int32:
		return reflect.ValueOf(int32(123)), nil
	case reflect.Int64:
		return reflect.ValueOf(int64(123)), nil
	case reflect.Float32:
		return reflect.ValueOf(float32(123)), nil
	case reflect.Float64:
		return reflect.ValueOf(float64(1.123)), nil
	case reflect.Bool:
		return reflect.ValueOf(true), nil

	case reflect.Uint:
		return reflect.ValueOf(uint(123)), nil

	case reflect.Uint8:
		return reflect.ValueOf(uint8(123)), nil

	case reflect.Uint16:
		return reflect.ValueOf(uint16(123)), nil

	case reflect.Uint32:
		return reflect.ValueOf(uint32(123)), nil

	case reflect.Uint64:
		return reflect.ValueOf(uint64(123)), nil

	case reflect.Map:
		v := reflect.MakeMap(t)
		for i := 0; i < 1; i++ {
			keyInstance := reflect.New(t.Key()).Elem().Interface()
			key, err := f.getFakedValue(keyInstance)
			if err != nil {
				return reflect.Value{}, err
			}

			valueInstance := reflect.New(t.Elem()).Elem().Interface()
			val, err := f.getFakedValue(valueInstance)
			if err != nil {
				return reflect.Value{}, err
			}
			val = val.Convert(v.Type().Elem())
			v.SetMapIndex(key, val)
		}
		return v, nil
	default:
		err := fmt.Errorf(constants.Nosupportforkindv, t)
		return reflect.Value{}, err
	}
}

func WithMaxDepth(depth int) Option {
	return func(f *faker) {
		f.maxDepth = depth
	}
}

func FakeObject(obj interface{}, opts ...Option) error {
	reflectType := reflect.TypeOf(obj)

	if reflectType.Kind() != reflect.Ptr {
		return fmt.Errorf(constants.Objectisnotapointer)
	}

	if reflect.ValueOf(obj).IsNil() {
		return fmt.Errorf(constants.Objectisnils, reflectType.Elem().String())
	}
	f := &faker{
		maxDepth: 12,
	}
	for _, o := range opts {
		o(f)
	}

	rval := reflect.ValueOf(obj)
	finalValue, err := f.getFakedValue(obj)
	if err != nil {
		return err
	}

	rval.Elem().Set(finalValue.Elem().Convert(reflectType.Elem()))
	return nil
}
