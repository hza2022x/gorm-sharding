package sharding

// Config sharding配置
type Config struct {
	// DoubleWrite 为true则数据会双写到主表和分片表.
	DoubleWrite bool

	// ShardingKey 分片键
	ShardingKey string

	// InstanceShards ext：实例数
	InstanceShards uint

	// DbShards ext：总的db分片数
	DbShards uint

	// TableShards 每库分表数
	TableShards uint

	// TableFormat 后缀格式
	TableFormat string

	// ShardingAlgorithm sharding算法，如取模算法
	//	 取模算法示例:
	//		func(value interface{}) (suffix string, err error) {
	//			if uid, ok := value.(int64);ok {
	//				return fmt.Sprintf("_%02d", user_id % 64), nil
	//			}
	//			return "", errors.New("invalid user_id")
	//		}
	ShardingAlgorithm func(columnValue interface{}) (suffix string, err error)

	// ShardingSuffixs 用于生成表后缀的函数
	// 用于支持Migrator并生成PrimaryKey。例如，该函数获取 mod all 分片后缀。
	//
	// func () (suffixs []string) {
	// 	numberOfShards := 5
	// 	for i := 0; i < numberOfShards; i++ {
	// 		suffixs = append(suffixs, fmt.Sprintf("_%02d", i%numberOfShards))
	// 	}
	// }
	ShardingSuffixs func() (suffixs []string)

	// ShardingAlgorithmByPrimaryKey specifies a function to generate the sharding
	// table's suffix by the primary key. Used when no sharding key specified.
	// For example, this function use the Snowflake library to generate the suffix.
	//
	// 	func(id int64) (suffix string) {
	//		return fmt.Sprintf("_%02d", snowflake.ParseInt64(id).Node())
	//	}
	ShardingAlgorithmByPrimaryKey func(id int64) (suffix string)

	// PrimaryKeyGenerator specifies the primary key generate algorithm.
	// Used only when insert and the record does not contains an id field.
	// Options are PKSnowflake, PKPGSequence and PKCustom.
	// When use PKCustom, you should also specify PrimaryKeyGeneratorFn.
	PrimaryKeyGenerator int

	// PrimaryKeyGeneratorFn specifies a function to generate the primary key.
	// When use auto-increment like generator, the tableIdx argument could ignored.
	// For example, this function use the Snowflake library to generate the primary key.
	//
	// 	func(tableIdx int64) int64 {
	//		return nodes[tableIdx].Generate().Int64()
	//	}
	PrimaryKeyGeneratorFn func(tableIdx int64) int64
}
