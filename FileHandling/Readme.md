File handling with Go

- Open and Create
- Move Copy delete
- Read/Write
- Status ex. last modified. 
- Watchers

os package
https://golang.org/pkg/os/

File handling
 - Json
 - XML
 - CSV

 file, err := os.Open("filename") // for read access
 if err != nil {
     log.Fatal(err)
 }


Example Read and Write: 
data := make([]byte, 100)
count, err := file.Read(data)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("read %d bytes: %q\n", count, data[:count])


See [Example](../Hydra/filehandler/main.go)


File system permissions notation

copy contents from one file to another

	_, err = io.Copy(fNew, fOld)

ioutil package

func ReadAll(r io.Reader) ([]byte error)

func main() {
    r := strings.NewReader("Go is a genteral-purpose language designed with systems programming in mind)
    b, err := ioutil.ReadAll(r)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s", b)
}
ReadFile(filename string) ([]byte, error) {

}

func NewScanner(r io.Reader) *Scanner 
// returns a new Scanner to read from r.  
func(*Scanner) Scan() bool


## Json
Marshal takes a struct and creates json
Json Example 
See [Example](../jsonEx/main.go)


Unmarshal takes json from input and creates a struct. 

## XML

See [Example](../xmlEx/main.go)
www.w3school.com/xml/dom_nodes.asp
golang.org/pkg/encoding/xml

func Marshal(v interface{})([]byte, error)

- Encoding
- Decoding
- Writing 

## XML

See [Example](../csvEx/main.go)
 
https://golang.org/pkg/encoding/csv

func Marshal(v interface{})([]byte, error)

- Encoding
- Decoding
- Writing 

NewReader 
//takes io.Reader interface type

func NewReader(r io.Reader) *Reader