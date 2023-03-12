/*
 * Copyright 2023 CloudWeGo Authors
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

// Code generated by hertz generator.

package message

import (
	"context"
	"offer_tiktok/biz/pack"
	"offer_tiktok/pkg/errno"

	message "offer_tiktok/biz/model/social/message"

	message_service "offer_tiktok/biz/service/message"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// MessageChat .
// @router /douyin/message/chat/ [GET]
func MessageChat(ctx context.Context, c *app.RequestContext) {
	var err error
	var req message.DouyinMessageChatRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := pack.BuildBaseResp(err)
		c.JSON(consts.StatusOK, message.DouyinMessageChatResponse{
			StatusCode:  resp.StatusCode,
			StatusMsg:   resp.StatusMsg,
			MessageList: []*message.Message{},
		})
		return
	}

	messages, err := message_service.NewMessageService(ctx, c).GetMessageChat(&req)
	if err != nil {
		resp := pack.BuildBaseResp(err)
		c.JSON(consts.StatusOK, message.DouyinMessageChatResponse{
			StatusCode:  resp.StatusCode,
			StatusMsg:   resp.StatusMsg,
			MessageList: messages,
		})
		return
	}

	c.JSON(consts.StatusOK, message.DouyinMessageChatResponse{
		StatusCode:  errno.SuccessCode,
		StatusMsg:   errno.SuccessMsg,
		MessageList: messages,
	})
}

// MessageAction .
// @router /douyin/message/action/ [POST]
func MessageAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req message.DouyinMessageActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := pack.BuildBaseResp(err)
		c.JSON(consts.StatusOK, message.DouyinMessageActionResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	err = message_service.NewMessageService(ctx, c).MessageAction(&req)
	if err != nil {
		resp := pack.BuildBaseResp(err)
		c.JSON(consts.StatusOK, message.DouyinMessageActionResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}
	c.JSON(consts.StatusOK, message.DouyinMessageActionResponse{
		StatusCode: errno.SuccessCode,
		StatusMsg:  errno.SuccessMsg,
	})
}
