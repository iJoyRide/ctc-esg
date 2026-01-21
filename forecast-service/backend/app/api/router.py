"""
API registry, groups all API routers.
Groups endpoints by version, v1
"""
from fastapi import APIRouter
from app.api.v1.health import router as health_router
from app.api.v1.gin import router as gin_router
api_router = APIRouter(prefix="/api/v1")

api_router.include_router(health_router, tags=["fastapi"])
api_router.include_router(gin_router, tags=["gin"])
