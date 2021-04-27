package core

import (
	"fmt"
	"io/ioutil"
	"log"
	DB "ormt/library/database"
	"ormt/template"
	"os"
	"path/filepath"
	"strings"
)

type Result struct {
	Field   string
	Type    string
	Null    string
	Key     string
	Default string
	Extra   string
}

var Schema = map[string]string{
	"int":      "int",
	"integer":  "int",
	"smallint": "int",
	"tinyint":  "int",
	"bigint":   "int64",
	"bool":     "bool",
	"float":    "float",
	"double":   "float64",
	"numeric":  "float64",
	"longtext": "string",
	"varchar":  "string",
	"char":     "string",
	"text":     "string",
	"tinytext": "string",
	"TINYBLOB": "string",
	"date":     "time.Time",
	"datetime": "time.Time",
}

func Generate() {
	tableDescMap := getTablesStruct()
	export(tableDescMap)
}

func export(tableDescMap map[string][]Result) {
	tmd := make([]map[string]string, 0)
	for key, value := range tableDescMap {
		tmp := map[string]string{
			"{{DbName}}":       key,
			"{{DbNameStruct}}": CovetString(key),
			"{{SturctInfo}}":   makeInfo(value),
		}
		tmd = append(tmd, tmp)
	}
	for _, value := range tmd {
		file := template.ModelTpl
		for key, val := range value {
			file = strings.ReplaceAll(file, key, val)
		}
		if strings.Index(file, "time.Time") > -1 {
			file = strings.Replace(file, "{{Import}}", "import \"time\"", 1)
		}
		fileName := value["{{DbName}}"] + ".go"
		WriteFile("./models", fileName, file, 0755)
	}
}

//获取表结构
func getTablesStruct() map[string][]Result {
	db := DB.GetDB()
	tables := make([]string, 0)
	db.Raw("show tables").Scan(&tables)
	resMap := map[string][]Result{}
	var res []Result
	for _, value := range tables {
		db.Raw("DESCRIBE " + value).Scan(&res)
		resMap[value] = res
	}
	return resMap
}

// 字符首字母大写
func ToUpper(str string) string {
	var upperStr string
	vv := []rune(str) // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 { // 后文有介绍
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				fmt.Println("Not begins with lowercase letter,")
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}

func UcFirst(str string) string {
	return strings.ToUpper(str[:1]) + str[1:]
}

func LcFirst(str string) string {
	return strings.ToLower(str[:1]) + str[1:]
}

func makeInfo(info []Result) string {
	str := ""
	for _, value := range info {
		endInt := 0
		sqlType := value.Type
		str += "  " + UcFirst(value.Field)
		endInt = strings.Index(sqlType, "(")
		if endInt < 0 {
			endInt = strings.Index(sqlType, " ")
		}
		if endInt > 0 {
			sqlType = sqlType[0:endInt]
		}
		if Schema[sqlType] != "" {
			sqlType = Schema[sqlType]
		}
		str += "  " + sqlType
		str += "  " + "`json:\"" + value.Field + "\"`" + "\n"
	}
	return str
}

func CovetString(str string) string {
	arrStr := strings.Split(str, "_")
	strNew := ""
	for _, value := range arrStr {
		strNew += UcFirst(value)
	}
	return strNew
}

//写入文件
func WriteFile(fileDir string, fileName string, file string, mode os.FileMode) error {
	_, err := os.Stat(fileDir)
	if err != nil {
		err = os.Mkdir(fileDir, mode)
		if err != nil {
			log.Fatalln(err.Error() + ": " + fileDir)
		}
	}
	fn := filepath.Join(fileDir, fileName)
	err = ioutil.WriteFile(fn, []byte(file), mode)
	if err != nil {
		log.Fatalln(err.Error() + ": " + fn)
	}
	fmt.Println("success create :" + fn)
	return err
}
