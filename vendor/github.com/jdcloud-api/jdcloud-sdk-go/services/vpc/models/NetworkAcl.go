// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package models


type NetworkAcl struct {

    /* networkAcl ID (Optional) */
    NetworkAclId string `json:"networkAclId"`

    /* networkAcl名称 (Optional) */
    NetworkAclName string `json:"networkAclName"`

    /* 私有网络 ID (Optional) */
    VpcId string `json:"vpcId"`

    /* networkAcl规则列表 (Optional) */
    NetworkAclRules []NetworkAclRule `json:"networkAclRules"`

    /* networkAcl绑定的子网列表 (Optional) */
    SubnetIds []string `json:"subnetIds"`

    /* 描述,允许输入UTF-8编码下的全部字符，不超过256字符 (Optional) */
    Description string `json:"description"`

    /* networkAcl创建时间 (Optional) */
    CreatedTime string `json:"createdTime"`
}
