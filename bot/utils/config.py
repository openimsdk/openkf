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

import yaml
from colorama import Fore

sys.path.append(os.path.dirname(os.path.dirname(__file__)))

class KBConfig:
    '''
    KBConfig is a class that loads the configuration file.
    Use load_config to load the configuration file, and use get_* to get the configuration.
    '''
    def __init__(self, file: str) -> None:
        self._config = None
        
        with open(file, encoding="utf-8") as fp:
            self._config = yaml.safe_load(fp.read())

    ###############################
    # OpenKF bot config
    ###############################
    def get_app_version(self) -> str:
        return self._config['app']['version']

    def get_app_host(self) -> str:
        return self._config['app']['host'] if self._config['app']['host'] else "0.0.0.0"

    def get_app_port(self) -> str:
        return self._config['app']['port']

    def get_app_debug(self) -> str:
        return self._config['app']['debug']
    
    def get_app_file_dir(self) -> str:
        return self._config['app']['file_dir']

    def get_app_log_path(self) -> str:
        return self._config['app']['log_path']

    def get_app_doc(self) -> str:
        return self._config['app']['doc']
    
    def get_app_token(self) -> str:
        return self._config['app']['token']
    
    def get_openai_api_key(self) -> str:
        return self._config['app']['openai_api_key']

    ###############################
    # LLM langchain config
    ###############################
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
    
    ###############################
    # LLM fastchat config
    ###############################
    def get_fastchat_models_model_path(self) -> str:
        return str(self._config['fastchat']['models']['model_path'])
    
    def get_fastchat_models_llm_model_name(self) -> str:
        return str(self._config['fastchat']['models']['llm_model_name'])
    
    def get_fastchat_models_embedding_model_name(self) -> str:
        return str(self._config['fastchat']['models']['embedding_model_name'])
    
    def get_fastchat_controller_host(self) -> str:
        return str(self._config['fastchat']['controller']['host'])
    
    def get_fastchat_controller_port(self) -> str:
        return str(self._config['fastchat']['controller']['port'])
    
    def get_fastchat_model_worker_host(self) -> str:
        return str(self._config['fastchat']['model_worker']['host'])
    
    def get_fastchat_model_worker_port(self) -> str:
        return str(self._config['fastchat']['model_worker']['port'])
    
    def get_fastchat_model_worker_device(self) -> str:
        return str(self._config['fastchat']['model_worker']['device'])
    
    def get_fastchat_model_worker_limit_worker_concurrency(self) -> str:
        return str(self._config['fastchat']['model_worker']['limit_worker_concurrency'])
    
    def get_fastchat_model_worker_num_gpus(self) -> str:
        return str(self._config['fastchat']['model_worker']['num_gpus'])
   
    def get_fastchat_model_worker_max_gpu_memory(self) -> str:
        return str(self._config['fastchat']['model_worker']['max_gpu_memory'])
    
    def get_fastchat_openai_api_server_host(self) -> str:
        return str(self._config['fastchat']['openai_api_server']['host'])
    
    def get_fastchat_openai_api_server_port(self) -> str:
        return str(self._config['fastchat']['openai_api_server']['port'])
    
    ###############################
    # Milvus config
    ###############################
    def get_milvus_host(self) -> str:
        return self._config['milvus']['host']
    
    def get_milvus_port(self) -> str:
        return self._config['milvus']['port']
    
    def get_milvus_top_k(self) -> str:
        return self._config['milvus']['top_k']
    
    def get_milvus_score_threshold(self) -> str:
        return self._config['milvus']['score_threshold']