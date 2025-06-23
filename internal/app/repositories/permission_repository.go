package repositories

import (
	"gorm.io/gorm"
	"homestay.com/nguyenduy/internal/app/models"
)

type PermissionRepository interface {
	GetAllPermissions() ([]models.Permission, error)
	GetPermissionByID(id string) (models.Permission, error)
	CreatePermission(permission models.Permission) (models.Permission, error)
	UpdatePermission(permission models.Permission) (models.Permission, error)
	DeletePermission(id string) error
	GetPermissionsByUserID(userID string) ([]models.Permission, error)
	GetPermissionsByRoleID(roleID string) ([]models.Permission, error)
}

type permissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) PermissionRepository {
	return &permissionRepository{db: db}
}

func (r *permissionRepository) GetAllPermissions() ([]models.Permission, error) {
	var permissions []models.Permission
	err := r.db.Find(&permissions).Error
	return permissions, err
}

func (r *permissionRepository) GetPermissionByID(id string) (models.Permission, error) {
	var permission models.Permission
	err := r.db.First(&permission, id).Error
	return permission, err
}

func (r *permissionRepository) CreatePermission(permission models.Permission) (models.Permission, error) {
	err := r.db.Create(&permission).Error
	return permission, err
}

func (r *permissionRepository) UpdatePermission(permission models.Permission) (models.Permission, error) {
	err := r.db.Save(&permission).Error
	return permission, err
}

func (r *permissionRepository) DeletePermission(id string) error {
	return r.db.Delete(&models.Permission{}, id).Error
}

func (r *permissionRepository) GetPermissionsByUserID(userID string) ([]models.Permission, error) {
	var user models.User
	var permissions []models.Permission
	if err := r.db.Preload("Role.Permissions").First(&user, userID).Error; err != nil {
		return nil, err
	}
	permissions = user.Role.Permissions
	return permissions, nil
}

func (r *permissionRepository) GetPermissionsByRoleID(roleID string) ([]models.Permission, error) {
	var role models.Role
	if err := r.db.Preload("Permissions").First(&role, roleID).Error; err != nil {
		return nil, err
	}
	return role.Permissions, nil
}
