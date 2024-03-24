import sys
import asyncio
import nats


async def message_handler(msg):
    subject = msg.subject
    data = msg.data.decode()
    print(f"Received a message on '{subject}': {data}")
    sys.stdout.flush()


async def run_message():
    print(f"Trying to connect to the NATS")
    sys.stdout.flush()
    nc = await nats.connect(servers=["nats://message:4222"])

    await nc.subscribe("example", cb=message_handler)

    # Keep the connection open indefinitely
    await nc.flush()


def run():
    print("Running the consumer")
    sys.stdout.flush()
    loop = asyncio.get_event_loop()
    loop.run_until_complete(run_message())
    loop.run_forever()
