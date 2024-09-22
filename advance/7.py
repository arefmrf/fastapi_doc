from dataclasses import dataclass, field
from typing import List

from fastapi import FastAPI


@dataclass
class Item:
    name: str
    price: float
    description: str | None = None
    tax: float | None = None


app = FastAPI()


@app.post("/items/")
async def create_item(item: Item):
    return item



@dataclass
class Item2:
    name: str
    price: float
    tags: List[str] = field(default_factory=list)
    description: str | None = None
    tax: float | None = None


@app.get("/items/next", response_model=Item2)
async def read_next_item():
    return {
        "name": "Island In The Moon",
        "price": 12.99,
        "description": "A place to be playin' and havin' fun",
        "tags": ["breater"],
    }

###################################################################################################
@dataclass
class Item3:
    name: str
    description: str | None = None


@dataclass
class Author:
    name: str
    items: List[Item3] = field(default_factory=list)

@app.post("/authors/{author_id}/items/", response_model=Author)  #
async def create_author_items(author_id: str, items: List[Item3]):  #
    return {"name": author_id, "items": items}


@app.get("/authors/", response_model=List[Author])
def get_authors():  #
    return [  #
        {
            "name": "Breaters",
            "items": [
                {
                    "name": "Island In The Moon",
                    "description": "A place to be playin' and havin' fun",
                },
                {"name": "Holy Buddies"},
            ],
        },
        {
            "name": "System of an Up",
            "items": [
                {
                    "name": "Salt",
                    "description": "The kombucha mushroom people's favorite",
                },
                {"name": "Pad Thai"},
                {
                    "name": "Lonely Night",
                    "description": "The mostests lonliest nightiest of allest",
                },
            ],
        },
    ]