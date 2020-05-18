package servicediscovery

import (
	common2 "github.com/breathbath/api-go-wrapper/pkg/api/common"
)

type getServiceEndpointsResponse struct {
	Status  common2.Status
	Records []ServiceEndpoints `json:"records"`
}

type ServiceEndpoints struct {
	Cafa        Endpoint `json:"cafa"`
	Pim         Endpoint `json:"pim"`
	Wms         Endpoint `json:"wms"`
	Promotion   Endpoint `json:"promotion"`
	Reports     Endpoint `json:"reports"`
	Json        Endpoint `json:"json"`
	Assignments Endpoint `json:"assignments"`
}
type Endpoint struct {
	Url           string `json:"url"`
	Documentation string `json:"documentation"`
}
