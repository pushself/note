##数据结构
[_zval_struct](#_zval_struct)
[_zend_value](#_zend_value)
[_zend_function](#_zend_function)
[_zend_execute_data](#_zend_execute_data)
[_zend_array](#_zend_array)

```
struct _zend_array {
	zend_refcounted_h gc;
	union {
		struct {
			ZEND_ENDIAN_LOHI_4(
				zend_uchar    flags,
				zend_uchar    nApplyCount,
				zend_uchar    nIteratorsCount,
				zend_uchar    consistency)
		} v;
		uint32_t flags;
	} u;
	uint32_t          nTableMask;
	Bucket           *arData;
	uint32_t          nNumUsed;
	uint32_t          nNumOfElements;
	uint32_t          nTableSize;
	uint32_t          nInternalPointer;
	zend_long         nNextFreeElement;
	dtor_func_t       pDestructor;
};
```

####_zval_struct
######1、PHP7中是使用zval结构存储变量信息的
```
struct _zval_struct {
	zend_value        value;			/* value */ //存储的是变量的值
	union {
		struct {
			ZEND_ENDIAN_LOHI_4(
				zend_uchar    type,			/* active type */
				zend_uchar    type_flags,
				zend_uchar    const_flags,
				zend_uchar    reserved)	    /* call info for EX(This) */
		} v;
		uint32_t type_info;
	} u1;
	union {
		uint32_t     next;                 /* hash collision chain */
		uint32_t     cache_slot;           /* literal cache slot */
		uint32_t     lineno;               /* line number (for ast nodes) */
		uint32_t     num_args;             /* arguments number for EX(This) */
		uint32_t     fe_pos;               /* foreach position */
		uint32_t     fe_iter_idx;          /* foreach iterator index */
		uint32_t     access_flags;         /* class constant access flags */
		uint32_t     property_guard;       /* single property guard */
		uint32_t     extra;                /* not further specified */
	} u2;
};
```
####_zend_value
```
typedef union _zend_value {
	zend_long         lval;				/* long value */
	double            dval;				/* double value */
	zend_refcounted  *counted;
	zend_string      *str;
	zend_array       *arr;
	zend_object      *obj;
	zend_resource    *res;
	zend_reference   *ref;
	zend_ast_ref     *ast;
	zval             *zv;
	void             *ptr;
	zend_class_entry *ce;
	zend_function    *func;
	struct {
		uint32_t w1;
		uint32_t w2;
	} ww;
} zend_value;
```
####_zend_function
```
union _zend_function {
	zend_uchar type;	/* MUST be the first element of this struct! */
	uint32_t   quick_arg_flags;

	struct {
		zend_uchar type;  /* never used */
		zend_uchar arg_flags[3]; /* bitset of arg_info.pass_by_reference */
		uint32_t fn_flags;
		zend_string *function_name;
		zend_class_entry *scope;
		union _zend_function *prototype;
		uint32_t num_args;
		uint32_t required_num_args;
		zend_arg_info *arg_info;
	} common;

	zend_op_array op_array;
	zend_internal_function internal_function;
};
```
####_zend_execute_data
```
struct _zend_execute_data {
	const zend_op       *opline;           /* executed opline                */
	zend_execute_data   *call;             /* current call                   */
	zval                *return_value;
	zend_function       *func;             /* executed function              */
	zval                 This;             /* this + call_info + num_args    */
	zend_execute_data   *prev_execute_data;
	zend_array          *symbol_table;
#if ZEND_EX_USE_RUN_TIME_CACHE
	void               **run_time_cache;   /* cache op_array->run_time_cache */
#endif
#if ZEND_EX_USE_LITERALS
	zval                *literals;         /* cache op_array->literals       */
#endif
};
```
####返回值
	RETURN_NULL()	返回null
	RETURN_LONG(l)	返回整型
	RETURN_DOUBLE(d) 返回浮点型
	RETURN_STR(s)	返回一个字符串。参数是一个zend_string * 指针
	RETURN_STRING(s)	返回一个字符串。参数是一个char * 指针
	RETURN_STRINGL(s, l) 返回一个字符串。第二个参数是字符串长度。
	RETURN_EMPTY_STRING()	返回一个空字符串。
	RETURN_ARR(r)	返回一个数组。参数是zend_array *指针。
	RETURN_OBJ(r) 返回一个对象。参数是zend_object *指针。
	RETURN_ZVAL(zv, copy, dtor) 返回任意类型。参数是 zval *指针。
	RETURN_FALSE	返回false
	RETURN_TRUE	返回true

####获取类型
	define Z_TYPE(zval)                zval_get_type(&(zval))
	define Z_TYPE_P(zval_p)            Z_TYPE(*(zval_p))
####设置变量类型
	define Z_TYPE_INFO(zval)       (zval).u1.type_info
	define Z_TYPE_INFO_P(zval_p)       Z_TYPE_INFO(*(zval_p))
####操作相关
	zend_hash_num_elements(Z_ARRVAL_P(zv)) //计算数组大小
	
	
####函数返回数组
```c
zval result;
zval var_value;
array_init(&var_value);//初始化数组
add_index_long(&var_value, 0, 1);//0：索引  1：值
add_assoc_stringl_ex(&var_value, "a", 1, "b", 1);

array_init(&result);
add_next_index_zval(&result, &var_value);
RETURN_ARR(result.value.arr);
```



##ZEND_API `Zend/zend_API.h`

####数组
####函数