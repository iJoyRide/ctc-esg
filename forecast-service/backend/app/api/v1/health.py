"""
Service level endpoints
Owns all endpoints related to health service
"""
from fastapi import APIRouter
from app.services.gin import gin_service

router = APIRouter()

@router.get("/health")
async def health():
    gin_health = await gin_service.health_check()
    return {"status": gin_health}
