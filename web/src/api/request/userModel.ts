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
