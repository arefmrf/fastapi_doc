from enum import Enum
from typing import Union, Optional, Annotated

from fastapi import FastAPI, Query, Path, Body
from pydantic import BaseModel

app = FastAPI()


class Item(BaseModel):
    name: str
    price: float
    is_offer: Union[bool, None] = None
    is_have: str | None = None
    is_has: Optional[str] = None
    tags: list = []


@app.get("/")
def read_root():
    return {"Hello": "World"}


@app.get("/items/{item_id}")
def read_item(item_id: int, q: Union[str, None] = None):
    return {"item_id": item_id, "q": q}


@app.put("/items/{item_id}")
def update_item(item_id: int, item: Item):
    return {
        "item_name": item.name,
        "item_id": item_id,
        "price": item.price,
        "is_offer": item.is_offer,
        "is_have": item.is_have,
        "is_has": item.is_has
    }


@app.get("/items2/{item_id}")
async def read_items2(
        *,  # to have argument with default value first
        q: Annotated[str | None, Query(alias="item-query")] = None,
        item_id: Annotated[int, Path(title="The ID of the item to get", gt=0, le=1000)]
):
    results = {"item_id": item_id}
    if q:
        results.update({"q": q})
    return results


class Item2(BaseModel):
    name: str
    description: str | None = None
    price: float
    tax: float | None = None


class User2(BaseModel):
    username: str
    full_name: str | None = None


@app.put("/items3/{item_id}")
async def update_item3(
        item_id: int,
        item: Item2,
        user: User2,
        importance: Annotated[int, Body()]  # to add single item in body
):
    results = {"item_id": item_id, "item": item, "user": user, "importance": importance}
    return results


@app.put("/items4/{item_id}")
async def update_item3(
        item_id: int,
        # to have body like {"item": ...}
        item: Annotated[Item2, Body(embed=True)],
):
    results = {"item_id": item_id, "item": item}
    return results