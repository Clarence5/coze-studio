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

export type Int64 = string | number;

export interface OpenSpace {
  /** 空间 id */
  id?: string;
  /** 空间名称 */
  name?: string;
  /** 空间图标 url */
  icon_url?: string;
  /** 当前用户角色, 枚举值: owner, admin, member */
  role_type?: string;
  /** 工作空间类型, 枚举值: personal, team */
  workspace_type?: string;
}

export interface OpenSpaceData {
  workspaces?: Array<OpenSpace>;
  /** 空间总数 */
  total_count?: Int64;
}

/** *  plagyground 开放api idl文件
 * */
export interface OpenSpaceListRequest {
  page_num?: Int64;
  page_size?: Int64;
}

export interface OpenSpaceListResponse {
  data?: OpenSpaceData;
  code: Int64;
  msg: string;
}
/* eslint-enable */
