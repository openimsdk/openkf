<!--
 Copyright © 2023 OpenKF & OpenIM open source community. All rights reserved.

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
-->

<template>
    <div class="chat">
        <div class="chat-header">
            <div class="header-title">
                <span class="header-title-left">OpenKF</span>
                <div class="header-title-right">
                    Unsatisfied?
                    <swap-icon
                        style="color: white; margin-left: 7px"
                        size="medium"
                    />
                    <div class="header-title-right-prompts">
                        <i>Click here to change a customer service</i>
                    </div>
                </div>
            </div>
            <div class="header-middle">
                <OpenKFLogo class="logo" />
                <div class="middle-right">
                    <div class="name">Judy</div>
                    <div class="meet">Nice to meet you</div>
                </div>
            </div>
            <div class="header-footer">
                <chevron-left-icon style="color: white" size="2em" />
                <view-list-icon style="color: white" size="large" />
            </div>
        </div>
        <div class="chat-message-box">
            <div class="message-box-time">01/01/2023 3.33 p.m</div>
            <div class="kf">
                <Message
                    :content="'adadad'"
                    :position="'left'"
                    :time="'3.3 p.m'"
                    :avatarSrc="LoginBG"
                />
                <template v-for="message in messageList" :key="message">
                    <Message
                        :content="message"
                        :position="'right'"
                        :time="'3.3 p.m'"
                        :avatarSrc="LoginBG"
                    />
                </template>
            </div>
        </div>
        <div class="chat-input-box-wrapper">
            <div class="chat-input-box">
                <div
                    contentEditable="true"
                    ref="inputRef"
                    class="input-div"
                    @input="handleInput"
                    @blur="handleInput"
                ></div>
                <div class="chat-kits">
                    <check-icon
                        size="2em"
                        style="color: #41ac44"
                        @click="send"
                    />
                    <photo-icon size="2em" style="color: #41ac44" />
                    <add-icon size="2em" style="color: #41ac44" />
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import {
    SwapIcon,
    ViewListIcon,
    ChevronLeftIcon,
    CheckIcon,
    PhotoIcon,
    AddIcon,
} from 'tdesign-icons-vue-next';
import { ref, onMounted, reactive } from 'vue';
import OpenKFLogo from '@/assets/openkf-logo.svg';
import LoginBG from '@/assets/opekf-login-bg-light.png'; // change next time
import Message from '@/components/Message.vue';
import { getSDK } from '@/utils/open-im-sdk-wasm';
const IMSDK = getSDK();

const messageList = reactive<string[]>([
    '1111111111111111111111111111111111111111111111111111111111111111111111111111111',
]);
const inputRef = ref('');
const message = ref('');
onMounted(() => {
    inputRef.value.focus();
});
const send = () => {
    // IMSDK.createCardMessage({
    //     userID: '123',
    //     nickname: '456',
    //     faceURL: '123',
    //     ex: '123456',
    // })
    //     .then(({ data }) => {
    //         // 调用成功
    //         console.log('success', data);
    //     })
    //     .catch(({ errCode, errMsg }) => {
    //         // 调用失败
    //         console.log('error', errCode);
    //     });
    messageList.push(message.value);
    console.log('发送了一条消息...');
};
const handleInput = (e: Event) => {
    console.log(e.target.innerText);
    message.value = e.target?.innerText;
};
</script>

<style lang="less" scoped>
.chat {
    width: 502px;
    // height: 100vh;
    background: #ebebeb;
    box-shadow: -15px 12px 12px 1px rgba(0, 0, 0, 0.16);
    border-radius: 39px 39px 39px 39px;
    opacity: 1;
    margin: 0 auto;
    padding-bottom: 1px; // delete  after
    .chat-header {
        width: 502px;
        height: 163px;
        background: #529f54;
        box-shadow: 0px 5px 6px 1px rgba(0, 0, 0, 0.16);
        border-radius: 39px 39px 39px 39px;
        opacity: 1;
        padding: 10px 34px;
        position: relative;
        margin-bottom: 17px;
        .header-title {
            // width: 43px;
            // height: 11px;
            font-size: 10px;
            font-family: Arial-Regular, Arial;
            font-weight: 400;
            color: #ffffff;
            -webkit-background-clip: text;
            display: flex;
            justify-content: space-between;
            align-items: center;
            .header-title-right {
                display: flex;
                align-items: center;
                font-size: 10px;
                font-family: Arial-Italic, Arial;
                font-weight: normal;
                color: #ffffff;
                position: relative;
                .header-title-right-prompts {
                    position: absolute;
                    top: 100%;
                    left: 0;
                    width: 115px;
                    font-size: 10px;
                    font-family: Arial-Italic, Arial;
                    font-weight: normal;
                    color: #ffffff;
                    line-height: 0.9;
                    text-decoration: underline;
                }
            }
        }
        .header-middle {
            position: absolute;
            left: 103px;
            top: 48px;
            display: flex;
            .logo {
                width: 88px;
                height: 81px;
                margin-right: 18px;
            }
            .middle-right {
                display: flex;
                flex-direction: column;
                justify-content: center;
                font-size: 22px;
                font-family: Arial-Regular, Arial;
                font-weight: 400;
                color: #ffffff;
                .meet {
                    font-size: 10px;
                }
            }
        }
        .header-footer {
            margin-top: 85px;
            display: flex;
            justify-content: space-between;
        }
    }
    .chat-message-box {
        height: 615px;
        padding: 0 24px;
        overflow: scroll;
        .message-box-time {
            height: 16px;
            font-size: 14px;
            font-family: Arial-Regular, Arial;
            font-weight: 400;
            color: #858585;
            white-space: nowrap;
            text-align: center;
            margin-bottom: 10px;
        }
    }
    .chat-input-box-wrapper {
        width: 442px;
        min-height: 20px;
        padding: 12px 24px;
        background: #fff;
        margin: 0 auto;
        border-radius: 13px 13px 13px 13px;
        padding-right: 0px;
        margin-bottom: 20px;

        .chat-input-box {
            display: flex;
            align-items: flex-start;
            .input-div {
                width: 292px;
                min-height: 20px;
                flex-grow: 0;
                flex-shrink: 0;
                overflow: auto;
                outline: none;
            }
            .chat-kits {
                flex-grow: 1;
                :deep(.t-icon) {
                    margin: 0 5px;
                    cursor: pointer;
                }
                :deep(.t-icon:first-child) {
                    margin-left: 15px;
                }
            }
        }
    }
}
/*滚动条*/
::-webkit-scrollbar {
    width: 0px;
    // border-radius: 0px;
}

::-webkit-scrollbar-thumb {
    background: transparent;
    // border-radius: 0px;
}
</style>
