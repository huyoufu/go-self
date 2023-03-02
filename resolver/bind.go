package resolver

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const defaultDatePattern = "2006-01-02"

func Bind(bean interface{}, form map[string][]string) error {
	//将数据绑定给bean bean必须是指针类型的 不然会报错的
	typ := reflect.TypeOf(bean).Elem()
	val := reflect.ValueOf(bean).Elem()

	for i := 0; i < typ.NumField(); i++ {
		//遍历字段集合
		fieldType := typ.Field(i)
		fieldVal := val.Field(i)
		if !fieldVal.CanSet() {
			//如果该字段是私有的 直接放弃
			continue
		}
		//返回字段的类型
		fieldTypeKind := fieldType.Type.Kind()
		//进行设置值
		//结构体类型特殊
		if fieldTypeKind == reflect.Struct && !isTimeType(fieldType.Type) {
			//正常的结构体类型
			//fmt.Println("结构体类型:")
			//递归调用

			Bind(fieldVal.Addr().Interface(), form)
		}
		//先去map集合中获取对应的值
		lower := firstLower(fieldType.Name)
		value, exist := getValues(form, lower)
		if !exist {
			continue
		}

		if fieldTypeKind == reflect.Slice {
			//fmt.Println("切片类型")
			//切片类型 先判断给出值对应的值 切片个数
			//获取切片的类型
			sliceType := fieldType.Type.Elem()
			valueLen := len(value)
			//创建一个对应类型的切片
			slice_value := reflect.MakeSlice(fieldType.Type, valueLen, valueLen)
			for i := 0; i < valueLen; i++ {
				if err := setWithProperType(sliceType.Kind(), value[i], slice_value.Index(i)); err != nil {
					return err
				}
			}
			//fmt.Println(slice_value)
			fieldVal.Set(slice_value)
		} else if fieldTypeKind == reflect.Struct && isTimeType(fieldType.Type) {
			//fmt.Println("日期类型")
			time_value, e := time.Parse(defaultDatePattern, value[0])
			if e != nil {
				return e
			}
			fieldVal.Set(reflect.ValueOf(time_value))
		} else {
			//普通类型
			//fmt.Println("普通类型")

			setWithProperType(fieldTypeKind, value[0], fieldVal)
		}
	}
	return nil
}
func isTimeType(fieldType reflect.Type) bool {
	return fieldType.AssignableTo(reflect.TypeOf(time.Time{}))
}

func getValues(form map[string][]string, key string) (vals []string, exist bool) {
	vals, exist = form[key]
	if !exist {
		//不存在就找相识的
		for _, k := range getKeys(form) {
			if strings.Contains(k, key) {
				vals, exist = form[k]
				return
			}
		}
		return nil, false
	}
	return

}
func getKeys(form map[string][]string) []string {
	var result []string
	for key := range form {
		result = append(result, key)
	}
	return result
}
func firstLower(src string) string {
	firstLetter := substr(src, 0, 1)
	lastLetter := substr(src, 1, len(src))

	firstLetter = strings.ToLower(firstLetter)
	return firstLetter + lastLetter
}
func substr(str string, start int, end int) string {
	rs := []rune(str)
	return string(rs[start:end])
}
func setWithProperType(typeKind reflect.Kind, val string, valueField reflect.Value) error {
	switch typeKind {
	case reflect.Int:
		return setIntField(val, 0, valueField)
	case reflect.Int8:
		return setIntField(val, 8, valueField)
	case reflect.Int16:
		return setIntField(val, 16, valueField)
	case reflect.Int32:
		return setIntField(val, 32, valueField)
	case reflect.Int64:
		return setIntField(val, 64, valueField)
	case reflect.Uint:
		return setUintField(val, 0, valueField)
	case reflect.Uint8:
		return setUintField(val, 8, valueField)
	case reflect.Uint16:
		return setUintField(val, 16, valueField)
	case reflect.Uint32:
		return setUintField(val, 32, valueField)
	case reflect.Uint64:
		return setUintField(val, 64, valueField)
	case reflect.Bool:
		return setBoolField(val, valueField)
	case reflect.Float32:
		return setFloatField(val, 32, valueField)
	case reflect.Float64:
		return setFloatField(val, 64, valueField)
	case reflect.String:
		valueField.SetString(val)
	default:
		return errors.New("unknown type")
	}
	return nil
}
func setIntField(value string, bitSize int, field reflect.Value) error {
	if value == "" {
		value = "0"
	}
	intVal, err := strconv.ParseInt(value, 10, bitSize)
	if err == nil {
		field.SetInt(intVal)
	}
	return err
}

func setUintField(value string, bitSize int, field reflect.Value) error {
	if value == "" {
		value = "0"
	}
	uintVal, err := strconv.ParseUint(value, 10, bitSize)
	if err == nil {
		field.SetUint(uintVal)
	}
	return err
}

func setBoolField(value string, field reflect.Value) error {
	if value == "" {
		value = "false"
	}
	boolVal, err := strconv.ParseBool(value)
	if err == nil {
		field.SetBool(boolVal)
	}
	return err
}

func setFloatField(value string, bitSize int, field reflect.Value) error {
	if value == "" {
		value = "0.0"
	}
	floatVal, err := strconv.ParseFloat(value, bitSize)
	if err == nil {
		field.SetFloat(floatVal)
	}
	return err
}
