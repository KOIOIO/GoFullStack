from app.extensions import db
import bcrypt
from bson.objectid import ObjectId
from datetime import datetime

class User:
    collection = db['users']

    @staticmethod
    def create_user(username, password, email):
        """创建用户（注册）"""
        # 密码加密
        hashed_pwd = bcrypt.hashpw(password.encode('utf-8'), bcrypt.gensalt())
        user_data = {
            "username": username,
            "password": hashed_pwd.decode('utf-8'),
            "email": email,
            "confirm_password": hashed_pwd.decode('utf-8'),
            "created_at": datetime.utcnow(),
            "updated_at":datetime.utcnow()
        }
        result = User.collection.insert_one(user_data)
        return str(result.inserted_id)

    @staticmethod
    def get_user_by_username(username):
        """通过用户名查询用户"""
        return User.collection.find_one({"username": username})

    @staticmethod
    def get_user_by_email(email):
        """通过邮箱查询用户"""
        return User.collection.find_one({"email": email})

    @staticmethod
    def verify_password(plain_pwd, hashed_pwd):
        """验证密码"""
        return bcrypt.checkpw(plain_pwd.encode('utf-8'), hashed_pwd.encode('utf-8'))