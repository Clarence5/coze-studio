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

package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"code.byted.org/flow/opencoze/backend/domain/workflow/entity/vo"
	"code.byted.org/flow/opencoze/backend/types/errno"
)

type cancelSignalStoreImpl struct {
	redis *redis.Client
}

const (
	workflowExecutionCancelChannelKey = "workflow:cancel:signal:%d"
	workflowExecutionCancelStatusKey  = "workflow:cancel:status:%d"
)

func (c *cancelSignalStoreImpl) EmitWorkflowCancelSignal(ctx context.Context, wfExeID int64) (err error) {
	defer func() {
		if err != nil {
			err = vo.WrapIfNeeded(errno.ErrRedisError, err)
		}
	}()

	signalChannel := fmt.Sprintf(workflowExecutionCancelChannelKey, wfExeID)
	statusKey := fmt.Sprintf(workflowExecutionCancelStatusKey, wfExeID)
	// Define a reasonable expiration for the status key, e.g., 24 hours
	expiration := 24 * time.Hour

	// set a kv to redis to indicate cancellation status
	err = c.redis.Set(ctx, statusKey, "cancelled", expiration).Err()
	if err != nil {
		return fmt.Errorf("failed to set workflow cancel status for wfExeID %d after publishing signal: %w", wfExeID, err)
	}

	// Publish a signal to Redis
	err = c.redis.Publish(ctx, signalChannel, "").Err()
	if err != nil {
		return fmt.Errorf("failed to publish workflow cancel signal for wfExeID %d: %w", wfExeID, err)
	}

	return nil
}

func (c *cancelSignalStoreImpl) SubscribeWorkflowCancelSignal(ctx context.Context, wfExeID int64) (<-chan *redis.Message, func(), error) {
	// Subscribe to Redis channel specific to this workflow execution
	channelName := fmt.Sprintf(workflowExecutionCancelChannelKey, wfExeID)
	pubSub := c.redis.Subscribe(ctx, channelName)

	// Verify subscription was successful
	_, err := pubSub.Receive(ctx) // Wait for subscription confirmation
	if err != nil {
		_ = pubSub.Close() // Cleanup on error
		return nil, nil, vo.WrapError(errno.ErrRedisError, fmt.Errorf("failed to subscribe to cancel signal: %w", err))
	}

	closeFn := func() {
		_ = pubSub.Close()
	}

	return pubSub.Channel(redis.WithChannelSize(1)), closeFn, nil
}

func (c *cancelSignalStoreImpl) GetWorkflowCancelFlag(ctx context.Context, wfExeID int64) (bool, error) {
	// Construct Redis key for workflow cancellation status
	key := fmt.Sprintf(workflowExecutionCancelStatusKey, wfExeID)

	// Check if the key exists in Redis
	count, err := c.redis.Exists(ctx, key).Result()
	if err != nil {
		return false, vo.WrapError(errno.ErrRedisError, fmt.Errorf("failed to check cancellation status in Redis: %w", err))
	}

	// If key exists (count == 1), return true; otherwise return false
	return count == 1, nil
}
