import time 
import struct 

package_size = 100
Time = time.time()




with open("sample.bin", "rb") as file1:
    while True:
        chunk = file1.read(package_size)
        if not chunk:
            break


        packet_id, port, payload = struct.unpack("!I H 94s", chunk)


end_time = time.time()
elapsed_time = end_time - Time
print(f"Elapsed time: {elapsed_time:.6f} seconds")

