import pydantic
from pydantic import BaseModel
from typing import List

class OpenAIMessage(BaseModel):
    role: str = "user"
    content: str = "hello"

class OpenAIChatMsgIn(BaseModel):
    model: str = "gpt-3.5-turbo"
    messages: List[OpenAIMessage]
    temperature: float = 0.7
    n: int = 1
    max_tokens: int = 1024
    stop: List[str] = []
    stream: bool = False
    presence_penalty: int = 0
    frequency_penalty: int = 0

class KnowledgeBaseMsgIn(BaseModel):
    query: str = "What is openim"