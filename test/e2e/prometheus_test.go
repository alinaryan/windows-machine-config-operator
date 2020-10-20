package e2e

import (
	"context"
	framework "github.com/operator-framework/operator-sdk/pkg/test"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"testing"

	monitoringv1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	"github.com/stretchr/testify/require"
)

func testPrometheus(t *testing.T) {
	testCtx, err := NewTestContext(t)
	require.noError(t, err)

	// check that service exists
	_, err = testCtx.kubeclient.CoreV1().Services("openshift-windows-machine-config-operator").Get(context.TODO(), "windows-machine-config-operator-metrics", metav1.GetOptions{})
	require.NoError(t, err)

	// check that SM existS
	windowsServiceMonitor := &monitoringv1.ServiceMonitor{}
	err = framework.Global.Client.Get(context.TODO(), types.NamespacedName{Namespace: "openshift-windows-machine-config-operator", Name: "windows-machine-config-operator-metrics"}, windowsServiceMonitor)
	require.NoError(t, err)
	// check that endpoints exists
	windowsEndpoints, err := testCtx.kubeclient.CoreV1().Endpoints("openshift-windows-machine-config-operator").Get(context.TODO(), "windows-machine-config-operator-metrics", metav1.GetOptions{})
	require.NoError(t, err)
	require.Equal(t, len(windowsEndpoints.Subsets[0].Addresses), gc.nodes)

	// check Port name matches
	require.Equal(t, windowsEndpoints.Subsets[0].Ports[0].Name, "windows-metrics")
}
