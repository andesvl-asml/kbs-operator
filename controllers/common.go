/*
Copyright Confidential Containers Contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

const (

	// KbsFinalizerName for KbsConfig
	KbsFinalizerName = "kbsconfig.confidentialcontainers.org/finalizer"

	// KBS Deployment name
	KbsDeploymentName = "kbs-deployment"

	// KBS operator default namespace
	KbsOperatorNamespace = "kbs-operator-system"

	// Default KBS image name
	DefaultKbsImageName = "europe-docker.pkg.dev/asml-dpng-dev-shared/dpng/ghcr.io/confidential-containers/key-broker-service:latest"

	// Default AS image name
	DefaultAsImageName = "europe-docker.pkg.dev/asml-dpng-dev-shared/dpng/ghcr.io/confidential-containers/attestation-service:latest"

	// Default RVPS image name
	DefaultRvpsImageName = "europe-docker.pkg.dev/asml-dpng-dev-shared/dpng/ghcr.io/confidential-containers/reference-value-provider-service:latest"

	// KBS service name
	KbsServiceName = "kbs-service"

	// Default KBS Resources Path
	kbsResourcesPath = "/opt/confidential-containers/kbs/repository/default"

	// Default KBS config path
	kbsDefaultConfigPath = "/etc"

	// Default AS config path
	asDefaultConfigPath = "/etc"

	// Default RVPS config path
	rvpsDefaultConfigPath = "/etc"

	// Default RVPS reference values Path
	rvpsReferenceValuesPath = "/opt/confidential-containers/rvps"
)

func contains(list []string, s string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}

// Add remove method to remove element from slice
func remove(slice []string, s string) []string {
	var result []string
	for _, v := range slice {
		if v != s {
			result = append(result, v)
		}
	}
	return result
}
