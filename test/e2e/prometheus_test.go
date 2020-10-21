package e2e

import (
	"context"
	framework "github.com/operator-framework/operator-sdk/pkg/test"
	"go/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	core "k8s.io/api/core/v1"
	monitoringv1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
)

func testPrometheus(t *testing.T){
	testCtx, err := NewTestContext(t)
	require.noError(t, err)

	// check that service exists
	_, err = testCtx.kubeclient.CoreV1().Services("openshift-windows-machine-config-operator").Get(context.TODO(), "windows-machine-config-operator-metrics", metav1.GetOptions{})
	require.NoError(t, err)

	// check that SM existS
	windowsServiceMonitor := &monitoringv1.ServiceMonitor{}
	framework.Global.Client.Get(context.TODO()

	// check that endpoints exists
	windowsEndpoints, err := testCtx.kubeclient.CoreV1().Endpoints("openshift-windows-machine-config-operator").Get(context.TODO(), "windows-machine-config-operator-metrics", metav1.GetOptions{})
	require.NoError(t, err)
	require.Equal(t, len(windowsEndpoints.Subsets[0].Addresses), gc.nodes)

	// check Port name matches
	require.Equal(t, windowsEndpoints.Subsets[0].Ports[0].Name, "windows-metrics")
}