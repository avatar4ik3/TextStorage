package models

type Repository struct {
	store *Store
}

func NewRepository(store *Store) *Repository {
	return &Repository{
		store: store,
	}
}

func (this *Repository) AddText(m *Text) (*Text, error) {
	return m, this.store.db.QueryRow("INSERT INTO texts (value) VALUES ($1) RETURNING id", m.Value).Scan(&m.Id)
}

func (this *Repository) RemoveText(id uint64) error {
	r, err := this.store.db.Query("Delete From texts where id = ($1)", id)
	for r.Next() != false {

	}
	return err
}

func (this *Repository) GetText(id uint64) (*Text, error) {
	m := &Text{Id: id}
	return m, this.store.db.QueryRow("Select 1 From texts where id = ($1)", id).Scan(&m.Id, &m.Value)
}

func (this *Repository) AllTexts() ([]Text, error) {
	m := []Text{}
	rows, err := this.store.db.Query("Select * From texts")
	for rows.Next() != false {
		model := Text{}
		rows.Scan(&model.Id, &model.Value)
		m = append(m, model)
	}
	return m, err
}

// func (this *Repository) AddGroup(m *Group) (*Group, error) {
// 	return m, this.store.db.QueryRow("INSERT INTO groups (name) VALUES ($1) RETURNING id", m.name).Scan(&m.id)
// }

// func (this *Repository) RemoveGroup(m *Group) error {
// 	return this.store.db.QueryRow("Delete From groups where id = ($1)", m.id).Scan(new(int))
// }

// func (this *Repository) GetGroup(m *Group) (*Group, error) {
// 	return m, this.store.db.QueryRow("Select * From groups where id = ($1)", m.id).Scan(&m)
// }

// func (this *Repository) AddRelation(m *Relation) (*Relation, error) {
// 	return m, this.store.db.QueryRow(
// 		"INSERT INTO relations (groupId,textId) VALUES ($1,$2) RETURNING id",
// 		m.groupId,
// 		m.textId,
// 	).Scan(&m.id)
// }

// func (this *Repository) RemoveRelation(m *Relation) error {
// 	return this.store.db.QueryRow("Delete From relations where id = ($1)", m.id).Scan(new(int))
// }
