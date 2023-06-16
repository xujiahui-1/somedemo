# function
## Reader
带有缓存的读取，从硬盘中读取数据时先写入缓存，
- NewReaderSize():传入reader和size，返回一个带有缓冲的Reader
  - NewReaderSize() 封装bufio.Reader，给他安排一个Size指定缓存大小,如果size小于16，则为16
  - 传入io.Reader,返回bufio.Reader 获得一个带缓存的bufio.Reader对象
- NewReader():传入reader给予其默认4096缓冲的Reader
- ReadString():传入一个字符，在缓冲中读取直到读取到该字符为止
- Reset() :重制reader
- Read():将缓冲中的数据读取到传入的字节数组中，如果缓冲中的数据别读取完，返回的err=io.EOF
- Peek(n):引用缓存的前n个切片，注意不是读取，只是引用，n不能大于总缓存，
- ReadByte():读取缓冲中的第一个字节
- UnreadByte():取消最近一次读取的最后一个字节
- ReadRune():读取一个UTF8的unicode值，返回编码值，长度 和err
- UnReadRune():吐出一个UTF8的unicode值。如果最近一个读取不是用的ReadRune，则报错
- ReadLine():读取一行，读到换行符 不推荐时使用 ，建议使用readBytes
- ReadSlies():也是读取分割，不推荐时使用
- Buffer():只有读取一次后，才能使用
- Discard(n):抛弃前n个bytes
- Size():返回的是Reader中buff的大小。而不是存了多少数据的大小
## Writer
在将数据写入磁盘的时候先写入缓存，积攒多了再统一通过IO写入磁盘，减少了频繁的io操作
- 实打实