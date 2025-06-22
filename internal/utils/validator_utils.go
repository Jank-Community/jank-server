// Package utils 提供参数验证工具
// 创建者：Done-0
// 创建时间：2025-05-10
package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"

	"jank.com/jank_blog/pkg/enums"
)

// ValidErrRes 验证错误结果结构体
type ValidErrRes struct {
	Error bool        // 是否存在错误
	Field string      // 错误字段名
	Tag   string      // 错误标签
	Value interface{} // 错误值
}

// NewValidator 全局验证器实例
var NewValidator = validator.New()

func init() {
	err := NewValidator.RegisterValidation("auditStatus", auditStatusValidation)
	if err != nil {
		fmt.Printf("自定义验证器注册失败: %v\n", err)
	}
}

// Validator 参数验证器
// 参数：
//   - data: 待验证的数据
//
// 返回值：
//   - []ValidErrRes: 验证错误结果数组
func Validator(data interface{}) []ValidErrRes {
	var Errors []ValidErrRes
	errs := NewValidator.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var el ValidErrRes
			el.Error = true
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Value()

			Errors = append(Errors, el)
		}
	}
	return Errors
}

// auditStatusValidation 自定义验证器，检查字段值是否为 "approved" 或 "rejected"
func auditStatusValidation(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	switch value {
	case string(enums.AuditApproved), string(enums.AuditRejected):
		return true // 允许的值
	default:
		return false // 不允许的值
	}
}
