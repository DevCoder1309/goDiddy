import socket
from decode_test import decode_response

def build_query(domain):
    # 1. DNS Header (12 bytes)
    # ID: 0xAAAA (random-ish ID), Flags: 0x0100 (standard query, recursion desired)
    # Questions: 1, Answers: 0, Authority: 0, Additional: 0
    header = b'\xaa\xbb\x01\x00\x00\x01\x00\x00\x00\x00\x00\x00'
    
    # 2. Query Section (QNAME)
    # Convert 'example.com' into length-prefixed labels: '\x07example\x03com\x00'
    qname = b''.join([bytes([len(part)]) + part.encode() for part in domain.split('.')]) + b'\x00'
    
    # 3. Type (A = 0x0001) and Class (IN = 0x0001)
    qtype_qclass = b'\x00\x05\x00\x01'
    
    return header + qname + qtype_qclass

def query_dns(domain, server_ip="127.0.0.1"):
    query_packet = build_query(domain)
    
    # Create UDP socket
    sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    sock.settimeout(3.0)
    
    try:
        # Send to DNS server on port 8080
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
    query_dns("www.google.com")