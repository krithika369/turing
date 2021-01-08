// Code generated by mockery v2.1.0. DO NOT EDIT.

package mocks

import (
	merlinclient "github.com/gojek/merlin/client"
	client "github.com/gojek/mlp/client"

	mock "github.com/stretchr/testify/mock"

	models "github.com/gojek/turing/api/turing/models"
)

// AlertService is an autogenerated mock type for the AlertService type
type AlertService struct {
	mock.Mock
}

// Delete provides a mock function with given fields: alert, authorEmail, dashboardURL
func (_m *AlertService) Delete(alert models.Alert, authorEmail string, dashboardURL string) error {
	ret := _m.Called(alert, authorEmail, dashboardURL)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Alert, string, string) error); ok {
		r0 = rf(alert, authorEmail, dashboardURL)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindByID provides a mock function with given fields: id
func (_m *AlertService) FindByID(id uint) (*models.Alert, error) {
	ret := _m.Called(id)

	var r0 *models.Alert
	if rf, ok := ret.Get(0).(func(uint) *models.Alert); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Alert)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDashboardURL provides a mock function with given fields: alert, project, environment, router, routerVersion
func (_m *AlertService) GetDashboardURL(alert *models.Alert, project *client.Project, environment *merlinclient.Environment, router *models.Router, routerVersion *models.RouterVersion) (string, error) {
	ret := _m.Called(alert, project, environment, router, routerVersion)

	var r0 string
	if rf, ok := ret.Get(0).(func(*models.Alert, *client.Project, *merlinclient.Environment, *models.Router, *models.RouterVersion) string); ok {
		r0 = rf(alert, project, environment, router, routerVersion)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Alert, *client.Project, *merlinclient.Environment, *models.Router, *models.RouterVersion) error); ok {
		r1 = rf(alert, project, environment, router, routerVersion)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: _a0
func (_m *AlertService) List(_a0 string) ([]*models.Alert, error) {
	ret := _m.Called(_a0)

	var r0 []*models.Alert
	if rf, ok := ret.Get(0).(func(string) []*models.Alert); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Alert)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: alert, authorEmail, dashboardURL
func (_m *AlertService) Save(alert models.Alert, authorEmail string, dashboardURL string) (*models.Alert, error) {
	ret := _m.Called(alert, authorEmail, dashboardURL)

	var r0 *models.Alert
	if rf, ok := ret.Get(0).(func(models.Alert, string, string) *models.Alert); ok {
		r0 = rf(alert, authorEmail, dashboardURL)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Alert)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Alert, string, string) error); ok {
		r1 = rf(alert, authorEmail, dashboardURL)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: alert, authorEmail, dashboardURL
func (_m *AlertService) Update(alert models.Alert, authorEmail string, dashboardURL string) error {
	ret := _m.Called(alert, authorEmail, dashboardURL)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Alert, string, string) error); ok {
		r0 = rf(alert, authorEmail, dashboardURL)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
