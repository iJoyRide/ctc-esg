"""
Service level endpoints
Owns all endpoints related to health service
"""
from fastapi import APIRouter

router = APIRouter()

@router.get("/health")
def health():
    return {"status": "ok"}
