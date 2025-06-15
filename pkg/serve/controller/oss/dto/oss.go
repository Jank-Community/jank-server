// Package dto 提供对象存储相关的数据传输对象定义
// 创建者：Done-0
// 创建时间：2025-05-10
package dto

// UploadOneFileRequest 上传文件请求
// @Param BucketName  formData file   true  "要上传的文件"
// @Param UploadFile  formData string true  "存储桶名称"
type UploadOneFileRequest struct {
	BucketName string `json:"bucket_name" xml:"bucket_name" form:"bucket_name" validate:"required,min=1,max=63"`
	UploadFile string `json:"upload_file" xml:"upload_file" form:"upload_file"`
}

// DownloadOneFileRequest 下载文件请求
// @Param BucketName    query string true  "存储桶名称"
// @Param ObjectName    query string true  "对象名称"
type DownloadOneFileRequest struct {
	BucketName string `json:"bucket_name" xml:"bucket_name" form:"bucket_name" query:"bucket_name" validate:"required,min=1,max=63"`
	ObjectName string `json:"object_name" xml:"object_name" form:"object_name" query:"object_name" validate:"required,min=1,max=1024"`
}

// DeleteOneFileRequest 删除文件请求
// @Param BucketName body string true  "存储桶名称"
// @Param ObjectName body string true  "对象名称"
type DeleteOneFileRequest struct {
	BucketName string `json:"bucket_name" xml:"bucket_name" form:"bucket_name" validate:"required,min=1,max=63"`
	ObjectName string `json:"object_name" xml:"object_name" form:"object_name" validate:"required,min=1,max=1024"`
}

// ListAllObjectsRequest 列出对象请求
// @Param BucketName  query string true  "存储桶名称"
// @Param Prefix      query string false "对象名称前缀(可选，如果不指定则列出所有对象)"
type ListAllObjectsRequest struct {
	BucketName string `json:"bucket_name" xml:"bucket_name" form:"bucket_name" query:"bucket_name" validate:"required,min=1,max=63"`
	Prefix     string `json:"prefix" xml:"prefix" form:"prefix" query:"prefix" validate:"omitempty,max=1024"`
}
