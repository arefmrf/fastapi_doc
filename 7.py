import time

from fastapi import FastAPI, Request
from fastapi.middleware.cors import CORSMiddleware

app = FastAPI()


@app.middleware("http")
async def add_process_time_header(request: Request, call_next):
    start_time = time.perf_counter()
    response = await call_next(request)
    process_time = time.perf_counter() - start_time
    response.headers["X-Process-Time"] = str(process_time)
    return response


origins = [
    "http://localhost.tiangolo.com",
    "https://localhost.tiangolo.com",
    "http://localhost",
    "http://localhost:8080",
]

# allow_origins cannot be set to ['*'] for credentials to be allowed, origins must be specified.
app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    # allow_origin_regex='https://.*\.example\.org',
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
    # expose_headers - Indicate any response headers that should be made accessible to the browser. Defaults to [],
    # max_age - Sets a maximum time in seconds for browsers to cache CORS responses. Defaults to 600
)

@app.get("/")
async def main():
    return {"message": "Hello World"}