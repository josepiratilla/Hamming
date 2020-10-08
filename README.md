# Hamming
In order to improbe my go skills I put myself this challenge. I'd like to deal with:
* Bit manipulation
* Concurrency
* Channels in concurrency
* IO Reader/writer
## Initial challenge
The plan is to have different modules connected to each other.
* **hamming**: will have the options to encode and decode. The idea is to have a stream as an input and output and see how can I connect it to a cocurrent thread.
* **channel**: also as a stream in and out, it will randomly modify the data transmitted. Initially it will be a clean channel, but latter it can add random errors, or maybe even bursts or errors. Should I have interlacing to avoid it? 
* **system**: The system will generate a hamming encoder, a channel, a hammind decoder and have all three interconnected and running concurrently. Then it will send some data, receive the response and check the error rate.
* **main**: The idea is to have a cli where we can select the parameters of the coding, the channel, the amount of data and others for the simulation. 