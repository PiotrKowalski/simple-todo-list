package todo_list

type Status string

var (
	StatusCompleted   Status = "completed"
	StatusUnCompleted Status = "uncompleted"
)

type TodoList struct {
	Id   string `bson:"_id" json:"id,omitempty"`
	Name string `json:"name" bson:"name"`

	Entries []entry `json:"entries" bson:"entries"`
}

func (t TodoList) IsAggregateRoot() {}

func (t TodoList) GetId() string {
	return t.Id
}

func (t TodoList) GetName() string {
	return t.Name
}

func (t TodoList) GetEntries() []entry {
	return t.Entries
}

func (t TodoList) AddEntry(e entry) {
	t.Entries = append(t.Entries, e)
}

type entry struct {
	Id          uint   `json:"id" bson:"id"`
	Description string `json:"description" bson:"description"`
	Status      Status `json:"status" bson:"status"`
}

func (e entry) GetId() uint {
	return e.Id
}

func (e entry) GetDescription() string {
	return e.Description
}

func (e entry) GetStatus() Status {
	return e.Status
}
