provider "n0stack" {
	endpoint = "192.168.1.31:20180"
}

resource "n0stack_blockstorage" "example" {
	blockstorage_name = "test_storage"
	annotations = {
		annotation1 = "hoge"
		"n0core/provisioning/block_storage/request_node_name" = "archlinux"
		"n0core/provisioning/block_storage/url"               = "file:///var/lib/n0core/block_storage/test_storage"
	}
	labels = {
		label1 = "hoge"
	}
	request_bytes = 100
	limit_bytes = 100
}
