Simple Benchmark scripts to test serialization and deserialization in difference languages.
The program takes in two arguments, number of iterations and number of person objects to serialize/deserialize.

Below experiments completed on a Macbook Pro (Apple M2 Pro), 16GB RAM, MacOS Ventura 13.5.2

### Python

```
python3 ser_der.py -i 30 -p 1000

Number of persons: 1000. Iterations: 30
Avg Serialization time (30 iterations): 442.219106 milliseconds
Size of serialized data: 5644051 bytes
Avg Deserialization time (30 iterations): 61.056574 milliseconds
```