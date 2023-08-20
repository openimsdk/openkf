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

import utils.config as config
import utils.log as log
from utils import banner
import constants.constants as constants

logger = None


def main():
    kf_config = config.KBConfig("config.yaml")
    global logger
    log.KFLog.init_logger(
        kf_config.get_app_log_path(), constants.LOG_LEVEL_DEBUG
        if kf_config.get_app_debug() else constants.LOG_LEVEL_INFO)
    logger = log.KFLog.get_logger()
    banner.kf_banner(kf_config.get_app_version(), kf_config.get_app_debug(),
                     kf_config.get_app_log_path())


if __name__ == "__main__":
    main()