/*
Copyright 2020 Clastix Labs.

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

package v1alpha1

import (
	"sort"
	"strings"
)

type IngressClassList []string

func (n IngressClassList) Len() int {
	return len(n)
}

func (n IngressClassList) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n IngressClassList) Less(i, j int) bool {
	return strings.ToLower(n[i]) < strings.ToLower(n[j])
}

func (n IngressClassList) IsStringInList(value string) (ok bool) {
	sort.Sort(n)
	i := sort.SearchStrings(n, value)
	ok = i < n.Len() && n[i] == value
	return
}
