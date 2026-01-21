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
        if not self.base_url:
            logger.error("GIN_SERVICE_URL is missing or not set in .env file.")
        
        self.client = httpx.AsyncClient(base_url=self.base_url, timeout=30.0)
        
        
    async def close(self):
        await self.client.aclose()
    
    async def health_check(self):
        response = await self.client.get("/gin_health")
        response.raise_for_status()
        return response.json()

_gin_service: GinService | None = None

def get_gin_service() -> GinService:
    global _gin_service
    if _gin_service is None:
        _gin_service = GinService()
    return _gin_service
    
