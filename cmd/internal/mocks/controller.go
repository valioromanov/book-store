// Code generated by MockGen. DO NOT EDIT.
// Source: controller.go

// Package mocks is a generated GoMock package.
package mocks

import (
	repository "book-store/pkg/repository"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockBookRepository is a mock of BookRepository interface.
type MockBookRepository struct {
	ctrl     *gomock.Controller
	recorder *MockBookRepositoryMockRecorder
}

// MockBookRepositoryMockRecorder is the mock recorder for MockBookRepository.
type MockBookRepositoryMockRecorder struct {
	mock *MockBookRepository
}

// NewMockBookRepository creates a new mock instance.
func NewMockBookRepository(ctrl *gomock.Controller) *MockBookRepository {
	mock := &MockBookRepository{ctrl: ctrl}
	mock.recorder = &MockBookRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBookRepository) EXPECT() *MockBookRepositoryMockRecorder {
	return m.recorder
}

// DeleteBook mocks base method.
func (m *MockBookRepository) DeleteBook(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBook", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBook indicates an expected call of DeleteBook.
func (mr *MockBookRepositoryMockRecorder) DeleteBook(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBook", reflect.TypeOf((*MockBookRepository)(nil).DeleteBook), arg0)
}

// FindById mocks base method.
func (m *MockBookRepository) FindById(id int) (repository.BookDB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", id)
	ret0, _ := ret[0].(repository.BookDB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockBookRepositoryMockRecorder) FindById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockBookRepository)(nil).FindById), id)
}

// InsertBook mocks base method.
func (m *MockBookRepository) InsertBook(arg0 repository.BookDB) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertBook", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertBook indicates an expected call of InsertBook.
func (mr *MockBookRepositoryMockRecorder) InsertBook(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertBook", reflect.TypeOf((*MockBookRepository)(nil).InsertBook), arg0)
}

// UpdateBook mocks base method.
func (m *MockBookRepository) UpdateBook(arg0 repository.BookDB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBook", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBook indicates an expected call of UpdateBook.
func (mr *MockBookRepositoryMockRecorder) UpdateBook(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBook", reflect.TypeOf((*MockBookRepository)(nil).UpdateBook), arg0)
}
