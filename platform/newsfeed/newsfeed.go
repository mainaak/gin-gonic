package newsfeed

type Getter interface {
	GetAll() []Item
}

type Adder interface {
	Add(item Item)
}

type Item struct {
	Title string `json:"title"`
	Post string `json:"post"`
}

//This will act like a Database
type Repo struct {
	Items []Item
}

//A constructor that returna de-referenced value of the pointer &Repo
func New() *Repo {
	return &Repo{
		Items: []Item{},
	}
}

//A method that will only be accessible for variables of type *Repo (De-referenced value of &Repo)
func (r *Repo) Add(item Item){
	r.Items = append(r.Items, item)
}

//A method that will only be accessible for variables of type *Repo (De-referenced value of &Repo)
func (r *Repo) GetAll() []Item{
	return r.Items
}