package elasticsearch

import (
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/husnialhamdani/eck-custom-resources/apis/es.eck/v1alpha1"
	"github.com/husnialhamdani/eck-custom-resources/utils"
	ctrl "sigs.k8s.io/controller-runtime"
)

func DeleteIndexLifecyclePolicy(esClient *elasticsearch.Client, indexLifecyclePolicyName string) (ctrl.Result, error) {
	res, err := esClient.ILM.DeleteLifecycle(indexLifecyclePolicyName)
	if err != nil || res.IsError() {
		return utils.GetRequeueResult(), err
	}
	return ctrl.Result{}, nil
}

func UpsertIndexLifecyclePolicy(esClient *elasticsearch.Client, indexLifecyclePolicy v1alpha1.IndexLifecyclePolicy) (ctrl.Result, error) {
	res, err := esClient.ILM.PutLifecycle(
		indexLifecyclePolicy.Name,
		esClient.ILM.PutLifecycle.WithBody(strings.NewReader(indexLifecyclePolicy.Spec.Body)),
	)

	if err != nil || res.IsError() {
		return utils.GetRequeueResult(), GetClientErrorOrResponseError(err, res)
	}

	return ctrl.Result{}, nil
}
