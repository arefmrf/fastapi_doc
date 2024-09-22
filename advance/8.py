from fastapi import FastAPI
from fastapi.middleware.httpsredirect import HTTPSRedirectMiddleware
from fastapi.middleware.trustedhost import TrustedHostMiddleware
from fastapi.middleware.gzip import GZipMiddleware

from unicorn import UnicornMiddleware

app = FastAPI()

app.add_middleware(UnicornMiddleware, some_config="rainbow")


app.add_middleware(HTTPSRedirectMiddleware)
"""
Enforces that all incoming requests must either be https or wss.
Any incoming request to http or ws will be redirected to the secure scheme instead.
"""


app.add_middleware(
    TrustedHostMiddleware,
    allowed_hosts=["example.com", "*.example.com"]
)
"""
Enforces that all incoming requests have a correctly set Host header, in order to guard against HTTP Host Header attacks.
"""


app.add_middleware(GZipMiddleware, minimum_size=1000, compresslevel=5)
"""
Handles GZip responses for any request that includes "gzip" in the Accept-Encoding header.
The middleware will handle both standard and streaming responses.
"""


@app.get("/")
async def main():
    return {"message": "Hello World"}