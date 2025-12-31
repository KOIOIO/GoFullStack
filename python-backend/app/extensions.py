from pymongo import MongoClient
import yaml
import os

# 加载配置
config_path = os.path.join(os.path.dirname(os.path.dirname(__file__)), 'config/config.yaml')
with open(config_path, 'r', encoding='utf-8') as f:
    full_config = yaml.safe_load(f)
    mongo_config = full_config['Mongo']
    jwt_config = full_config['JWT']

# 构建MongoDB连接URI
mongo_uri = f"mongodb://{mongo_config['User']}:{mongo_config['Password']}@{','.join(mongo_config['Hosts'])}/"
client = MongoClient(
    mongo_uri,
    authSource=mongo_config['AuthSource'],
    authMechanism=mongo_config['AuthMechanism']
)
db = client[mongo_config['Database']]