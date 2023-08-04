export interface CreateBotParam {
    bot_addr: string;
    bot_port: number;
    bot_token: string;
    nickname: string;
    avatar?: string;
    description?: string;
}

export interface UpdateBotParam {
    uuid: string;
    bot_addr?: string;
    bot_port?: number;
    bot_token?: string;
    nickname?: string;
    avatar?: string;
    description?: string;
}

export interface DeleteBotParam {
    uuid: string;
}