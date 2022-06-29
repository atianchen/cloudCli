package utils

import (
	"reflect"
	"strings"
)

/**
 * 自动映射工具
 * 将map映射到struct实例
 * @tag 标识映射关系的tag
 * @author jensen.chen
 * @date 2022/6/29
 */
func MapToStruct(data map[string]interface{}, target interface{}, tagName string) {
	targetType := reflect.TypeOf(target).Elem()
	targetVal := reflect.ValueOf(target).Elem()
	fieldNum := targetType.NumField()
	for i := 0; i < fieldNum; i++ {
		field := targetType.Field(i)
		rs, _ := data[fieldConfigName(field, tagName)]
		if rs != nil {
			/**
			如果类型是Struct Pointer，则创建实例，递归调用
			*/
			switch field.Type.Kind() {
			case reflect.Pointer:
				{
					nv := reflect.New(field.Type.Elem()).Interface()
					MapToStruct(rs.(map[string]interface{}), nv, tagName)
					targetVal.Field(i).Set(reflect.ValueOf(nv))
				}
			case reflect.String:
				{
					targetVal.Field(i).SetString(rs.(string))
				}
			case reflect.Int, reflect.Int32, reflect.Int64, reflect.Int8:
				{
					targetVal.Field(i).SetInt(int64(rs.(int)))
				}
			case reflect.Float64, reflect.Float32:
				{
					targetVal.Field(i).SetFloat(rs.(float64))
				}
			}
		}
	}
}

/**
获取与FIELD对应的配置项名
*/
func fieldConfigName(field reflect.StructField, tagName string) string {
	tag := field.Tag
	if tag == "" {
		return strings.ToLower(field.Name)
	} else {
		tagVal := tag.Get(tagName)
		if tagVal == "" {
			return strings.ToLower(field.Name)
		} else {
			return tagVal
		}
	}
}
