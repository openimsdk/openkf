import sys
import os

from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware

from apis.chat import knowledge_base_chat
from apis.docs import doc_upload
import utils.config as config

# Add parent dir to sys.path so we can import from parent dir
sys.path.append(os.path.dirname(os.path.dirname(__file__)))

kb_config = None

def get_api(config: config.KBConfig) -> FastAPI:
    app = FastAPI()
    
    global kb_config
    kb_config = config
    
    app.add_middleware(
        CORSMiddleware,
        allow_origins=["*"],
        allow_credentials=True,
        allow_methods=["*"],
        allow_headers=["*"],
    )

    app.post("/api/v1/ask")(knowledge_base_chat)
    app.post("/api/v1/doc/upload")(doc_upload)
    
    return app