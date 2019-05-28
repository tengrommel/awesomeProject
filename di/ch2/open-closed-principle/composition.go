package open_closed_principle

import "database/sql"

type rowConvert struct {}

// populate the supplied Person from *sql.Row or *sql.Rows object
func (d *rowConvert)populate(in *Person, scan func(dest ...interface{}) error) error {
	return scan(in.Name, in.Email)
}

type LoadPerson struct {
	// compose the row converter into this loader
	rowConvert
}

func (loader *LoadPerson)ByID(id int) (Person, error) {
	row := loader.loadFromDB(id)
	person := Person{}
	err := loader.populate(&person, row.Scan)
	return person, err
}

func (loader *LoadPerson)loadFromDB(id int) *sql.Row {
	return nil
}

type LoadAll struct {
	rowConvert
}

func (loader *LoadAll)All() ([]Person, error) {
	rows := loader.loadAllFromDB()
	defer rows.Close()
	var output []Person
	for rows.Next() {
		person := Person{}
		// call the composed "abstract class"
		err := loader.populate(&person, rows.Scan)
		if err != nil {
			return nil, err
		}
	}
	return output, nil
}

func (loader *LoadAll)loadAllFromDB() *sql.Rows {
	return nil
}
