package main

import (
	"context"
	"google.golang.org/grpc"
        "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	pdeployment "github.com/onokatio/terraform-provider-n0stack/n0proto.go/deployment/v0"
)

func resource_n0stack_image() *schema.Resource {
	return &schema.Resource{
		Create: resource_n0stack_image_create,
		Read: resource_n0stack_image_read,
		Update: resource_n0stack_image_update,
		Delete: resource_n0stack_image_delete,

		Schema: map[string]*schema.Schema{
			"image_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"blockstorage_name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resource_n0stack_image_create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(Config)
	conn, err := grpc.Dial(config.endpoint, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pdeployment.NewImageServiceClient(conn)

	{
		request := pdeployment.ApplyImageRequest{
			Name: d.Get("image_name").(string) ,
		}
		_, err = client.ApplyImage(context.Background(), &request)
		if err != nil {
			return err
		}
	}

	{
		request := pdeployment.RegisterBlockStorageRequest{
			ImageName: d.Get("image_name").(string) ,
			BlockStorageName: d.Get("blockstorage_name").(string),
			Tags: interfaceList2stringList(d.Get("tags").([]interface{})),
		}
		_, err = client.RegisterBlockStorage(context.Background(), &request)
		if err != nil {
			return err
		}
	}

	d.SetId(d.Get("image_name").(string))

	return resource_n0stack_image_read(d, meta)
}

func resource_n0stack_image_read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(Config)
	conn, err := grpc.Dial(config.endpoint, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pdeployment.NewImageServiceClient(conn)

	request := pdeployment.GetImageRequest{
		Name: d.Get("image_name").(string) ,
	}
	res, err := client.GetImage(context.Background(), &request)
	if err != nil {
		return err
	}

	d.Set("image_name",res.Name)
	d.Set("tags",res.Tags)
	//d.Set("blockstorage_name",res.RegisteredBlockStorages)
	return nil
}

func resource_n0stack_image_update(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resource_n0stack_image_delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(Config)
	conn, err := grpc.Dial(config.endpoint, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pdeployment.NewImageServiceClient(conn)

	request := pdeployment.DeleteImageRequest{
		Name: d.Get("image_name").(string) ,
	}
	_, err = client.DeleteImage(context.Background(), &request)
	if err != nil {
		return err
	}

	return resource_n0stack_image_read(d, meta)
}
