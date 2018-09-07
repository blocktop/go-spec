package spec

import (
	"fmt"
	"strings"
)

// URI is a string that identifies a resource in a blockchain. The canonical form is:
//  <blockchain type>://<resource type>/<component type>/<version>/<id>
// For example:
//  blckit://block/luckyblock/v1/0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef
//  blckit://transaction/exchange/v2/fedcba9876543210fedcba9876543210fedcba9876543210fedcba9876543210
type URI struct {
	uri            string
	blockchainType string
	resourceType   string
	componentType  string
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
	u.SetBlockchainType(p.GetBlockchainType())
	u.SetResourceType(p.GetResourceType())
	u.SetComponentType(p.GetComponentType())
	u.SetVersion(p.GetVersion())
	return u
}

func NewURIParts(blockchainType string, resourceType string, componentType string, version string, id string) *URI {
	u := &URI{}
	u.SetBlockchainType(blockchainType)
	u.SetResourceType(resourceType)
	u.SetComponentType(componentType)
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
	u.blockchainType = parts[0][:len(parts[0])] //omit the colon
	u.resourceType = parts[2]
	u.componentType = parts[3]
	u.version = parts[4]
	u.id = parts[5]
	u.formatValue()
	return true
}

func (u *URI) formatValue() {
	u.uri = fmt.Sprintf("%s://%s/%s/%s/%s", u.blockchainType, u.resourceType, u.componentType, u.version, u.id)
}

func (u *URI) GetBlockchainType() string {
	return u.blockchainType
}
func (u *URI) SetBlockchainType(blockchainType string) {
	u.blockchainType = strings.ToLower(blockchainType)
	u.formatValue()
}
func (u *URI) GetResourceType() string {
	return u.resourceType
}
func (u *URI) SetResourceType(resourceType string) {
	u.resourceType = strings.ToLower(resourceType)
	u.formatValue()
}
func (u *URI) GetComponentType() string {
	return u.componentType
}
func (u *URI) SetComponentType(componentType string) {
	u.componentType = strings.ToLower(componentType)
	u.formatValue()
}
func (u *URI) GetVersion() string {
	return u.version
}
func (u *URI) SetVersion(version string) {
	u.version = strings.ToLower(version)
	u.formatValue()
}
func (u *URI) GetID() string {
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
