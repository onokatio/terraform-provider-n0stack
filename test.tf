provider "n0stack" {
	endpoint = "192.168.1.31:20180"
}

resource "n0stack_blockstorage" "FetchUbuntuDisk" {
	blockstorage_name = "ubuntu-bionic"
	annotations = {
		"n0core/provisioning/block_storage/request_node_name" = "archlinux"
	}
	labels = {
		label1 = "hoge"
	}
	request_bytes = 1073741824
	limit_bytes = 10737418240
	source_url = "http://localhost/bionic-server-cloudimg-amd64.img"
}

resource "n0stack_image" "RegisterUbuntuAsImageLatest" {
	image_name = "ubuntu-bionic-image"
	tags = ["latest"]
	blockstorage_name = "ubuntu-bionic"

	depends_on = [n0stack_blockstorage.FetchUbuntuDisk]
}


resource "n0stack_blockstorage" "CreateNewDiskFromImage" {
	image_name = "ubuntu-bionic-image"
	tag = "latest"
	blockstorage_name = "test-disk"
	annotations = {
		"n0core/provisioning/block_storage/request_node_name" = "archlinux"
	}
	request_bytes = 1073741824
	limit_bytes = 10737418240

	depends_on = [n0stack_image.RegisterUbuntuAsImageLatest]
}

resource "n0stack_network" "CreateNetwork" {
	name = "test-network"
	ipv4_cidr = "192.168.100.0/24"
	annotations = {
		"n0core/provisioning/block_storage/request_node_name" = "archlinux"
	}
}

resource "n0stack_virtualmachine" "CreateVM" {
	name = "test-vm"
	annotations = {
		"n0core/provisioning/virtual_machine/request_node_name" = "vm-host1"
	}
	request_cpu_milli_core = 10
	limit_cpu_milli_core = 1000
	request_memory_bytes = 536870912
	limit_memory_bytes = 536870912
	block_storage_names = ["test-disk"]
	nics = [
		{
			network_name = "test-network"
			ipv4_address = "192.168.100.1"
		},
	]
	depends_on = [
		n0stack_blockstorage.CreateNewDiskFromImage,
		n0stack_network.CreateNetwork,
	]
}
