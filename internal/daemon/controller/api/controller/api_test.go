package controller_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/davidsbond/orca/internal/daemon/controller/api/controller"
	"github.com/davidsbond/orca/internal/daemon/controller/api/controller/mocks"
	"github.com/davidsbond/orca/internal/daemon/controller/database/worker"
	controllersvcv1 "github.com/davidsbond/orca/internal/proto/orca/private/controller/service/v1"
	workerv1 "github.com/davidsbond/orca/internal/proto/orca/private/worker/v1"
)

func TestAPI_RegisterWorker(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name         string
		Request      *controllersvcv1.RegisterWorkerRequest
		Expected     *controllersvcv1.RegisterWorkerResponse
		ExpectsError bool
		ExpectedCode codes.Code
		Setup        func(svc *mocks.Service)
	}{
		{
			Name: "registration failed",
			Request: &controllersvcv1.RegisterWorkerRequest{
				Worker: &workerv1.Worker{
					Id:               "test-id",
					AdvertiseAddress: "localhost:1337",
					Workflows:        []string{"TestWorkflow"},
					Tasks:            []string{"TestTask"},
				},
			},
			ExpectsError: true,
			ExpectedCode: codes.Internal,
			Setup: func(svc *mocks.Service) {
				w := worker.Worker{
					ID:               "test-id",
					AdvertiseAddress: "localhost:1337",
					Workflows:        []string{"TestWorkflow"},
					Tasks:            []string{"TestTask"},
				}

				svc.EXPECT().
					RegisterWorker(mock.Anything, w).
					Return(errors.New("some failure")).
					Once()
			},
		},
		{
			Name: "registration successful",
			Request: &controllersvcv1.RegisterWorkerRequest{
				Worker: &workerv1.Worker{
					Id:               "test-id",
					AdvertiseAddress: "localhost:1337",
					Workflows:        []string{"TestWorkflow"},
					Tasks:            []string{"TestTask"},
				},
			},
			Expected: &controllersvcv1.RegisterWorkerResponse{},
			Setup: func(svc *mocks.Service) {
				w := worker.Worker{
					ID:               "test-id",
					AdvertiseAddress: "localhost:1337",
					Workflows:        []string{"TestWorkflow"},
					Tasks:            []string{"TestTask"},
				}

				svc.EXPECT().
					RegisterWorker(mock.Anything, w).
					Return(nil).
					Once()
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			svc := mocks.NewService(t)
			if tc.Setup != nil {
				tc.Setup(svc)
			}

			actual, err := controller.NewAPI(svc).RegisterWorker(t.Context(), tc.Request)
			if tc.ExpectsError {
				assert.Error(t, err)
				assert.EqualValues(t, tc.ExpectedCode, status.Code(err))
				assert.Nil(t, actual)
				return
			}

			assert.NoError(t, err)
			assert.EqualValues(t, tc.Expected, actual)
		})
	}
}

func TestAPI_DeregisterWorker(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name         string
		Request      *controllersvcv1.DeregisterWorkerRequest
		Expected     *controllersvcv1.DeregisterWorkerResponse
		ExpectsError bool
		ExpectedCode codes.Code
		Setup        func(svc *mocks.Service)
	}{
		{
			Name: "deregistration failed",
			Request: &controllersvcv1.DeregisterWorkerRequest{
				WorkerId: "test-id",
			},
			ExpectsError: true,
			ExpectedCode: codes.Internal,
			Setup: func(svc *mocks.Service) {
				svc.EXPECT().
					DeregisterWorker(mock.Anything, "test-id").
					Return(errors.New("some failure")).
					Once()
			},
		},
		{
			Name: "deregistration successful",
			Request: &controllersvcv1.DeregisterWorkerRequest{
				WorkerId: "test-id",
			},
			Expected: &controllersvcv1.DeregisterWorkerResponse{},
			Setup: func(svc *mocks.Service) {
				svc.EXPECT().
					DeregisterWorker(mock.Anything, "test-id").
					Return(nil).
					Once()
			},
		},
		{
			Name: "worker not found",
			Request: &controllersvcv1.DeregisterWorkerRequest{
				WorkerId: "test-id",
			},
			ExpectsError: true,
			ExpectedCode: codes.NotFound,
			Setup: func(svc *mocks.Service) {
				svc.EXPECT().
					DeregisterWorker(mock.Anything, "test-id").
					Return(controller.ErrWorkerNotFound).
					Once()
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			svc := mocks.NewService(t)
			if tc.Setup != nil {
				tc.Setup(svc)
			}

			actual, err := controller.NewAPI(svc).DeregisterWorker(t.Context(), tc.Request)
			if tc.ExpectsError {
				assert.Error(t, err)
				assert.EqualValues(t, tc.ExpectedCode, status.Code(err))
				assert.Nil(t, actual)
				return
			}

			assert.NoError(t, err)
			assert.EqualValues(t, tc.Expected, actual)
		})
	}
}
