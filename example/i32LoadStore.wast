(module
 (memory $0 1)
 (export "memory" (memory $0))
 (export "_Z1fv" (func $_Z1fv))
 (func $_Z1fv (; 1 ;) (result i32)
  (i32.store (i32.const 0) (i32.const 42))
  (i32.load (i32.const 0))
  )
)
