package main

import (
	v1 "github.com/giantswarm/terraform-controller/pkg/apis/terraformcontroller.cattle.io/v1"
	controllergen "github.com/rancher/wrangler/pkg/controller-gen"
	"github.com/rancher/wrangler/pkg/controller-gen/args"
)

func main() {
	controllergen.Run(args.Options{
		OutputPackage: "github.com/giantswarm/terraform-controller/pkg/generated",
		Boilerplate:   "hack/boilerplate.go.txt",
		Groups: map[string]args.Group{
			"terraformcontroller.cattle.io": {
				Types: []interface{}{
					v1.Module{},
					v1.State{},
					v1.Execution{},
				},
				GenerateTypes: true,
			},
		},
	})
}
