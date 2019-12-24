package main

import (
	"context"
	"google.golang.org/grpc"
        "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ppool "github.com/onokatio/terraform-provider-n0stack/n0proto.go/pool/v0"
)

func resource_n0stack_network() *schema.Resource {
	return &schema.Resource{
		Create: resource_n0stack_network_create,
		Read: resource_n0stack_network_read,
		Update: resource_n0stack_network_update,
		Delete: resource_n0stack_network_delete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ipv4_cidr": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ipv6_cidr": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"annotations": {
				Type:     schema.TypeMap,
				Required: true,
				Elem: schema.TypeString,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: schema.TypeString,
			},
		},
	}
}

func resource_n0stack_network_create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(Config)
	conn, err := grpc.Dial(config.endpoint, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := ppool.NewNetworkServiceClient(conn)

	request := ppool.ApplyNetworkRequest{
		Name: d.Get("name").(string),
		Ipv4Cidr: d.Get("ipv4_cidr").(string),
		Ipv6Cidr: d.Get("ipv6_cidr").(string),
		Annotations: interfaceMap2stringMap(d.Get("annotations").(map[string]interface{})),
		Labels: interfaceMap2stringMap(d.Get("labels").(map[string]interface{})),
	}
	_, err = client.ApplyNetwork(context.Background(), &request)
	if err != nil {
		return err
	}

	d.SetId(d.Get("name").(string))

	return resource_n0stack_network_read(d, meta)
}

func resource_n0stack_network_read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(Config)
	conn, err := grpc.Dial(config.endpoint, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := ppool.NewNetworkServiceClient(conn)

	request := ppool.GetNetworkRequest{
		Name: d.Get("name").(string) ,
	}
	res, err := client.GetNetwork(context.Background(), &request)
	if err != nil {
		return err
	}

	d.Set("name",res.Name)
	d.Set("ipv4_cidr",res.Ipv4Cidr)
	d.Set("ipv6_cidr",res.Ipv6Cidr)
	d.Set("annotations", res.Annotations)
	d.Set("labels", res.Labels)
	return nil
}

func resource_n0stack_network_update(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resource_n0stack_network_delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(Config)
	conn, err := grpc.Dial(config.endpoint, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := ppool.NewNetworkServiceClient(conn)

	request := ppool.DeleteNetworkRequest{
		Name: d.Get("name").(string) ,
	}
	_, err = client.DeleteNetwork(context.Background(), &request)
	if err != nil {
		return err
	}

	return resource_n0stack_network_read(d, meta)
}
