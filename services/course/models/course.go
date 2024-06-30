package models

type WriteCourse struct {
	ID          string        `bson:"id" json:"id,omitempty"`
	Title       string        `bson:"title" json:"title,omitempty"`
	Description string        `bson:"description" json:"description,omitempty"`
	CreatorID   string        `bson:"creator_id" json:"creator_id,omitempty"`
	Likes       int32         `bson:"likes" json:"likes,omitempty"`
	Students    []string      `bson:"students" json:"students,omitempty"`
	Topics      []string      `bson:"topics" json:"topics,omitempty"`
	Modules     []WriteModule `bson:"modules" json:"modules,omitempty"`
	UpdatedAt   string        `bson:"updated_at" json:"updated_at,omitempty"`
	CreatedAt   string        `bson:"created_at" json:"created_at,omitempty"`
}

// Module represents a module within a course
type WriteModule struct {
	ID    string      `bson:"id" json:"id,omitempty"`
	Title string      `bson:"title" json:"title,omitempty"`
	Tasks []WriteTask `bson:"tasks" json:"tasks,omitempty"`
}

// Task represents a task within a module
type WriteTask struct {
	ID   string `bson:"id" json:"id,omitempty"`
	Task string `bson:"task" json:"task,omitempty"`
	Type string `bson:"type" json:"type,omitempty"`
	XP   int32  `bson:"xp" json:"xp,omitempty"`
}

type ReadCourse struct {
	ID          string       `bson:"id" json:"id"`
	Title       string       `bson:"title" json:"title"`
	Description string       `bson:"description" json:"description"`
	CreatorID   string       `bson:"creator_id" json:"creator_id"`
	Likes       int32        `bson:"likes" json:"likes"`
	Students    []string     `bson:"students" json:"students"`
	Topics      []string     `bson:"topics" json:"topics"`
	Modules     []ReadModule `bson:"modules" json:"modules"`
	UpdatedAt   string       `bson:"updated_at" json:"updated_at"`
	CreatedAt   string       `bson:"created_at" json:"created_at"`
}

// Module represents a module within a course
type ReadModule struct {
	ID    string     `bson:"id" json:"id"`
	Title string     `bson:"title" json:"title"`
	Tasks []ReadTask `bson:"tasks" json:"tasks"`
}

// Task represents a task within a module
type ReadTask struct {
	ID   string `bson:"id" json:"id"`
	Task string `bson:"task" json:"task"`
	Type string `bson:"type" json:"type"`
	XP   int32  `bson:"xp" json:"xp"`
}
