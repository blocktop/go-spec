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

// MessageProtocol is a string that identifies a messaging protocol. The canonical form is:
//  /<blockchain type>/<resource type>/<component type>/<version>
// For example:
//  /blocktop/block/luckyblock/v1
//  /blocktop/transaction/exchange/v2
type MessageProtocol struct {
	protocol       string
	blockchainName string
	resourceType   string
	componentName  string
	version        string
}

const (
	ResourceTypeBlock       string = "block"
	ResourceTypeTransaction string = "transaction"
)

func NewProtocol(value string) *MessageProtocol {
	p := &MessageProtocol{}
	if !p.SetValue(value) {
		return nil
	}
	return p
}

func NewProtocolParts(blockchainName string, resourceType string, componentName string, version string) *MessageProtocol {
	p := &MessageProtocol{}
	p.SetBlockchainName(blockchainName)
	p.SetResourceType(resourceType)
	p.SetComponentName(componentName)
	p.SetVersion(version)

	return p
}

func NewProtocolMarshalled(blockchainName string, m Marshalled) *MessageProtocol {
	p := &MessageProtocol{}
	p.SetBlockchainName(blockchainName)
	p.SetResourceType(m.ResourceType())
	p.SetComponentName(FullyQualifiedName(m))
	p.SetVersion(m.Version())

	return p
}

func (p *MessageProtocol) String() string {
	return p.protocol
}

func (p *MessageProtocol) SetValue(value string) (ok bool) {
	parts := strings.Split(strings.ToLower(value), "/")
	if len(parts) != 5 {
		return false
	}
	p.blockchainName = parts[1]
	p.resourceType = parts[2]
	p.componentName = parts[3]
	p.version = parts[4]

	p.formatValue()
	return true
}

func (p *MessageProtocol) formatValue() {
	p.protocol = fmt.Sprintf("/%s/%s/%s/%s", p.blockchainName, p.resourceType, p.componentName, p.version)
}

func (p *MessageProtocol) BlockchainName() string {
	return p.blockchainName
}
func (p *MessageProtocol) SetBlockchainName(blockchainName string) {
	p.blockchainName = strings.ToLower(blockchainName)
	p.formatValue()
}
func (p *MessageProtocol) ResourceType() string {
	return p.resourceType
}
func (p *MessageProtocol) SetResourceType(resourceType string) {
	p.resourceType = strings.ToLower(resourceType)
	p.formatValue()
}
func (p *MessageProtocol) ComponentName() string {
	return p.componentName
}
func (p *MessageProtocol) SetComponentName(componentName string) {
	p.componentName = strings.ToLower(componentName)
	p.formatValue()
}
func (p *MessageProtocol) Version() string {
	return p.version
}
func (p *MessageProtocol) SetVersion(version string) {
	p.version = strings.ToLower(version)
	p.formatValue()
}

func (p *MessageProtocol) IsValid() bool {
	if p.resourceType == ResourceTypeBlock ||
		p.resourceType == ResourceTypeTransaction {
		return true
	}

	return false
}
