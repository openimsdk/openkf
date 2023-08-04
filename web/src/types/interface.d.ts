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

export interface MenuRoute {
    path: string;
    name: string;
    icon: string;
}

export interface Platform {
    order: number;
    avatar: string;
    name: string;
    description: string;
    is_enable: boolean;
    tags: string[];
}
