package api

import (
	"github.com/grafana/grafana/pkg/api/dtos"
	"github.com/grafana/grafana/pkg/api/response"
	"github.com/grafana/grafana/pkg/models"
)

var DemoCatalog = dtos.Catalog{
	{
		Name: "cart",
		Components: []dtos.CatalogComponent{
			{
				Name: "cart-api",
				Labels: map[string]string{
					"k8s.grafana.com/service": "cart",
					"k8s.grafana.com/teams":   "cart,billing",
					"owner":                   "brian",
				},
				Teams:     []string{"cart", "billing"},
				Namespace: "sock-shop",
				Address:   "cart-api.sock-shop.svc.cluster.local",
				Pods: []dtos.CatalogPod{
					{
						Name:   "cart-api-asdf-1",
						Status: dtos.PodStatusGreen,
					},
					{
						Name:   "cart-api-asdf-2",
						Status: dtos.PodStatusGreen,
					},
					{
						Name:   "cart-api-asdf-3",
						Status: dtos.PodStatusRed,
					},
				},
				External: false,
			},
			{
				Name: "cart-db",
				Labels: map[string]string{
					"k8s.grafana.com/service": "cart",
					"k8s.grafana.com/teams":   "cart,billing,dba",
					"owner":                   "brian",
				},
				Teams:     []string{"cart", "billing"},
				Namespace: "sock-shop",
				Address:   "cart-db.sock-shop.svc.cluster.local",
				Pods: []dtos.CatalogPod{
					{
						Name:   "cart-db-asdf-1",
						Status: dtos.PodStatusGreen,
					},
					{
						Name:   "cart-db-asdf-2",
						Status: dtos.PodStatusGreen,
					},
					{
						Name:   "cart-db-asdf-3",
						Status: dtos.PodStatusRed,
					},
				},
				External: false,
			},
			{
				Name: "cart-redis",
				Labels: map[string]string{
					"k8s.grafana.com/service": "cart",
					"k8s.grafana.com/teams":   "cart,billing,dba",
					"owner":                   "brian",
				},
				Teams:     []string{"cart", "billing"},
				Namespace: "sock-shop",
				Address:   "example.redis.io",
				Pods:      nil,
				External:  true,
			},
		},
	},
	{
		Name: "frontend",
		Components: []dtos.CatalogComponent{
			{
				Name: "nginx",
				Labels: map[string]string{
					"k8s.grafana.com/service": "frontend",
					"k8s.grafana.com/teams":   "frontend",
				},
				Teams:     []string{"frontend"},
				Namespace: "sock-shop",
				Address:   "nginx.sock-shop.svc.cluster.local",
				Pods: []dtos.CatalogPod{
					{
						Name:   "nginx-asdf-1",
						Status: dtos.PodStatusGreen,
					},
					{
						Name:   "cart-api-asdf-2",
						Status: dtos.PodStatusGreen,
					},
					{
						Name:   "cart-api-asdf-3",
						Status: dtos.PodStatusGreen,
					},
				},
				External: false,
			},
			{
				Name: "varnish",
				Labels: map[string]string{
					"k8s.grafana.com/service": "frontend",
					"k8s.grafana.com/teams":   "frontend",
				},
				Teams:     []string{"frontend"},
				Namespace: "sock-shop",
				Address:   "varnish.sock-shop.svc.cluster.local",
				Pods: []dtos.CatalogPod{
					{
						Name:   "varnish-asdf-1",
						Status: dtos.PodStatusGreen,
					},
				},
				External: false,
			},
		},
	},
}

// HandleGetCatalog gets the full list of catalog services
func (hs *HTTPServer) HandleGetCatalog(c *models.ReqContext) response.Response {
	return response.JSON(200, DemoCatalog)
}
