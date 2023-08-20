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

import yaml
from colorama import Fore


class KBConfig:
    '''
    KBConfig is a class that loads the configuration file.
    Use load_config to load the configuration file, and use get_* to get the configuration.
    '''
    def __init__(self, file: str) -> None:
        self._config = None
        
        with open(file, encoding="utf-8") as fp:
            self._config = yaml.safe_load(fp.read())

    def get_app_version(self) -> str:
        return self._config['app']['version']

    def get_app_host(self) -> str:
        return self._config['app']['host'] if self._config['app']['host'] else "0.0.0.0"

    def get_app_port(self) -> str:
        return self._config['app']['port']

    def get_app_debug(self) -> str:
        return self._config['app']['debug']

    def get_app_log_path(self) -> str:
        return self._config['app']['log_path']

    def get_app_doc(self) -> str:
        return self._config['app']['doc']

    def get_inference_device(self) -> str:
        return self._config['inference']['device']

    def get_model_history(self) -> str:
        return self._config['model']['history']

    def get_model_temperature(self) -> str:
        return self._config['model']['temperature']

    def get_model_top_k(self) -> str:
        return self._config['model']['top_k']

    def get_model_top_p(self) -> str:
        return self._config['model']['top_p']