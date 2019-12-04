package types

import (
	"database/sql"
	"encoding/json"
)

type NullString sql.NullString

func (s NullString) MarshalJSON() ([]byte, error) {
	if s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.String)
}

type EUsers struct {
	EUSERID        string     `json:"euserid"`
	KHID           string     `json:"khid"`
	TENTRUYCAP     string     `json:"tentruycap"`
	MATKHAU        string     `json:"matkhau"`
	KHOA           NullString `json:"khoa"`
	THOIGIANKHOA   NullString `json:"thoigiankhoa"`
	TRUYCAPGANNHAT NullString `json:"truycapgannhat"`
	DATELASTMAINT  NullString `json:"datelastmaint"`
	NOTE           NullString `json:"note"`
	KHID_FBE       NullString `json:"khid_fbe"`
}
