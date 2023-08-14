export interface CommunityInfo {
    avatar: string;
    name: string;
    email: string;
    description: string;
}

export type CreateCommunityParam = CommunityInfo;

export interface GetCommunityInfoParam {
    uuid: string;
}

export interface UpdateCommunityParam {
    name: string;
    email: string;
    description: string;
    avatar: string;
}