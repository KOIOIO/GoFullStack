from flask import Blueprint, request, jsonify
import jwt
import yaml
import os
from datetime import datetime, timedelta
from app.models.user import User
from app.extensions import jwt_config

auth_bp = Blueprint('auth', __name__)

@auth_bp.route('/register', methods=['POST'])
def register():
    """用户注册接口"""
    data = request.get_json()
    username = data.get('username')
    password = data.get('password')
    email = data.get('email')

    # 校验参数
    if not all([username, password, email]):
        return jsonify({"code": 400, "message": "缺少必要参数"}), 400

    # 检查用户名/邮箱是否已存在
    if User.get_user_by_username(username):
        return jsonify({"code": 400, "message": "用户名已存在"}), 400
    if User.get_user_by_email(email):
        return jsonify({"code": 400, "message": "邮箱已存在"}), 400

    # 创建用户
    user_id = User.create_user(username, password, email)
    return jsonify({"code": 200, "message": "success", "user_id": user_id}), 200

@auth_bp.route('/login', methods=['POST'])
def login():
    """用户登录接口（生成JWT Token）"""
    data = request.get_json()
    username = data.get('username')
    password = data.get('password')

    # 校验参数
    if not all([username, password]):
        return jsonify({"code": 400, "message": "缺少必要参数"}), 400

    # 查找用户并验证密码
    user = User.get_user_by_username(username)
    if not user or not User.verify_password(password, user['password']):
        return jsonify({"code": 401, "message": "用户名或密码错误"}), 401

    # 生成JWT Token
    expire_time = datetime.utcnow() + timedelta(hours=jwt_config['ExpireHours'])
    token = jwt.encode(
        {"user_id": str(user['_id']), "exp": expire_time},
        jwt_config['SecretKey'],
        algorithm="HS256"
    )
    return jsonify({"code": 200, "message": "success", "token": token}), 200
