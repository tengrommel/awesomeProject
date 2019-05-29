package optimizing_for_humans

type WideFormatter interface {
	TOCSV(pets []Pet) ([]byte, error)
	TOGOB(pets []Pet) ([]byte, error)
	TOJSON(pets []Pet) ([]byte, error)
}