from enum import Enum
from typing import Annotated

from fastapi import Query, Body, Header
from pydantic import BaseModel, Field

from main import app


class ModelName(str, Enum):
    alex1 = "alex"
    resi1 = "resi"
    leni1 = "leni"


@app.get("/user/{name}")
def read_item(
        name: ModelName,

        q: list[str] | None = Query(default=None),

        w: Annotated[list[str] | None, Query(
            alias="item-query",
            title="Query string",
            description="Query s",
            deprecated=True,
            include_in_schema=False  # not show in openapi
            )
        ] = None,

        e: Annotated[
            str | None, Query(min_length=3, max_length=50, pattern="^a.*z$")
        ] = None,
        r: Annotated[
            str | None, Query(min_length=3, max_length=50, pattern="^a.*z$")
        ] = ...  # ... means required
):
    return {
        "enum_name": name.name,
        "enum_value": name.value,
        "w": w,
        "q": q,
        "e": e,
        "r": r,
    }


class Itemqqq(BaseModel):
    name: str
    description: str | None = Field(
        default=None, title="The description of the item", max_length=300
    )
    price: float = Field(gt=0, description="The price must be greater than zero")
    tax: float | None = None
    model_config = {
        "json_schema_extra": {
            "examples": [
                {
                    "name": "Foo",
                    "description": "A very nice Item",
                    "price": 35.4,
                    "tax": 3.2,
                }
            ]
        }
    }


@app.put("/qqq/")
async def update_item(
        item_id: int,
        item: Annotated[Itemqqq, Body(embed=True)],
        items: list[Itemqqq] | None = None
):
    results = {"item_id": item_id, "item": item, "items": items}
    return results


@app.get("/items/")
async def read_items(
        user_agent: Annotated[str | None, Header()] = None,
        user_agent2: Annotated[str | None, Header(convert_underscores=False)] = None,
        x_token: Annotated[list[str] | None, Header()] = None
):
    return {
        "User-Agent": user_agent,
        "User-Agent2": user_agent2,
        "X-token": x_token,
    }
