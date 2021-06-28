package catalog

import (
	"github.com/grafana/grafana/pkg/registry"
)

const ServiceName = "Catalog"

type Service struct {
}

func (s *Service) Init() error {
	return nil
}

func init() {
	registry.Register(&registry.Descriptor{
		Name:         ServiceName,
		Instance:     &Service{},
		InitPriority: registry.High,
	})
}
