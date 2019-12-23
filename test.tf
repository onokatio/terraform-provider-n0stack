provider "n0stack" {
	endpoint = "192.168.1.31:20180"
}

resource "n0stack_blockstorage" "FetchUbuntuDisk" {
	blockstorage_name = "ubuntu-bionic"
	annotations = {
		"n0core/provisioning/block_storage/request_node_name" = "archlinux"
		"n0core/provisioning/block_storage/url"               = "file:///var/lib/n0core/block_storage/ubuntu-bionic"
		"n0core/provisioning/block_storage/fetch_from"        = "https://cloud-images.ubuntu.com/bionic/current/bionic-server-cloudimg-amd64.img"
	}
	labels = {
		label1 = "hoge"
	}
	request_bytes = 1073741824
	limit_bytes = 10737418240
	source_url = "https://cloud-images.ubuntu.com/bionic/current/bionic-server-cloudimg-amd64.img"
}

resource "n0stack_image" "RegisterUbuntuAsImageLatest" {
	image_name = "ubuntu-bionic-image"
	tags = ["latest"]
	blockstorage_name = "ubuntu-bionic"
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
}
