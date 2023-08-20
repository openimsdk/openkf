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

import os
import datetime

from loguru import logger
import constants.constants as constants


class KFLog:
    log_path = datetime.datetime.now().strftime("logs/%Y-%m-%d.log")
    log_level = constants.LOG_LEVEL_INFO
    
    @classmethod
    def init_logger(self, log_path: str, log_level: str):
        KFLog.log_level = log_level
        KFLog.log_path = os.path.join(log_path, datetime.datetime.now().strftime("%Y-%m-%d.log"))
        logger.add(KFLog.log_path, rotation="10 MB", retention="1 day", level=self.log_level)

    @classmethod
    def get_logger(self):
        return logger