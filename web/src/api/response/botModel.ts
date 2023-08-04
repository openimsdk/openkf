export interface GetBotInfoResponse {
    uuid: string;
    bot_addr: string;
    bot_port: number;
    bot_token: string;
    nickname: string;
    avatar: string;
    description: string;
}

export interface GetBotInfoListResponse {
    page: number;
    page_size: number;
    total: number;
    list: GetBotInfoResponse[];
}