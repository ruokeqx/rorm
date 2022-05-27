package dialect
// 兼容不同数据库 抽象出不同数据库差异的部分
import "reflect"

var dialectsMap = map[string]Dialect{}

// Dialect 接口 go语言类型转化为数据库类型 表是否存在的sql的语句
type Dialect interface {
	DataTypeOf(typ reflect.Value) string
	TableExistSQL(tableName string) (string, []interface{})
}

func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}

func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectsMap[name]
	return
}