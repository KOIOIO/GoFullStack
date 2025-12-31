from flask import Blueprint, request, jsonify
from app.models.user import User
from app.extensions import jwt_config
import jwt
import datetime

auth_bp = Blueprint('auth', __name__)

@auth_bp.route('/register', methods=['POST'])
def register():
    data = request.get_json()
    username = data.get('username')
    password = data.get('password')
    email = data.get('email')

    if not username or not password or not email:
        return jsonify({"message": "Missing required fields"}), 400

    if User.get_user_by_username(username):
        return jsonify({"message": "Username already exists"}), 409

    if User.get_user_by_email(email):
        return jsonify({"message": "Email already exists"}), 409

    user_id = User.create_user(username, password, email)
    return jsonify({"message": "User created successfully", "user_id": user_id}), 201

@auth_bp.route('/login', methods=['POST'])
def login():
    data = request.get_json()
    username = data.get('username')
    password = data.get('password')

    if not username or not password:
        return jsonify({"message": "Missing required fields"}), 400

    user = User.get_user_by_username(username)
    if not user or not User.verify_password(password, user['password']):
        return jsonify({"message": "Invalid username or password"}), 401

    # Generate JWT
    token_payload = {
        "user_id": str(user['_id']),
        "username": user['username'],
        "exp": datetime.datetime.utcnow() + datetime.timedelta(hours=jwt_config['ExpireHours'])
    }
    
    token = jwt.encode(token_payload, jwt_config['SecretKey'], algorithm="HS256")

    return jsonify({"token": token, "message": "Login successful"}), 200
