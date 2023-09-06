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

import { CbEvents } from "@/utils/open-im-sdk-wasm/constant";
import { OpenIM } from '@/api/openim';
import { onMounted, onUnmounted } from "vue";
import {
    ConversationItem,
    WSEvent,
    MessageItem,
  } from "@/utils/open-im-sdk-wasm/types/entity";
import useSessionStore from "@/store/openim_session";
import useMessageStore from "@/store/openim_message";
import { conversationSort } from "@/utils/common/im";
import { NotifyPlugin } from "tdesign-vue-next";

export function useGlobalEvent() {
    const sessionStore = useSessionStore();
    const messageStore = useMessageStore();


    const setIMListener = () => {
        console.log("setIMListener");

        // message
        OpenIM.on(CbEvents.OnRecvNewMessage, newMessageHandler);
        OpenIM.on(CbEvents.OnRecvNewMessages, newMessageHandler);

        // session
        OpenIM.on(CbEvents.OnConversationChanged, conversationChangeHandler)
        OpenIM.on(CbEvents.OnNewConversation, newConversationHandler);
        OpenIM.on(
            CbEvents.OnTotalUnreadMessageCountChanged,
            totalUnreadChangeHandler
        );

    };
    const disposeIMListener = () => {
        console.log("disposeIMListener");

        // message
        OpenIM.off(CbEvents.OnRecvNewMessage, newMessageHandler);
        OpenIM.off(CbEvents.OnRecvNewMessages, newMessageHandler);

        // conversation
        OpenIM.off(CbEvents.OnConversationChanged, conversationChangeHandler)
        OpenIM.off(CbEvents.OnNewConversation, newConversationHandler);
        OpenIM.off(
            CbEvents.OnTotalUnreadMessageCountChanged,
            totalUnreadChangeHandler
        );
    };

    // message
    const newMessageHandler = ({ data }: any) => {
        const parsedData = data;

        // message array
        if (Array.isArray(parsedData)) {
            parsedData.map((message) => handleNewMessage(message));
            return;
        }
        handleNewMessage(parsedData);
    };
    const handleNewMessage = (message: MessageItem) => {
        console.log('receve new message', message);
        NotifyPlugin.info({
            title: 'New Messages: ' + message.senderNickname,
            content: message.textElem?.content,
            duration: 2000,
            placement: 'top-right',
            closeBtn: true,
        });
        messageStore.pushNewMessage(message);
    };

    // session
    const conversationChangeHandler = ({ data }: WSEvent<ConversationItem[]>) => {
        let filterArr: ConversationItem[] = [];
        const changes: ConversationItem[] = data;
        const chids = changes.map((ch) => ch.conversationID);
        filterArr = sessionStore.storeConversationList.filter(
          (tc) => !chids.includes(tc.conversationID)
        );
        const idx = changes.findIndex(
          (c) =>
            c.conversationID ===
            sessionStore.storeCurrentConversation.conversationID
        );
        if (idx !== -1) sessionStore.updateCurrentConversation(changes[idx]);
        const result = [...changes, ...filterArr];
        sessionStore.updateConversationList(conversationSort(result));
    };
    const newConversationHandler = ({ data }: WSEvent<ConversationItem[]>) => {
        console.log('receve new conversation', data);
        const news: ConversationItem[] = data;
        const result = [...news, ...sessionStore.storeConversationList];
        sessionStore.updateConversationList(conversationSort(result));
    };
    const totalUnreadChangeHandler = ({ data }: WSEvent<number>) => {
        console.log('receve new unread count', data);
        sessionStore.updateUnReadCount(data);
    };

    onMounted(() => {
        setIMListener();
        sessionStore.getConversationListFromReq();
        sessionStore.getUnReadCountFromReq();
    });
    onUnmounted(() => {
        disposeIMListener();
    });
}