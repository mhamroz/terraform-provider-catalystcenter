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

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

//template:begin imports
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

//template:end imports

//template:begin types
type Building struct {
	Id         types.String  `tfsdk:"id"`
	Name       types.String  `tfsdk:"name"`
	ParentName types.String  `tfsdk:"parent_name"`
	Country    types.String  `tfsdk:"country"`
	Address    types.String  `tfsdk:"address"`
	Latitude   types.Float64 `tfsdk:"latitude"`
	Longitude  types.Float64 `tfsdk:"longitude"`
}

//template:end types

//template:begin getPath
func (data Building) getPath() string {
	return "/dna/intent/api/v1/site"
}

//template:end getPath

//template:begin toBody
func (data Building) toBody(ctx context.Context, state Building) string {
	body := ""
	body, _ = sjson.Set(body, "type", "building")
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "site.building.name", data.Name.ValueString())
	}
	if !data.ParentName.IsNull() {
		body, _ = sjson.Set(body, "site.building.parentName", data.ParentName.ValueString())
	}
	if !data.Country.IsNull() {
		body, _ = sjson.Set(body, "site.building.country", data.Country.ValueString())
	}
	if !data.Address.IsNull() {
		body, _ = sjson.Set(body, "site.building.address", data.Address.ValueString())
	}
	if !data.Latitude.IsNull() {
		body, _ = sjson.Set(body, "site.building.latitude", data.Latitude.ValueFloat64())
	}
	if !data.Longitude.IsNull() {
		body, _ = sjson.Set(body, "site.building.longitude", data.Longitude.ValueFloat64())
	}
	return body
}

//template:end toBody

//template:begin fromBody
func (data *Building) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.0.name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("response.0.additionalInfo.0.attributes.country"); value.Exists() {
		data.Country = types.StringValue(value.String())
	} else {
		data.Country = types.StringNull()
	}
	if value := res.Get("response.0.additionalInfo.0.attributes.address"); value.Exists() {
		data.Address = types.StringValue(value.String())
	} else {
		data.Address = types.StringNull()
	}
	if value := res.Get("response.0.additionalInfo.0.attributes.latitude"); value.Exists() {
		data.Latitude = types.Float64Value(value.Float())
	} else {
		data.Latitude = types.Float64Null()
	}
	if value := res.Get("response.0.additionalInfo.0.attributes.longitude"); value.Exists() {
		data.Longitude = types.Float64Value(value.Float())
	} else {
		data.Longitude = types.Float64Null()
	}
}

//template:end fromBody

//template:begin updateFromBody
func (data *Building) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.0.name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("response.0.additionalInfo.0.attributes.country"); value.Exists() && !data.Country.IsNull() {
		data.Country = types.StringValue(value.String())
	} else {
		data.Country = types.StringNull()
	}
	if value := res.Get("response.0.additionalInfo.0.attributes.address"); value.Exists() && !data.Address.IsNull() {
		data.Address = types.StringValue(value.String())
	} else {
		data.Address = types.StringNull()
	}
	if value := res.Get("response.0.additionalInfo.0.attributes.latitude"); value.Exists() && !data.Latitude.IsNull() {
		data.Latitude = types.Float64Value(value.Float())
	} else {
		data.Latitude = types.Float64Null()
	}
	if value := res.Get("response.0.additionalInfo.0.attributes.longitude"); value.Exists() && !data.Longitude.IsNull() {
		data.Longitude = types.Float64Value(value.Float())
	} else {
		data.Longitude = types.Float64Null()
	}
}

//template:end updateFromBody
