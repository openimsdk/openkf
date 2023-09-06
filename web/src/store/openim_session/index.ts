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

import { OpenIM } from '@/api/openim';
import {
  ConversationItem,
} from "@/utils/open-im-sdk-wasm/types/entity";

import { defineStore } from "pinia";
import store from "../index";

interface StateType {
  conversationList: ConversationItem[];
  currentConversation: ConversationItem;
  unReadCount: number;
}

const useStore = defineStore("session", {
  state: (): StateType => ({
    conversationList: [],
    currentConversation: {} as ConversationItem,
    unReadCount: 0,
  }),
  getters: {
    storeConversationList: (state) => state.conversationList,
    storeCurrentConversation: (state) => state.currentConversation,
    storeUnReadCount: (state) => state.unReadCount,
  },
  actions: {
    async getConversationListFromReq(): Promise<boolean> {
      try {
        const { data } = await OpenIM.getAllConversationList<ConversationItem[]>();
        this.conversationList = data;
        console.log('getConversationListFromReq', data);
        return true;
      } catch (error) {
        return false;
      }
    },
    async getUnReadCountFromReq() {
      const { data } = await OpenIM.getTotalUnreadMsgCount<number>();
      this.unReadCount = data;
    },
    updateUnReadCount(data: number) {
      this.unReadCount = data;
    },
    updateCurrentConversation(item: ConversationItem) {
      this.currentConversation = { ...item };
    },
    updateConversationList(list: ConversationItem[]) {
      this.conversationList = [...list];
    },
    delConversationByCID(conversationID: string) {
      const idx = this.conversationList.findIndex(
        (cve) => cve.conversationID === conversationID
      );
      if (idx !== -1) {
        this.conversationList.splice(idx, 1);
      }
    },
  },
});

export default function useSessionStore() {
  return useStore(store);
}
