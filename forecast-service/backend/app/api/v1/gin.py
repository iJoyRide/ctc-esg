"""
Service level endpoints
Owns all endpoints related to gin service
"""

from fastapi import APIRouter

router = APIRouter()

@router.get("/health")
def health():
    return {"gin status": gin_health}
