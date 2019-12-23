package main

import (
	"context"
	"google.golang.org/grpc"
        "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	pprovisioning "github.com/onokatio/terraform-provider-n0stack/n0proto.go/provisioning/v0"
)

func resource_n0stack_blockstorage() *schema.Resource {
	return &schema.Resource{
		Create: resource_n0stack_blockstorage_create,
		Read: resource_n0stack_blockstorage_read,
		Update: resource_n0stack_blockstorage_update,
		Delete: resource_n0stack_blockstorage_delete,

		Schema: map[string]*schema.Schema{
			"image_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tag": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"blockstorage_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"annotations": {
				Type:     schema.TypeMap,
				Required: true,
				Elem: schema.TypeString,
			},
			"labels": {
				Type:     schema.TypeMap,
				Required: true,
				Elem: schema.TypeString,
			},
			"request_bytes": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"limit_bytes": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"source_url": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func interfaceMap2stringMap(input map[string]interface{})  map[string]string {
	output := make(map[string]string)
	for key, value := range input {
		output[key] = value.(string)
	}
	return output
}

func resource_n0stack_blockstorage_create(d *schema.ResourceData, meta interface{}) error {
	conn, err := grpc.Dial("192.168.1.31:20180", grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pprovisioning.NewBlockStorageServiceClient(conn)
	request := pprovisioning.CreateBlockStorageRequest{
		Name: d.Get("blockstorage_name").(string) ,
		Annotations: interfaceMap2stringMap(d.Get("annotations").(map[string]interface{})),
		Labels: interfaceMap2stringMap(d.Get("labels").(map[string]interface{})),
		RequestBytes: uint64(d.Get("request_bytes").(int)) ,
		LimitBytes: uint64(d.Get("limit_bytes").(int)),
		//SourceUrl: d.Get("source_url").(string),
	}
	res, err := client.CreateBlockStorage(context.Background(), &request)
	if err != nil {
		return err
	}
	d.SetId(res.Name)
	return resource_n0stack_blockstorage_read(d, meta)
}

func resource_n0stack_blockstorage_read(d *schema.ResourceData, meta interface{}) error {
	conn, err := grpc.Dial("192.168.1.31:20180", grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pprovisioning.NewBlockStorageServiceClient(conn)
	request := pprovisioning.GetBlockStorageRequest{
		Name: d.Get("blockstorage_name").(string) ,
	}
	res, err := client.GetBlockStorage(context.Background(), &request)
	if err != nil {
		return err
	}
	d.Set("blockstorage_name", res.Name)
	d.Set("annotations", res.Annotations)
	d.Set("labels", res.Labels)
	d.Set("request_bytes", res.RequestBytes)
	d.Set("limit_bytes", res.LimitBytes)
	return nil
}

func resource_n0stack_blockstorage_update(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resource_n0stack_blockstorage_delete(d *schema.ResourceData, meta interface{}) error {
	conn, err := grpc.Dial("192.168.1.31:20180", grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pprovisioning.NewBlockStorageServiceClient(conn)
	request := pprovisioning.DeleteBlockStorageRequest{
		Name: d.Get("blockstorage_name").(string) ,
	}
	_, err = client.DeleteBlockStorage(context.Background(), &request)
	if err != nil {
		return err
	}
	return nil
}
