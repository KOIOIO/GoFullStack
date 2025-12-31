from app import create_app

app = create_app()

if __name__ == '__main__':
    # 修改port参数为8000，host设为0.0.0.0可允许外部访问（可选）
    app.run(debug=True, host='0.0.0.0', port=8000)