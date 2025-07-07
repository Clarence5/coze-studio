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

package repository

import (
	"context"

	"gorm.io/gorm"

	"code.byted.org/flow/opencoze/backend/api/model/crossdomain/database"
	"code.byted.org/flow/opencoze/backend/api/model/table"
	"code.byted.org/flow/opencoze/backend/domain/memory/database/entity"
	"code.byted.org/flow/opencoze/backend/domain/memory/database/internal/dal"
	"code.byted.org/flow/opencoze/backend/domain/memory/database/internal/dal/query"
	"code.byted.org/flow/opencoze/backend/infra/contract/idgen"
)

func NewAgentToDatabaseDAO(db *gorm.DB, idGen idgen.IDGenerator) AgentToDatabaseDAO {
	return dal.NewAgentToDatabaseDAO(db, idGen)
}

type AgentToDatabaseDAO interface {
	BatchCreate(ctx context.Context, relations []*database.AgentToDatabase) ([]int64, error)
	BatchDelete(ctx context.Context, basicRelations []*database.AgentToDatabaseBasic) error
	ListByAgentID(ctx context.Context, agentID int64, tableType table.TableType) ([]*database.AgentToDatabase, error)
}

func NewDraftDatabaseDAO(db *gorm.DB, idGen idgen.IDGenerator) DraftDAO {
	return dal.NewDraftDatabaseDAO(db, idGen)
}

type DraftDAO interface {
	Get(ctx context.Context, id int64) (*entity.Database, error)
	List(ctx context.Context, filter *entity.DatabaseFilter, page *entity.Pagination, orderBy []*database.OrderBy) ([]*entity.Database, int64, error)
	MGet(ctx context.Context, ids []int64) ([]*entity.Database, error)

	CreateWithTX(ctx context.Context, tx *query.QueryTx, database *entity.Database, draftID, onlineID int64, physicalTableName string) (*entity.Database, error)
	UpdateWithTX(ctx context.Context, tx *query.QueryTx, database *entity.Database) (*entity.Database, error)
	DeleteWithTX(ctx context.Context, tx *query.QueryTx, id int64) error
	BatchDeleteWithTX(ctx context.Context, tx *query.QueryTx, ids []int64) error
}

func NewOnlineDatabaseDAO(db *gorm.DB, idGen idgen.IDGenerator) OnlineDAO {
	return dal.NewOnlineDatabaseDAO(db, idGen)
}

type OnlineDAO interface {
	Get(ctx context.Context, id int64) (*entity.Database, error)
	MGet(ctx context.Context, ids []int64) ([]*entity.Database, error)
	List(ctx context.Context, filter *entity.DatabaseFilter, page *entity.Pagination, orderBy []*database.OrderBy) ([]*entity.Database, int64, error)

	UpdateWithTX(ctx context.Context, tx *query.QueryTx, database *entity.Database) (*entity.Database, error)
	CreateWithTX(ctx context.Context, tx *query.QueryTx, database *entity.Database, draftID, onlineID int64, physicalTableName string) (*entity.Database, error)
	DeleteWithTX(ctx context.Context, tx *query.QueryTx, id int64) error
	BatchDeleteWithTX(ctx context.Context, tx *query.QueryTx, ids []int64) error
}
