package models

import "database/sql"

type Repository struct {
	store *Store
}

func NewRepository(store *Store) *Repository {
	return &Repository{
		store: store,
	}
}

func (this *Repository) AddText(value string, description string) (*Text, error) {
	m := &Text{Value: value, Description: description}
	return m, this.store.db.QueryRow("INSERT INTO texts (value,description) VALUES ($1,$2) RETURNING id", value, description).Scan(&m.Id)
}

func (this *Repository) RemoveText(id uint64) error {
	r, err := this.store.db.Query("Delete From texts where id = ($1)", id)
	for r.Next() != false {

	}
	return err
}

func (this *Repository) GetText(id uint64) (*Text, error) {
	m := &Text{Id: id}
	return m, this.store.db.QueryRow("Select 1 From texts where id = ($1)", id).Scan(&m.Id, &m.Description, &m.Value)
}

func getAll[T any](s *Store, table string, apply func(row *sql.Rows) (T, error)) ([]T, error) {
	res := []T{}
	rows, err := s.db.Query("Select * From " + table)
	for rows.Next() != false {
		model, err2 := apply(rows)
		if err2 != nil {
			return nil, err2
		}
		res = append(res, model)
	}
	return res, err
}

func (this *Repository) AllTexts() ([]Text, error) {
	return getAll(this.store, "texts", func(row *sql.Rows) (Text, error) {
		m := Text{}
		return m, row.Scan(&m.Id, &m.Description, &m.Value)
	})
}

func (this *Repository) AllGroups() ([]Group, error) {
	return getAll(this.store, "texts", func(row *sql.Rows) (Group, error) {
		m := Group{}
		return m, row.Scan(&m.id, &m.description)
	})
}

func (this *Repository) AllRelations() ([]Relation, error) {
	return getAll(this.store, "texts", func(row *sql.Rows) (Relation, error) {
		m := Relation{}
		return m, row.Scan(&m.textId, &m.groupId)
	})
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
