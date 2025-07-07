/*
 * Copyright 2025 coze-dev Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
 
/* eslint-disable */
/* tslint:disable */
// @ts-nocheck

import * as table from './table';
import * as base from './base';
import * as project_memory from './project_memory';

export type Int64 = string | number;

export interface DelProfileMemoryRequest {
  user_id?: Int64;
  bot_id?: string;
  keywords?: Array<string>;
  connector_id?: Int64;
  /** 引用信息 */
  ref_info?: table.RefInfo;
  project_id?: string;
  Base?: base.Base;
}

export interface DelProfileMemoryResponse {
  BaseResp: base.BaseResp;
}

export interface GetProfileMemoryRequest {
  user_id?: Int64;
  bot_id?: string;
  keywords?: Array<string>;
  connector_id?: Int64;
  version?: string;
  /** 引用信息 */
  ref_info?: table.RefInfo;
  ext?: string;
  project_id?: string;
  ProjectVersion?: Int64;
  VariableChannel?: project_memory.VariableChannel;
  Base?: base.Base;
}

export interface GetProfileMemoryResponse {
  memories?: Array<KVItem>;
  BaseResp: base.BaseResp;
}

export interface GetSysVariableConfRequest {
  Base?: base.Base;
}

export interface GetSysVariableConfResponse {
  conf?: Array<VariableInfo>;
  group_conf?: Array<GroupVariableInfo>;
  BaseResp: base.BaseResp;
}

export interface GroupVariableInfo {
  group_name?: string;
  group_desc?: string;
  group_ext_desc?: string;
  var_info_list?: Array<VariableInfo>;
  sub_group_info?: Array<GroupVariableInfo>;
}

export interface KVItem {
  keyword?: string;
  value?: string;
  create_time?: Int64;
  update_time?: Int64;
  is_system?: boolean;
  prompt_disabled?: boolean;
  schema?: string;
}

export interface SetKvMemoryReq {
  bot_id: string;
  user_id?: Int64;
  data: Array<KVItem>;
  connector_id?: Int64;
  /** 引用信息 */
  ref_info?: table.RefInfo;
  project_id?: string;
  ProjectVersion?: Int64;
  Base?: base.Base;
}

export interface SetKvMemoryResp {
  BaseResp?: base.BaseResp;
}

export interface VariableInfo {
  key?: string;
  default_value?: string;
  description?: string;
  sensitive?: string;
  must_not_use_in_prompt?: string;
  can_write?: string;
  example?: string;
  ext_desc?: string;
  group_name?: string;
  group_desc?: string;
  group_ext_desc?: string;
  EffectiveChannelList?: Array<string>;
}
/* eslint-enable */
