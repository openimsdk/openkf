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

import sys
import os 

sys.path.append(os.path.dirname(os.path.dirname(__file__)))


def kf_banner(version: str, debug: bool, log_path: str) -> None:
    print(f'''
--------------------------------------------------------------------------------------------------------------------
_______                    ______ ___________   ________      _____ 
__  __ \______________________  //_/__  ____/   ___  __ )_______  /_
_  / / /__  __ \  _ \_  __ \_  ,<  __  /_       __  __  |  __ \  __/
/ /_/ /__  /_/ /  __/  / / /  /| | _  __/       _  /_/ // /_/ / /_  
\____/ _  .___/\___//_/ /_//_/ |_| /_/          /_____/ \____/\__/  
       /_/                                                                                         

APP Mode:
- Version: %s
- Debug: %s
- Log path: %s

Github repo: https://github.com/OpenIMSDK/OpenKF. 
Official website: https://www.openim.online/en
OpenKF Slack: https://join.slack.com/t/openimsdk/shared_invite/zt-1tmoj26uf-_FDy3dowVHBiGvLk9e5Xkg (OpenKF clannels)
🎉🎉 Welcome to your contribution :)

Copyright © 2023 OpenIM open source community. All rights reserved.
Licensed under the Apache License (the "License");
--------------------------------------------------------------------------------------------------------------------
''' % (version, debug, log_path))
          