#容器生态涉及技术

##容器技术
* ####核心知识
*  what
* why
* how
	1. 架构
	2. 镜像
	3. 容器
	4. 网络
	5. 存储
* ####进阶知识
multi-host
容器网络
监控
数据管理
日志管理
安全性
##容器平台:
* ####容器编排
> docker swarm
kubernetes
mesos+marathon
* ####容器管理平台
> rancher
containerShip
* ####基于容器的PaaS
>deis
flynn
dokku


##容器生态系统
* ####容器核心技术
	* 容器规范
		* 容器规范  `Open Container Initiative（OCI） 的组织，其目是制定开放的容器规范`
		* 容器runtime`runtime 需要跟操作系统 kernel 紧密协作，为容器提供运行环境。容器只有在 runtime 中才能运行（Ixc、runc、rkt =》runc Docker自己开发的符合oci规范，Docker默认runtime）;runtime中运行的容器即为一个容器镜像的实例`
		* 容器管理工具 `管理工具用于用户与runtime交互`
			* Ixd:与Ixc对应
			* docker engine：我们一般所说的Docker指的就是这个
			* rkt cli：rkt管理工具
		* 容器定义工具  `定义：容器内容和属性，这样容器能够被保存 共享  重建`
			* docker image：docker容器的模板，runtime依据docker image创建容器
			* dockerfile：若干命令的文本文件，通过这些命令创建出docker image
			* ACI（App container Image）：用于rkt容器的image格式
		* Registries    `存放image`
			* Docker Registry ：构建私有Registry
			* Docker hub：公众托管，有现成image
			* Quay.io：公众托管，Docker Hub类似服务
		* 容器OS `专门运行容器的操作系统（CoreOS、atomic 和 ubuntu core ）`



* ####容器平台技术 `让容器作为集群在分布式环境运行`
	* 容器编排引擎 `容器管理、调度、集群定义和服务发现等`
		1. docker swarm 是 Docker 开发的容器编排引擎
		2. kubernetes 是 Google 领导开发的开源容器编排引擎，同时支持 Docker 和 CoreOS 容器
		3. mesos 是一个通用的集群资源调度平台，mesos 与 marathon 一起提供容器编排引擎功能
	* 容器管理平台 `抽象了编排引擎的底层实现细节,架构在容器编排引擎之上`
	* 基于容器的pass `提供了开发、部署和管理应用的平台，使用户不必关心底层基础设施`

* ####容器支持技术
	* 容器网络 `容器与容器，容器与其他实体之间的连通性和隔离性`
		1. docker network原生解方案
		2. flannel、weave、calico
	* 服务发现 `容器会在host之间迁移，动态环境下，让 client 能够知道如何访问容器提供的服务`
		1. etcd、consul、zookeeper
	* 监控
		1. docker ps/top/stats 是 Docker 原生的命令行监控工具，也提供了stats API，通过HTTP请求获取状态信息
		2. sysdig、cAdvisor/Heapster 和 Weave Scope
	* 数据管理 `保证持久化数据也能够动态迁移`
		1. flocker
	* 日志管理 `问题排查和事件管理`
		1. docker logs
		1. logspout
	* 安全性 `对容器镜像进行扫描，发现潜在的漏洞`
		1. openSCAP
* ####Registry
* ####容器 OS



