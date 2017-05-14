package tableColumnsMdl

import (
	"log"
	"time"

	"bitbucket.org/restapi/db"
)

func GetTableColumn(schemaName string, tableName string) (tableColumns TableColumns, err error) {

	start := time.Now()

	rows, err := db.GetDB().Query("SELECT COLUMN_NAME,DATA_TYPE,CHARACTER_MAXIMUM_LENGTH,IS_NULLABLE,ORDINAL_POSITION,COLUMN_DEFAULT FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ?", schemaName, tableName)
	//rows, err := db.GetDB().Query("select cal_id,roster_id,doctor_id,doctor_name,calendar_from_time,calendar_to_time,site_id,calendar_date,calendar_time from calendar2 ")
	if err != nil {
		log.Println("users.go: All() err = ", err)
	}

	Response := TableColumns{}

	for rows.Next() {

		tableColumn := TableColumn{}
		rows.Scan(&tableColumn.COLUMNNAME, &tableColumn.DATATYPE, &tableColumn.CHARACTERMAXIMUMLENGTH, &tableColumn.ISNULLABLE, &tableColumn.ORDINALPOSITION, &tableColumn.COLUMNDEFAULT)
		//calendar.CalendarFromTimeInTime,err := time.Parse(layout, calendar.CalendarFromTime)
		Response.TableColumns = append(Response.TableColumns, tableColumn)
	}

	log.Printf("calendarMdl: sql with normal way duration = %s", time.Since(start))

	return Response, err
}

/*
rows, err := db.GetDB().Query("SELECT COLUMN_NAME,ORDINAL_POSITION,COLUMN_DEFAULT,IS_NULLABLE,DATA_TYPE,CHARACTER_MAXIMUM_LENGTH FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ?", schemaName, tableName)
//rows, err := db.GetDB().Query("select cal_id,roster_id,doctor_id,doctor_name,calendar_from_time,calendar_to_time,site_id,calendar_date,calendar_time from calendar2 ")
if err != nil {
	log.Println("users.go: All() err = ", err)
}

Response := TableColumns{}

for rows.Next() {

	tableColumn := TableColumn{}
	rows.Scan(&tableColumn.COLUMNNAME, &tableColumn.ORDINALPOSITION, &tableColumn.COLUMNDEFAULT, &tableColumn.ISNULLABLE, &tableColumn.DATATYPE, &tableColumn.CHARACTERMAXIMUMLENGTH)
	//calendar.CalendarFromTimeInTime,err := time.Parse(layout, calendar.CalendarFromTime)
	output, _ := json.Marshal(tableColumn)
	fmt.Println(string(output))
	Response.TableColumns = append(Response.TableColumns, tableColumn)
}
*/
