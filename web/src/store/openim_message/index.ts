import { OpenIM } from '@/api/openim';
import { MessageItem } from "open-im-sdk-wasm/lib/types/entity";
import { GetAdvancedHistoryMsgParams } from "open-im-sdk-wasm/lib/types/params";
import { defineStore } from "pinia";
import store from "../index";

interface StateType {
  historyMessageList: MessageItem[];
  hasMore: boolean;
}

interface IAdvancedMessageResponse {
    lastMinSeq: number;
    isEnd: boolean;
    messageList: MessageItem[];
}

type GetHistoryMessageListFromReqResp = {
  messageIDList: string[];
  lastMinSeq: number;
};

const useStore = defineStore("message", {
  state: (): StateType => ({
    historyMessageList: [],
    hasMore: true,
  }),
  getters: {
    storeHistoryMessageList: (state) => state.historyMessageList,
    storeHistoryMessageHasMore: (state) => state.hasMore,
  },
  actions: {
    async getHistoryMessageListFromReq(
      params: GetAdvancedHistoryMsgParams
    ): Promise<GetHistoryMessageListFromReqResp> {
      const isFirstPage =
        params.startClientMsgID === "" || params.lastMinSeq === 0;
      try {
        const { data } = await OpenIM.getAdvancedHistoryMessageList<IAdvancedMessageResponse>(params);
        this.historyMessageList = [
          ...data.messageList,
          ...(isFirstPage ? [] : this.historyMessageList),
        ];
        this.hasMore = !data.isEnd && data.messageList.length === 20;
        return {
          messageIDList: data.messageList.map(
            (message: MessageItem) => message.clientMsgID
          ),
          lastMinSeq: data.lastMinSeq,
        };
      } catch (error) {
        console.log("Get history message failed", error);
        this.hasMore = false;
        return {
          messageIDList: [],
          lastMinSeq: 0,
        };
      }
    },
    pushNewMessage(message: MessageItem) {
      this.historyMessageList.push(message);
    },
    clearHistoryMessage() {
      this.historyMessageList = [];
    },
    resetHistoryMessageList() {
      this.historyMessageList = [];
      this.hasMore = true;
    },
  },
});

export default function useMessageStore() {
  return useStore(store);
}
