# Protocol Buffers

* Language and platform neutral mechanism for serializing structured data
* Fast / Efficient
* Extensible and backward compatible
* great for commumication between microservices
* [Tutorial](https://developers.google.com/protocol-buffers)
<a href="https://developers.google.com/protocol-buffers">Tutorial</a>



## How Protocol Buffers Work

1. Write a .proto file to define data structure
2. use protocol buffer compiler to generate code from .proto file to endoce and decode data
3. Generated code provides functions / objects to read and write dtata streams
4. .proto file can be updated recompiled and deployed without affecting existing services compiled against old code. 

## Backward compatiblity

Add new fieilds but do not remove fields


keyword message is a type 
 ex. 
   message Person {
       required string name = 1; // must always be preasent or will fail
       required int32 id = 2;
       optional string email = 3; // can be ignored Not available in proto3

       enum PhoneType {
           MOBILE = 0;
           HOME = 1;
           OWRK = 2;
       }

       message PhoneNumber {
           required string number = 1;
           optional PhoneType type =2 [default = HOME];

       }
       repeated PhoneNumber phone = 4; // repeated means a list
   }


   ## Compare to XML
    - Simpler
    - 3/10 times smaller
    - 20-100 times faster
    - less ambiguous
    - genrate data access classes easier to use. 

    # Version 3

        good support for Golang; 
        Not compatible with proto2

 ## Example

 ```proto3

 message SearchRequest {
     string query = 1;
     int32 page_number = 2;
     int32 result_per_page = 3

 }       
 ```

 # Go Support

 - Get proto package, download protoc and protoc-gen-go
 - Write proto file with data structure
 - Compile protofile to get a Go File
 - Use Go file in code with Proto package

 ## Protocol Buffers Go Support
[Go Tutorial](https://developers.google.com/protocol-buffers/docs/gotutorial)
[Go Doc](https://godoc.org/github.com/golang/protobuf/proto#Buffer)
[Practical Guid to Protocol Buffers for Golang](http://www.minaandrawos.com/category/protocol-buffers/)


## File Structure

```proto3
syntax = "proto3";
package = tutorial;

message Person {
    string name = 1;
    int32 id = 2; //Unique ID 
    string email = 3;

    enum PhoneType {
        MOBILE = 0;
        HOME = 1;
        WORK = 2;
    }

    message PhoneNumber {
        string number = 1;
        PhoneType type =2;
    }
    repeated PhoneNumber phones = 4;
}
```

## Compile

1. Install compiler
2. go get -u github.com/golang/protobuf/protoc-gen-go
   It is installed in $GOBIN = $GOPATH/bin
3 run compiler

 use --go_out option to create Go Code

```sh
protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto
```


#Example Output 

[Ship.pb.go](./protout/Ship.pb.go)



