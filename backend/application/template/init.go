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

package template

import (
	"context"

	"gorm.io/gorm"

	"code.byted.org/flow/opencoze/backend/domain/template/repository"
	"code.byted.org/flow/opencoze/backend/infra/contract/idgen"
	"code.byted.org/flow/opencoze/backend/infra/contract/storage"
)

type ServiceComponents struct {
	DB      *gorm.DB
	IDGen   idgen.IDGenerator
	Storage storage.Storage
}

func InitService(ctx context.Context, components *ServiceComponents) *ApplicationService {

	tRepo := repository.NewTemplateDAO(components.DB, components.IDGen)

	ApplicationSVC.templateRepo = tRepo
	ApplicationSVC.storage = components.Storage

	return ApplicationSVC
}
