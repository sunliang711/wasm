(module
 (type $0 (func (param i32 i32) (result i32)))
 (type $1 (func (result i32)))
 (type $2 (func (param i32) (result i32)))
 (type $3 (func))
 (import "env" "memory" (memory $0 256))
 (data (get_global $gimport$1) "max")
 (import "env" "memoryBase" (global $gimport$1 i32))
 (import "env" "_printf" (func $fimport$2 (param i32 i32) (result i32)))
 (global $global$0 (mut i32) (i32.const 0))
 (global $global$1 (mut i32) (i32.const 0))
 (global $global$2 i32 (i32.const 1))
 (global $global$3 i32 (i32.const 6))
 (global $global$4 i32 (i32.const 2))
 (global $global$5 i32 (i32.const 3))
 (global $global$6 i32 (i32.const 4))
 (global $global$7 i32 (i32.const 5))
 (export "__post_instantiate" (func $6))
 (export "_height" (func $0))
 (export "_max" (func $5))
 (export "_predefined1" (func $1))
 (export "_predefined2" (func $2))
 (export "_predefined3" (func $3))
 (export "_predefined4" (func $4))
 (export "fp$_height" (global $global$2))
 (export "fp$_max" (global $global$3))
 (export "fp$_predefined1" (global $global$4))
 (export "fp$_predefined2" (global $global$5))
 (export "fp$_predefined3" (global $global$6))
 (export "fp$_predefined4" (global $global$7))
 (func $0 (; 1 ;) (type $1) (result i32)
  (i32.const 1)
 )
 (func $1 (; 2 ;) (type $2) (param $0 i32) (result i32)
  (i32.const 2)
 )
 (func $2 (; 3 ;) (type $2) (param $0 i32) (result i32)
  (i32.const 3)
 )
 (func $3 (; 4 ;) (type $2) (param $0 i32) (result i32)
  (i32.const 4)
 )
 (func $4 (; 5 ;) (type $2) (param $0 i32) (result i32)
  (i32.const 5)
 )
 (func $5 (; 6 ;) (type $0) (param $0 i32) (param $1 i32) (result i32)
  (local $2 i32)
  (set_local $2
   (get_global $global$0)
  )
  (set_global $global$0
   (i32.add
    (get_global $global$0)
    (i32.const 16)
   )
  )
  (if
   (i32.le_s
    (get_local $0)
    (get_local $1)
   )
   (block
    (drop
     (call $fimport$2
      (get_global $gimport$1)
      (get_local $2)
     )
    )
    (set_local $0
     (get_local $1)
    )
   )
  )
  (set_global $global$0
   (get_local $2)
  )
  (get_local $0)
 )
 (func $6 (; 7 ;) (type $3)
  (set_global $global$0
   (i32.add
    (get_global $gimport$1)
    (i32.const 16)
   )
  )
  (set_global $global$1
   (i32.add
    (get_global $global$0)
    (i32.const 5242880)
   )
  )
 )
 ;; custom section "dylink", size 7
)

