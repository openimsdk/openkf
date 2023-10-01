import os

from constants.request import OpenAIChatMsgIn, KnowledgeBaseMsgIn
from fastapi.responses import StreamingResponse
import db.milvus as milvus
import utils.config as config
import router.router as router
import constants.response as response

from fastapi import File, Form, Body, Query, UploadFile
from langchain.embeddings import OpenAIEmbeddings

async def doc_upload(file: UploadFile):
    config = router.kb_config
    
    # Save to temp file
    file_content = await file.read()
    
    # Create path
    if not os.path.exists(config.get_app_file_dir()):
        os.makedirs(config.get_app_file_dir())
    
    filepath = os.path.join(config.get_app_file_dir(), file.filename)
    
    try:
        with open(filepath, "wb") as f:
            f.write(file_content)
    except Exception as e:
        return response.BaseResponse(code=500, msg=f"{file.filename} Upload error: {e}")
    
    # Get db
    db = milvus.MilvusDBService(
        embeddings=OpenAIEmbeddings(
            model=config.get_fastchat_models_embedding_model_name()),
        host=config.get_milvus_host(),
        port=config.get_milvus_port(),
        top_k=config.get_milvus_top_k(),
        score_threshold=config.get_milvus_score_threshold()
    )

    db.insert(filepath)

    return response.BaseResponse(code=200, msg=f"{file.filename} Upload success")