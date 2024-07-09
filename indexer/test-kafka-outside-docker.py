# test_kafka_connection.py

from kafka import KafkaProducer
from kafka.errors import KafkaError
import sys

def test_kafka_connection(broker_address):
    try:
        # Create a Kafka producer
        print("trying to send message to kafka from outside container")
        producer = KafkaProducer(bootstrap_servers=[broker_address])
        
        # Send a test message
        future = producer.send('test-topic', b'This is a test message')
        print("sending .....")
        # Block until a single message is sent (or timeout)
        result = future.get(timeout=3)
        
        print(f"Successfully sent message to {broker_address}: {result}")
        
        # Close the producer connection
        producer.close()
        
        return 0
    except KafkaError as e:
        print(f"Failed to connect to Kafka broker at {broker_address}: {e}")
        return 1

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python test_kafka_connection.py <broker_address>")
        sys.exit(1)
    
    broker_address = sys.argv[1]
    sys.exit(test_kafka_connection(broker_address))
