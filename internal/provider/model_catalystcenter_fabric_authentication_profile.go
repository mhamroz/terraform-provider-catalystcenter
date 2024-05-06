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
type FabricAuthenticationProfile struct {
	Id                         types.String `tfsdk:"id"`
	SiteNameHierarchy          types.String `tfsdk:"site_name_hierarchy"`
	AuthenticationTemplateName types.String `tfsdk:"authentication_template_name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data FabricAuthenticationProfile) getPath() string {
	return "/dna/intent/api/v1/business/sda/authentication-profile"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data FabricAuthenticationProfile) toBody(ctx context.Context, state FabricAuthenticationProfile) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.SiteNameHierarchy.IsNull() {
		body, _ = sjson.Set(body, "0.siteNameHierarchy", data.SiteNameHierarchy.ValueString())
	}
	if !data.AuthenticationTemplateName.IsNull() {
		body, _ = sjson.Set(body, "0.authenticateTemplateName", data.AuthenticationTemplateName.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *FabricAuthenticationProfile) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("siteNameHierarchy"); value.Exists() {
		data.SiteNameHierarchy = types.StringValue(value.String())
	} else {
		data.SiteNameHierarchy = types.StringNull()
	}
	if value := res.Get("authenticateTemplateName"); value.Exists() {
		data.AuthenticationTemplateName = types.StringValue(value.String())
	} else {
		data.AuthenticationTemplateName = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *FabricAuthenticationProfile) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("siteNameHierarchy"); value.Exists() && !data.SiteNameHierarchy.IsNull() {
		data.SiteNameHierarchy = types.StringValue(value.String())
	} else {
		data.SiteNameHierarchy = types.StringNull()
	}
	if value := res.Get("authenticateTemplateName"); value.Exists() && !data.AuthenticationTemplateName.IsNull() {
		data.AuthenticationTemplateName = types.StringValue(value.String())
	} else {
		data.AuthenticationTemplateName = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *FabricAuthenticationProfile) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.AuthenticationTemplateName.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
