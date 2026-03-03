package users

import (
	"context"
	"io"
	"log"
	"testing"

	"github.com/solanoize/goblog/internal/auth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type mockRepository struct {
	mock.Mock
}

// Create implements [Repository].
func (m *mockRepository) Create(ctx context.Context, user User) (User, error) {
	panic("unimplemented")
}

// Delete implements [Repository].
func (m *mockRepository) Delete(ctx context.Context, id uint) error {
	panic("unimplemented")
}

// FindByEmail implements [Repository].
func (m *mockRepository) FindByEmail(ctx context.Context, email string) (User, error) {
	panic("unimplemented")
}

// Update implements [Repository].
func (m *mockRepository) Update(ctx context.Context, user User) (User, error) {
	panic("unimplemented")
}

func (m *mockRepository) FindByID(ctx context.Context, id uint) (User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return User{}, args.Error(1)
	}
	return args.Get(0).(User), args.Error(1)
}

func NewMockRepository() *mockRepository {
	return &mockRepository{}
}

func TestUserService_GetUserName(t *testing.T) {
	logger := log.New(io.Discard, "", 0)
	mockRepo := &mockRepository{}
	authService := auth.NewService(logger)

	ctx := context.Background()

	mockRepo.
		On("FindByID", ctx, uint(1)).
		Return(User{Model: gorm.Model{
			ID: 1,
		}, Username: "Budi"}, nil)

	service := NewService(mockRepo, authService, logger)

	responseContract, err := service.GetByID(ctx, 1)

	assert.NoError(t, err)
	assert.Equal(t, "Budi", responseContract.Username)
	assert.Equal(t, false, responseContract.IsAdmin)

	mockRepo.AssertExpectations(t)
}
