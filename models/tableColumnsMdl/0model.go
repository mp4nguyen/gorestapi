package tableColumnsMdl

type TableColumn struct {
	COLUMNNAME             string `json:"columnName"`
	ORDINALPOSITION        int    `json:"ordinalPosition"`
	COLUMNDEFAULT          string `json:"columnDefault"`
	ISNULLABLE             string `json:"isNullable"`
	DATATYPE               string `json:"dataType"`
	CHARACTERMAXIMUMLENGTH int    `json:"chracterMaxiumLength"`
}

type TableColumns struct {
	TableColumns []TableColumn `json:"calendars"`
}
