from flask import Flask, jsonify
from src.enums.http_code import HttpCode
from src.helpers.utils import error_dict

def app_instance() -> Flask:
  app = Flask(__name__)

  @app.errorhandler(HttpCode.NOT_FOUND.value)
  def not_found(e: Exception):
      return jsonify(error_dict('Resource not found.')), HttpCode.NOT_FOUND.value

  @app.errorhandler(HttpCode.SERVER_ERROR.value)
  def server_error(e: Exception):
      return jsonify(error_dict('An internal error has ocurred.')), HttpCode.SERVER_ERROR.value

  return app
