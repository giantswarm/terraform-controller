/*
Copyright 2019 Rancher Labs, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by main. DO NOT EDIT.

package terraformcontroller

import (
	clientset "github.com/giantswarm/terraform-controller/pkg/generated/clientset/versioned"
	v1 "github.com/giantswarm/terraform-controller/pkg/generated/controllers/terraformcontroller.cattle.io/v1"
	informers "github.com/giantswarm/terraform-controller/pkg/generated/informers/externalversions/terraformcontroller.cattle.io"
	"github.com/rancher/wrangler/pkg/generic"
)

type Interface interface {
	V1() v1.Interface
}

type group struct {
	controllerManager *generic.ControllerManager
	informers         informers.Interface
	client            clientset.Interface
}

// New returns a new Interface.
func New(controllerManager *generic.ControllerManager, informers informers.Interface,
	client clientset.Interface) Interface {
	return &group{
		controllerManager: controllerManager,
		informers:         informers,
		client:            client,
	}
}

func (g *group) V1() v1.Interface {
	return v1.New(g.controllerManager, g.client.TerraformcontrollerV1(), g.informers.V1())
}
