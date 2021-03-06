/*
Copyright 2019 The Knative Authors

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

package resources

import (
	"crypto/md5"
	"fmt"
)

// CopyMap makes a copy of the map.
func CopyMap(a map[string]string) map[string]string {
	ret := make(map[string]string, len(a))
	for k, v := range a {
		ret[k] = v
	}
	return ret
}

// UnionMaps returns a map constructed from the union of `a` and `b`,
// where value from `b` wins.
func UnionMaps(a, b map[string]string) map[string]string {
	out := make(map[string]string, len(a)+len(b))

	for k, v := range a {
		out[k] = v
	}
	for k, v := range b {
		out[k] = v
	}
	return out
}

// FilterMap creates a copy of the provided map, filtering out the elements
// that match `filter`.
// nil `filter` is accepted.
func FilterMap(in map[string]string, filter func(string) bool) map[string]string {
	ret := make(map[string]string, len(in))
	for k, v := range in {
		if filter != nil && filter(k) {
			continue
		}
		ret[k] = v
	}
	return ret
}

// The longest name supported by the K8s is 63.
// These constants
const (
	longest = 63
	md5Len  = 32
	head    = longest - md5Len
)

// ChildName generates a name for the resource based upong the parent resource and suffix.
// If the concatenated name is longer than K8s permits the name is hashed and truncated to permit
// construction of the resource, but still keeps it unique.
func ChildName(parent, suffix string) string {
	n := parent
	if len(parent) > (longest - len(suffix)) {
		n = fmt.Sprintf("%s%x", parent[:head-len(suffix)], md5.Sum([]byte(parent)))
	}
	return n + suffix
}
