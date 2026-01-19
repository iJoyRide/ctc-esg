from fastapi import FastAPI
from app.api.router import api_router
from app.services.gin import gin_service
import logging

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

app = FastAPI()

app.include_router(api_router)

@app.on_event("startup")
async def startup():
    
    try:
        result = await gin_service.health_check()
        logger.info(f"Gin service is healthy: {result}")
    except Exception as e:
        logger.error(f"❌ Gin service health check failed: {e}")
        logger.error("FastAPI will continue, but Gin service may be unavailable")
        
        
@app.on_event("shutdown")
async def shutdown():
    await gin_service.close()