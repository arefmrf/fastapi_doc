from contextlib import asynccontextmanager

from fastapi import FastAPI
from fastapi.testclient import TestClient


def fake_answer_to_everything_ml_model(x: float):
    return x * 42


ml_models = {}


@asynccontextmanager
async def lifespan(app: FastAPI):
    # Load the ML model
    ml_models["answer_to_everything"] = fake_answer_to_everything_ml_model
    yield
    # Clean up the ML models and release the resources
    ml_models.clear()


app = FastAPI(lifespan=lifespan)


@app.get("/predict")
async def predict(x: float):
    result = ml_models["answer_to_everything"](x)
    return {"result": result}


items = {}


@app.on_event("startup")
async def startup_event():
    items["foo"] = {"name": "Fighters"}
    items["bar"] = {"name": "Tenders"}


@app.get("/items/{item_id}")
async def read_items(item_id: str):
    return items[item_id]


def test_read_items():
    with TestClient(app) as client:
        response = client.get("/items/foo")
        assert response.status_code == 200
        assert response.json() == {"name": "Fighters"}


############################################################################################################

from typing import Annotated

from fastapi import Depends, FastAPI
from fastapi.testclient import TestClient

app = FastAPI()


async def common_parameters(q: str | None = None, skip: int = 0, limit: int = 100):
    return {"q": q, "skip": skip, "limit": limit}


@app.get("/items/")
async def read_items(commons: Annotated[dict, Depends(common_parameters)]):
    return {"message": "Hello Items!", "params": commons}


@app.get("/users/")
async def read_users(commons: Annotated[dict, Depends(common_parameters)]):
    return {"message": "Hello Users!", "params": commons}


client = TestClient(app)


async def override_dependency(q: str | None = None):
    return {"q": q, "skip": 5, "limit": 10}


app.dependency_overrides[common_parameters] = override_dependency


def test_override_in_items():
    response = client.get("/items/")
    assert response.status_code == 200
    assert response.json() == {
        "message": "Hello Items!",
        "params": {"q": None, "skip": 5, "limit": 10},
    }


def test_override_in_items_with_q():
    response = client.get("/items/?q=foo")
    assert response.status_code == 200
    assert response.json() == {
        "message": "Hello Items!",
        "params": {"q": "foo", "skip": 5, "limit": 10},
    }


def test_override_in_items_with_params():
    response = client.get("/items/?q=foo&skip=100&limit=200")
    assert response.status_code == 200
    assert response.json() == {
        "message": "Hello Items!",
        "params": {"q": "foo", "skip": 5, "limit": 10},
    }