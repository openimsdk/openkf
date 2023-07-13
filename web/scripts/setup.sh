#!/bin/bash
# Copyright © 2023 OpenIM open source community. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Download packages
npm install --save
echo "[0] NPM install finished..."

# Prepare for open-im-sdk-wasm
cp -r node_modules/open-im-sdk-wasm/assets/* public/
mkdir -p src/utils/open-im-sdk-wasm
cp -r node_modules/open-im-sdk-wasm/lib/* src/utils/open-im-sdk-wasm
echo "[1] Copy open-im-sdk-wasm finished..."

# Open openim support
# Ref: https://doc.rentsoft.cn/#/js_v2/sdk_integrate/sdk_use_wasm
FILE_PATH="src/utils/open-im-sdk-wasm/api/index.js"
sed -i "s/\/\/ import IMWorker from '.\/worker?worker';/import IMWorker from '.\/worker?worker';/g" "$FILE_PATH"
# sed -i 's/worker = new Worker(new URL('\''\.\/worker\.js'\'', import.meta.url));/\/\/ worker = new Worker(new URL('\''\.\/worker\.js'\'', import.meta.url));/g' "$FILE_PATH"
sed -i "s/\/\/ worker = new IMWorker();/worker = new IMWorker();/g" "$FILE_PATH"
sed -i "s/worker = new Worker(new URL/\/\/ worker = new Worker(new URL/g" "$FILE_PATH" # Update: close webpack5 support
echo "[2] Open openim support finished..."
