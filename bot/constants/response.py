import pydantic
from pydantic import BaseModel

class BaseResponse(BaseModel):
    code: int = pydantic.Field(200, description="HTTP status code")
    msg: str = pydantic.Field("success", description="HTTP status message")
    data: dict = pydantic.Field({}, description="HTTP response data")