FROM python:3.11-slim

WORKDIR /app

RUN pip install uv

COPY pyproject.toml uv.lock .
RUN uv sync

COPY train.py .
COPY Dataset1.xlsx .

CMD ["uv", "run", "train.py"]