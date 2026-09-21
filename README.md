 Is Go Actually Faster Than Python? (Packet Parsing Benchmark)

So, I wanted to run a quick test to see if Go could actually outrun Python when chewing through a massive binary data file. 

I threw a dataset of 41,666,666 records (100 bytes a pop) at both languages to see how they handled heavy raw byte parsing. Turns out Go *is* faster, but getting it there took a bit of troubleshooting.

 The Results

* Go (Optimized):** ~0.65 seconds
* Python (`struct.unpack`):** ~1.39 seconds
* Go (Naive/Reflection):** ~85.0 seconds (yep, really)

---

 What Happened Here?

 1. Python Out-of-the-Box
Even though Python is interpreted, its built-in `struct` module hands the heavy lifting down to C extensions. Reading chunks in 100-byte blocks (`!I H 94s`) let Python tear through the file in **1.39 seconds** right out of the gate.

 2. The Go "Reflection Trap"
My first Go attempt used `binary.Read` to map structs dynamically. Because Go had to use runtime reflection on all 41 million iterations, it crawled like a snail and took **85 seconds**. 

 3. Fixing Go
To fix it, I cut out reflection entirely:
* Buffered Readers (`bufio`):** Stopped making millions of individual system calls by pulling blocks into RAM efficiently.
* Direct Byte Slicing:** Switched to manual memory slicing (`binary.LittleEndian`) so the Go compiler could output pure, raw machine instructions. 

That dropped Go's time down to a blistering **0.65 seconds**, beating the Python baseline. 

---

 Running It

Because the raw binary file is nearly 500 MB (which exceeds GitHub's limits), `sample.bin` is excluded from version control. However, I've included the data generator script in the repository so you can build the exact same dataset locally.

1. To Generate
 Go  to file - python_data_maker.py ( I know so creative for a name)- run using either run button of python run python_data_maker.py).
