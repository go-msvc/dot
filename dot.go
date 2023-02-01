package dot

import (
	"strconv"
	"strings"

	"github.com/go-msvc/errors"
	"github.com/stewelarend/logger"
)

func New(v interface{}) Value {
	return Value{v: v}
}

type Value struct {
	v interface{}
}

var log = logger.New() //.WithLevel(logger.LevelDebug)

func (value Value) Value() interface{} {
	return value.v
}

func (value *Value) Set(name string, namedValue interface{}) error {
	newValue, err := set(value.v, name, namedValue)
	if err != nil {
		return err
	}
	value.v = newValue
	return nil
}

func set(value interface{}, name string, namedValue interface{}) (interface{}, error) {
	if name == "" {
		return namedValue, nil
	}

	//name must start with a '.' for an object reference
	//or [] for array element(s)
	log.Debugf("===== set((%T)%+v, %s = (%T)%+v) =====", value, value, name, namedValue, namedValue)

	if name[0] == '.' {
		//object item reference e.g. ".x" or ".x[10]" or ".x.y" or ".x.y[10]" etc...
		//get the first part "x"
		name = name[1:]
		remain := ""

		sepIndex := strings.IndexAny(name, ".[")
		log.Debugf("(%s).sepIndex=%d", name, sepIndex)
		if sepIndex >= 0 {
			remain = name[sepIndex:]
			name = name[:sepIndex]
			log.Debugf("name(%s) remain(%s)", name, remain)
		}

		if value == nil {
			value = map[string]interface{}{}
		}
		objValue, ok := value.(map[string]interface{})
		if !ok {
			return nil, errors.Errorf("invalid reference(%s): %T is not an object", name, value)
		}

		newValue, err := set(objValue[name], remain, namedValue)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to set(%T, %s)", objValue[name], remain)
		}
		objValue[name] = newValue
		log.Debugf("%s = (%T)%+v", name, newValue, newValue)
		log.Debugf("objValue[%s]=(%T)%+v", name, objValue[name], objValue[name])
		return objValue, nil
	} //if obj item ref

	if name[0] == '[' {
		//expect name to start with [...] followed optionally by more
		//get the first part "[...]"
		remain := ""
		sepIndex := strings.IndexAny(name, "]")
		log.Debugf("(%s).sepIndex=%d", name, sepIndex)
		if sepIndex < 0 {
			return nil, errors.Errorf("missing ']' in reference")
		}
		remain = name[sepIndex+1:]
		name = name[:sepIndex+1]
		log.Debugf("name(%s) remain(%s)", name, remain)

		arrValue, ok := value.([]interface{})
		if !ok {
			return nil, errors.Errorf("invalid reference(%s): %T is not an array", name, value)
		}
		if name == "[]" {
			return namedValue, nil //replace the whole array
		}
		i64, err := strconv.ParseInt(name[1:len(name)-1], 10, 64)
		if err != nil {
			return nil, errors.Errorf("array index(%s) is not integer", name)
		}
		if i64 < 0 || int(i64) >= len(arrValue) {
			return nil, errors.Errorf("array index(%s) out of range 0..%d not integer", name, i64)
		}
		newValue, err := set(arrValue[int(i64)], remain, namedValue)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to set %s=(%T)%+v", remain, namedValue, namedValue)
		}
		arrValue[int(i64)] = newValue
		log.Debugf("%s = (%T)%+v", name, arrValue[int(i64)], arrValue[int(i64)])
		return arrValue, nil
	} //if arr reference
	return nil, errors.Errorf("invalid name(%s)", name)
} //set()

func (value Value) Get(name string) (interface{}, error) {
	return get(value.v, name)
}

func get(value interface{}, name string) (interface{}, error) {
	if name == "" {
		return value, nil
	}

	//name must start with a '.' for an object reference
	//or [] for array element(s)
	log.Debugf("===== get((%T)%+v, %s) =====", value, value, name)

	if name[0] == '.' {
		//object item reference e.g. ".x" or ".x[10]" or ".x.y" or ".x.y[10]" etc...
		//get the first part "x"
		name = name[1:]
		remain := ""

		sepIndex := strings.IndexAny(name, ".[")
		log.Debugf("(%s).sepIndex=%d", name, sepIndex)
		if sepIndex >= 0 {
			remain = name[sepIndex:]
			name = name[:sepIndex]
			log.Debugf("name(%s) remain(%s)", name, remain)
		}

		objValue, ok := value.(map[string]interface{})
		if !ok {
			return nil, errors.Errorf("invalid reference(%s): %T is not an object", name, value)
		}
		v, ok := objValue[name]
		if !ok {
			return nil, errors.Errorf("%s not found", name)
		}
		log.Debugf("%s = (%T)%+v", name, v, v)
		return get(v, remain)
	} //if obj item ref

	if name[0] == '[' {
		//expect name to start with [...] followed optionally by more
		//get the first part "[...]"
		remain := ""
		sepIndex := strings.IndexAny(name, "]")
		log.Debugf("(%s).sepIndex=%d", name, sepIndex)
		if sepIndex < 0 {
			return nil, errors.Errorf("missing ']' in reference")
		}
		remain = name[sepIndex+1:]
		name = name[:sepIndex+1]
		log.Debugf("name(%s) remain(%s)", name, remain)

		arrValue, ok := value.([]interface{})
		if !ok {
			return nil, errors.Errorf("invalid reference(%s): %T is not an array", name, value)
		}
		if name == "[]" {
			return arrValue, nil //the whole array
		}
		i64, err := strconv.ParseInt(name[1:len(name)-1], 10, 64)
		if err != nil {
			return nil, errors.Errorf("array index(%s) is not integer", name)
		}
		if i64 < 0 || int(i64) >= len(arrValue) {
			return nil, errors.Errorf("array index(%s) out of range 0..%d not integer", name, i64)
		}
		v := arrValue[int(i64)]
		log.Debugf("%s = (%T)%+v", name, v, v)

		return get(v, remain)
	} //if arr reference
	return nil, errors.Errorf("invalid name(%s)", name)
} //get()
