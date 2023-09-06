/**
 * Copyright © 2023 OpenKF & OpenIM open source community. All rights reserved.
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

import {
    ConversationItem,
} from "@/utils/open-im-sdk-wasm/types/entity";

export const conversationSort = (conversationList: ConversationItem[]) => {
    const arr: string[] = [];
    const filterArr = conversationList.filter(
        (c) => !arr.includes(c.conversationID) && arr.push(c.conversationID)
    );
    filterArr.sort((a, b) => {
        if (a.isPinned === b.isPinned) {
            if (!(a.latestMsgSendTime && b.latestMsgSendTime)) {
                return 0;
            }
            const aCompare =
                a.draftTextTime! > a.latestMsgSendTime!
                    ? a.draftTextTime!
                    : a.latestMsgSendTime!;
            const bCompare =
                b.draftTextTime! > b.latestMsgSendTime!
                    ? b.draftTextTime!
                    : b.latestMsgSendTime!;
            if (aCompare > bCompare) {
                return -1;
            } else if (aCompare < bCompare) {
                return 1;
            } else {
                return 0;
            }
        } else if (a.isPinned && !b.isPinned) {
            return -1;
        } else {
            return 1;
        }
    });
    return filterArr;
};