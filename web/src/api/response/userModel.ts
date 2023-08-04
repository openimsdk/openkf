export interface TokenResponse {
    token: string;
    expire_time_seconds: number;
}

export interface UserLoginResponse {
    uuid: string;
    kf_token: TokenResponse;
    im_token: TokenResponse;
}

export interface GetUserInfoResponse {
    uuid: string;
    email: string;
    nickname: string;
    avatar: string;
    description: string;
    is_enable: boolean;
    is_admin: boolean;
    created_at: string;
}

export interface GetUserInfoListResponse {
    page: number;
    page_size: number;
    total: number;
    list: GetUserInfoResponse[];
}