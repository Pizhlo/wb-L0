// Code generated by MockGen. DO NOT EDIT.
// Source: store.go

// Package mock_postgres is a generated GoMock package.
package mock_postgres

import (
	context "context"
	reflect "reflect"

	models "github.com/Pizhlo/wb-L0/internal/model"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockRepo is a mock of Repo interface.
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo.
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance.
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// GetAll mocks base method.
func (m *MockRepo) GetAll(ctx context.Context) ([]models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].([]models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockRepoMockRecorder) GetAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockRepo)(nil).GetAll), ctx)
}

// GetDeliveryByOrderID mocks base method.
func (m *MockRepo) GetDeliveryByOrderID(ctx context.Context, orderId uuid.UUID) (*models.Delivery, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeliveryByOrderID", ctx, orderId)
	ret0, _ := ret[0].(*models.Delivery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeliveryByOrderID indicates an expected call of GetDeliveryByOrderID.
func (mr *MockRepoMockRecorder) GetDeliveryByOrderID(ctx, orderId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeliveryByOrderID", reflect.TypeOf((*MockRepo)(nil).GetDeliveryByOrderID), ctx, orderId)
}

// GetItemByTrackNumber mocks base method.
func (m *MockRepo) GetItemByTrackNumber(ctx context.Context, trackNumber string) ([]models.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItemByTrackNumber", ctx, trackNumber)
	ret0, _ := ret[0].([]models.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItemByTrackNumber indicates an expected call of GetItemByTrackNumber.
func (mr *MockRepoMockRecorder) GetItemByTrackNumber(ctx, trackNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItemByTrackNumber", reflect.TypeOf((*MockRepo)(nil).GetItemByTrackNumber), ctx, trackNumber)
}

// GetOrderByID mocks base method.
func (m *MockRepo) GetOrderByID(ctx context.Context, id uuid.UUID) (*models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderByID", ctx, id)
	ret0, _ := ret[0].(*models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderByID indicates an expected call of GetOrderByID.
func (mr *MockRepoMockRecorder) GetOrderByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderByID", reflect.TypeOf((*MockRepo)(nil).GetOrderByID), ctx, id)
}

// GetPaymentByOrderID mocks base method.
func (m *MockRepo) GetPaymentByOrderID(ctx context.Context, orderId uuid.UUID) (*models.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPaymentByOrderID", ctx, orderId)
	ret0, _ := ret[0].(*models.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPaymentByOrderID indicates an expected call of GetPaymentByOrderID.
func (mr *MockRepoMockRecorder) GetPaymentByOrderID(ctx, orderId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPaymentByOrderID", reflect.TypeOf((*MockRepo)(nil).GetPaymentByOrderID), ctx, orderId)
}

// SaveDelivery mocks base method.
func (m *MockRepo) SaveDelivery(ctx context.Context, delivery models.Delivery) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveDelivery", ctx, delivery)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveDelivery indicates an expected call of SaveDelivery.
func (mr *MockRepoMockRecorder) SaveDelivery(ctx, delivery interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveDelivery", reflect.TypeOf((*MockRepo)(nil).SaveDelivery), ctx, delivery)
}

// SaveItems mocks base method.
func (m *MockRepo) SaveItems(ctx context.Context, items []models.Item) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveItems", ctx, items)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveItems indicates an expected call of SaveItems.
func (mr *MockRepoMockRecorder) SaveItems(ctx, items interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveItems", reflect.TypeOf((*MockRepo)(nil).SaveItems), ctx, items)
}

// SaveOrder mocks base method.
func (m *MockRepo) SaveOrder(ctx context.Context, order models.Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveOrder", ctx, order)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveOrder indicates an expected call of SaveOrder.
func (mr *MockRepoMockRecorder) SaveOrder(ctx, order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveOrder", reflect.TypeOf((*MockRepo)(nil).SaveOrder), ctx, order)
}

// SavePayment mocks base method.
func (m *MockRepo) SavePayment(ctx context.Context, payment models.Payment) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SavePayment", ctx, payment)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SavePayment indicates an expected call of SavePayment.
func (mr *MockRepoMockRecorder) SavePayment(ctx, payment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SavePayment", reflect.TypeOf((*MockRepo)(nil).SavePayment), ctx, payment)
}
