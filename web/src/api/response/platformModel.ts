/**
 * Copyright © 2023 OpenKF & OpenIM open source community. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

export interface GetSlackConfigResponse {
    bot_token: string,
    app_token: string,
    app_id: string, 
    client_id: string,
    client_secret: string,
    signing_secret: string,
    verification_token: string,
}

export interface GetSlackCustomerResponse {
    id: number,
    created_at: string,
    updated_at: string,
    uuid: string,
    email: string,
    nickname: string,
    avatar: string,
    description: string,
    is_enable: boolean,
    first_name: string,
    last_name: string,
    real_name: string,
    real_name_normalized: string,
    display_name: string,
    display_name_normalized: string,
    skype: string,
    phone: string,
    image_24: string,
    image_32: string,
    image_48: string,
    image_72: string,
    image_192: string,
    image_512: string,
    image_original: string,
    title: string,
    status_expiration: number,
    team: string,
}