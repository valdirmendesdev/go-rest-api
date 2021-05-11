package services

import (
	"github.com/jinzhu/gorm"
)

type CommentRepository struct {
	DB *gorm.DB
}

type Comment struct {
	gorm.Model
	Slug   string
	Body   string
	Author string
}

type CommentService interface {
	GetComment(ID uint) (Comment, error)
	GetCommentBySlug(slug string) ([]Comment, error)
	PostComment(comment Comment) (Comment, error)
	UpdateComment(ID uint, newComment Comment)(Comment, error)
	DeleteComment(ID uint) error
	GetAllComments() ([]Comment, error)
}

func NewCommentService(db *gorm.DB) *CommentRepository {
	return &CommentRepository{
		DB: db,
	}
}

func (r *CommentRepository) GetComment(ID uint) (Comment, error) {
	var comment Comment
	if result := r.DB.First(&comment, ID); result.Error != nil {
		return Comment{}, result.Error
	}

	return comment, nil
}

func (r *CommentRepository) GetCommentBySlug(slug string) ([]Comment, error) {
	var comments []Comment
	if result := r.DB.First(&comments).Where("slug = ?", slug); result.Error != nil {
		return []Comment{}, result.Error
	}
	return comments, nil
}

func (r *CommentRepository) PostComment(comment Comment) (Comment, error) {
	if result := r.DB.Save(&comment); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

func (r *CommentRepository) UpdateComment(ID uint, newComment Comment) (Comment, error) {
	comment, err := r.GetComment(ID)
	if err != nil {
		return Comment{}, err
	}
	if result := r.DB.Model(&comment).Updates(newComment); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

func (r *CommentRepository) DeleteComment(ID uint) error {
	if result := r.DB.Delete(&Comment{}, ID); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *CommentRepository) GetAllComments() ([]Comment, error) {
	var comments []Comment
	if result := r.DB.Find(&comments); result.Error != nil {
		return []Comment{}, result.Error
	}
	return comments, nil
}