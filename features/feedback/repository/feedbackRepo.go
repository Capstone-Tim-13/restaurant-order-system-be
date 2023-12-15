package repository

import (
	"capstone/features/feedback"

	"gorm.io/gorm"
)

type FeedbackRepositoryImpl struct {
	db *gorm.DB
}

func NewFeedbackRepository(db *gorm.DB) feedback.Repository {
	return &FeedbackRepositoryImpl{db: db}
}

func (r *FeedbackRepositoryImpl) Save(feedback *feedback.Feedback) (*feedback.Feedback, error) {
	result := r.db.Create(feedback)
	if result.Error != nil {
		return nil, result.Error
	}

	return feedback, nil
}

func (r *FeedbackRepositoryImpl) Update(feedbacks *feedback.Feedback, id int) (*feedback.Feedback, error) {
	result := r.db.Where("id = ?", id).Updates(feedback.Feedback{Rating: feedbacks.Rating, Review: feedbacks.Review})
	if result.Error != nil {
		return nil, result.Error
	}

	return feedbacks, nil
}

func (r *FeedbackRepositoryImpl) FindById(id int) (*feedback.Feedback, error) {
	feedback := feedback.Feedback{}

	result := r.db.Preload("User").First(&feedback, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &feedback, nil
}

func (r *FeedbackRepositoryImpl) FindAll() ([]feedback.Feedback, error) {
	feedback := []feedback.Feedback{}

	result := r.db.Preload("User").Find(&feedback)
	if result.Error != nil {
		return nil, result.Error
	}

	return feedback, nil
}

func (r *FeedbackRepositoryImpl) Delete(id int) error {
	result := r.db.Table("feedbacks").Where("id = ?", id).Unscoped().Delete(id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
