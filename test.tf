provider "n0stack" {
}

resource "n0stack_blockstorage" "example" {
	blockstorage_name = "test_storage"
	annotations = {
		annotation1 = "hoge"
		"n0core/provisioning/block_storage/request_node_name" = "archlinux"
	}
	labels = {
		label1 = "hoge"
	}
	request_bytes = 100
	limit_bytes = 100
}
