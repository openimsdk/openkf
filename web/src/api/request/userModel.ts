import { CommunityInfo } from './communityModel';

export interface UserInfo {
    avatar: string;
    email: string;
    nickname: string;
    password: string;
}

export interface RegisterAdminParam {
    code: string;
    community_info: CommunityInfo;
    user_info: UserInfo;
}

export interface RegisterStaffParam {
    community_id: number;
    user_info: UserInfo;
}

export interface AccountLoginParam {
    email: string;
    password: string;
}

export interface GetUserInfoParam {
    uuid: string;
}

export interface CreateUserParam {
    user_info: UserInfo;
}

export interface UpdateUserEnableStatusParam {
    uuid: string;
    is_enable: boolean;
}

export interface DeleteUserParam {
    uuid: string;
}

export interface UpdateUserInfoParam {
    avatar: string;
    email: string;
    nickname: string;
    description: string;
}

export interface GetUserStatisticsParam {
    type: string;
    start_timestamp: number;
    end_timestamp: number;
}