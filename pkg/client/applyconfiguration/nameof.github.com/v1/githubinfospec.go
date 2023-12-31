/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// GithubInfoSpecApplyConfiguration represents an declarative configuration of the GithubInfoSpec type for use
// with apply.
type GithubInfoSpecApplyConfiguration struct {
	Username  *string `json:"username,omitempty"`
	Link      *string `json:"link,omitempty"`
	RepoCount *int32  `json:"repoCount,omitempty"`
}

// GithubInfoSpecApplyConfiguration constructs an declarative configuration of the GithubInfoSpec type for use with
// apply.
func GithubInfoSpec() *GithubInfoSpecApplyConfiguration {
	return &GithubInfoSpecApplyConfiguration{}
}

// WithUsername sets the Username field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Username field is set to the value of the last call.
func (b *GithubInfoSpecApplyConfiguration) WithUsername(value string) *GithubInfoSpecApplyConfiguration {
	b.Username = &value
	return b
}

// WithLink sets the Link field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Link field is set to the value of the last call.
func (b *GithubInfoSpecApplyConfiguration) WithLink(value string) *GithubInfoSpecApplyConfiguration {
	b.Link = &value
	return b
}

// WithRepoCount sets the RepoCount field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RepoCount field is set to the value of the last call.
func (b *GithubInfoSpecApplyConfiguration) WithRepoCount(value int32) *GithubInfoSpecApplyConfiguration {
	b.RepoCount = &value
	return b
}
