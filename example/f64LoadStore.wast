(module
 (memory $0 1)
 (export "memory" (memory $0))
 (export "_Z1fv" (func $_Z1fv))
 (func $_Z1fv (; 1 ;) (result f64)
  (f64.store offset=80 (i32.const 0) (f64.const 11.1))
  (f64.load offset=80 (i32.const 0))
  )
)
