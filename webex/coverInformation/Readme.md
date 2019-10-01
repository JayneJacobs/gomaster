# Web template
html.Templates protects against code injection and cross-site scripting
{{}} Substitution block
{{.}} Indicates the current object
    * {{.Name}} Indicates a field called main in the current object. 

{{range .mylist} <p>{{.ID}}</p>{{end}} => loop through a slice of objects inside the current object and then shos a field namce called ID on each object. 


{{if condition}}true{{else}}false{{end}} => if statements in Go templates

[bootstrap](http://getbootstrap.com)

Modify index.html to point to local versions of js and css files.

Modify cover.css file to point to the png file in assets folder



#Rest API


# Backend
Login Page 
Form values 
use ParseForm
func(r *Request) ParseForm() error

Form url.Values //Populated
type Values map[string][]string
ex. 
v.Get(name)
# Test



# 
