// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"errors"
	"time"

	gojson "github.com/goccy/go-json"
)

// StoreUpdateAction uses action as discriminator attribute
type StoreUpdateAction interface{}

func mapDiscriminatorStoreUpdateAction(input interface{}) (StoreUpdateAction, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["action"].(string)
		if !ok {
			return nil, errors.New("Error processing discriminator field 'action'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}
	switch discriminator {
	case "addDistributionChannel":
		new := StoreAddDistributionChannelAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addSupplyChannel":
		new := StoreAddSupplyChannelAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeDistributionChannel":
		new := StoreRemoveDistributionChannelAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeSupplyChannel":
		new := StoreRemoveSupplyChannelAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomField":
		new := StoreSetCustomFieldAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomType":
		new := StoreSetCustomTypeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDistributionChannels":
		new := StoreSetDistributionChannelsAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLanguages":
		new := StoreSetLanguagesAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setName":
		new := StoreSetNameAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setSupplyChannels":
		new := StoreSetSupplyChannelsAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// Store is of type BaseResource
type Store struct {
	Version              int                `json:"version"`
	SupplyChannels       []ChannelReference `json:"supplyChannels,omitempty"`
	Name                 *LocalizedString   `json:"name,omitempty"`
	LastModifiedBy       *LastModifiedBy    `json:"lastModifiedBy,omitempty"`
	LastModifiedAt       time.Time          `json:"lastModifiedAt"`
	Languages            []string           `json:"languages,omitempty"`
	Key                  string             `json:"key"`
	ID                   string             `json:"id"`
	DistributionChannels []ChannelReference `json:"distributionChannels"`
	Custom               *CustomFields      `json:"custom,omitempty"`
	CreatedBy            *CreatedBy         `json:"createdBy,omitempty"`
	CreatedAt            time.Time          `json:"createdAt"`
}

// StoreAddDistributionChannelAction implements the interface StoreUpdateAction
type StoreAddDistributionChannelAction struct {
	DistributionChannel *ChannelResourceIdentifier `json:"distributionChannel"`
}

// MarshalJSON override to set the discriminator value
func (obj StoreAddDistributionChannelAction) MarshalJSON() ([]byte, error) {
	type Alias StoreAddDistributionChannelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addDistributionChannel", Alias: (*Alias)(&obj)})
}

// StoreAddSupplyChannelAction implements the interface StoreUpdateAction
type StoreAddSupplyChannelAction struct {
	SupplyChannel *ChannelResourceIdentifier `json:"supplyChannel"`
}

// MarshalJSON override to set the discriminator value
func (obj StoreAddSupplyChannelAction) MarshalJSON() ([]byte, error) {
	type Alias StoreAddSupplyChannelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addSupplyChannel", Alias: (*Alias)(&obj)})
}

// StoreDraft is a standalone struct
type StoreDraft struct {
	SupplyChannels       []ChannelResourceIdentifier `json:"supplyChannels,omitempty"`
	Name                 *LocalizedString            `json:"name"`
	Languages            []string                    `json:"languages,omitempty"`
	Key                  string                      `json:"key"`
	DistributionChannels []ChannelResourceIdentifier `json:"distributionChannels,omitempty"`
	Custom               *CustomFieldsDraft          `json:"custom,omitempty"`
}

// StoreKeyReference implements the interface KeyReference
type StoreKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value
func (obj StoreKeyReference) MarshalJSON() ([]byte, error) {
	type Alias StoreKeyReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "store", Alias: (*Alias)(&obj)})
}

// StorePagedQueryResponse is a standalone struct
type StorePagedQueryResponse struct {
	Total   int     `json:"total,omitempty"`
	Results []Store `json:"results"`
	Offset  int     `json:"offset"`
	Limit   int     `json:"limit"`
	Count   int     `json:"count"`
}

// StoreReference implements the interface Reference
type StoreReference struct {
	ID  string `json:"id"`
	Obj *Store `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StoreReference) MarshalJSON() ([]byte, error) {
	type Alias StoreReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "store", Alias: (*Alias)(&obj)})
}

// StoreRemoveDistributionChannelAction implements the interface StoreUpdateAction
type StoreRemoveDistributionChannelAction struct {
	DistributionChannel *ChannelResourceIdentifier `json:"distributionChannel"`
}

// MarshalJSON override to set the discriminator value
func (obj StoreRemoveDistributionChannelAction) MarshalJSON() ([]byte, error) {
	type Alias StoreRemoveDistributionChannelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeDistributionChannel", Alias: (*Alias)(&obj)})
}

// StoreRemoveSupplyChannelAction implements the interface StoreUpdateAction
type StoreRemoveSupplyChannelAction struct {
	SupplyChannel *ChannelResourceIdentifier `json:"supplyChannel"`
}

// MarshalJSON override to set the discriminator value
func (obj StoreRemoveSupplyChannelAction) MarshalJSON() ([]byte, error) {
	type Alias StoreRemoveSupplyChannelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeSupplyChannel", Alias: (*Alias)(&obj)})
}

// StoreResourceIdentifier implements the interface ResourceIdentifier
type StoreResourceIdentifier struct {
	Key string `json:"key,omitempty"`
	ID  string `json:"id,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StoreResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias StoreResourceIdentifier
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "store", Alias: (*Alias)(&obj)})
}

// StoreSetCustomFieldAction implements the interface StoreUpdateAction
type StoreSetCustomFieldAction struct {
	Value interface{} `json:"value,omitempty"`
	Name  string      `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj StoreSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias StoreSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

// StoreSetCustomTypeAction implements the interface StoreUpdateAction
type StoreSetCustomTypeAction struct {
	Type   *TypeResourceIdentifier `json:"type,omitempty"`
	Fields interface{}             `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StoreSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias StoreSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

// StoreSetDistributionChannelsAction implements the interface StoreUpdateAction
type StoreSetDistributionChannelsAction struct {
	DistributionChannels []ChannelResourceIdentifier `json:"distributionChannels,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StoreSetDistributionChannelsAction) MarshalJSON() ([]byte, error) {
	type Alias StoreSetDistributionChannelsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDistributionChannels", Alias: (*Alias)(&obj)})
}

// StoreSetLanguagesAction implements the interface StoreUpdateAction
type StoreSetLanguagesAction struct {
	Languages []string `json:"languages,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StoreSetLanguagesAction) MarshalJSON() ([]byte, error) {
	type Alias StoreSetLanguagesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLanguages", Alias: (*Alias)(&obj)})
}

// StoreSetNameAction implements the interface StoreUpdateAction
type StoreSetNameAction struct {
	Name *LocalizedString `json:"name,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StoreSetNameAction) MarshalJSON() ([]byte, error) {
	type Alias StoreSetNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setName", Alias: (*Alias)(&obj)})
}

// StoreSetSupplyChannelsAction implements the interface StoreUpdateAction
type StoreSetSupplyChannelsAction struct {
	SupplyChannels []ChannelResourceIdentifier `json:"supplyChannels,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StoreSetSupplyChannelsAction) MarshalJSON() ([]byte, error) {
	type Alias StoreSetSupplyChannelsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setSupplyChannels", Alias: (*Alias)(&obj)})
}

// StoreUpdate is a standalone struct
type StoreUpdate struct {
	Version int                 `json:"version"`
	Actions []StoreUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *StoreUpdate) UnmarshalJSON(data []byte) error {
	type Alias StoreUpdate
	if err := gojson.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorStoreUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}
