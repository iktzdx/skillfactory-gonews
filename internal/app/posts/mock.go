package posts

import (
	"github.com/stretchr/testify/mock"

	"github.com/iktzdx/skillfactory-gonews/pkg/storage"
)

type MockBoundaryRepoPort struct {
	mock.Mock
}

func (mockRepo *MockBoundaryRepoPort) FindPostByID(id int) (storage.Data, error) {
	args := mockRepo.Called(id)

	return args.Get(0).(storage.Data), args.Error(1) //nolint:forcetypeassert,wrapcheck
}

func (mockRepo *MockBoundaryRepoPort) List(opts storage.SearchOpts) (storage.BulkData, error) {
	args := mockRepo.Called(opts)

	return args.Get(0).(storage.BulkData), args.Error(1) //nolint:forcetypeassert,wrapcheck
}