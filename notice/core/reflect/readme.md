# 反射

变量、interface{}和reflect.Value是可以相互转换的

```
rVal := reflect.ValueOf(b)
iVal := rVal.Interface()
v, ok := iVal.(Struct_Type)
```