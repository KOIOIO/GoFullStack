from flask import Flask
from flask_cors import CORS
from app.routes.auth import auth_bp

def create_app():
    app = Flask(__name__)
    CORS(app)  # Enable CORS for all routes
    
    # 注册蓝图
    app.register_blueprint(auth_bp)
    
    return app