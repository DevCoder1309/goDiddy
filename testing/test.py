import socket
from decode_test import decode_response

def build_query():
    header = b'\xaa\xbb\x01\x00\x00\x02\x00\x00\x00\x00\x00\x00'

    qname1 = b''.join([bytes([len(p)]) + p.encode() for p in "google.com".split('.')]) + b'\x00'
    qname2 = b''.join([bytes([len(p)]) + p.encode() for p in "www.google.com".split('.')]) + b'\x00'

    question1 = qname1 + b'\x00\x01\x00\x01'
    question2 = qname2 + b'\x00\x05\x00\x01'

    return header + question1 + question2

def query_dns(server_ip="127.0.0.1"):
    query_packet = build_query()

    sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    sock.settimeout(3.0)

    try:
        sock.sendto(query_packet, (server_ip, 8080))
        response, _ = sock.recvfrom(512)
        decode_response(response)
        print(f"Successfully received {len(response)} bytes from {server_ip}")
        return response

    except socket.timeout:
        print("The DNS request timed out.")

    finally:
        sock.close()

if __name__ == "__main__":
    query_dns()