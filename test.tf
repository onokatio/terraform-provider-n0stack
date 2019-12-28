provider "n0stack" {
	endpoint = "172.16.14.10:20180"
}

resource "n0stack_blockstorage" "MyDiskFromImage" {
	image_name = "baseimage-ubuntu"
	tag = "latest"
	blockstorage_name = "MyDisk"
	annotations = {
		"n0core/provisioning/block_storage/request_node_name" = "n0node00"
	}
	request_bytes = 1073741824
	limit_bytes = 10737418240
}

resource "n0stack_network" "MyNetwork" {
	name = "MyNetwork"
	ipv4_cidr = "192.168.100.0/24"
	annotations = {
		"n0core/provisioning/virtual_machine/vlan_id" = "126"
	}
}

resource "n0stack_virtualmachine" "CreateVM" {
	name = "test-vm"
	annotations = {
		"n0core/provisioning/virtual_machine/request_node_name" = "n0node00"
	}
	request_cpu_milli_core = 10
	limit_cpu_milli_core = 1000
	request_memory_bytes = 536870912
	limit_memory_bytes = 536870912
	block_storage_names = [n0stack_blockstorage.MyDiskFromImage.blockstorage_name]

	nics {
		network_name = n0stack_network.MyNetwork.name
		ipv4_address = "192.168.100.1"
	}
}
