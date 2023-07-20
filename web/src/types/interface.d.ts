export interface BaseInfo {
    User: UserInfo;
    Community: CommunityInfo;
}

export interface UserInfo {
    avatar: string;
    name: string;
    email: string;
    description: string;
}

export interface CommunityInfo {
    avatar: string;
    name: string;
    email: string;
    description: string;
}

// export interface MenuRoute {}
