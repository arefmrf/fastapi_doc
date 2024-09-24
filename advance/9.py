from fastapi import FastAPI, Request

app = FastAPI()


@app.get("/app")
def read_main():
    return {"message": "Hello World from main app"}


subapi = FastAPI()


@subapi.get("/sub")
def read_sub():
    return {"message": "Hello World from sub API"}


app.mount("/subapi", subapi)


#######################################################################

app = FastAPI(root_path="/api/v1")  # or fastapi run main.py --root-path /api/v1


@app.get("/app")
def read_main(request: Request):
    return {"message": "Hello World", "root_path": request.scope.get("root_path")}


app = FastAPI(
    servers=[
        {"url": "https://stag.example.com", "description": "Staging environment"},
        {"url": "https://prod.example.com", "description": "Production environment"},
    ],
    root_path="/api/v1",
)