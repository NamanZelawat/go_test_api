import sys
from kafka import KafkaConsumer


def run():
    print("Running the consumer")
    sys.stdout.flush()
    # Define Kafka broker address and topic
    bootstrap_servers = 'kafka:9092'
    topic = 'my-topic'

    # Create Kafka consumer
    consumer = KafkaConsumer(topic,
                            bootstrap_servers=bootstrap_servers,
                            auto_offset_reset='earliest',
                            enable_auto_commit=True,
                            group_id='my-group')

    # Start consuming messages
    try:
        for message in consumer:
            print(f"Received message: {message.value.decode('utf-8')}")
            sys.stdout.flush()
    except KeyboardInterrupt:
        print("Consumer stopped")
        sys.stdout.flush()
    finally:
        consumer.close()
