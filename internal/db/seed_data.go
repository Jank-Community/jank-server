package db

import (
	"log"

	"jank.com/jank_blog/internal/global"
	model "jank.com/jank_blog/internal/model/account"
)

// 预定义角色常量
const (
	ROLE_ADMIN = "admin" // 超级管理员
	ROLE_USER  = "user"  // 普通用户
)

// seedDefaultData 初始化默认的种子数据
// 返回值：
//   - error: 创建过程中的错误
func seedDefaultData() error {
	if err := seedRoles(); err != nil {
		return err
	}

	if err := seedPermissions(); err != nil {
		return err
	}

	if err := seedRolePermissions(); err != nil {
		return err
	}

	return nil
}

// seedRoles 初始化默认角色
// 返回值：
//   - error: 创建过程中的错误
func seedRoles() error {
	var count int64
	if err := global.DB.Model(&model.Role{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		log.Println("角色数据已存在，跳过初始化...")
		return nil
	}

	roles := []model.Role{
		{
			Name:        ROLE_ADMIN,
			Description: "超级管理员，拥有所有权限",
			Status:      true,
		},
		{
			Name:        ROLE_USER,
			Description: "普通用户，具有基本权限",
			Status:      true,
		},
	}

	if err := global.DB.Create(&roles).Error; err != nil {
		global.SysLog.Errorf("初始化角色数据失败: %v", err)
		return err
	}

	log.Println("初始化角色数据成功...")
	global.SysLog.Info("初始化角色数据成功...")

	return nil
}

// seedPermissions 初始化默认权限
// 返回值：
//   - error: 创建过程中的错误
func seedPermissions() error {
	var count int64
	if err := global.DB.Model(&model.Permission{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		log.Println("权限数据已存在，跳过初始化...")
		return nil
	}

	permissions := []model.Permission{
		// 文章管理权限
		{
			Key:         "post:create",
			Name:        "创建文章",
			Description: "创建新文章",
			Status:      true,
		},
		{
			Key:         "post:read",
			Name:        "查看文章",
			Description: "查看文章内容",
			Status:      true,
		},
		{
			Key:         "post:update",
			Name:        "编辑文章",
			Description: "编辑现有文章",
			Status:      true,
		},
		{
			Key:         "post:delete",
			Name:        "删除文章",
			Description: "删除文章",
			Status:      true,
		},

		// 分类管理权限
		{
			Key:         "category:manage",
			Name:        "管理分类",
			Description: "创建、编辑和删除分类",
			Status:      true,
		},

		// 评论管理权限
		{
			Key:         "comment:create",
			Name:        "发表评论",
			Description: "发表评论",
			Status:      true,
		},
		{
			Key:         "comment:moderate",
			Name:        "管理评论",
			Description: "审核、编辑和删除评论",
			Status:      true,
		},

		// 用户管理权限
		{
			Key:         "user:manage",
			Name:        "管理用户",
			Description: "管理用户账户",
			Status:      true,
		},

		// 系统管理权限
		{
			Key:         "system:settings",
			Name:        "系统设置",
			Description: "管理系统配置",
			Status:      true,
		},
	}

	if err := global.DB.Create(&permissions).Error; err != nil {
		global.SysLog.Errorf("初始化权限数据失败: %v", err)
		return err
	}

	log.Println("初始化权限数据成功...")
	global.SysLog.Info("初始化权限数据成功...")

	return nil
}

// seedRolePermissions 初始化角色权限关联
// 返回值：
//   - error: 创建过程中的错误
func seedRolePermissions() error {
	var count int64
	if err := global.DB.Model(&model.RolePermission{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		log.Println("角色权限关联数据已存在，跳过初始化...")
		return nil
	}

	var adminRole model.Role
	var userRole model.Role
	var permissions []model.Permission
	var userPermissions []model.Permission

	if err := global.DB.Where("name = ?", ROLE_ADMIN).First(&adminRole).Error; err != nil {
		global.SysLog.Errorf("获取管理员角色失败: %v", err)
		return err
	}

	if err := global.DB.Where("name = ?", ROLE_USER).First(&userRole).Error; err != nil {
		global.SysLog.Errorf("获取普通用户角色失败: %v", err)
		return err
	}

	if err := global.DB.Find(&permissions).Error; err != nil {
		global.SysLog.Errorf("获取权限列表失败: %v", err)
		return err
	}

	if err := global.DB.Where("key IN ?", []string{"post:read", "comment:create"}).Find(&userPermissions).Error; err != nil {
		global.SysLog.Errorf("获取普通用户权限列表失败: %v", err)
		return err
	}

	var rolePermissions []model.RolePermission
	// 为管理员角色分配所有权限
	for _, permission := range permissions {
		rolePermissions = append(rolePermissions, model.RolePermission{
			RoleID:       adminRole.ID,
			PermissionID: permission.ID,
		})
	}
	// 为普通用户分配基本权限
	for _, permission := range userPermissions {
		rolePermissions = append(rolePermissions, model.RolePermission{
			RoleID:       userRole.ID,
			PermissionID: permission.ID,
		})
	}

	if err := global.DB.Create(&rolePermissions).Error; err != nil {
		global.SysLog.Errorf("初始化角色权限关联数据失败: %v", err)
		return err
	}

	log.Println("初始化角色权限关联数据成功...")
	global.SysLog.Info("初始化角色权限关联数据成功...")

	return nil
}
