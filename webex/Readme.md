## Http Handler Procedure. 

Step 1: Start by writing a http handler function with signature ( w http.ResponseWriter, r *http.Request)
Step 2: Map the function to a URL path ex "/" http.HandleFunc("/", funcName)
Step 3: call http.ListenAndServe(addr string, handler(or nil)) 

http.ListenAndServe(addr string, handler) 
nil handler is the default hansler. 
[Example](webEx.go)

# Handler object

[Example](handlers/handleEx.go)


# type Server

uncomment this block
<a href="./webEx.go#L17"> Example</a>


#Query 

?field1=value1&field2=value2...


URL type implements func(*URL)Query

Comment out 
uncomment this block
<a href="./webEx.go#L17"> L17-26</a>


uncomment 
uncomment this block
<a href="./webEx.go#L17"> L13-16</a>