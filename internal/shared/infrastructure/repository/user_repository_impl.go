package repository

import (
	"app/internal/shared/domain/entity"
	"app/internal/shared/domain/repository"
	"app/pkg"
	"context"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *entity.User) error {
	result := r.db.WithContext(ctx).Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *userRepository) GetByID(ctx context.Context, id string) (*entity.User, error) {
	var user entity.User
	result := r.db.WithContext(ctx).Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	result := r.db.WithContext(ctx).Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	var user entity.User
	if err := r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(ctx context.Context, filter entity.FilterUser, user *entity.User) error {
	result := r.db.WithContext(ctx).Where("deleted_at IS NULL AND id = ?", filter.ID).Updates(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Delete deletes a user (soft delete)
func (r *userRepository) Delete(ctx context.Context, id string) error {
	if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.User{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) List(ctx context.Context, filter entity.FilterUser) ([]*entity.User, int, error) {
	scopes := []func(db *gorm.DB) *gorm.DB{
		func(db *gorm.DB) *gorm.DB {
			return db.Where("deleted_at IS NULL")
		},
	}

	if filter.ID != "" {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("id = ?", filter.ID)
		})
	}
	if filter.Email != "" {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("email = ?", filter.Email)
		})
	}
	if filter.Username != "" {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("username = ?", filter.Username)
		})
	}
	if filter.Name != "" {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("name LIKE ?", "%"+filter.Name+"%")
		})
	}
	if filter.IsActive != nil {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("is_active = ?", *filter.IsActive)
		})
	}

	if filter.Role != "" {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("role = ?", filter.Role)
		})
	}

	if len(filter.Roles) > 0 {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("role IN ?", filter.Roles)
		})
	}

	var users []*entity.User
	err := r.db.WithContext(ctx).
		Scopes(pkg.Paginate(filter.Offset, filter.PerPage, r.db)).
		Scopes(scopes...).
		Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	var totalRows int64
	if err := r.db.WithContext(ctx).Model(&entity.User{}).Scopes(scopes...).Count(&totalRows).Error; err != nil {
		return nil, 0, err
	}

	return users, int(totalRows), nil
}
