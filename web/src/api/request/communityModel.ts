export interface CommunityInfo {
    avatar: string;
    name: string;
    email: string;
}

export type CreateCommunityParam = CommunityInfo;

export interface GetCommunityInfoParam {
    uuid: string;
}
