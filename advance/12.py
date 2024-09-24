from functools import lru_cache
from typing import Annotated

from fastapi import FastAPI, Depends
from fastapi.testclient import TestClient
from pydantic_settings import BaseSettings, SettingsConfigDict


class Settings(BaseSettings):
    app_name: str = "Awesome API"
    admin_email: str
    items_per_user: int = 50
    model_config = SettingsConfigDict(env_file=".env")


###
settings = Settings()
# or
@lru_cache
def get_settings():
    return Settings()
###


app = FastAPI()


@app.get("/info")
async def info():
    return {
        "app_name": settings.app_name,
        "admin_email": settings.admin_email,
        "items_per_user": settings.items_per_user,
    }


@app.get("/info2")
async def info(settings2: Annotated[Settings, Depends(get_settings)]):
    return {
        "app_name": settings2.app_name,
        "admin_email": settings2.admin_email,
        "items_per_user": settings2.items_per_user,
    }


### and for test

client = TestClient(app)


def get_settings_override():
    return Settings(admin_email="testing_admin@example.com")


app.dependency_overrides[get_settings] = get_settings_override


def test_app():
    response = client.get("/info")
    data = response.json()
    assert data == {
        "app_name": "Awesome API",
        "admin_email": "testing_admin@example.com",
        "items_per_user": 50,
    }
