// Copyright 2018 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package common

import (
	api "github.com/kubeflow/pipelines/backend/api/go_client"
	"github.com/kubeflow/pipelines/backend/src/apiserver/model"
	"github.com/kubeflow/pipelines/backend/src/common/util"
)

const (
	Experiment      model.ResourceType = "Experiment"
	Job             model.ResourceType = "Job"
	Run             model.ResourceType = "Run"
	Pipeline        model.ResourceType = "pipeline"
	PipelineVersion model.ResourceType = "PipelineVersion"
	Namespace       model.ResourceType = "Namespace"
)

const (
	RbacKubeflowGroup    = "kubeflow.org"
	RbacPipelinesGroup   = "pipelines.kubeflow.org"
	RbacPipelinesVersion = "v1beta1"

	RbacResourceTypePipelines      = "pipelines"
	RbacResourceTypeExperiments    = "experiments"
	RbacResourceTypeRuns           = "runs"
	RbacResourceTypeJobs           = "jobs"
	RbacResourceTypeViewers        = "viewers"
	RbacResourceTypeVisualizations = "visualizations"

	RbacResourceVerbArchive   = "archive"
	RbacResourceVerbUpdate    = "update"
	RbacResourceVerbCreate    = "create"
	RbacResourceVerbDelete    = "delete"
	RbacResourceVerbDisable   = "disable"
	RbacResourceVerbEnable    = "enable"
	RbacResourceVerbGet       = "get"
	RbacResourceVerbList      = "list"
	RbacResourceVerbRetry     = "retry"
	RbacResourceVerbTerminate = "terminate"
	RbacResourceVerbUnarchive = "unarchive"
)

const (
	Owner   model.Relationship = "Owner"
	Creator model.Relationship = "Creator"
)

const (
	GoogleIAPUserIdentityHeader    string = "x-goog-authenticated-user-email"
	GoogleIAPUserIdentityPrefix    string = "accounts.google.com:"
	AuthorizationBearerTokenHeader string = "Authorization"
	AuthorizationBearerTokenPrefix string = "Bearer "
)

const DefaultTokenReviewAudience string = "pipelines.kubeflow.org"

const (
	DefaultArtifactBucket         string = "mlpipeline"
	DefaultArtifactEndpoint       string = "minio-service.kubeflow:9000"
	DefaultArtifactEndpointScheme string = "http://"
	DefaultArtifactImage          string = "minio/mc"
)

const (
	ArtifactItemsAnnotation          string = "tekton.dev/artifact_items"
	ArtifactBucketAnnotation         string = "tekton.dev/artifact_bucket"
	ArtifactEndpointAnnotation       string = "tekton.dev/artifact_endpoint"
	ArtifactEndpointSchemeAnnotation string = "tekton.dev/artifact_endpoint_scheme"
)

// For backward compatibility. Remove after 0.3 release
const DefaultArtifactScript string = "push_artifact() {\n" +
	"    tar -cvzf $1.tgz $2\n" +
	"    mc cp $1.tgz storage/$ARTIFACT_BUCKET/artifacts/$PIPELINERUN/$PIPELINETASK/$1.tgz\n" +
	"}\n" +
	"push_log() {" +
	"    cat /var/log/containers/$PODNAME*$NAMESPACE*step-main*.log > step-main.log\n" +
	"    push_artifact main-log step-main.log\n" +
	"}\n" +
	"strip_eof() {\n" +
	"    awk 'NF' $2 | head -c -1 > $1_temp_save && cp $1_temp_save $2\n" +
	"}\n" +
	"mc config host add storage ${ARTIFACT_ENDPOINT_SCHEME}${ARTIFACT_ENDPOINT} $AWS_ACCESS_KEY_ID $AWS_SECRET_ACCESS_KEY\n"

// Default Tekton 0.25.0+ Terminate Status.
const DefaultTerminateStatus string = "Cancelled"

const DefaultPath4InternalResults string = "/tekton/home/tep-results"

func ToModelResourceType(apiType api.ResourceType) (model.ResourceType, error) {
	switch apiType {
	case api.ResourceType_EXPERIMENT:
		return Experiment, nil
	case api.ResourceType_JOB:
		return Job, nil
	case api.ResourceType_PIPELINE_VERSION:
		return PipelineVersion, nil
	case api.ResourceType_NAMESPACE:
		return Namespace, nil
	default:
		return "", util.NewInvalidInputError("Unsupported resource type: %s", api.ResourceType_name[int32(apiType)])
	}
}

func ToModelRelationship(r api.Relationship) (model.Relationship, error) {
	switch r {
	case api.Relationship_CREATOR:
		return Creator, nil
	case api.Relationship_OWNER:
		return Owner, nil
	default:
		return "", util.NewInvalidInputError("Unsupported resource relationship: %s", api.Relationship_name[int32(r)])
	}
}
