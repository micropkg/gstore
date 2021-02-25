package gstore

var s = make(map[string][]string)

//Store :
type Store []string

//Open : open store object
func Open(id string) Store {
	s[id] = nil
	return s[id]
}

//Get Values
func (x Store) Get() []string {
	return x
}

//Append Value
func (x Store) Append(v ...string) {
	x = append(x, v...)
}

//Clear Values
func (x Store) Clear() {
	x = nil
}
