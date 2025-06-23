package repositories

import (
	"gorm.io/gorm"
	model "homestay.com/nguyenduy/internal/app/models"
)

type RoleRepository interface {
	Create(role *model.Role) error
	GetAll() ([]model.Role, error)
	GetByID(id uint) (*model.Role, error)
	Update(id uint, role *model.Role) error
	Delete(id uint) error
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) Create(role *model.Role) error {
	return r.db.Create(role).Error
}

func (r *roleRepository) GetAll() ([]model.Role, error) {
	var roles []model.Role
	if err := r.db.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *roleRepository) GetByID(id uint) (*model.Role, error) {
	var role model.Role
	if err := r.db.First(&role, id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) Update(id uint, role *model.Role) error {
	return r.db.Model(&model.Role{}).Where("id = ?", id).Updates(role).Error
}

func (r *roleRepository) Delete(id uint) error {
	return r.db.Delete(&model.Role{}, id).Error
}
