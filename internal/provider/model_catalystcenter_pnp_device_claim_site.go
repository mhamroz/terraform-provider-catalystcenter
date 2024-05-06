// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type PnPDeviceClaimSite struct {
	Id               types.String                         `tfsdk:"id"`
	DeviceId         types.String                         `tfsdk:"device_id"`
	SiteId           types.String                         `tfsdk:"site_id"`
	Type             types.String                         `tfsdk:"type"`
	ImageId          types.String                         `tfsdk:"image_id"`
	ImageSkip        types.Bool                           `tfsdk:"image_skip"`
	ConfigId         types.String                         `tfsdk:"config_id"`
	ConfigParameters []PnPDeviceClaimSiteConfigParameters `tfsdk:"config_parameters"`
	RfProfile        types.String                         `tfsdk:"rf_profile"`
	StaticIp         types.String                         `tfsdk:"static_ip"`
	SubnetMask       types.String                         `tfsdk:"subnet_mask"`
	Gateway          types.String                         `tfsdk:"gateway"`
	VlanId           types.String                         `tfsdk:"vlan_id"`
	IpInterfaceName  types.String                         `tfsdk:"ip_interface_name"`
	SensorProfile    types.String                         `tfsdk:"sensor_profile"`
}

type PnPDeviceClaimSiteConfigParameters struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data PnPDeviceClaimSite) getPath() string {
	return "/dna/intent/api/v1/onboarding/pnp-device/site-claim"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data PnPDeviceClaimSite) toBody(ctx context.Context, state PnPDeviceClaimSite) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.DeviceId.IsNull() {
		body, _ = sjson.Set(body, "deviceId", data.DeviceId.ValueString())
	}
	if !data.SiteId.IsNull() {
		body, _ = sjson.Set(body, "siteId", data.SiteId.ValueString())
	}
	if !data.Type.IsNull() {
		body, _ = sjson.Set(body, "type", data.Type.ValueString())
	}
	if !data.ImageId.IsNull() {
		body, _ = sjson.Set(body, "imageInfo.imageId", data.ImageId.ValueString())
	}
	if !data.ImageSkip.IsNull() {
		body, _ = sjson.Set(body, "imageInfo.skip", data.ImageSkip.ValueBool())
	}
	if !data.ConfigId.IsNull() {
		body, _ = sjson.Set(body, "configInfo.configId", data.ConfigId.ValueString())
	}
	if len(data.ConfigParameters) > 0 {
		body, _ = sjson.Set(body, "configInfo.configParameters", []interface{}{})
		for _, item := range data.ConfigParameters {
			itemBody := ""
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "key", item.Name.ValueString())
			}
			if !item.Value.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "value", item.Value.ValueString())
			}
			body, _ = sjson.SetRaw(body, "configInfo.configParameters.-1", itemBody)
		}
	}
	if !data.RfProfile.IsNull() {
		body, _ = sjson.Set(body, "rfProfile", data.RfProfile.ValueString())
	}
	if !data.StaticIp.IsNull() {
		body, _ = sjson.Set(body, "staticIP", data.StaticIp.ValueString())
	}
	if !data.SubnetMask.IsNull() {
		body, _ = sjson.Set(body, "subnetMask", data.SubnetMask.ValueString())
	}
	if !data.Gateway.IsNull() {
		body, _ = sjson.Set(body, "gateway", data.Gateway.ValueString())
	}
	if !data.VlanId.IsNull() {
		body, _ = sjson.Set(body, "vlanId", data.VlanId.ValueString())
	}
	if !data.IpInterfaceName.IsNull() {
		body, _ = sjson.Set(body, "ipInterfaceName", data.IpInterfaceName.ValueString())
	}
	if !data.SensorProfile.IsNull() {
		body, _ = sjson.Set(body, "sensorProfile", data.SensorProfile.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *PnPDeviceClaimSite) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("deviceId"); value.Exists() {
		data.DeviceId = types.StringValue(value.String())
	} else {
		data.DeviceId = types.StringNull()
	}
	if value := res.Get("siteId"); value.Exists() {
		data.SiteId = types.StringValue(value.String())
	} else {
		data.SiteId = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("imageInfo.imageId"); value.Exists() {
		data.ImageId = types.StringValue(value.String())
	} else {
		data.ImageId = types.StringNull()
	}
	if value := res.Get("imageInfo.skip"); value.Exists() {
		data.ImageSkip = types.BoolValue(value.Bool())
	} else {
		data.ImageSkip = types.BoolNull()
	}
	if value := res.Get("configInfo.configId"); value.Exists() {
		data.ConfigId = types.StringValue(value.String())
	} else {
		data.ConfigId = types.StringNull()
	}
	if value := res.Get("configInfo.configParameters"); value.Exists() && len(value.Array()) > 0 {
		data.ConfigParameters = make([]PnPDeviceClaimSiteConfigParameters, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := PnPDeviceClaimSiteConfigParameters{}
			if cValue := v.Get("key"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			} else {
				item.Name = types.StringNull()
			}
			if cValue := v.Get("value"); cValue.Exists() {
				item.Value = types.StringValue(cValue.String())
			} else {
				item.Value = types.StringNull()
			}
			data.ConfigParameters = append(data.ConfigParameters, item)
			return true
		})
	}
	if value := res.Get("rfProfile"); value.Exists() {
		data.RfProfile = types.StringValue(value.String())
	} else {
		data.RfProfile = types.StringNull()
	}
	if value := res.Get("staticIP"); value.Exists() {
		data.StaticIp = types.StringValue(value.String())
	} else {
		data.StaticIp = types.StringNull()
	}
	if value := res.Get("subnetMask"); value.Exists() {
		data.SubnetMask = types.StringValue(value.String())
	} else {
		data.SubnetMask = types.StringNull()
	}
	if value := res.Get("gateway"); value.Exists() {
		data.Gateway = types.StringValue(value.String())
	} else {
		data.Gateway = types.StringNull()
	}
	if value := res.Get("vlanId"); value.Exists() {
		data.VlanId = types.StringValue(value.String())
	} else {
		data.VlanId = types.StringNull()
	}
	if value := res.Get("ipInterfaceName"); value.Exists() {
		data.IpInterfaceName = types.StringValue(value.String())
	} else {
		data.IpInterfaceName = types.StringNull()
	}
	if value := res.Get("sensorProfile"); value.Exists() {
		data.SensorProfile = types.StringValue(value.String())
	} else {
		data.SensorProfile = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *PnPDeviceClaimSite) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("deviceId"); value.Exists() && !data.DeviceId.IsNull() {
		data.DeviceId = types.StringValue(value.String())
	} else {
		data.DeviceId = types.StringNull()
	}
	if value := res.Get("siteId"); value.Exists() && !data.SiteId.IsNull() {
		data.SiteId = types.StringValue(value.String())
	} else {
		data.SiteId = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("imageInfo.imageId"); value.Exists() && !data.ImageId.IsNull() {
		data.ImageId = types.StringValue(value.String())
	} else {
		data.ImageId = types.StringNull()
	}
	if value := res.Get("imageInfo.skip"); value.Exists() && !data.ImageSkip.IsNull() {
		data.ImageSkip = types.BoolValue(value.Bool())
	} else {
		data.ImageSkip = types.BoolNull()
	}
	if value := res.Get("configInfo.configId"); value.Exists() && !data.ConfigId.IsNull() {
		data.ConfigId = types.StringValue(value.String())
	} else {
		data.ConfigId = types.StringNull()
	}
	for i := range data.ConfigParameters {
		keys := [...]string{"key"}
		keyValues := [...]string{data.ConfigParameters[i].Name.ValueString()}

		var r gjson.Result
		res.Get("configInfo.configParameters").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("key"); value.Exists() && !data.ConfigParameters[i].Name.IsNull() {
			data.ConfigParameters[i].Name = types.StringValue(value.String())
		} else {
			data.ConfigParameters[i].Name = types.StringNull()
		}
		if value := r.Get("value"); value.Exists() && !data.ConfigParameters[i].Value.IsNull() {
			data.ConfigParameters[i].Value = types.StringValue(value.String())
		} else {
			data.ConfigParameters[i].Value = types.StringNull()
		}
	}
	if value := res.Get("rfProfile"); value.Exists() && !data.RfProfile.IsNull() {
		data.RfProfile = types.StringValue(value.String())
	} else {
		data.RfProfile = types.StringNull()
	}
	if value := res.Get("staticIP"); value.Exists() && !data.StaticIp.IsNull() {
		data.StaticIp = types.StringValue(value.String())
	} else {
		data.StaticIp = types.StringNull()
	}
	if value := res.Get("subnetMask"); value.Exists() && !data.SubnetMask.IsNull() {
		data.SubnetMask = types.StringValue(value.String())
	} else {
		data.SubnetMask = types.StringNull()
	}
	if value := res.Get("gateway"); value.Exists() && !data.Gateway.IsNull() {
		data.Gateway = types.StringValue(value.String())
	} else {
		data.Gateway = types.StringNull()
	}
	if value := res.Get("vlanId"); value.Exists() && !data.VlanId.IsNull() {
		data.VlanId = types.StringValue(value.String())
	} else {
		data.VlanId = types.StringNull()
	}
	if value := res.Get("ipInterfaceName"); value.Exists() && !data.IpInterfaceName.IsNull() {
		data.IpInterfaceName = types.StringValue(value.String())
	} else {
		data.IpInterfaceName = types.StringNull()
	}
	if value := res.Get("sensorProfile"); value.Exists() && !data.SensorProfile.IsNull() {
		data.SensorProfile = types.StringValue(value.String())
	} else {
		data.SensorProfile = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *PnPDeviceClaimSite) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.SiteId.IsNull() {
		return false
	}
	if !data.Type.IsNull() {
		return false
	}
	if !data.ImageId.IsNull() {
		return false
	}
	if !data.ImageSkip.IsNull() {
		return false
	}
	if !data.ConfigId.IsNull() {
		return false
	}
	if len(data.ConfigParameters) > 0 {
		return false
	}
	if !data.RfProfile.IsNull() {
		return false
	}
	if !data.StaticIp.IsNull() {
		return false
	}
	if !data.SubnetMask.IsNull() {
		return false
	}
	if !data.Gateway.IsNull() {
		return false
	}
	if !data.VlanId.IsNull() {
		return false
	}
	if !data.IpInterfaceName.IsNull() {
		return false
	}
	if !data.SensorProfile.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
