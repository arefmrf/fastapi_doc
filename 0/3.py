from typing import Annotated

from fastapi import Form

from main import app


@app.post("/items/", status_code=201)
async def create_item(name: str):
    return {"name": name}


# FORM DATA
@app.post("/login/")
async def login(username: Annotated[str, Form()], password: Annotated[str, Form()]):
    return {"username": username}
