/*
Copyright 2021.

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

package v1beta4

import (
	"encoding/json"

	"github.com/emqx/emqx-operator/apis/apps/v2alpha2"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// ConvertTo converts this version to the Hub version (v1).
func (src *Rebalance) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v2alpha2.Rebalance)

	b, err := json.Marshal(src)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, dst); err != nil {
		return err
	}

	dst.Spec.InstanceKind = "EMQX"

	// +kubebuilder:docs-gen:collapse=rote conversion
	return nil
}

// ConvertFrom converts from the Hub version (v1) to this version.
func (dst *Rebalance) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v2alpha2.Rebalance)

	b, err := json.Marshal(src)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, dst); err != nil {
		return err
	}

	// +kubebuilder:docs-gen:collapse=rote conversion
	return nil
}