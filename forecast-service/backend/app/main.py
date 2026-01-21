from fastapi import FastAPI
from app.api.router import api_router
from app.services.gin import get_gin_service, _gin_service
import logging

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

app = FastAPI()

app.include_router(api_router)

@app.on_event("startup")
async def startup():
    
    try:
        gin = get_gin_service()
        result = await gin.health_check()
        logger.info(f"Gin service is healthy: {result}")
    except Exception as e:
        logger.error(f"❌ Gin service health check failed: {e}")
        logger.error("FastAPI will continue, but Gin service may be unavailable")
        
        
@app.on_event("shutdown")
async def shutdown():
    if _gin_service:
        await _gin_service.close()