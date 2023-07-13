export interface TokenResponse {
    token: string;
    expire_time_seconds: number;
}

export interface UserLoginResponse {
    uuid: string;
    kf_token: TokenResponse;
    im_token: TokenResponse;
}
