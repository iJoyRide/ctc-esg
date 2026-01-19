import os
import httpx
import logging
from dotenv import load_dotenv

load_dotenv()

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)


class GinService:
    def __init__(self):
        self.base_url = os.getenv("GIN_SERVICE_URL")
        self.client = httpx.AsyncClient(base_url=self.base_url, timeout=30.0)
        
    async def close(self):
        await self.client.aclose()
    
    async def health_check(self):
        response = await self.client.get("/health")
        response.raise_for_status()
        return response.json()

gin_service = GinService()
