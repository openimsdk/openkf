import type {
    TokenResponse,
    GetUserInfoResponse,
    UserLoginResponse,
} from '@/api/response/userModel';
import type { GetCommunityInfoResponse } from '@/api/response/communityModel';
import { defineStore } from 'pinia';
import { getMyCommunityInfo } from '@/api/index/community';
import { getMyInfo } from '@/api/index/user';
import store from "../index";

const InitUserInfo: GetUserInfoResponse = {
    uuid: '',
    email: '',
    nickname: '',
    avatar: '',
    description: '',
    is_enable: false,
    is_admin: false,
    created_at: '',
};

const InitCommunityInfo: GetCommunityInfoResponse = {
    uuid: '',
    name: '',
    email: '',
    description: '',
    avatar: '',
};

const InitToken: TokenResponse = {
    token: '',
    expire_time_seconds: 0,
};

const InitKFToken: TokenResponse = {
    token: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX3V1aWQiOiI1NTUyNDhhMWQxNDA5YTdhYmI1ODMwZmRhZDVkIiwiY29tbXVuaXR5X3V1aWQiOiI4NjNiOTExMjdmNGYwODhlZDQ2MjZjMzU3MTYyIiwiaXNzIjoib3BlbmtmIiwiZXhwIjoxNjkxNzI5MzgwLCJuYmYiOjE2OTE2NDI5ODB9.JomPDoVlrRYHWr03yIWqHbPaIQfuXdACEkhw4Ccwm0c',
    expire_time_seconds: 0,
};
const InitIMToken: TokenResponse = {
    token: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiI1NTUyNDhhMWQxNDA5YTdhYmI1ODMwZmRhZDVkIiwiUGxhdGZvcm1JRCI6NSwiZXhwIjoxNjk5NDE4OTgwLCJuYmYiOjE2OTE2NDI2ODAsImlhdCI6MTY5MTY0Mjk4MH0.btYIWiNc21o7sowLddv-6k9qEKQdiWlTFnQraX3eAYE',
    expire_time_seconds: 0,
};

const useStore = defineStore('user', {
    state: () => ({
        kf_token: { ...InitKFToken },
        im_token: { ...InitIMToken },
        userInfo: { ...InitUserInfo },
        communityInfo: { ...InitCommunityInfo },
    }),
    getters: {},
    actions: {
        async StoreToken(info: UserLoginResponse) {
            this.kf_token = info.kf_token;
            this.im_token = info.im_token;
        },
        async StoreInfo() {
            // fetch user info and community info
            const [userInfo, communityInfo] = await Promise.all([
                getMyInfo()
                    .then(res => {
                        // console.log('getMyInfo success', res);
                        return res;
                    })
                    .catch(res => {
                        console.log('getMyInfo error', res);
                        return InitUserInfo;
                    }), // if error, set to default value
                getMyCommunityInfo()
                    .then(res => {
                        // console.log('getMyCommunityInfo success', res);
                        return res;
                    })
                    .catch(res => {
                        console.log('getMyCommunityInfo error', res);
                        return InitCommunityInfo;
                    }),
            ]);
            this.userInfo = userInfo;
            this.communityInfo = communityInfo;
        },
        async logout() {
            this.kf_token = { ...InitToken };
            this.im_token = { ...InitToken };
            this.userInfo = { ...InitUserInfo };
            this.communityInfo = { ...InitCommunityInfo };
        },
    },
});

export default function useUserStore() {
    return useStore(store);
  }
  