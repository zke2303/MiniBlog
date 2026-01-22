package service

import (
	"context"

	"mini-blog/internal/dto/errmsg"
	"mini-blog/internal/dto/request"
	"mini-blog/internal/model"
	"mini-blog/internal/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BlogService Blog的业务逻辑层
type BlogService struct {
	db   *gorm.DB
	repo repository.IBlogRepository
}

// NewBlogService 创建 BlogService 实例对象
func NewBlogService(repo repository.IBlogRepository, db *gorm.DB) *BlogService {
	return &BlogService{
		db:   db,
		repo: repo,
	}
}

// Create 创建 Blog 实例, 并保存到数据库中
func (svc *BlogService) Create(ctx context.Context, userID string, req request.CreateBlogRequest) (string, error) {
	// 1.校验 title 是否已经被使用

	// 2.生成 uuid
	id, err := uuid.NewV7()
	if err != nil {
		return "", errmsg.UUIDGeneratorErr
	}

	// 3.解析 userID 为 uuid
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return "", errmsg.InternalErr
	}

	// 3.构建 Blog 对象
	blog := model.Blog{
		ID:       id,
		Title:    req.Title,
		AuthorID: userUUID,
		Content:  req.Content,
	}

	// 4.使用事务, 创建数据库记录
	err = svc.db.Transaction(func(tx *gorm.DB) error {
		if err := svc.repo.Create(ctx, tx, blog); err != nil {
			return err
		}
		return nil
	})
	// 5.判断事务是否发送错误
	if err != nil {
		return "", errmsg.InternalErr.Wrap(err)
	}

	// 6.返回 blog id
	return blog.ID.String(), nil
}

func (svc *BlogService) ListBlogs(ctx context.Context, req request.BlogsPageQuery) ([]model.Blog, error) {
	// 1.防止非法参数
	offset := 0
	limit := 10
	if req.Size < 100 && req.Size > 0 {
		limit = req.Size
	}
	if req.Page > 1 {
		offset = (req.Page - 1) * limit
	}

	return svc.repo.ListBlogs(ctx, svc.db, offset, limit, req)
}
