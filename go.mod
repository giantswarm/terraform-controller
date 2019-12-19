module github.com/rancher/terraform-controller

go 1.12

require (
	github.com/docker/go-units v0.4.0
	github.com/pkg/errors v0.8.1
	github.com/rancher/wrangler v0.2.0
	github.com/rancher/wrangler-api v0.2.0
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.3.0
	github.com/urfave/cli v1.20.0
	k8s.io/api v0.0.0
	k8s.io/apiextensions-apiserver v0.0.0
	k8s.io/apimachinery v0.0.0
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/gengo v0.0.0
)

replace (
	github.com/matryer/moq => github.com/rancher/moq v0.0.0-20190404221404-ee5226d43009
	k8s.io/api => k8s.io/api v0.0.0-20191114100352-16d7abae0d2a
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20191114105449-027877536833
	k8s.io/apimachinery => k8s.io/apimachinery v0.16.5-beta.1
	k8s.io/client-go => k8s.io/client-go v0.0.0-20191114101535-6c5935290e33
	k8s.io/gengo => k8s.io/gengo v0.0.0-20191120174120-e74f70b9b27e
)
