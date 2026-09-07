import socket
import struct


def decode_response(response):
    # ----------------
    # Header
    # ----------------
    transaction_id, flags, qdcount, ancount, nscount, arcount = struct.unpack(
        "!6H", response[:12]
    )

    print("HEADER")
    print("ID:", transaction_id)
    print("Flags:", hex(flags))
    print("Questions:", qdcount)
    print("Answers:", ancount)
    print("Authority:", nscount)
    print("Additional:", arcount)

    # ----------------
    # Question
    # ----------------
    offset = 12

    labels = []

    while True:
        length = response[offset]
        offset += 1

        if length == 0:
            break

        labels.append(
            response[offset:offset + length].decode()
        )
        offset += length

    qname = ".".join(labels)

    qtype, qclass = struct.unpack(
        "!HH", response[offset:offset + 4]
    )
    offset += 4

    print("\nQUESTION")
    print("Name:", qname)
    print("Type:", qtype)
    print("Class:", qclass)

    # ----------------
    # Answer
    # ----------------
    for i in range(ancount):

        labels = []

        while True:
            length = response[offset]
            offset += 1

            if length == 0:
                break

            labels.append(
                response[offset:offset + length].decode()
            )
            offset += length

        name = ".".join(labels)

        answer_type, answer_class = struct.unpack(
            "!HH", response[offset:offset + 4]
        )
        offset += 4

        ttl = struct.unpack(
            "!I", response[offset:offset + 4]
        )[0]
        offset += 4

        rdlength = struct.unpack(
            "!H", response[offset:offset + 2]
        )[0]
        offset += 2

        rdata = response[offset:offset + rdlength]
        offset += rdlength

        # A record
        if answer_type == 1 and rdlength == 4:
            ip = socket.inet_ntoa(rdata)
        elif answer_type == 28 and rdlength == 16:
            ip = socket.inet_ntop(socket.AF_INET6, rdata)
        else:
            ip = rdata

        print("\nANSWER")
        print("Name:", name)
        print("Type:", answer_type)
        print("Class:", answer_class)
        print("TTL:", ttl)
        print("RDLength:", rdlength)
        print("RData:", ip)