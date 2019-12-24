# terraform-provider-n0stack

## これはなに

n0stackのリソースをTerraformから扱えるもの

## めざすもの

トラコンの本戦問題を作る際に、手軽にVMを作成・削除できるようにする。  
あくまで、作成や削除のためにyamlを書かなければいけなかった、
操作的な志向のn0cliの機能限定版として、
宣言的な志向でリソースを作成、破壊することに特化する。

## やらないこと

- n0stackの機能の網羅
- VMの起動やシャットダウンなど、作成削除にかかわらない操作

# インストール方法

1. [Github Release](https://github.com/onokatio/terraform-provider-n0stack/releases/)から実行ファイルをダウンロードし、`terraform-provider-n0stack`という名前へ変更する
2. `terraform-provider-n0stack`がカレントディレクトリにある状態で、`terraform init`を行う

今後そのディレクトリ内の`.tf`ファイルからn0stackのproviderが使えるようになる。既存のディレクトリへ実行ファイルをコピーして`terraform init`でも可能。

# 使い方

## ざっくりイメージを作る

```
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
```

## (作ったイメージを使って) VMを作る

```
provider "n0stack" {
	endpoint = "192.168.1.31:20180"
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
		"n0core/provisioning/virtual_machine/request_node_name" = "archlinux"
	}
	request_cpu_milli_core = 10
	limit_cpu_milli_core = 1000
	request_memory_bytes = 536870912
	limit_memory_bytes = 536870912
	block_storage_names = ["test-disk"]

	nics {
		network_name = "test-network"
		ipv4_address = "192.168.100.1"
	}

	depends_on = [
		n0stack_blockstorage.CreateNewDiskFromImage,
		n0stack_network.CreateNetwork,
	]
}
```

## より詳しい使い方

To written.

# 既知の問題

- CRUDのうち、Updateを実装していないため、ラベルを途中で変更してapplyなどはできない。基本的にはリソースの全作成、全削除のみ。  
  つまり、既にリソースが存在する場合Terraformは一切書き込み操作を行わない。

- `terraform destroy`を行う際に、なぜか依存関係をterraformが認識しない。  
  イメージのもととなったブロックストレージをイメージから切り離す前に、イメージそのものを削除してしまったりすることが多々ある。  
  状態がわけわからなくなったら`n0cli`から手動で削除がオススメ
