from typing import Union
from flask import Request, jsonify, make_response, Response
from src.helpers.utils import error_dict
from src.enums.http_code import HttpCode

class AcmeChallengeController():
  def token(self, request: Request) -> Response:
    domain_name: Union[str, None] = request.args.get('domain_name', default = None)

    if domain_name is None:
        return make_response(
          jsonify(error_dict("The field 'domain_name' is required.")),
          HttpCode.UNPROCESSABLE_ENTITY.value,
        )

    return jsonify({
      'success': domain_name,
    })

  def challenge(self, request: Request) -> Response:
    return jsonify({
      'success': 'test',
    })