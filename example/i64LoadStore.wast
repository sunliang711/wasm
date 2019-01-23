(module
 (memory $0 1)
 (export "memory" (memory $0))
 (export "_Z1fv" (func $_Z1fv))
 (func $_Z1fv (; 1 ;) (result i64)
  (i64.store offset=8 (i32.const 0) (i64.const 99))
  (i64.load offset=8 (i32.const 0))
  )
)
