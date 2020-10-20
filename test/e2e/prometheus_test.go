package e2e

import (
	"context"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"testing"

	monitoringv1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	framework "github.com/operator-framework/operator-sdk/pkg/test"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func testPrometheus(t *testing.T) {
	testCtx, err := NewTestContext(t)
	require.NoError(t, err)

	err = framework.AddToFrameworkScheme(monitoringv1.AddToScheme, &monitoringv1.ServiceMonitor{})
	require.NoError(t, err)

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
	require.Equal(t, int32(len(windowsEndpoints.Subsets[0].Addresses)), gc.numberOfNodes)

	// check Nodes in the targetRef of Endpoints are same as the Windows Nodes bootstrapped using WMCO
	err = testIfNodesPresent(windowsEndpoints)
	require.NoError(t, err)

	// check Port name matches
	require.Equal(t, windowsEndpoints.Subsets[0].Ports[0].Name, "metrics")
	// check Port matches the defined port
	require.Equal(t, windowsEndpoints.Subsets[0].Ports[0].Port, int32(9182))

}

func testIfNodesPresent(windowsEndpoints *v1.Endpoints) error {
	for _, node := range gc.nodes {
		foundNode := false
		for _, endpointAddress := range windowsEndpoints.Subsets[0].Addresses {
			if node.Name == endpointAddress.TargetRef.Name {
				foundNode = true
				break
			}
		}
		if !foundNode {
			return errors.New("target node not found in Endpoints object ")
		}
	}
	return nil
}
