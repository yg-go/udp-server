
import socket
import struct

def makeHeader(src, dst, com, datasize):
    return struct.pack("bbhi", src, dst, com, datasize)

def makeIQData(num):
    l = []
    for i in range(num):
        l.append(struct.pack("ff", i, i))

    ret = struct.pack("i",num)
    ret += struct.pack("q", 480000000)
    ret += struct.pack("i", 0)
    ret += struct.pack("i", num)

    for i in l :
        ret += i

    return ret

def StartUDPServer():
    MULTICAST_TTL = 2
    GROUP_IP = "239.0.0.253"
    MCAST_PORT = 10051

    sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM, socket.IPPROTO_UDP)
    sock.setsockopt(socket.IPPROTO_IP, socket.IP_MULTICAST_TTL, MULTICAST_TTL)

    # iq = makeIQData(5)
    iq = makeIQData(5)
    header = makeHeader(0, 0, 0x12, len(iq))
    payload = header + iq
    for i in range(1):
        sock.sendto(payload, (GROUP_IP, MCAST_PORT))

StartUDPServer()
