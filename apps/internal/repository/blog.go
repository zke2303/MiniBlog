package repository

import (
	"context"

	"mini-blog/internal/dto/errmsg"
	"mini-blog/internal/model"

	"gorm.io/gorm"
)

// IBlogRepository BlogRepository的接口
type IBlogRepository interface {
	Create(ctx context.Context, db *gorm.DB, blog model.Blog) error
}

// BlogRepository BlogReposioty的实现
type BlogRepository struct{}

// NewBlogRepository 创建 BlogRepository 实例对象
func NewBlogRepository() IBlogRepository {
	return &BlogRepository{}
}

// Create 创建一篇 Blog
func (repo *BlogRepository) Create(ctx context.Context, db *gorm.DB, blog model.Blog) error {
	res := db.WithContext(ctx).Create(&blog)
	if res.Error != nil {
		return errmsg.InternalErr.Wrap(res.Error)
	}

	return nil
}
