package api

import (
	"github.com/TovarischSuhov/pipeline-manager/internal/pkg/gitlab/models"
	"github.com/prometheus/common/config"
)

func GetPipelinesList(projectID string) ([]models.Pipeline, error) {
	uri = config.URL + apiPrefix + "/" + projectID + "/pipelines"
}
