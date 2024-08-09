from dotenv import load_dotenv
import os
from src.helpers.utils import str2bool
from src.routes.acme_challenge import acme_challenge_bp
from src.app.app import app_instance

app = app_instance()
app.register_blueprint(acme_challenge_bp)

if __name__ == '__main__':
    load_dotenv()

    app.run(
        host = os.getenv('APP_URL'),
        port = os.getenv('APP_PORT'),
        debug = str2bool(os.getenv('APP_DEBUG')),
    )
