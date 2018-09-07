package spec

import (
	"fmt"
	"strings"
)

// MessageProtocol is a string that identifies a messaging protocol. The canonical form is:
//  /<blockchain type>/<resource type>/<component type>/<version>
// For example:
//  /blckit/block/luckyblock/v1
//  /blckit/transaction/exchange/v2
type MessageProtocol struct {
	protocol       string
	blockchainType string
	resourceType   string
	componentType  string
	version        string
}

const (
	ResourceTypeBlock string = "block"
	ResourceTypeTransaction string = "transaction"
)

func NewProtocol(value string) *MessageProtocol {
	p := &MessageProtocol{}
	if !p.SetValue(value) {
		return nil
	}
	return p
}

func NewProtocolParts(blockchainType string, resourceType string, componentType string, version string) *MessageProtocol {
	p := &MessageProtocol{}
	p.SetBlockchainType(blockchainType)
	p.SetResourceType(resourceType)
	p.SetComponentType(componentType)
	p.SetVersion(version)

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
	p.blockchainType = parts[1]
	p.resourceType = parts[2]
	p.componentType = parts[3]
	p.version = parts[4]

	p.formatValue()
	return true
}

func (p *MessageProtocol) formatValue() {
	p.protocol = fmt.Sprintf("/%s/%s/%s/%s", p.blockchainType, p.resourceType, p.componentType, p.version)
}

func (p *MessageProtocol) GetBlockchainType() string {
	return p.blockchainType
}
func (p *MessageProtocol) SetBlockchainType(blockchainType string) {
	p.blockchainType = strings.ToLower(blockchainType)
	p.formatValue()
}
func (p *MessageProtocol) GetResourceType() string {
	return p.resourceType
}
func (p *MessageProtocol) SetResourceType(resourceType string) {
	p.resourceType = strings.ToLower(resourceType)
	p.formatValue()
}
func (p *MessageProtocol) GetComponentType() string {
	return p.componentType
}
func (p *MessageProtocol) SetComponentType(componentType string) {
	p.componentType = strings.ToLower(componentType)
	p.formatValue()
}
func (p *MessageProtocol) GetVersion() string {
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
