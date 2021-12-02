package access

var allEntities = []Access{
	{ID: 1, Role_ID: 5, Resource_ID: 3},
	{ID: 2, Role_ID: 6, Resource_ID: 3},
	{ID: 3, Role_ID: 3, Resource_ID: 5},
	{ID: 4, Role_ID: 2, Resource_ID: 1},
	{ID: 5, Role_ID: 2, Resource_ID: 3},
	{ID: 6, Role_ID: 7, Resource_ID: 2},
	{ID: 7, Role_ID: 8, Resource_ID: 25},
}

type Access struct {
	ID          uint64 `json:"id"`
	Role_ID     uint64 `json:"role_id"`
	Resource_ID uint64 `json:"resource_id"`
}
