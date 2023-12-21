package mocks_test

import (
	"testing"

	"capstone/features/feedback"
	"capstone/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestDeleteFeedback(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockFeedbackRepository(ctrl)
	repo.EXPECT().Delete(1).Return(nil)

	err := repo.Delete(1)
	assert.NoError(t, err)
}

func TestFindAllFeedback(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockFeedbackRepository(ctrl)
	expectedFeedbacks := []feedback.Feedback{
		// ... populate with sample feedbacks for testing
	}
	repo.EXPECT().FindAll().Return(expectedFeedbacks, nil)

	feedbacks, err := repo.FindAll()
	assert.NoError(t, err)
	assert.Len(t, feedbacks, len(expectedFeedbacks))
}

func TestFindByIdFeedback(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockFeedbackRepository(ctrl)
	expectedFeedback := &feedback.Feedback{
		// ... populate with a sample feedback for testing
	}
	repo.EXPECT().FindById(1).Return(expectedFeedback, nil)

	feedback, err := repo.FindById(1)
	assert.NoError(t, err)
	assert.Equal(t, expectedFeedback, feedback)
}

func TestSaveFeedback(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockFeedbackRepository(ctrl)
	expectedFeedback := &feedback.Feedback{
		// ... populate with a sample feedback for testing
	}
	repo.EXPECT().Save(expectedFeedback).Return(expectedFeedback, nil)

	feedback, err := repo.Save(expectedFeedback)
	assert.NoError(t, err)
	assert.Equal(t, expectedFeedback, feedback)
}

func TestUpdateFeedback(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockFeedbackRepository(ctrl)
	feedbackToUpdate := &feedback.Feedback{
		// ... populate with a sample feedback for testing
	}
	repo.EXPECT().Update(feedbackToUpdate, 1).Return(feedbackToUpdate, nil)

	feedback, err := repo.Update(feedbackToUpdate, 1)
	assert.NoError(t, err)
	assert.Equal(t, feedbackToUpdate, feedback)
}

// Additional tests can be added to handle error scenarios, etc.
