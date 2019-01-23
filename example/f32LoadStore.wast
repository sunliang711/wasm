(module
 (memory $0 1)
 (export "memory" (memory $0))
 (export "_Z1fv" (func $_Z1fv))
 (func $_Z1fv (; 1 ;) (result f32)
  (f32.store offset=8 (i32.const 0) (f32.const 1.1))
  (f32.load offset=8 (i32.const 0))
  )
)
