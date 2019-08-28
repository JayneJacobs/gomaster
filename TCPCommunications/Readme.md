# TCP Communications


<div>
<table>
<tr>
<th align="left">

<h5>Client: </h5>
<p>
<ul class="left">
    <li> Dial function</li>
    <li> Conn interfaxe</li>
    <li> Conn reads and writes</li>
</ul>
</p>
<h5>Server:</h5>
<p>
<ul>
    <li> Listen function</li>
      Type Listener
        Accept()(Conn, error)
        Close() error
        Addr() Addr
    <li> Listener interface</li>
    <li> Listener.Accept</li>
    <li> Conn interface again</li>
</ul>
</p>
</th>
<th>
<img src="tcpstreamCom.png">
TCP Client-Server
</th>
</tr>
<tr>
<th align="left">

<h5>Client: </h5>
<p>
<ul class="left">
    <li> Dial function</li>
    <li> Conn interfaxe</li>
    <li> Conn reads and writes</li>
</ul>
</p>
<h5>Server:</h5>
<p>
<ul>
    <li> Listen function</li>
    <li> ListenPacket function</li>
    <li> PacketConn interface</li>
    <li> ReadFrom WriteTo functions</li>
</ul>
</p>
</th>
<th>
<img src="udpstreamCom.png">
UDPClient-Server
</th>
</tr>
</table>
</div>

##Conn interface:

Supports:
 - ioreader interace / buffio interface
    Read(b []byte) (n int, err error)
- iowrite interface 
    Write(b []byte)(n int, err error)
- Close Interface
    Close() error


implemented by TCPConn type
Supports concurrency - multiple go routines can use the connection at the same time. 


## UDP 


One way; No handshakes
- Client Server
- Base64 Encoding

Packet Oriented

  
