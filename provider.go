package main

import (
        "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
                ResourcesMap: map[string]*schema.Resource{
			"n0stack_blockstorage": resource_n0stack_blockstorage(),
			"n0stack_image": resource_n0stack_image(),
			"n0stack_network": resource_n0stack_network(),
		},
        }
	provider.ConfigureFunc = func (d *schema.ResourceData) (interface{}, error) {
		endpoint := d.Get("endpoint").(string)
		config := Config{
			endpoint: endpoint,
		}
		return config, nil
	}
        return provider
}

type Config struct {
	endpoint string
}

