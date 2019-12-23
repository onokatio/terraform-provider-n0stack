package main

import (
        "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
        return &schema.Provider{
		Schema: map[string]*schema.Schema{},
                ResourcesMap: map[string]*schema.Resource{
			"n0stack_blockstorage": resource_n0stack_blockstorage(),
		},
        }
}
