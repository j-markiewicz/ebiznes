from ollama import chat
from fastapi import FastAPI, Request, Response

app = FastAPI()


@app.post("/")
async def proxy(req: Request):
    content = (await req.body()).decode("utf-8")
    response = chat(
        model="tinyllama",
        messages=[
            {
                "role": "user",
                "content": content,
            },
        ],
    )

    return Response(response.message.content)
