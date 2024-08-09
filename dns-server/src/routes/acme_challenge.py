from flask import Blueprint, request, Response
from src.controllers.acme_challenge_controller import AcmeChallengeController

acme_challenge_controller = AcmeChallengeController()

acme_challenge_bp = Blueprint('acme_challenge', __name__, url_prefix = '/_acme_challenge')

@acme_challenge_bp.route('', methods=['GET'])
def token() -> Response:
  return acme_challenge_controller.token(request)

@acme_challenge_bp.route('', methods=['POST'])
def challenge() -> Response:
  return acme_challenge_controller.challenge(request)
