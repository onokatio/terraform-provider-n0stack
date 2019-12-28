package main

import (
	"context"
	"google.golang.org/grpc"
        "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	pprovisioning "github.com/n0stack/n0stack/n0proto.go/provisioning/v0"
)

func resource_n0stack_virtualmachine() *schema.Resource {
	return &schema.Resource{
		Create: resource_n0stack_virtualmachine_create,
		Read: resource_n0stack_virtualmachine_read,
		Update: resource_n0stack_virtualmachine_update,
		Delete: resource_n0stack_virtualmachine_delete,

		Schema: map[string]*schema.Schema{
			"name": {
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
			"request_cpu_milli_core": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"limit_cpu_milli_core": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"request_memory_bytes": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"limit_memory_bytes": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"block_storage_names": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ssh_authorized_keys": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"nics": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"network_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"ipv4_address": {
							Type:     schema.TypeString,
							Required: true,
						},
						"ipv6_address": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resource_n0stack_virtualmachine_create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(Config)
	conn, err := grpc.Dial(config.endpoint, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pprovisioning.NewVirtualMachineServiceClient(conn)

	nics := make([]*(pprovisioning.VirtualMachineNIC),0)
	for _, value := range (d.Get("nics").([]interface{})) {
		nic := pprovisioning.VirtualMachineNIC{}
		element := value.(map[string]interface{})
		nic.NetworkName = element["network_name"].(string)
		nic.Ipv4Address = element["ipv4_address"].(string)
		nic.Ipv6Address = element["ipv6_address"].(string)
		nics = append(nics, &nic)
	}

	request := pprovisioning.CreateVirtualMachineRequest{
		Name: d.Get("name").(string),
		Annotations: interfaceMap2stringMap(d.Get("annotations").(map[string]interface{})),
		Labels: interfaceMap2stringMap(d.Get("labels").(map[string]interface{})),
		RequestCpuMilliCore: uint32(d.Get("request_cpu_milli_core").(int)),
		LimitCpuMilliCore: uint32(d.Get("limit_cpu_milli_core").(int)),
		RequestMemoryBytes: uint64(d.Get("request_memory_bytes").(int)),
		LimitMemoryBytes: uint64(d.Get("limit_memory_bytes").(int)),
		BlockStorageNames: interfaceList2stringList(d.Get("block_storage_names").([]interface{})),
		SshAuthorizedKeys: interfaceList2stringList(d.Get("ssh_authorized_keys").([]interface{})),
		Nics: nics,
	}
	_, err = client.CreateVirtualMachine(context.Background(), &request)
	if err != nil {
		return err
	}

	d.SetId(d.Get("name").(string))

	return resource_n0stack_virtualmachine_read(d, meta)
}

func resource_n0stack_virtualmachine_read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(Config)
	conn, err := grpc.Dial(config.endpoint, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pprovisioning.NewVirtualMachineServiceClient(conn)

	request := pprovisioning.GetVirtualMachineRequest{
		Name: d.Get("name").(string) ,
	}
	res, err := client.GetVirtualMachine(context.Background(), &request)
	if err != nil {
		return err
	}

	d.Set("name",res.Name)
	return nil
}

func resource_n0stack_virtualmachine_update(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resource_n0stack_virtualmachine_delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(Config)
	conn, err := grpc.Dial(config.endpoint, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pprovisioning.NewVirtualMachineServiceClient(conn)

	request := pprovisioning.DeleteVirtualMachineRequest{
		Name: d.Get("name").(string) ,
	}
	_, err = client.DeleteVirtualMachine(context.Background(), &request)
	if err != nil {
		return err
	}

	return resource_n0stack_virtualmachine_read(d, meta)
}
