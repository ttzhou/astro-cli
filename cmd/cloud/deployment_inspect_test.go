package cloud

import (
	"fmt"
	"testing"

	astroplatformcore_mocks "github.com/astronomer/astro-cli/astro-client-platform-core/mocks"
	"github.com/stretchr/testify/mock"

	testUtil "github.com/astronomer/astro-cli/pkg/testing"
	"github.com/stretchr/testify/assert"
)

func TestNewDeploymentInspectCmd(t *testing.T) {
	expectedHelp := "Inspect an Astro Deployment configuration, which can be useful if you manage deployments as code or use Deployment configuration templates. This command returns the Deployment's configuration as a YAML or JSON output, which includes information about resources, such as cluster ID, region, and Airflow API URL, as well as scheduler and worker queue configurations."
	testUtil.InitTestConfig(testUtil.LocalPlatform)
	mockPlatformCoreClient := new(astroplatformcore_mocks.ClientWithResponsesInterface)
	platformCoreClient = mockPlatformCoreClient

	t.Run("-h prints help", func(t *testing.T) {
		cmdArgs := []string{"inspect", "-h"}
		resp, err := execDeploymentCmd(cmdArgs...)
		assert.NoError(t, err)
		assert.Contains(t, resp, expectedHelp)
	})
	t.Run("returns deployment in yaml format when a deployment name was provided", func(t *testing.T) {
		mockPlatformCoreClient.On("ListDeploymentsWithResponse", mock.Anything, mock.Anything, mock.Anything).Return(&mockListDeploymentsResponse, nil).Times(1)
		mockPlatformCoreClient.On("GetDeploymentWithResponse", mock.Anything, mock.Anything, mock.Anything).Return(&deploymentResponse, nil).Times(1)
		mockPlatformCoreClient.On("GetClusterWithResponse", mock.Anything, mock.Anything, mock.Anything).Return(&mockGetClusterResponse, nil).Times(1)

		cmdArgs := []string{"inspect", "-n", "test"}
		resp, err := execDeploymentCmd(cmdArgs...)
		assert.NoError(t, err)
		assert.Contains(t, resp, deploymentResponse.JSON200.Namespace)
		assert.Contains(t, resp, deploymentResponse.JSON200.Name)
		assert.Contains(t, resp, deploymentResponse.JSON200.RuntimeVersion)
		mockPlatformCoreClient.AssertExpectations(t)
	})
	t.Run("returns deployment in yaml format when a deployment id was provided", func(t *testing.T) {
		mockPlatformCoreClient.On("ListDeploymentsWithResponse", mock.Anything, mock.Anything, mock.Anything).Return(&mockListDeploymentsResponse, nil).Times(1)
		mockPlatformCoreClient.On("GetDeploymentWithResponse", mock.Anything, mock.Anything, mock.Anything).Return(&deploymentResponse, nil).Times(1)
		mockPlatformCoreClient.On("GetClusterWithResponse", mock.Anything, mock.Anything, mock.Anything).Return(&mockGetClusterResponse, nil).Times(1)
		cmdArgs := []string{"inspect", "test-id-1"}
		resp, err := execDeploymentCmd(cmdArgs...)
		assert.NoError(t, err)
		assert.Contains(t, resp, deploymentResponse.JSON200.Namespace)
		assert.Contains(t, resp, deploymentResponse.JSON200.Name)
		assert.Contains(t, resp, deploymentResponse.JSON200.RuntimeVersion)
		mockPlatformCoreClient.AssertExpectations(t)
	})
	t.Run("returns deployment template in yaml format when a deployment id was provided", func(t *testing.T) {
		mockPlatformCoreClient.On("ListDeploymentsWithResponse", mock.Anything, mock.Anything, mock.Anything).Return(&mockListDeploymentsResponse, nil).Times(1)
		mockPlatformCoreClient.On("GetDeploymentWithResponse", mock.Anything, mock.Anything, mock.Anything).Return(&deploymentResponse, nil).Times(1)
		mockPlatformCoreClient.On("GetClusterWithResponse", mock.Anything, mock.Anything, mock.Anything).Return(&mockGetClusterResponse, nil).Times(1)
		cmdArgs := []string{"inspect", "test-id-1", "--template"}
		resp, err := execDeploymentCmd(cmdArgs...)
		assert.NoError(t, err)
		assert.Contains(t, resp, deploymentResponse.JSON200.RuntimeVersion)
		mockPlatformCoreClient.AssertExpectations(t)
	})
	t.Run("returns a deployment's specific field", func(t *testing.T) {
		mockPlatformCoreClient.On("ListDeploymentsWithResponse", mock.Anything, mock.Anything, mock.Anything).Return(&mockListDeploymentsResponse, nil).Times(1)
		mockPlatformCoreClient.On("GetDeploymentWithResponse", mock.Anything, mock.Anything, mock.Anything).Return(&deploymentResponse, nil).Times(1)
		mockPlatformCoreClient.On("GetClusterWithResponse", mock.Anything, mock.Anything, mock.Anything).Return(&mockGetClusterResponse, nil).Times(1)
		cmdArgs := []string{"inspect", "-n", "test", "-k", "metadata.cluster_id"}
		_, err := execDeploymentCmd(cmdArgs...)
		assert.NoError(t, err)
		mockPlatformCoreClient.AssertExpectations(t)
	})
	t.Run("returns empty workload identity when show-workload-identity flag is not passed", func(t *testing.T) {
		mockPlatformCoreClient.On("ListDeploymentsWithResponse", mock.Anything, mock.Anything, mock.Anything).Return(&mockListDeploymentsResponse, nil).Times(1)
		mockPlatformCoreClient.On("GetDeploymentWithResponse", mock.Anything, mock.Anything, mock.Anything).Return(&deploymentResponse, nil).Times(1)
		mockPlatformCoreClient.On("GetClusterWithResponse", mock.Anything, mock.Anything, mock.Anything).Return(&mockGetClusterResponse, nil).Times(1)
		cmdArgs := []string{"inspect", "-n", "test"}
		out, err := execDeploymentCmd(cmdArgs...)
		assert.NoError(t, err)
		assert.Contains(t, out, `workload_identity: ""`)
		mockPlatformCoreClient.AssertExpectations(t)
	})
	t.Run("returns workload identity when show-workload-identity flag is passed", func(t *testing.T) {
		mockPlatformCoreClient.On("ListDeploymentsWithResponse", mock.Anything, mock.Anything, mock.Anything).Return(&mockListDeploymentsResponse, nil).Times(1)
		mockPlatformCoreClient.On("GetDeploymentWithResponse", mock.Anything, mock.Anything, mock.Anything).Return(&deploymentResponse, nil).Times(1)
		mockPlatformCoreClient.On("GetClusterWithResponse", mock.Anything, mock.Anything, mock.Anything).Return(&mockGetClusterResponse, nil).Times(1)
		cmdArgs := []string{"inspect", "-n", "test", "--show-workload-identity"}
		out, err := execDeploymentCmd(cmdArgs...)
		assert.NoError(t, err)
		assert.Contains(t, out, fmt.Sprintf(`workload_identity: %s`, mockWorkloadIdentity))
		mockPlatformCoreClient.AssertExpectations(t)
	})
	t.Run("returns an error when getting workspace fails", func(t *testing.T) {
		testUtil.InitTestConfig(testUtil.Initial)
		expectedOut := "Usage:\n"
		cmdArgs := []string{"inspect", "-n", "doesnotexist"}
		resp, err := execDeploymentCmd(cmdArgs...)
		assert.Error(t, err)
		assert.NotContains(t, resp, expectedOut)
	})
}
