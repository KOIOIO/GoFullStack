FROM python:3.12-slim

WORKDIR /app

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

COPY . .

# Replace localhost with mongo service name in config
RUN sed -i 's/127.0.0.1:27017/mongo:27017/g' config/config.yaml

EXPOSE 8000

CMD ["python", "run.py"]
