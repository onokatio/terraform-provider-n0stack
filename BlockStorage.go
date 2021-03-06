package main

import (
	"context"
	"google.golang.org/grpc"
        "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	pprovisioning "github.com/n0stack/n0stack/n0proto.go/provisioning/v0"
	pdeployment "github.com/n0stack/n0stack/n0proto.go/deployment/v0"
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
				Optional: true,
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
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resource_n0stack_blockstorage_create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(Config)
	conn, err := grpc.Dial(config.endpoint, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	if(d.Get("source_url").(string) != ""){
		client := pprovisioning.NewBlockStorageServiceClient(conn)
		request := pprovisioning.FetchBlockStorageRequest{
			Name: d.Get("blockstorage_name").(string) ,
			Annotations: interfaceMap2stringMap(d.Get("annotations").(map[string]interface{})),
			Labels: interfaceMap2stringMap(d.Get("labels").(map[string]interface{})),
			RequestBytes: uint64(d.Get("request_bytes").(int)) ,
			LimitBytes: uint64(d.Get("limit_bytes").(int)),
			SourceUrl: d.Get("source_url").(string),
		}
		res, err := client.FetchBlockStorage(context.Background(), &request)
		if err != nil {
			return err
		}
		d.SetId(res.Name)
	} else if(d.Get("image_name").(string) != ""){
		client := pdeployment.NewImageServiceClient(conn)
		request := pdeployment.GenerateBlockStorageRequest{
			ImageName: d.Get("image_name").(string) ,
			BlockStorageName: d.Get("blockstorage_name").(string) ,
			Annotations: interfaceMap2stringMap(d.Get("annotations").(map[string]interface{})),
			RequestBytes: uint64(d.Get("request_bytes").(int)) ,
			LimitBytes: uint64(d.Get("limit_bytes").(int)),
			Tag: d.Get("tag").(string),
		}
		res, err := client.GenerateBlockStorage(context.Background(), &request)
		if err != nil {
			return err
		}
		d.SetId(res.Name)
	} else {
		client := pprovisioning.NewBlockStorageServiceClient(conn)
		request := pprovisioning.CreateBlockStorageRequest{
			Name: d.Get("blockstorage_name").(string) ,
			Annotations: interfaceMap2stringMap(d.Get("annotations").(map[string]interface{})),
			Labels: interfaceMap2stringMap(d.Get("labels").(map[string]interface{})),
			RequestBytes: uint64(d.Get("request_bytes").(int)) ,
			LimitBytes: uint64(d.Get("limit_bytes").(int)),
		}
		res, err := client.CreateBlockStorage(context.Background(), &request)
		if err != nil {
			return err
		}
		d.SetId(res.Name)
	}

	return resource_n0stack_blockstorage_read(d, meta)
}

func resource_n0stack_blockstorage_read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(Config)
	conn, err := grpc.Dial(config.endpoint, grpc.WithInsecure())
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
	return resource_n0stack_blockstorage_read(d, meta)
}

func resource_n0stack_blockstorage_delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(Config)
	conn, err := grpc.Dial(config.endpoint, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pprovisioning.NewBlockStorageServiceClient(conn)
	{
		request := pprovisioning.DeleteBlockStorageRequest{
			Name: d.Get("blockstorage_name").(string) ,
		}
		_, err = client.DeleteBlockStorage(context.Background(), &request)
		if err != nil {
			return err
		}
	}
	{
		request := pprovisioning.PurgeBlockStorageRequest{
			Name: d.Get("blockstorage_name").(string) ,
		}
		_, err = client.PurgeBlockStorage(context.Background(), &request)
		if err != nil {
			return err
		}
	}

	return nil
}
