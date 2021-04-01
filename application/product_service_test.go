package application_test

import (
	"github.com/felipeesc/hexagonal-go/application"
	mock_application "github.com/felipeesc/hexagonal-go/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	product := mock_application.NewMockProductInterface(ctrl)
	persistece := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistece.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()
	service := application.ProductService{Persistence: persistece}

	result, err := service.Get("abc")
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	product := mock_application.NewMockProductInterface(ctrl)
	persistece := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistece.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()
	service := application.ProductService{Persistence: persistece}

	result, err := service.Create("product 1", 10)
	require.Nil(t, err)
	require.Equal(t, product, result)

}

func TestProductService_Enable(t *testing.T) {

}

func TestProductService_Disable(t *testing.T) {

}
