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

# 既知の問題

- CRUDのうち、Updateを実装していないため、ラベルを途中で変更してapplyなどはできない。基本的にはリソースの全作成、全削除のみ。  
  つまり、既にリソースが存在する場合Terraformは一切書き込み操作を行わない。

- `terraform destroy`を行う際に、なぜか依存関係をterraformが認識しない。  
  イメージのもととなったブロックストレージをイメージから切り離す前に、イメージそのものを削除してしまったりすることが多々ある。  
  状態がわけわからなくなったら`n0cli`から手動で削除がオススメ

# 使い方

## ざっくりイメージを作る

```hcl
provider "n0stack" {
	endpoint = "192.168.1.31:20180"
}

resource "n0stack_blockstorage" "Ubuntu-Cloudimage-Disk" {
	blockstorage_name = "ubuntu-bionic"
	annotations = {
		"n0core/provisioning/block_storage/request_node_name" = "archlinux"
	}
	request_bytes = 1073741824
	limit_bytes = 10737418240
	source_url = "http://example/bionic-server-cloudimg-amd64.img"
}

resource "n0stack_image" "Ubuntu-Cloudimage-Image" {
	image_name = "ubuntu-image"
	tags = ["bionic"]
	blockstorage_name = n0stack_blockstorage.Ubuntu-Cloudimage-Disk.blockstorage_name
}
```

## (作ったイメージを使って) VMを作る

```hcl
provider "n0stack" {
	endpoint = "192.168.1.31:20180"
}

resource "n0stack_blockstorage" "MyDisk" {
	image_name = "ubuntu-image"
	tag = "latest"
	blockstorage_name = "MyDisk"
	annotations = {
		"n0core/provisioning/block_storage/request_node_name" = "archlinux"
	}
	request_bytes = 1073741824
	limit_bytes = 10737418240
}

resource "n0stack_network" "MyNetwork" {
	name = "MyNetwork"
	ipv4_cidr = "192.168.100.0/24"
	annotations = {
		"n0core/provisioning/block_storage/request_node_name" = "archlinux"
	}
}

resource "n0stack_virtualmachine" "MyVM" {
	name = "test-vm"
	annotations = {
		"n0core/provisioning/virtual_machine/request_node_name" = "archlinux"
	}
	request_cpu_milli_core = 10
	limit_cpu_milli_core = 1000
	request_memory_bytes = 536870912
	limit_memory_bytes = 536870912
	block_storage_names = [n0stack_blockstorage.blockstorage_name]

	nics {
		network_name = n0stack_network.name
		ipv4_address = "192.168.100.1"
	}
}
```

## リソース

ほとんどの場合、キーとバリューの意味はn0cli doコマンドのyamlと同じです。

### プロバイダ設定

```hcl
provider "n0stack" {
	endpoint = "192.168.1.31:20180"
}
```

n0stack apiサーバーの設定をここでする。

###  n0stack_blockstorage

ブロックストレージを定義します。  

- `image_name`(string) : 任意です。指定した場合、作成時にそのイメージを元にブロックストレージを作成します。
- `tag`([]string): `image_name`を指定した場合、必要です。
- `source_url`(string): 指定すると、作成時にそのURLのイメージを元に、新しくブロックストレージを生成します。
- `blockstorage_name`(string): ブロックストレージの名前。
- `annotations`([string]string): アノテーションを追加できます。配置するノードの指定など。
- `label`([string]string): ラベル。
- `limit_bytes`(int): ブロックストレージが使用できる最大容量。単位はバイト。
- `request_bytes`(int): 普段実際に使う可能性のある容量。単位はバイト。

`image_name`と`source_url`を同時に指定することはできません。
どちらも指定しないと、空のブロックストレージを設定します。  

例: 

```hcl
resource "n0stack_blockstorage" "MyDisk" {
	image_name = "Ubuntu18.04"
	blockstorage_name = "MyDisk"
	annotations = {
		"n0core/provisioning/block_storage/request_node_name" = "archlinux"
	}
	request_bytes = 1073741824
	limit_bytes = 10737418240
}
```
### n0stack_network

ネットワークを定義します。

- `name`(string): ネットワークの名前を定義します。
- `ipv4_cidr`(string): IPv4アドレスを定義します。構文は`x.x.x.x/yy`
- `ipv6_cidr`(string): IPv6アドレスを定義します。構文は`x::x/yy`
- `annotations`([string]string): アノテーションを追加できます。vlan idの指定など。
- `label`([string]string): ラベル。

例:

```hcl
resource "n0stack_network" "MyNetwork" {
	name = "MyNetwork"
	ipv4_cidr = "192.168.100.0/24"
	annotations = {
		"n0core/provisioning/virtual_machine/vlan_id" = "100"
	}
}
```

### n0stack_virtualmachine

仮想マシンを定義します。

- `name`(string): 仮想マシンの名前を定義します。
- `annotations`([string]string): アノテーションを追加できます。vlan idの指定など。
- `label`([string]string): ラベル。
- `request_cpu_milli_core`(int): 普段利用する可能性のあるcpuミリ秒。
- `limit_cpu_milli_core`(int): 利用できる最大cpuミリ秒。
- `request_memory_bytes`(int): 普段利用する可能性のあるメモリ容量。単位はバイト。
- `limit_memory_bytes`(int): 利用できる最大メモリ容量。単位はバイト。
- `block_storage_names`([]string): 接続するブロックストレージ名。
- `nics`(object): vmに刺さるnicの定義。
  - `network_name`(string): NICへ接続するネットワークの定義。
  - `ipv4_address`(string): NICへ設定する固定IPv4アドレスの定義。
  - `ipv6_address`(string): NICへ設定する固定IPv6アドレスの定義。

例:

```hcl
resource "n0stack_virtualmachine" "MyVM" {
	name = "MyVM"
	annotations = {
		"n0core/provisioning/virtual_machine/request_node_name" = "archlinux"
	}
	request_cpu_milli_core = 10
	limit_cpu_milli_core = 1000
	request_memory_bytes = 536870912
	limit_memory_bytes = 536870912
	block_storage_names = [n0stack_blockstorage.MyDisk.blockstorage_name]

	nics {
		network_name = n0stack_network.MyNetwork.network_name
		ipv4_address = "192.168.100.1"
	}
}

### n0stack_image

- `image_name`(string): 定義するイメージの名前
- `blockstorage_name`(string): イメージに、tagsのタグとしてブロックストレージをイメージに登録する。
- `tags`([]string): 登録するタグ。複数指定できる。

例:

```hcl
resource "n0stack_image" "MyImage" {
	image_name = "MyImage"
	blockstorage_name = "MyDisk"
	tags = ["latest"]
}
```

※注: 本来であれば、イメージに別のタグとして複数のブロックストレージを追加できますが、現在非対応です。
