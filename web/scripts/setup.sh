#!/bin/bash
# Copyright © 2023 OpenIMSDK open source community. All rights reserved.
# Licensed under the MIT License (the "License");
# you may not use this file except in compliance with the License.

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
echo "[2] Open openim support finished..."
