"""
Service level endpoints
Owns all endpoints related to health service
"""
from fastapi import APIRouter, HTTPException
from app.services.gin import get_gin_service

router = APIRouter()

@router.get("/gin_health")
async def health():
    try:
        gin = get_gin_service()
        gin_health = await gin.health_check()
        return {"status": gin_health}
    except Exception as e:
        raise HTTPException(
            status_code=503,
            detail=f"Gin service is unavailable, {e}"
        )

