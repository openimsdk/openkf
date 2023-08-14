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