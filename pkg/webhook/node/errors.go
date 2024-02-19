// Copyright 2020-2023 Project Capsule Authors.
// SPDX-License-Identifier: Apache-2.0

package node

import (
	"fmt"
	"strings"

	capsulev1beta2 "github.com/projectcapsule/capsule/pkg/api"
)

func appendForbiddenError(spec *capsulev1beta2.ForbiddenListSpec) (forbidden string) {
	forbidden += "Forbidden are "
	if len(spec.Exact) > 0 {
		forbidden += fmt.Sprintf("one of the following (%s)", strings.Join(spec.Exact, ", "))
		if len(spec.Regex) > 0 {
			forbidden += " or "
		}
	}

	if len(spec.Regex) > 0 {
		forbidden += fmt.Sprintf("matching the regex %s", spec.Regex)
	}

	return
}

type nodeLabelForbiddenError struct {
	spec *capsulev1beta2.ForbiddenListSpec
}

func NewNodeLabelForbiddenError(forbiddenSpec *capsulev1beta2.ForbiddenListSpec) error {
	return &nodeLabelForbiddenError{
		spec: forbiddenSpec,
	}
}

func (f nodeLabelForbiddenError) Error() string {
	return fmt.Sprintf("Unable to update node as some labels are marked as forbidden by system administrator. %s", appendForbiddenError(f.spec))
}

type nodeAnnotationForbiddenError struct {
	spec *capsulev1beta2.ForbiddenListSpec
}

func NewNodeAnnotationForbiddenError(forbiddenSpec *capsulev1beta2.ForbiddenListSpec) error {
	return &nodeAnnotationForbiddenError{
		spec: forbiddenSpec,
	}
}

func (f nodeAnnotationForbiddenError) Error() string {
	return fmt.Sprintf("Unable to update node as some annotations are marked as forbidden by system administrator. %s", appendForbiddenError(f.spec))
}
