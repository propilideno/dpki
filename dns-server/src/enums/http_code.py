from enum import Enum

class HttpCode(Enum):
  NOT_FOUND: int = 404
  SERVER_ERROR: int = 500
  UNPROCESSABLE_ENTITY: int = 422
