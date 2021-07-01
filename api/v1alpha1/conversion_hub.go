// Copyright 2020-2021 Clastix Labs
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	capsulev1beta1 "github.com/clastix/capsule/api/v1beta1"
)

func (t *Tenant) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*capsulev1beta1.Tenant)

	println(dst)

	return nil
}

func (t *Tenant) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*capsulev1beta1.Tenant)

	println(src)

	return nil
}
