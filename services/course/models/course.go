package models

type Course struct {
	ID          string   `bson:"id"`
	Title       string   `bson:"title"`
	Description string   `bson:"description"`
	CreatorID   string   `bson:"creator_id"`
	Likes       int32    `bson:"likes"`
	Students    []string `bson:"students"`
	Topics      []string `bson:"topics"`
	Modules     []Module `bson:"modules"`
	UpdatedAt   string   `bson:"updated_at"`
	CreatedAt   string   `bson:"created_at"`
}

// Module represents a module within a course
type Module struct {
	ID    string `bson:"id"`
	Title string `bson:"title"`
	Tasks []Task `bson:"tasks"`
}

// Task represents a task within a module
type Task struct {
	ID   string `bson:"id"`
	Task string `bson:"task"`
	Type string `bson:"type"`
	XP   int32  `bson:"xp"`
}
