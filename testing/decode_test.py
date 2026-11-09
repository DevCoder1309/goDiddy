import socket
import struct


def decode_name(response, offset):
    labels = []
    original_offset = None

    while True:
        length = response[offset]

        if length & 0xC0 == 0xC0:
            pointer = ((length & 0x3F) << 8) | response[offset + 1]

            if original_offset is None:
                original_offset = offset + 2

            offset = pointer
            continue

        offset += 1

        if length == 0:
            break

        labels.append(response[offset:offset + length].decode())
        offset += length

    return ".".join(labels), original_offset if original_offset is not None else offset


def decode_response(response):
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

    offset = 12

    for _ in range(qdcount):
        qname, offset = decode_name(response, offset)

        qtype, qclass = struct.unpack("!HH", response[offset:offset + 4])
        offset += 4

        print("\nQUESTION")
        print("Name:", qname)
        print("Type:", qtype)
        print("Class:", qclass)

    for i in range(ancount):
        name, offset = decode_name(response, offset)

        answer_type, answer_class = struct.unpack(
            "!HH", response[offset:offset + 4]
        )
        offset += 4

        ttl = struct.unpack("!I", response[offset:offset + 4])[0]
        offset += 4

        rdlength = struct.unpack("!H", response[offset:offset + 2])[0]
        offset += 2

        rdata = response[offset:offset + rdlength]
        offset += rdlength

        if answer_type == 1 and rdlength == 4:
            decoded_data = socket.inet_ntoa(rdata)

        elif answer_type == 28 and rdlength == 16:
            decoded_data = socket.inet_ntop(socket.AF_INET6, rdata)

        elif answer_type == 5:
            decoded_data, _ = decode_name(response, offset - rdlength)

        else:
            decoded_data = rdata

        print("\nANSWER")
        print("Name:", name)
        print("Type:", answer_type)
        print("Class:", answer_class)
        print("TTL:", ttl)
        print("RDLength:", rdlength)
        print("RData:", decoded_data)