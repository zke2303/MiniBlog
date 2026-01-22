package repository

import (
	"context"

	"mini-blog/internal/dto/errmsg"
	"mini-blog/internal/dto/request"
	"mini-blog/internal/model"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// IBlogRepository BlogRepository的接口
type IBlogRepository interface {
	Create(ctx context.Context, db *gorm.DB, blog model.Blog) error
	ListBlogs(ctx context.Context, db *gorm.DB, offset, limit int, req request.BlogsPageQuery) ([]model.Blog, error)
}

// BlogRepository BlogReposioty的实现
type BlogRepository struct {
	rdb *redis.Client
}

// NewBlogRepository 创建 BlogRepository 实例对象
func NewBlogRepository(rdb *redis.Client) IBlogRepository {
	return &BlogRepository{
		rdb: rdb,
	}
}

// Create 创建一篇 Blog
func (repo *BlogRepository) Create(ctx context.Context, db *gorm.DB, blog model.Blog) error {
	res := db.WithContext(ctx).Create(&blog)
	if res.Error != nil {
		return errmsg.InternalErr.Wrap(res.Error)
	}
	return nil
}

// ListBlogs 分页查询
func (repo *BlogRepository) ListBlogs(ctx context.Context, db *gorm.DB, offset, limit int, req request.BlogsPageQuery) ([]model.Blog, error) {
	// 1.构建分页查询条件
	db = db.WithContext(ctx)

	if req.Title != "" {
		db = db.Where("title like ?", "%"+req.Title+"%")
	}

	if req.AuthorID != "" {
		db = db.Where("author_id = ?", req.AuthorID)
	}

	db = db.Offset(offset).Limit(limit)

	// 2.执行 sql 操作
	var blogs []model.Blog
	res := db.Find(&blogs)

	// 3.判断是否执行成功
	if res.Error != nil {
		return []model.Blog{}, errmsg.InternalErr.Wrap(res.Error)
	}

	return blogs, nil
}
