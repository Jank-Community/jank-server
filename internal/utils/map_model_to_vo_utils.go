// Package utils 提供模型对象到视图对象的映射工具
// 创建者：Done-0
// 创建时间：2025-05-10
package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
)

// MapModelToVO 将模型数据映射到对应的 VO，返回具体类型的指针
// 参数：
//   - modelData: 源模型数据
//   - voPtr: 目标视图对象指针
//
// 返回值：
//   - interface{}: 映射后的视图对象指针
//   - error: 映射过程中的错误
func MapModelToVO(modelData interface{}, voPtr interface{}) (interface{}, error) {
	modelVal := reflect.ValueOf(modelData)
	voVal := reflect.ValueOf(voPtr)

	// 检查 voPtr 是否为指向结构体的指针
	if voVal.Kind() != reflect.Ptr || voVal.IsNil() {
		return nil, fmt.Errorf("voPtr 必须为指向结构体的指针")
	}
	voVal = voVal.Elem()

	// 如果 modelData 是指针类型，解引用获取实际的结构体值
	if modelVal.Kind() == reflect.Ptr {
		modelVal = modelVal.Elem()
	}

	// 确保 modelData 和 voPtr 都是结构体类型
	if modelVal.Kind() != reflect.Struct || voVal.Kind() != reflect.Struct {
		return nil, fmt.Errorf("modelData 和 voPtr 必须为结构体类型")
	}

	numFields := modelVal.NumField()
	modelType := modelVal.Type()

	// 控制并发的最大数量
	maxConcurrency := 8
	sem := make(chan struct{}, maxConcurrency)

	var wg sync.WaitGroup
	var mu sync.Mutex

	// 遍历 modelData 中的字段并并行处理
	for i := 0; i < numFields; i++ {
		wg.Add(1)
		sem <- struct{}{} // 控制并发

		go func(i int) {
			defer wg.Done()
			defer func() { <-sem }() // 释放信号量

			modelField := modelVal.Field(i)
			fieldType := modelType.Field(i)
			voField := voVal.FieldByName(fieldType.Name)

			if voField.IsValid() && voField.CanSet() {
				// 赋值操作
				if modelField.Type().AssignableTo(voField.Type()) {
					mu.Lock()
					voField.Set(modelField)
					mu.Unlock()
				} else if modelField.Kind() == reflect.String && voField.Kind() == reflect.Slice && voField.Type().Elem().Kind() == reflect.Int64 {
					str := modelField.String()
					if str != "" {
						strArray := strings.Split(str, ",")
						intArray := make([]int64, len(strArray))
						for j, s := range strArray {
							if val, err := strconv.ParseInt(s, 10, 64); err == nil {
								intArray[j] = val
							}
						}
						mu.Lock()
						voField.Set(reflect.ValueOf(intArray))
						mu.Unlock()
					}
				} else if strings.HasSuffix(strings.ToLower(fieldType.Name), "id") && modelField.Kind() == reflect.Int64 && voField.Kind() == reflect.String {
					mu.Lock()
					voField.SetString(strconv.FormatInt(modelField.Int(), 10))
					mu.Unlock()
				}
			}

			// 处理嵌套结构体字段
			if modelField.Kind() == reflect.Struct {
				embeddedModelType := modelField.Type()
				embeddedNumFields := embeddedModelType.NumField()

				for j := 0; j < embeddedNumFields; j++ {
					embeddedField := embeddedModelType.Field(j)
					voField := voVal.FieldByName(embeddedField.Name)

					if voField.IsValid() && voField.CanSet() {
						mu.Lock()
						if embeddedField.Type.AssignableTo(voField.Type()) {
							voField.Set(modelField.Field(j))
						} else if strings.HasSuffix(embeddedField.Name, "ID") && modelField.Field(j).Kind() == reflect.Int64 && voField.Kind() == reflect.String {
							voField.SetString(strconv.FormatInt(modelField.Field(j).Int(), 10))
						} else if (embeddedField.Name == "GmtCreate" || embeddedField.Name == "GmtModified") && modelField.Field(j).Kind() == reflect.Int64 && voField.Kind() == reflect.String {
							timestamp := modelField.Field(j).Int()
							if timestamp > 0 {
								timeStr := time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
								voField.SetString(timeStr)
							} else {
								voField.SetString("")
							}
						}
						mu.Unlock()
					}
				}
			}
		}(i)
	}

	wg.Wait()

	return voVal.Addr().Interface(), nil
}
