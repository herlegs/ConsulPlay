module github.com/herlegs/ConsulPlay

replace github.com/hashicorp/vault/api => github.com/hashicorp/vault/api v0.0.0-20200718022110-340cc2fa263f

require (
	github.com/hashicorp/consul v1.8.5
	github.com/hashicorp/consul/api v1.7.0
	github.com/hashicorp/vault v1.5.5
	github.com/hashicorp/vault/api v1.0.5-0.20201104232315-52f641658ed7
	github.com/hashicorp/vault/sdk v0.1.14-0.20201104232315-52f641658ed7 // indirect
)

go 1.13
