import struct

packet_count = 5000000  # Let's push to 5 million packets (~500MB)
filename = "sample.bin"

print(f"Generating {filename} with {packet_count} packets...")

with open(filename, "wb") as f:
    for i in range(packet_count):
        packet_id = i
        port = 8080
        payload = b"A" * 94
        
        packed_data = struct.pack("!I H 94s", packet_id, port, payload)
        f.write(packed_data)

print("Generation complete!")