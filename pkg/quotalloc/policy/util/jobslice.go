/*
Copyright 2017 The Kubernetes Authors.

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

package util

import (
	"github.com/kubernetes-incubator/kube-arbitrator/pkg/quotalloc/cache"
)

type CPUJobSlice []*cache.QuotaAllocatorInfo

func (s CPUJobSlice) Len() int {
	return len(s)
}

func (s CPUJobSlice) Less(i, j int) bool {
	cpu1 := s[i].QuotaAllocator().Spec.Request.Resources["cpu"]
	cpu2 := s[j].QuotaAllocator().Spec.Request.Resources["cpu"]
	return cpu1.Cmp(cpu2) < 0
}

func (s CPUJobSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type MEMJobSlice []*cache.QuotaAllocatorInfo

func (s MEMJobSlice) Len() int {
	return len(s)
}

func (s MEMJobSlice) Less(i, j int) bool {
	mem1 := s[i].QuotaAllocator().Spec.Request.Resources["memory"]
	mem2 := s[j].QuotaAllocator().Spec.Request.Resources["memory"]
	return mem1.Cmp(mem2) < 0
}

func (s MEMJobSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type WeightJobSlice []*cache.QuotaAllocatorInfo

func (s WeightJobSlice) Len() int {
	return len(s)
}

func (s WeightJobSlice) Less(i, j int) bool {
	weight1 := s[i].QuotaAllocator().Spec.Weight
	weight2 := s[j].QuotaAllocator().Spec.Weight
	return weight1 > weight2
}

func (s WeightJobSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
