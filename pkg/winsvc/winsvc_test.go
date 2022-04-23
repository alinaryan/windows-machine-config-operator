//go:build windows

package winsvc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"
)

// These test functions have been created to define the expected behavior of the structs mocking the Windows
// service API. The expected behavior should match up with the behavior seen when using the Windows API, and
// if differences are seen between the two, the mock implementations should be modified to correct the difference.

func TestCreateService(t *testing.T) {
	testIO := []struct {
		name         string
		svcName      string
		svcExepath   string
		svcConfig    mgr.Config
		existingSvcs map[string]*FakeService
		expectErr    bool
	}{
		{
			name:    "new service with no other services",
			svcName: "svc-one",
			svcConfig: mgr.Config{
				Description: "testsvc",
			},
			svcExepath:   "testpath",
			existingSvcs: nil,
			expectErr:    false,
		},
		{
			name:    "new service with different existing service",
			svcName: "svc-one",
			svcConfig: mgr.Config{
				Description: "testsvc",
			},
			svcExepath:   "testpath",
			existingSvcs: map[string]*FakeService{"svc-two": {}},
			expectErr:    false,
		},
		{
			name:    "existing service",
			svcName: "svc-one",
			svcConfig: mgr.Config{
				Description: "testsvc",
			},
			svcExepath:   "testpath",
			existingSvcs: map[string]*FakeService{"svc-one": {}},
			expectErr:    true,
		},
	}
	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			testMgr := NewTestMgr(test.existingSvcs)
			_, err := testMgr.CreateService(test.svcName, test.svcExepath, test.svcConfig)
			if test.expectErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			// ensure all existing keys are preserved
			for svcName := range test.existingSvcs {
				assert.Contains(t, testMgr.svcList.svcs, svcName)
			}
			// ensure new service has been added to list with correct values
			require.Contains(t, testMgr.svcList.svcs, test.svcName)
			newSvcInterface := testMgr.svcList.svcs[test.svcName]
			newSvc, ok := newSvcInterface.(*FakeService)
			require.True(t, ok, "cannot cast service to correct type")
			assert.Equal(t, test.svcName, newSvc.name)
			assert.Equal(t, test.svcExepath, newSvc.config.BinaryPathName)
			assert.Equal(t, test.svcConfig.Description, newSvc.config.Description)
			assert.Equal(t, svc.Stopped, newSvc.status.State)
		})
	}
}

func TestListServices(t *testing.T) {
	testIO := []struct {
		name         string
		existingSvcs map[string]*FakeService
		expected     []string
	}{
		{
			name:         "no services",
			existingSvcs: map[string]*FakeService{},
			expected:     []string{},
		},
		{
			name:         "some services",
			existingSvcs: map[string]*FakeService{"svc-one": {}, "svc-two": {}},
			expected:     []string{"svc-one", "svc-two"},
		},
	}
	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			testMgr := NewTestMgr(test.existingSvcs)
			list, err := testMgr.ListServices()
			require.NoError(t, err)
			assert.ElementsMatch(t, test.expected, list)
		})
	}
}

func TestOpenService(t *testing.T) {
	testIO := []struct {
		name         string
		svcName      string
		existingSvcs map[string]*FakeService
		expected     *FakeService
		expectErr    bool
	}{
		{
			name:         "existing service",
			svcName:      "svc-one",
			existingSvcs: map[string]*FakeService{"svc-one": {config: mgr.Config{Description: "testsvc"}}},
			expectErr:    false,
		},
		{
			name:         "nonexistent service",
			svcName:      "svc-two",
			existingSvcs: map[string]*FakeService{"svc-one": {config: mgr.Config{Description: "testsvc"}}},
			expectErr:    true,
		},
	}
	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			testMgr := NewTestMgr(test.existingSvcs)
			s, err := testMgr.OpenService(test.svcName)
			if test.expectErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, test.existingSvcs[test.svcName], s)
		})
	}
}

func TestDeleteService(t *testing.T) {
	testIO := []struct {
		name         string
		svcName      string
		existingSvcs map[string]*FakeService
	}{
		{
			name:         "service exists",
			svcName:      "svc-one",
			existingSvcs: map[string]*FakeService{"svc-one": {name: "svc-one", config: mgr.Config{Description: "testsvc"}}},
		},
	}
	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			testMgr := NewTestMgr(test.existingSvcs)
			// First check that the service exists in the service list
			list, err := testMgr.ListServices()
			require.NoError(t, err)
			require.Contains(t, list, test.svcName)
			// Open the service to get a service handle, and then delete the service
			s, err := testMgr.OpenService(test.svcName)
			require.NoError(t, err)
			require.NoError(t, s.Delete())
			// Check that the service is no longer present in the service list
			list, err = testMgr.ListServices()
			require.NoError(t, err)
			assert.NotContains(t, list, test.svcName)
		})
	}
}