// Copyright © 2018 J. Strobus White.
// This file is part of the blocktop blockchain development kit.
//
// Blocktop is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Blocktop is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with blocktop. If not, see <http://www.gnu.org/licenses/>.

package spec

import (
	"fmt"
	"strings"
)

// URI is a string that identifies a resource in a blockchain. The canonical form is:
//  <blockchain type>://<resource type>/<component type>/<version>/<id>
// For example:
//  blocktop://block/luckyblock/v1/0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef
//  blocktop://transaction/exchange/v2/fedcba9876543210fedcba9876543210fedcba9876543210fedcba9876543210
type URI struct {
	uri            string
	blockchainName string
	resourceType   string
	componentName  string
	version        string
	id             string
}

func NewURI(value string) *URI {
	u := &URI{}
	if !u.SetValue(value) {
		return nil
	}
	return u
}

func NewURIFromProtocol(p *MessageProtocol, id string) *URI {
	u := &URI{}
	u.SetID(id)
	u.SetBlockchainName(p.BlockchainName())
	u.SetResourceType(p.ResourceType())
	u.SetComponentName(p.ComponentName())
	u.SetVersion(p.Version())
	return u
}

func NewURIParts(blockchainName string, resourceType string, componentName string, version string, id string) *URI {
	u := &URI{}
	u.SetBlockchainName(blockchainName)
	u.SetResourceType(resourceType)
	u.SetComponentName(componentName)
	u.SetVersion(version)
	u.SetID(id)

	return u
}

func (u *URI) String() string {
	return u.uri
}

func (u *URI) SetValue(value string) (ok bool) {
	parts := strings.Split(strings.ToLower(value), "/")
	if len(parts) != 6 {
		return false
	}
	u.blockchainName = parts[0][:len(parts[0])] //omit the colon
	u.resourceType = parts[2]
	u.componentName = parts[3]
	u.version = parts[4]
	u.id = parts[5]
	u.formatValue()
	return true
}

func (u *URI) formatValue() {
	u.uri = fmt.Sprintf("%s://%s/%s/%s/%s", u.blockchainName, u.resourceType, u.componentName, u.version, u.id)
}

func (u *URI) BlockchainName() string {
	return u.blockchainName
}
func (u *URI) SetBlockchainName(blockchainName string) {
	u.blockchainName = strings.ToLower(blockchainName)
	u.formatValue()
}
func (u *URI) ResourceType() string {
	return u.resourceType
}
func (u *URI) SetResourceType(resourceType string) {
	u.resourceType = strings.ToLower(resourceType)
	u.formatValue()
}
func (u *URI) ComponentName() string {
	return u.componentName
}
func (u *URI) SetComponentName(componentName string) {
	u.componentName = strings.ToLower(componentName)
	u.formatValue()
}
func (u *URI) Version() string {
	return u.version
}
func (u *URI) SetVersion(version string) {
	u.version = strings.ToLower(version)
	u.formatValue()
}
func (u *URI) ID() string {
	return u.id
}
func (u *URI) SetID(id string) {
	u.id = strings.ToLower(id)
	u.formatValue()
}

func (u *URI) IsValid() bool {
	if u.resourceType == ResourceTypeBlock ||
		u.resourceType == ResourceTypeTransaction {
		return true
	}

	return false
}
