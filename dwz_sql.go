package main

import "log"

func queryDwzWithLongUrl(longUrl string) (*DwzStruct, error) {
	var dwz = new(DwzStruct)
	querySql := "SELECT ID, TINY_URL, LONG_URL FROM t_tiny_url WHERE LONG_URL = ?"
	if err := db.QueryRow(querySql, longUrl).Scan(&dwz.Id, &dwz.TinyUrl, &dwz.LongUrl); err != nil {
		return nil, err
	}
	return dwz, nil
}

func queryDwzWithTinyUrl(tinyUrl string) (*DwzStruct, error) {
	var dwz = new(DwzStruct)
	querySql := "SELECT ID, TINY_URL, LONG_URL FROM t_tiny_url WHERE TINY_URL = ?"
	if err := db.QueryRow(querySql, tinyUrl).Scan(&dwz.Id, &dwz.TinyUrl, &dwz.LongUrl); err != nil {
		return nil, err
	}
	return dwz, nil
}

func saveDwz(dwz *DwzStruct) {
	log.Printf("Save tiny url: %+v\n", dwz)
	insertSql, err := db.Prepare("INSERT INTO t_tiny_url (ID, TINY_URL, LONG_URL, CREATE_TIME) VALUES (?, ?, ?, NOW())")
	checkErr(err)
	_, err = insertSql.Exec(dwz.Id, dwz.TinyUrl, dwz.LongUrl)
	checkErr(err)
}
