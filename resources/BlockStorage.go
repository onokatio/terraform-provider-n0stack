package main

import (
	"google.golang.org/grpc"
        "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	pprovisioning "github.com/onokatio/terraform-provider-n0stack/n0proto.go/provisioning/v0"
)

func resource_n0stack_blockstorage() *schema.Resource {
	return &schema.Resource{
		Create: resource_n0stack_blockstorage_create,
		Fetch: resource_n0stack_blockstorage_fetch,
		Read: resource_n0stack_blockstorage_read,
		Update: resource_n0stack_blockstorage_update,
		Delete: resource_n0stack_blockstorage_delete,

		Schema: map[string]*schema.Schema{
			"image_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tag": {
				Type:     schema.TypeString,
				Required: true,
			},
			"block_storage_name": {
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

func resource_n0stack_blockstorage_create(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resource_n0stack_blockstorage_fetch(d *schema.ResourceData, meta interface{}) error {
	conn, err := grpc.Dial("localhost:20180", grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pprovisioning.NewBlockStorageServiceClient(conn)
	request := pprovisioning.FetchBlockStorageRequest{
		Name: "hoge",
		Annotations: "hoge",
		Labels: "hoge",
		RequestBytes: 100,
		LimitBytes: 100,
		SourceUrl: "https://example.com",
	}
	res, err := client.FetchBlockStorage(context.Background(), &client)
	if err != nil {
		PrintGrpcError(err)
		return nil
	}
	//return resource_n0stack_blockstorage_read(d, meta)
	return nil
}

func resource_n0stack_blockstorage_read(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resource_n0stack_blockstorage_update(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resource_n0stack_blockstorage_delete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
