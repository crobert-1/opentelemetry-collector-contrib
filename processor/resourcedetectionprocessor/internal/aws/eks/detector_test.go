// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package eks

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/processor/processortest"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/aws/eks/internal/metadata"
)

const (
	clusterNameKey = "k8s.cluster.name"
	clusterName    = "my-cluster"
)

type MockDetectorUtils struct {
	mock.Mock
}

func (detectorUtils *MockDetectorUtils) getConfigMap(_ context.Context, namespace string, name string) (map[string]string, error) {
	args := detectorUtils.Called(namespace, name)
	return args.Get(0).(map[string]string), args.Error(1)
}

func (detectorUtils *MockDetectorUtils) getClusterName(_ context.Context) (string, error) {
	var reservations []*ec2.Reservation
	return detectorUtils.getClusterNameTagFromReservations(reservations), nil
}

func (detectorUtils *MockDetectorUtils) getClusterNameTagFromReservations(_ []*ec2.Reservation) string {
	return clusterName
}

func TestNewDetector(t *testing.T) {
	dcfg := CreateDefaultConfig()
	detector, err := NewDetector(processortest.NewNopCreateSettings(), dcfg)
	assert.NoError(t, err)
	assert.NotNil(t, detector)
}

// Tests EKS resource detector running in EKS environment
func TestEKS(t *testing.T) {
	detectorUtils := new(MockDetectorUtils)
	ctx := context.Background()

	t.Setenv("KUBERNETES_SERVICE_HOST", "localhost")
	detectorUtils.On("getConfigMap", authConfigmapNS, authConfigmapName).Return(map[string]string{clusterNameKey: clusterName}, nil)
	// Call EKS Resource detector to detect resources
	eksResourceDetector := &detector{utils: detectorUtils, err: nil, rb: metadata.NewResourceBuilder(metadata.DefaultResourceAttributesConfig())}
	res, _, err := eksResourceDetector.Detect(ctx)
	require.NoError(t, err)

	assert.Equal(t, map[string]interface{}{
		"cloud.provider": "aws",
		"cloud.platform": "aws_eks",
	}, res.Attributes().AsRaw(), "Resource object returned is incorrect")
}

// Tests EKS resource detector not running in EKS environment by verifying resource is not running on k8s
func TestNotEKS(t *testing.T) {
	eksResourceDetector := detector{logger: zap.NewNop()}
	r, _, err := eksResourceDetector.Detect(context.Background())
	require.NoError(t, err)
	assert.Equal(t, 0, r.Attributes().Len(), "Resource object should be empty")
}
