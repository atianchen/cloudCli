package utils

import "reflect"

/**
 *
 * @author jensen.chen
 * @date 2022/7/11
 */
func CopyProperties(data interface{}, source interface{}) error {
	_type := reflect.TypeOf(data).Elem()
	_val := reflect.ValueOf(data).Elem()
	_sourceVal := reflect.ValueOf(source)
	if reflect.TypeOf(source).Kind() == reflect.Pointer {
		_sourceVal = reflect.ValueOf(source).Elem()
	}
	fieldNum := _type.NumField()
	for i := 0; i < fieldNum; i++ {
		field := _type.Field(i)
		if _val.FieldByName(field.Name).IsValid() {
			switch field.Type.Kind() {
			case reflect.Pointer:
				{
					nv := reflect.New(field.Type.Elem()).Interface()
					CopyProperties(nv, _sourceVal.FieldByName(field.Name).Interface())
					_sourceVal.Field(i).Set(reflect.ValueOf(nv))
				}
			case reflect.String:
				{
					_val.Field(i).SetString(_sourceVal.FieldByName(field.Name).String())
				}
			case reflect.Int, reflect.Int32, reflect.Int64, reflect.Int8:
				{
					_val.Field(i).SetInt(_sourceVal.FieldByName(field.Name).Int())
				}
			case reflect.Float64, reflect.Float32:
				{
					_val.Field(i).SetFloat(_sourceVal.FieldByName(field.Name).Float())
				}
			case reflect.Bool:
				{
					_val.Field(i).SetBool(_sourceVal.FieldByName(field.Name).Bool())
				}
			case reflect.Uint, reflect.Uint32, reflect.Uint64, reflect.Uint8:
				{
					_val.Field(i).SetUint(_sourceVal.FieldByName(field.Name).Uint())
				}
			}
		}
	}
	return nil
}
