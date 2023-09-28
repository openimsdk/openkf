from typing import Awaitable
import json

from constants.request import OpenAIChatMsgIn, KnowledgeBaseMsgIn
from fastapi.responses import StreamingResponse
import db.milvus as milvus
import utils.config as config
import router.router as router
from constants.response import BaseResponse

import asyncio
from langchain.embeddings import OpenAIEmbeddings
from langchain.chat_models import ChatOpenAI
from langchain.callbacks import AsyncIteratorCallbackHandler
from langchain.prompts.chat import ChatPromptTemplate
from langchain import LLMChain

PROMPT_TEMPLATE = """
Answer the question concisely and professionally based on known information. If you cannot get an answer from it, say "I cannot answer the question based on known information" and do not allow fabrications to be added to your answer.

Knowledge base:{context}
Question:{question}
"""

# TODO: Support multi history chat
# Use default knowledge base and do not use history support
# Now do not use flow chat
def knowledge_base_chat(msg: KnowledgeBaseMsgIn):
    config = router.kb_config

    db = milvus.MilvusDBService(
        embeddings=OpenAIEmbeddings(
            model=config.get_fastchat_models_embedding_model_name()),
        host=config.get_milvus_host(),
        port=config.get_milvus_port(),
        top_k=config.get_milvus_top_k(),
        score_threshold=config.get_milvus_score_threshold()
    )

    llm = ChatOpenAI(
        streaming=False,
        verbose=True,
        openai_api_key=config.get_openai_api_key(),
        openai_api_base="http://" + config.get_fastchat_openai_api_server_host() + ":" +
        config.get_fastchat_openai_api_server_port() + "/v1",
        model_name=config.get_fastchat_models_llm_model_name(),
    )
    docs = db.search(msg.query)
    context = "\n".join([doc.page_content for doc in docs])

    chat_prompt = ChatPromptTemplate.from_messages([("human", PROMPT_TEMPLATE)])
    chain = LLMChain(prompt=chat_prompt, llm=llm)
    
    res = chain({"context": context, "question": msg.query})
        
    return BaseResponse(data=res)