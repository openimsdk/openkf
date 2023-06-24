// Copyright © 2023 OpenIMSDK open source community. All rights reserved.
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

import { ElLoading } from 'element-plus';

// global loading
let loadingCount = 0;
let loadingInstance = null;

export const startLoading = () => {
    loadingCount++;
    loadingInstance = ElLoading.service({
        lock: true,
        text: 'Loading...',
        background: 'rgba(0, 0, 0, 0.7)'
    });
    setTimeout(() => {
        loadingInstance.close();
    }, 3000)
}

export const stopLoading = () => {
    loadingCount--;
    if(loadingCount === 0) {
        loadingInstance.close();
    }
}