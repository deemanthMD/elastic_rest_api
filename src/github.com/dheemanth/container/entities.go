package container

type Job_history struct {
	Company string `json:"Name"`
	Role    string `json:"Role"`
}

type Employee struct {
	Name        string       `json:"Name"`
	Age         int          `json:"Age"`
	Designation string       `json:"Designation"`
	Email_id    string       `json:"Email_id"`
	Experience  float32      `json:"Experience"`
	Job_history *Job_history `json:"Job_history"`
}

type coredata struct {
	Index  string    `json:"_index"`
	Type   string    `json:"_type"`
	Id     string    `json:"_id"`
	Score  float32   `json:"_score"`
	Source *Employee `json:"_source"`
}

type hits struct {
	Total     int         `json:"total"`
	Max_score float32     `json:"max_score"`
	Hits      []*coredata `json:"hits"`
}

type ESResponse struct {
	Took      int    `json:"took"`
	Timed_out bool   `json:"timed_out"`
	Shards    string `json:"_shards"`
	Hits      *hits  `json:"hits"`
}

type Shards struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Skipped    int `json:"skipped"`
	Failed     int `json:"failed"`
}
