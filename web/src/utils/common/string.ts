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

export const getAvatarString = (name: string): string => {
    return name.length > 2 ? name.slice(0, 2) : name.slice(0, 1);
}

// To capitail the first
export const capitalizeFirstLetter = (str: string): string => {
    return str.replace(/\b[a-z]/g, (char) => char.toUpperCase());
}

// Change all letters to upper case
export const toUpperCase = (str: string): string => {
    return str.toUpperCase();
}