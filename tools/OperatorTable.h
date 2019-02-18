#define NONE                                    WAVM::IR::FunctionType()
#define LOAD(resultTypeId)                      WAVM::IR::FunctionType({WAVM::IR::ValueType::resultTypeId}, {WAVM::IR::ValueType::i32                                                                                })
#define STORE(valueTypeId)                      WAVM::IR::FunctionType({},                                  {WAVM::IR::ValueType::i32,           WAVM::IR::ValueType::valueTypeId                                    })
#define NULLARY(resultTypeId)                   WAVM::IR::FunctionType({WAVM::IR::ValueType::resultTypeId}, {                                                                                                        })
#define BINARY(operandTypeId, resultTypeId)     WAVM::IR::FunctionType({WAVM::IR::ValueType::resultTypeId}, {WAVM::IR::ValueType::operandTypeId, WAVM::IR::ValueType::operandTypeId                                  })
#define UNARY(operandTypeId, resultTypeId)      WAVM::IR::FunctionType({WAVM::IR::ValueType::resultTypeId}, {WAVM::IR::ValueType::operandTypeId                                                                      })
#define VECTORSELECT(vectorTypeId)              WAVM::IR::FunctionType({WAVM::IR::ValueType::vectorTypeId}, {WAVM::IR::ValueType::vectorTypeId,  WAVM::IR::ValueType::vectorTypeId, WAVM::IR::ValueType::vectorTypeId})
#define REPLACELANE(scalarTypeId, vectorTypeId) WAVM::IR::FunctionType({WAVM::IR::ValueType::vectorTypeId}, {WAVM::IR::ValueType::vectorTypeId,  WAVM::IR::ValueType::scalarTypeId                                   })
#define COMPAREEXCHANGE(valueTypeId)            WAVM::IR::FunctionType({WAVM::IR::ValueType::valueTypeId},  {WAVM::IR::ValueType::i32,           WAVM::IR::ValueType::valueTypeId,  WAVM::IR::ValueType::valueTypeId })
#define WAIT(valueTypeId)                       WAVM::IR::FunctionType({WAVM::IR::ValueType::i32},          {WAVM::IR::ValueType::i32,           WAVM::IR::ValueType::valueTypeId,  WAVM::IR::ValueType::f64         })
#define ATOMICRMW(valueTypeId)                  WAVM::IR::FunctionType({WAVM::IR::ValueType::valueTypeId},  {WAVM::IR::ValueType::i32,           WAVM::IR::ValueType::valueTypeId                                    })
#define BULKCOPY                                WAVM::IR::FunctionType({},                                  {WAVM::IR::ValueType::i32,           WAVM::IR::ValueType::i32,          WAVM::IR::ValueType::i32         })


#define ENUM_CONTROL_OPERATORS(visitOp)                                                                                                                     \
	visitOp(0x0002, OPCblock              , "block"                            , ControlStructureImm       ,1    , PARAMETRIC           , mvp                    )   \
	visitOp(0x0003, OPCloop               , "loop"                             , ControlStructureImm       ,1    , PARAMETRIC           , mvp                    )   \
	visitOp(0x0004, OPCif_                , "if"                               , ControlStructureImm       ,1    , PARAMETRIC           , mvp                    )   \
	visitOp(0x0005, OPCelse_              , "else"                             , NoImm                     ,1    , PARAMETRIC           , mvp                    )   \
	visitOp(0x000b, OPCend                , "end"                              , NoImm                     ,1    , PARAMETRIC           , mvp                    )   \
	visitOp(0xfb02, OPCtry_               , "try"                              , ControlStructureImm       ,1    , PARAMETRIC           , exceptionHandling      )   \
	visitOp(0xfb03, OPCcatch_             , "catch"                            , ExceptionTypeImm          ,1    , PARAMETRIC           , exceptionHandling      )   \
	visitOp(0xfb04, OPCcatch_all          , "catch_all"                        , NoImm                     ,1    , PARAMETRIC           , exceptionHandling      )

#define ENUM_PARAMETRIC_OPERATORS(visitOp)                                                                                                                  \
/* Control flow                                                                                                                                          */ \
	visitOp(0x0000, OPCunreachable        , "unreachable"                      , NoImm                     ,1    , PARAMETRIC           , mvp                    )   \
	visitOp(0x000c, OPCbr                 , "br"                               , BranchImm                 ,1    , PARAMETRIC           , mvp                    )   \
	visitOp(0x000d, OPCbr_if              , "br_if"                            , BranchImm                 ,1    , PARAMETRIC           , mvp                    )   \
	visitOp(0x000e, OPCbr_table           , "br_table"                         , BranchTableImm            ,1    , PARAMETRIC           , mvp                    )   \
	visitOp(0x000f, OPCreturn_            , "return"                           , NoImm                     ,1    , PARAMETRIC           , mvp                    )   \
	visitOp(0x0010, OPCcall               , "call"                             , FunctionImm               ,1    , PARAMETRIC           , mvp                    )   \
	visitOp(0x0011, OPCcall_indirect      , "call_indirect"                    , CallIndirectImm           ,1    , PARAMETRIC           , mvp                    )   \
/* Stack manipulation                                                                                                                                    */ \
	visitOp(0x001a, OPCdrop               , "drop"                             , NoImm                     ,1    , PARAMETRIC           , mvp                    )   \
	visitOp(0x001b, OPCselect             , "select"                           , NoImm                     ,1    , PARAMETRIC           , mvp                    )   \
/* Variables                                                                                                                                             */ \
	visitOp(0x0020, OPCget_local          , "get_local"                        , GetOrSetVariableImm       ,1    , PARAMETRIC           , mvp                    )   \
	visitOp(0x0021, OPCset_local          , "set_local"                        , GetOrSetVariableImm       ,1    , PARAMETRIC           , mvp                    )   \
	visitOp(0x0022, OPCtee_local          , "tee_local"                        , GetOrSetVariableImm       ,1    , PARAMETRIC           , mvp                    )   \
	visitOp(0x0023, OPCget_global         , "get_global"                       , GetOrSetVariableImm       ,1    , PARAMETRIC           , mvp                    )   \
	visitOp(0x0024, OPCset_global         , "set_global"                       , GetOrSetVariableImm       ,1    , PARAMETRIC           , mvp                    )   \
/* Table access                                                                                                                                          */ \
	visitOp(0x0025, OPCtable_get          , "table.get"                        , TableImm                  ,1    , PARAMETRIC           , referenceTypes         )   \
	visitOp(0x0026, OPCtable_set          , "table.set"                        , TableImm                  ,1    , PARAMETRIC           , referenceTypes         )   \
/* Exceptions                                                                                                                                            */ \
	visitOp(0xfb00, OPCthrow_             , "throw"                            , ExceptionTypeImm          ,1    , PARAMETRIC           , exceptionHandling      )   \
	visitOp(0xfb01, OPCrethrow            , "rethrow"                          , RethrowImm                ,1    , PARAMETRIC           , exceptionHandling      )

#define ENUM_NONCONTROL_NONPARAMETRIC_OPERATORS(visitOp)                                                                                                    \
	visitOp(0x0001, OPCnop                , "nop"                              , NoImm                     ,1    , NONE                 , mvp                    )   \
/* Memory                                                                                                                                                */ \
	visitOp(0x0028, OPCi32_load           , "i32.load"                         , LoadOrStoreImm            ,1    , LOAD(i32)            , mvp                    )   \
	visitOp(0x0029, OPCi64_load           , "i64.load"                         , LoadOrStoreImm            ,1    , LOAD(i64)            , mvp                    )   \
	visitOp(0x002a, OPCf32_load           , "f32.load"                         , LoadOrStoreImm            ,1    , LOAD(f32)            , mvp                    )   \
	visitOp(0x002b, OPCf64_load           , "f64.load"                         , LoadOrStoreImm            ,1    , LOAD(f64)            , mvp                    )   \
	visitOp(0x002c, OPCi32_load8_s        , "i32.load8_s"                      , LoadOrStoreImm            ,1    , LOAD(i32)            , mvp                    )   \
	visitOp(0x002d, OPCi32_load8_u        , "i32.load8_u"                      , LoadOrStoreImm            ,1    , LOAD(i32)            , mvp                    )   \
	visitOp(0x002e, OPCi32_load16_s       , "i32.load16_s"                     , LoadOrStoreImm            ,1    , LOAD(i32)            , mvp                    )   \
	visitOp(0x002f, OPCi32_load16_u       , "i32.load16_u"                     , LoadOrStoreImm            ,1    , LOAD(i32)            , mvp                    )   \
	visitOp(0x0030, OPCi64_load8_s        , "i64.load8_s"                      , LoadOrStoreImm            ,1    , LOAD(i64)            , mvp                    )   \
	visitOp(0x0031, OPCi64_load8_u        , "i64.load8_u"                      , LoadOrStoreImm            ,1    , LOAD(i64)            , mvp                    )   \
	visitOp(0x0032, OPCi64_load16_s       , "i64.load16_s"                     , LoadOrStoreImm            ,1    , LOAD(i64)            , mvp                    )   \
	visitOp(0x0033, OPCi64_load16_u       , "i64.load16_u"                     , LoadOrStoreImm            ,1    , LOAD(i64)            , mvp                    )   \
	visitOp(0x0034, OPCi64_load32_s       , "i64.load32_s"                     , LoadOrStoreImm            ,1    , LOAD(i64)            , mvp                    )   \
	visitOp(0x0035, OPCi64_load32_u       , "i64.load32_u"                     , LoadOrStoreImm            ,1    , LOAD(i64)            , mvp                    )   \
	visitOp(0x0036, OPCi32_store          , "i32.store"                        , LoadOrStoreImm            ,1    , STORE(i32)           , mvp                    )   \
	visitOp(0x0037, OPCi64_store          , "i64.store"                        , LoadOrStoreImm            ,1    , STORE(i64)           , mvp                    )   \
	visitOp(0x0038, OPCf32_store          , "f32.store"                        , LoadOrStoreImm            ,1    , STORE(f32)           , mvp                    )   \
	visitOp(0x0039, OPCf64_store          , "f64.store"                        , LoadOrStoreImm            ,1    , STORE(f64)           , mvp                    )   \
	visitOp(0x003a, OPCi32_store8         , "i32.store8"                       , LoadOrStoreImm            ,1    , STORE(i32)           , mvp                    )   \
	visitOp(0x003b, OPCi32_store16        , "i32.store16"                      , LoadOrStoreImm            ,1    , STORE(i32)           , mvp                    )   \
	visitOp(0x003c, OPCi64_store8         , "i64.store8"                       , LoadOrStoreImm            ,1    , STORE(i64)           , mvp                    )   \
	visitOp(0x003d, OPCi64_store16        , "i64.store16"                      , LoadOrStoreImm            ,1    , STORE(i64)           , mvp                    )   \
	visitOp(0x003e, OPCi64_store32        , "i64.store32"                      , LoadOrStoreImm            ,1    , STORE(i64)           , mvp                    )   \
	visitOp(0x003f, OPCmemory_size        , "memory.size"                      , MemoryImm                 ,1    , NULLARY(i32)         , mvp                    )   \
	visitOp(0x0040, OPCmemory_grow        , "memory.grow"                      , MemoryImm                 ,1    , UNARY(i32,i32)       , mvp                    )   \
/* Literals                                                                                                                                              */ \
	visitOp(0x0041, OPCi32_const          , "i32.const"                        , LiteralImm_I32           ,1    , NULLARY(i32)         , mvp                    )   \
	visitOp(0x0042, OPCi64_const          , "i64.const"                        , LiteralImm_I64           ,1    , NULLARY(i64)         , mvp                    )   \
	visitOp(0x0043, OPCf32_const          , "f32.const"                        , LiteralImm_F32           ,1    , NULLARY(f32)         , mvp                    )   \
	visitOp(0x0044, OPCf64_const          , "f64.const"                        , LiteralImm_F64           ,1    , NULLARY(f64)         , mvp                    )   \
/* Comparisons                                                                                                                                           */ \
	visitOp(0x0045, OPCi32_eqz            , "i32.eqz"                          , NoImm                     ,1    , UNARY(i32,i32)       , mvp                    )   \
	visitOp(0x0046, OPCi32_eq             , "i32.eq"                           , NoImm                     ,1    , BINARY(i32,i32)      , mvp                    )   \
	visitOp(0x0047, OPCi32_ne             , "i32.ne"                           , NoImm                     ,1    , BINARY(i32,i32)      , mvp                    )   \
	visitOp(0x0048, OPCi32_lt_s           , "i32.lt_s"                         , NoImm                     ,1    , BINARY(i32,i32)      , mvp                    )   \
	visitOp(0x0049, OPCi32_lt_u           , "i32.lt_u"                         , NoImm                     ,1    , BINARY(i32,i32)      , mvp                    )   \
	visitOp(0x004a, OPCi32_gt_s           , "i32.gt_s"                         , NoImm                     ,1    , BINARY(i32,i32)      , mvp                    )   \
	visitOp(0x004b, OPCi32_gt_u           , "i32.gt_u"                         , NoImm                     ,1    , BINARY(i32,i32)      , mvp                    )   \
	visitOp(0x004c, OPCi32_le_s           , "i32.le_s"                         , NoImm                     ,1    , BINARY(i32,i32)      , mvp                    )   \
	visitOp(0x004d, OPCi32_le_u           , "i32.le_u"                         , NoImm                     ,1    , BINARY(i32,i32)      , mvp                    )   \
	visitOp(0x004e, OPCi32_ge_s           , "i32.ge_s"                         , NoImm                     ,1    , BINARY(i32,i32)      , mvp                    )   \
	visitOp(0x004f, OPCi32_ge_u           , "i32.ge_u"                         , NoImm                     ,1    , BINARY(i32,i32)      , mvp                    )   \
	visitOp(0x0050, OPCi64_eqz            , "i64.eqz"                          , NoImm                     ,1    , UNARY(i64,i32)       , mvp                    )   \
	visitOp(0x0051, OPCi64_eq             , "i64.eq"                           , NoImm                     ,1    , BINARY(i64,i32)      , mvp                    )   \
	visitOp(0x0052, OPCi64_ne             , "i64.ne"                           , NoImm                     ,1    , BINARY(i64,i32)      , mvp                    )   \
	visitOp(0x0053, OPCi64_lt_s           , "i64.lt_s"                         , NoImm                     ,1    , BINARY(i64,i32)      , mvp                    )   \
	visitOp(0x0054, OPCi64_lt_u           , "i64.lt_u"                         , NoImm                     ,1    , BINARY(i64,i32)      , mvp                    )   \
	visitOp(0x0055, OPCi64_gt_s           , "i64.gt_s"                         , NoImm                     ,1    , BINARY(i64,i32)      , mvp                    )   \
	visitOp(0x0056, OPCi64_gt_u           , "i64.gt_u"                         , NoImm                     ,1    , BINARY(i64,i32)      , mvp                    )   \
	visitOp(0x0057, OPCi64_le_s           , "i64.le_s"                         , NoImm                     ,1    , BINARY(i64,i32)      , mvp                    )   \
	visitOp(0x0058, OPCi64_le_u           , "i64.le_u"                         , NoImm                     ,1    , BINARY(i64,i32)      , mvp                    )   \
	visitOp(0x0059, OPCi64_ge_s           , "i64.ge_s"                         , NoImm                     ,1    , BINARY(i64,i32)      , mvp                    )   \
	visitOp(0x005a, OPCi64_ge_u           , "i64.ge_u"                         , NoImm                     ,1    , BINARY(i64,i32)      , mvp                    )   \
	visitOp(0x005b, OPCf32_eq             , "f32.eq"                           , NoImm                     ,1    , BINARY(f32,i32)      , mvp                    )   \
	visitOp(0x005c, OPCf32_ne             , "f32.ne"                           , NoImm                     ,1    , BINARY(f32,i32)      , mvp                    )   \
	visitOp(0x005d, OPCf32_lt             , "f32.lt"                           , NoImm                     ,1    , BINARY(f32,i32)      , mvp                    )   \
	visitOp(0x005e, OPCf32_gt             , "f32.gt"                           , NoImm                     ,1    , BINARY(f32,i32)      , mvp                    )   \
	visitOp(0x005f, OPCf32_le             , "f32.le"                           , NoImm                     ,1    , BINARY(f32,i32)      , mvp                    )   \
	visitOp(0x0060, OPCf32_ge             , "f32.ge"                           , NoImm                     ,1    , BINARY(f32,i32)      , mvp                    )   \
	visitOp(0x0061, OPCf64_eq             , "f64.eq"                           , NoImm                     ,1    , BINARY(f64,i32)      , mvp                    )   \
	visitOp(0x0062, OPCf64_ne             , "f64.ne"                           , NoImm                     ,1    , BINARY(f64,i32)      , mvp                    )   \
	visitOp(0x0063, OPCf64_lt             , "f64.lt"                           , NoImm                     ,1    , BINARY(f64,i32)      , mvp                    )   \
	visitOp(0x0064, OPCf64_gt             , "f64.gt"                           , NoImm                     ,1    , BINARY(f64,i32)      , mvp                    )   \
	visitOp(0x0065, OPCf64_le             , "f64.le"                           , NoImm                     ,1    , BINARY(f64,i32)      , mvp                    )   \
	visitOp(0x0066, OPCf64_ge             , "f64.ge"                           , NoImm                     ,1    , BINARY(f64,i32)      , mvp                    )   \
/* i32 arithmetic                                                                                                                                        */ \
	visitOp(0x0067, OPCi32_clz            , "i32.clz"                          , NoImm                     ,1    , UNARY(i32,i32)       , mvp                    )   \
	visitOp(0x0068, OPCi32_ctz            , "i32.ctz"                          , NoImm                     ,1    , UNARY(i32,i32)       , mvp                    )   \
	visitOp(0x0069, OPCi32_popcnt         , "i32.popcnt"                       , NoImm                     ,1    , UNARY(i32,i32)       , mvp                    )   \
	visitOp(0x006a, OPCi32_add            , "i32.add"                          , NoImm                     ,1    , BINARY(i32,i32)      , mvp                    )   \
	visitOp(0x006b, OPCi32_sub            , "i32.sub"                          , NoImm                     ,1    , BINARY(i32,i32)      , mvp                    )   \
	visitOp(0x006c, OPCi32_mul            , "i32.mul"                          , NoImm                     ,1    , BINARY(i32,i32)      , mvp                    )   \
	visitOp(0x006d, OPCi32_div_s          , "i32.div_s"                        , NoImm                     ,1    , BINARY(i32,i32)      , mvp                    )   \
	visitOp(0x006e, OPCi32_div_u          , "i32.div_u"                        , NoImm                     ,1    , BINARY(i32,i32)      , mvp                    )   \
	visitOp(0x006f, OPCi32_rem_s          , "i32.rem_s"                        , NoImm                     ,1    , BINARY(i32,i32)      , mvp                    )   \
	visitOp(0x0070, OPCi32_rem_u          , "i32.rem_u"                        , NoImm                     ,1    , BINARY(i32,i32)      , mvp                    )   \
	visitOp(0x0071, OPCi32_and_           , "i32.and"                          , NoImm                     ,1    , BINARY(i32,i32)      , mvp                    )   \
	visitOp(0x0072, OPCi32_or_            , "i32.or"                           , NoImm                     ,1    , BINARY(i32,i32)      , mvp                    )   \
	visitOp(0x0073, OPCi32_xor_           , "i32.xor"                          , NoImm                     ,1    , BINARY(i32,i32)      , mvp                    )   \
	visitOp(0x0074, OPCi32_shl            , "i32.shl"                          , NoImm                     ,1    , BINARY(i32,i32)      , mvp                    )   \
	visitOp(0x0075, OPCi32_shr_s          , "i32.shr_s"                        , NoImm                     ,1    , BINARY(i32,i32)      , mvp                    )   \
	visitOp(0x0076, OPCi32_shr_u          , "i32.shr_u"                        , NoImm                     ,1    , BINARY(i32,i32)      , mvp                    )   \
	visitOp(0x0077, OPCi32_rotl           , "i32.rotl"                         , NoImm                     ,1    , BINARY(i32,i32)      , mvp                    )   \
	visitOp(0x0078, OPCi32_rotr           , "i32.rotr"                         , NoImm                     ,1    , BINARY(i32,i32)      , mvp                    )   \
/* i64 arithmetic                                                                                                                                        */ \
	visitOp(0x0079, OPCi64_clz            , "i64.clz"                          , NoImm                     ,1    , UNARY(i64,i64)       , mvp                    )   \
	visitOp(0x007a, OPCi64_ctz            , "i64.ctz"                          , NoImm                     ,1    , UNARY(i64,i64)       , mvp                    )   \
	visitOp(0x007b, OPCi64_popcnt         , "i64.popcnt"                       , NoImm                     ,1    , UNARY(i64,i64)       , mvp                    )   \
	visitOp(0x007c, OPCi64_add            , "i64.add"                          , NoImm                     ,1    , BINARY(i64,i64)      , mvp                    )   \
	visitOp(0x007d, OPCi64_sub            , "i64.sub"                          , NoImm                     ,1    , BINARY(i64,i64)      , mvp                    )   \
	visitOp(0x007e, OPCi64_mul            , "i64.mul"                          , NoImm                     ,1    , BINARY(i64,i64)      , mvp                    )   \
	visitOp(0x007f, OPCi64_div_s          , "i64.div_s"                        , NoImm                     ,1    , BINARY(i64,i64)      , mvp                    )   \
	visitOp(0x0080, OPCi64_div_u          , "i64.div_u"                        , NoImm                     ,1    , BINARY(i64,i64)      , mvp                    )   \
	visitOp(0x0081, OPCi64_rem_s          , "i64.rem_s"                        , NoImm                     ,1    , BINARY(i64,i64)      , mvp                    )   \
	visitOp(0x0082, OPCi64_rem_u          , "i64.rem_u"                        , NoImm                     ,1    , BINARY(i64,i64)      , mvp                    )   \
	visitOp(0x0083, OPCi64_and_           , "i64.and"                          , NoImm                     ,1    , BINARY(i64,i64)      , mvp                    )   \
	visitOp(0x0084, OPCi64_or_            , "i64.or"                           , NoImm                     ,1    , BINARY(i64,i64)      , mvp                    )   \
	visitOp(0x0085, OPCi64_xor_           , "i64.xor"                          , NoImm                     ,1    , BINARY(i64,i64)      , mvp                    )   \
	visitOp(0x0086, OPCi64_shl            , "i64.shl"                          , NoImm                     ,1    , BINARY(i64,i64)      , mvp                    )   \
	visitOp(0x0087, OPCi64_shr_s          , "i64.shr_s"                        , NoImm                     ,1    , BINARY(i64,i64)      , mvp                    )   \
	visitOp(0x0088, OPCi64_shr_u          , "i64.shr_u"                        , NoImm                     ,1    , BINARY(i64,i64)      , mvp                    )   \
	visitOp(0x0089, OPCi64_rotl           , "i64.rotl"                         , NoImm                     ,1    , BINARY(i64,i64)      , mvp                    )   \
	visitOp(0x008a, OPCi64_rotr           , "i64.rotr"                         , NoImm                     ,1    , BINARY(i64,i64)      , mvp                    )   \
/* f32 arithmetic                                                                                                                                        */ \
	visitOp(0x008b, OPCf32_abs            , "f32.abs"                          , NoImm                     ,1    , UNARY(f32,f32)       , mvp                    )   \
	visitOp(0x008c, OPCf32_neg            , "f32.neg"                          , NoImm                     ,1    , UNARY(f32,f32)       , mvp                    )   \
	visitOp(0x008d, OPCf32_ceil           , "f32.ceil"                         , NoImm                     ,1    , UNARY(f32,f32)       , mvp                    )   \
	visitOp(0x008e, OPCf32_floor          , "f32.floor"                        , NoImm                     ,1    , UNARY(f32,f32)       , mvp                    )   \
	visitOp(0x008f, OPCf32_trunc          , "f32.trunc"                        , NoImm                     ,1    , UNARY(f32,f32)       , mvp                    )   \
	visitOp(0x0090, OPCf32_nearest        , "f32.nearest"                      , NoImm                     ,1    , UNARY(f32,f32)       , mvp                    )   \
	visitOp(0x0091, OPCf32_sqrt           , "f32.sqrt"                         , NoImm                     ,1    , UNARY(f32,f32)       , mvp                    )   \
	visitOp(0x0092, OPCf32_add            , "f32.add"                          , NoImm                     ,1    , BINARY(f32,f32)      , mvp                    )   \
	visitOp(0x0093, OPCf32_sub            , "f32.sub"                          , NoImm                     ,1    , BINARY(f32,f32)      , mvp                    )   \
	visitOp(0x0094, OPCf32_mul            , "f32.mul"                          , NoImm                     ,1    , BINARY(f32,f32)      , mvp                    )   \
	visitOp(0x0095, OPCf32_div            , "f32.div"                          , NoImm                     ,1    , BINARY(f32,f32)      , mvp                    )   \
	visitOp(0x0096, OPCf32_min            , "f32.min"                          , NoImm                     ,1    , BINARY(f32,f32)      , mvp                    )   \
	visitOp(0x0097, OPCf32_max            , "f32.max"                          , NoImm                     ,1    , BINARY(f32,f32)      , mvp                    )   \
	visitOp(0x0098, OPCf32_copysign       , "f32.copysign"                     , NoImm                     ,1    , BINARY(f32,f32)      , mvp                    )   \
/* f64 arithmetic                                                                                                                                        */ \
	visitOp(0x0099, OPCf64_abs            , "f64.abs"                          , NoImm                     ,1    , UNARY(f64,f64)       , mvp                    )   \
	visitOp(0x009a, OPCf64_neg            , "f64.neg"                          , NoImm                     ,1    , UNARY(f64,f64)       , mvp                    )   \
	visitOp(0x009b, OPCf64_ceil           , "f64.ceil"                         , NoImm                     ,1    , UNARY(f64,f64)       , mvp                    )   \
	visitOp(0x009c, OPCf64_floor          , "f64.floor"                        , NoImm                     ,1    , UNARY(f64,f64)       , mvp                    )   \
	visitOp(0x009d, OPCf64_trunc          , "f64.trunc"                        , NoImm                     ,1    , UNARY(f64,f64)       , mvp                    )   \
	visitOp(0x009e, OPCf64_nearest        , "f64.nearest"                      , NoImm                     ,1    , UNARY(f64,f64)       , mvp                    )   \
	visitOp(0x009f, OPCf64_sqrt           , "f64.sqrt"                         , NoImm                     ,1    , UNARY(f64,f64)       , mvp                    )   \
	visitOp(0x00a0, OPCf64_add            , "f64.add"                          , NoImm                     ,1    , BINARY(f64,f64)      , mvp                    )   \
	visitOp(0x00a1, OPCf64_sub            , "f64.sub"                          , NoImm                     ,1    , BINARY(f64,f64)      , mvp                    )   \
	visitOp(0x00a2, OPCf64_mul            , "f64.mul"                          , NoImm                     ,1    , BINARY(f64,f64)      , mvp                    )   \
	visitOp(0x00a3, OPCf64_div            , "f64.div"                          , NoImm                     ,1    , BINARY(f64,f64)      , mvp                    )   \
	visitOp(0x00a4, OPCf64_min            , "f64.min"                          , NoImm                     ,1    , BINARY(f64,f64)      , mvp                    )   \
	visitOp(0x00a5, OPCf64_max            , "f64.max"                          , NoImm                     ,1    , BINARY(f64,f64)      , mvp                    )   \
	visitOp(0x00a6, OPCf64_copysign       , "f64.copysign"                     , NoImm                     ,1    , BINARY(f64,f64)      , mvp                    )   \
/* Conversions                                                                                                                                           */ \
	visitOp(0x00a7, OPCi32_wrap_i64       , "i32.wrap/i64"                     , NoImm                     ,1    , UNARY(i64,i32)       , mvp                    )   \
	visitOp(0x00a8, OPCi32_trunc_s_f32    , "i32.trunc_s/f32"                  , NoImm                     ,1    , UNARY(f32,i32)       , mvp                    )   \
	visitOp(0x00a9, OPCi32_trunc_u_f32    , "i32.trunc_u/f32"                  , NoImm                     ,1    , UNARY(f32,i32)       , mvp                    )   \
	visitOp(0x00aa, OPCi32_trunc_s_f64    , "i32.trunc_s/f64"                  , NoImm                     ,1    , UNARY(f64,i32)       , mvp                    )   \
	visitOp(0x00ab, OPCi32_trunc_u_f64    , "i32.trunc_u/f64"                  , NoImm                     ,1    , UNARY(f64,i32)       , mvp                    )   \
	visitOp(0x00ac, OPCi64_extend_s_i32   , "i64.extend_s/i32"                 , NoImm                     ,1    , UNARY(i32,i64)       , mvp                    )   \
	visitOp(0x00ad, OPCi64_extend_u_i32   , "i64.extend_u/i32"                 , NoImm                     ,1    , UNARY(i32,i64)       , mvp                    )   \
	visitOp(0x00ae, OPCi64_trunc_s_f32    , "i64.trunc_s/f32"                  , NoImm                     ,1    , UNARY(f32,i64)       , mvp                    )   \
	visitOp(0x00af, OPCi64_trunc_u_f32    , "i64.trunc_u/f32"                  , NoImm                     ,1    , UNARY(f32,i64)       , mvp                    )   \
	visitOp(0x00b0, OPCi64_trunc_s_f64    , "i64.trunc_s/f64"                  , NoImm                     ,1    , UNARY(f64,i64)       , mvp                    )   \
	visitOp(0x00b1, OPCi64_trunc_u_f64    , "i64.trunc_u/f64"                  , NoImm                     ,1    , UNARY(f64,i64)       , mvp                    )   \
	visitOp(0x00b2, OPCf32_convert_s_i32  , "f32.convert_s/i32"                , NoImm                     ,1    , UNARY(i32,f32)       , mvp                    )   \
	visitOp(0x00b3, OPCf32_convert_u_i32  , "f32.convert_u/i32"                , NoImm                     ,1    , UNARY(i32,f32)       , mvp                    )   \
	visitOp(0x00b4, OPCf32_convert_s_i64  , "f32.convert_s/i64"                , NoImm                     ,1    , UNARY(i64,f32)       , mvp                    )   \
	visitOp(0x00b5, OPCf32_convert_u_i64  , "f32.convert_u/i64"                , NoImm                     ,1    , UNARY(i64,f32)       , mvp                    )   \
	visitOp(0x00b6, OPCf32_demote_f64     , "f32.demote/f64"                   , NoImm                     ,1    , UNARY(f64,f32)       , mvp                    )   \
	visitOp(0x00b7, OPCf64_convert_s_i32  , "f64.convert_s/i32"                , NoImm                     ,1    , UNARY(i32,f64)       , mvp                    )   \
	visitOp(0x00b8, OPCf64_convert_u_i32  , "f64.convert_u/i32"                , NoImm                     ,1    , UNARY(i32,f64)       , mvp                    )   \
	visitOp(0x00b9, OPCf64_convert_s_i64  , "f64.convert_s/i64"                , NoImm                     ,1    , UNARY(i64,f64)       , mvp                    )   \
	visitOp(0x00ba, OPCf64_convert_u_i64  , "f64.convert_u/i64"                , NoImm                     ,1    , UNARY(i64,f64)       , mvp                    )   \
	visitOp(0x00bb, OPCf64_promote_f32    , "f64.promote/f32"                  , NoImm                     ,1    , UNARY(f32,f64)       , mvp                    )   \
	visitOp(0x00bc, OPCi32_reinterpret_f32, "i32.reinterpret/f32"              , NoImm                     ,1    , UNARY(f32,i32)       , mvp                    )   \
	visitOp(0x00bd, OPCi64_reinterpret_f64, "i64.reinterpret/f64"              , NoImm                     ,1    , UNARY(f64,i64)       , mvp                    )   \
	visitOp(0x00be, OPCf32_reinterpret_i32, "f32.reinterpret/i32"              , NoImm                     ,1    , UNARY(i32,f32)       , mvp                    )   \
	visitOp(0x00bf, OPCf64_reinterpret_i64, "f64.reinterpret/i64"              , NoImm                     ,1    , UNARY(i64,f64)       , mvp                    )   \
/* 8- and 16-bit sign extension operators                                                                                                                */ \
	visitOp(0x00c0, OPCi32_extend8_s             , "i32.extend8_s"             , NoImm                     ,1    , UNARY(i32,i32)       , extendedSignExtension  )   \
	visitOp(0x00c1, OPCi32_extend16_s            , "i32.extend16_s"            , NoImm                     ,1    , UNARY(i32,i32)       , extendedSignExtension  )   \
	visitOp(0x00c2, OPCi64_extend8_s             , "i64.extend8_s"             , NoImm                     ,1    , UNARY(i64,i64)       , extendedSignExtension  )   \
	visitOp(0x00c3, OPCi64_extend16_s            , "i64.extend16_s"            , NoImm                     ,1    , UNARY(i64,i64)       , extendedSignExtension  )   \
	visitOp(0x00c4, OPCi64_extend32_s            , "i64.extend32_s"            , NoImm                     ,1    , UNARY(i64,i64)       , extendedSignExtension  )   \
/* Reference type operators                                                                                                                              */ \
	visitOp(0x00d0, OPCref_null                  , "ref.null"                  , NoImm                     ,1    , NULLARY(nullref)     , referenceTypes         )   \
	visitOp(0x00d1, OPCref_isnull                , "ref.isnull"                , NoImm                     ,1    , UNARY(anyref,i32)    , referenceTypes         )   \
	visitOp(0x00d2, OPCref_func                  , "ref.func"                  , FunctionImm               ,1    , NULLARY(anyfunc)     , functionRefInstruction )   \
/* Saturating float->int truncation operators                                                                                                            */ \
	visitOp(0xfc00, OPCi32_trunc_s_sat_f32       , "i32.trunc_s:sat/f32"       , NoImm                     ,1    , UNARY(f32,i32)       , nonTrappingFloatToInt  )   \
	visitOp(0xfc01, OPCi32_trunc_u_sat_f32       , "i32.trunc_u:sat/f32"       , NoImm                     ,1    , UNARY(f32,i32)       , nonTrappingFloatToInt  )   \
	visitOp(0xfc02, OPCi32_trunc_s_sat_f64       , "i32.trunc_s:sat/f64"       , NoImm                     ,1    , UNARY(f64,i32)       , nonTrappingFloatToInt  )   \
	visitOp(0xfc03, OPCi32_trunc_u_sat_f64       , "i32.trunc_u:sat/f64"       , NoImm                     ,1    , UNARY(f64,i32)       , nonTrappingFloatToInt  )   \
	visitOp(0xfc04, OPCi64_trunc_s_sat_f32       , "i64.trunc_s:sat/f32"       , NoImm                     ,1    , UNARY(f32,i64)       , nonTrappingFloatToInt  )   \
	visitOp(0xfc05, OPCi64_trunc_u_sat_f32       , "i64.trunc_u:sat/f32"       , NoImm                     ,1    , UNARY(f32,i64)       , nonTrappingFloatToInt  )   \
	visitOp(0xfc06, OPCi64_trunc_s_sat_f64       , "i64.trunc_s:sat/f64"       , NoImm                     ,1    , UNARY(f64,i64)       , nonTrappingFloatToInt  )   \
	visitOp(0xfc07, OPCi64_trunc_u_sat_f64       , "i64.trunc_u:sat/f64"       , NoImm                     ,1    , UNARY(f64,i64)       , nonTrappingFloatToInt  )   \
/* Bulk memory operators                                                                                                                                 */ \
	visitOp(0xfc08, OPCmemory_init               , "memory.init"               , DataSegmentAndMemImm      ,1    , BULKCOPY             , bulkMemoryOperations   )   \
	visitOp(0xfc09, OPCmemory_drop               , "memory.drop"               , DataSegmentImm            ,1    , NONE                 , bulkMemoryOperations   )   \
	visitOp(0xfc0a, OPCmemory_copy               , "memory.copy"               , MemoryImm                 ,1    , BULKCOPY             , bulkMemoryOperations   )   \
	visitOp(0xfc0b, OPCmemory_fill               , "memory.fill"               , MemoryImm                 ,1    , BULKCOPY             , bulkMemoryOperations   )   \
	visitOp(0xfc0c, OPCtable_init                , "table.init"                , ElemSegmentAndTableImm    ,1    , BULKCOPY             , bulkMemoryOperations   )   \
	visitOp(0xfc0d, OPCtable_drop                , "table.drop"                , ElemSegmentImm            ,1    , NONE                 , bulkMemoryOperations   )   \
	visitOp(0xfc0e, OPCtable_copy                , "table.copy"                , TableImm                  ,1    , BULKCOPY             , bulkMemoryOperations   )   \
/* v128 operators                                                                                                                                        */ \
	visitOp(0xfd00, OPCv128_const                , "v128.const"                , LiteralImm_V128           ,1    , NULLARY(v128)        , simd                   )   \
	visitOp(0xfd01, OPCv128_load                 , "v128.load"                 , LoadOrStoreImm            ,1    , LOAD(v128)           , simd                   )   \
	visitOp(0xfd02, OPCv128_store                , "v128.store"                , LoadOrStoreImm            ,1    , STORE(v128)          , simd                   )   \
	visitOp(0xfd03, OPCi8x16_splat               , "i8x16.splat"               , NoImm                     ,1    , UNARY(i32,v128)      , simd                   )   \
	visitOp(0xfd04, OPCi16x8_splat               , "i16x8.splat"               , NoImm                     ,1    , UNARY(i32,v128)      , simd                   )   \
	visitOp(0xfd05, OPCi32x4_splat               , "i32x4.splat"               , NoImm                     ,1    , UNARY(i32,v128)      , simd                   )   \
	visitOp(0xfd06, OPCi64x2_splat               , "i64x2.splat"               , NoImm                     ,1    , UNARY(i64,v128)      , simd                   )   \
	visitOp(0xfd07, OPCf32x4_splat               , "f32x4.splat"               , NoImm                     ,1    , UNARY(f32,v128)      , simd                   )   \
	visitOp(0xfd08, OPCf64x2_splat               , "f64x2.splat"               , NoImm                     ,1    , UNARY(f64,v128)      , simd                   )   \
	visitOp(0xfd09, OPCi8x16_extract_lane_s      , "i8x16.extract_lane_s"      , LaneIndexImm              ,1    , UNARY(v128,i32)      , simd                   )   \
	visitOp(0xfd0a, OPCi8x16_extract_lane_u      , "i8x16.extract_lane_u"      , LaneIndexImm              ,1    , UNARY(v128,i32)      , simd                   )   \
	visitOp(0xfd0b, OPCi16x8_extract_lane_s      , "i16x8.extract_lane_s"      , LaneIndexImm              ,1    , UNARY(v128,i32)      , simd                   )   \
	visitOp(0xfd0c, OPCi16x8_extract_lane_u      , "i16x8.extract_lane_u"      , LaneIndexImm              ,1    , UNARY(v128,i32)      , simd                   )   \
	visitOp(0xfd0d, OPCi32x4_extract_lane        , "i32x4.extract_lane"        , LaneIndexImm              ,1    , UNARY(v128,i32)      , simd                   )   \
	visitOp(0xfd0e, OPCi64x2_extract_lane        , "i64x2.extract_lane"        , LaneIndexImm              ,1    , UNARY(v128,i64)      , simd                   )   \
	visitOp(0xfd0f, OPCf32x4_extract_lane        , "f32x4.extract_lane"        , LaneIndexImm              ,1    , UNARY(v128,f32)      , simd                   )   \
	visitOp(0xfd10, OPCf64x2_extract_lane        , "f64x2.extract_lane"        , LaneIndexImm              ,1    , UNARY(v128,f64)      , simd                   )   \
	visitOp(0xfd11, OPCi8x16_replace_lane        , "i8x16.replace_lane"        , LaneIndexImm              ,1    , REPLACELANE(i32,v128), simd                   )   \
	visitOp(0xfd12, OPCi16x8_replace_lane        , "i16x8.replace_lane"        , LaneIndexImm              ,1    , REPLACELANE(i32,v128), simd                   )   \
	visitOp(0xfd13, OPCi32x4_replace_lane        , "i32x4.replace_lane"        , LaneIndexImm              ,1    , REPLACELANE(i32,v128), simd                   )   \
	visitOp(0xfd14, OPCi64x2_replace_lane        , "i64x2.replace_lane"        , LaneIndexImm              ,1    , REPLACELANE(i64,v128), simd                   )   \
	visitOp(0xfd15, OPCf32x4_replace_lane        , "f32x4.replace_lane"        , LaneIndexImm              ,1    , REPLACELANE(f32,v128), simd                   )   \
	visitOp(0xfd16, OPCf64x2_replace_lane        , "f64x2.replace_lane"        , LaneIndexImm              ,1    , REPLACELANE(f64,v128), simd                   )   \
	visitOp(0xfd17, OPCv8x16_shuffle             , "v8x16.shuffle"             , ShuffleImm_16                ,1    , BINARY(v128,v128)    , simd                   )   \
/* v128 integer arithmetic                                                                                                                               */ \
	visitOp(0xfd18, OPCi8x16_add                 , "i8x16.add"                 , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd19, OPCi16x8_add                 , "i16x8.add"                 , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd1a, OPCi32x4_add                 , "i32x4.add"                 , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd1b, OPCi64x2_add                 , "i64x2.add"                 , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd1c, OPCi8x16_sub                 , "i8x16.sub"                 , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd1d, OPCi16x8_sub                 , "i16x8.sub"                 , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd1e, OPCi32x4_sub                 , "i32x4.sub"                 , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd1f, OPCi64x2_sub                 , "i64x2.sub"                 , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd20, OPCi8x16_mul                 , "i8x16.mul"                 , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd21, OPCi16x8_mul                 , "i16x8.mul"                 , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd22, OPCi32x4_mul                 , "i32x4.mul"                 , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
/*	visitOp(0xfd23, OPCi64x2_mul                 , "i64x2.mul"                 , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )*/ \
	visitOp(0xfd24, OPCi8x16_neg                 , "i8x16.neg"                 , NoImm                     ,1    , UNARY(v128,v128)     , simd                   )   \
	visitOp(0xfd25, OPCi16x8_neg                 , "i16x8.neg"                 , NoImm                     ,1    , UNARY(v128,v128)     , simd                   )   \
	visitOp(0xfd26, OPCi32x4_neg                 , "i32x4.neg"                 , NoImm                     ,1    , UNARY(v128,v128)     , simd                   )   \
	visitOp(0xfd27, OPCi64x2_neg                 , "i64x2.neg"                 , NoImm                     ,1    , UNARY(v128,v128)     , simd                   )   \
	visitOp(0xfd28, OPCi8x16_add_saturate_s      , "i8x16.add_saturate_s"      , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd29, OPCi8x16_add_saturate_u      , "i8x16.add_saturate_u"      , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd2a, OPCi16x8_add_saturate_s      , "i16x8.add_saturate_s"      , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd2b, OPCi16x8_add_saturate_u      , "i16x8.add_saturate_u"      , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd2c, OPCi8x16_sub_saturate_s      , "i8x16.sub_saturate_s"      , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd2d, OPCi8x16_sub_saturate_u      , "i8x16.sub_saturate_u"      , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd2e, OPCi16x8_sub_saturate_s      , "i16x8.sub_saturate_s"      , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd2f, OPCi16x8_sub_saturate_u      , "i16x8.sub_saturate_u"      , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd30, OPCi8x16_shl                 , "i8x16.shl"                 , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd31, OPCi16x8_shl                 , "i16x8.shl"                 , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd32, OPCi32x4_shl                 , "i32x4.shl"                 , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd33, OPCi64x2_shl                 , "i64x2.shl"                 , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd34, OPCi8x16_shr_s               , "i8x16.shr_s"               , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd35, OPCi8x16_shr_u               , "i8x16.shr_u"               , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd36, OPCi16x8_shr_s               , "i16x8.shr_s"               , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd37, OPCi16x8_shr_u               , "i16x8.shr_u"               , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd38, OPCi32x4_shr_s               , "i32x4.shr_s"               , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd39, OPCi32x4_shr_u               , "i32x4.shr_u"               , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd3a, OPCi64x2_shr_s               , "i64x2.shr_s"               , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd3b, OPCi64x2_shr_u               , "i64x2.shr_u"               , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
/* v128 bitwise                                                                                                                                          */ \
	visitOp(0xfd3c, OPCv128_and                  , "v128.and"                  , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd3d, OPCv128_or                   , "v128.or"                   , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd3e, OPCv128_xor                  , "v128.xor"                  , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd3f, OPCv128_not                  , "v128.not"                  , NoImm                     ,1    , UNARY(v128,v128)     , simd                   )   \
	visitOp(0xfd40, OPCv128_bitselect            , "v128.bitselect"            , NoImm                     ,1    , VECTORSELECT(v128)   , simd                   )   \
	visitOp(0xfd41, OPCi8x16_any_true            , "i8x16.any_true"            , NoImm                     ,1    , UNARY(v128,i32)      , simd                   )   \
	visitOp(0xfd42, OPCi16x8_any_true            , "i16x8.any_true"            , NoImm                     ,1    , UNARY(v128,i32)      , simd                   )   \
	visitOp(0xfd43, OPCi32x4_any_true            , "i32x4.any_true"            , NoImm                     ,1    , UNARY(v128,i32)      , simd                   )   \
	visitOp(0xfd44, OPCi64x2_any_true            , "i64x2.any_true"            , NoImm                     ,1    , UNARY(v128,i32)      , simd                   )   \
	visitOp(0xfd45, OPCi8x16_all_true            , "i8x16.all_true"            , NoImm                     ,1    , UNARY(v128,i32)      , simd                   )   \
	visitOp(0xfd46, OPCi16x8_all_true            , "i16x8.all_true"            , NoImm                     ,1    , UNARY(v128,i32)      , simd                   )   \
	visitOp(0xfd47, OPCi32x4_all_true            , "i32x4.all_true"            , NoImm                     ,1    , UNARY(v128,i32)      , simd                   )   \
	visitOp(0xfd48, OPCi64x2_all_true            , "i64x2.all_true"            , NoImm                     ,1    , UNARY(v128,i32)      , simd                   )   \
/* v128 comparisons                                                                                                                                      */ \
	visitOp(0xfd49, OPCi8x16_eq                  , "i8x16.eq"                  , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd4a, OPCi16x8_eq                  , "i16x8.eq"                  , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd4b, OPCi32x4_eq                  , "i32x4.eq"                  , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
/*	visitOp(0xfd4c, i64x2_eq                  , "i64x2.eq"                  , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )*/ \
	visitOp(0xfd4d, OPCf32x4_eq                  , "f32x4.eq"                  , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd4e, OPCf64x2_eq                  , "f64x2.eq"                  , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd4f, OPCi8x16_ne                  , "i8x16.ne"                  , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd50, OPCi16x8_ne                  , "i16x8.ne"                  , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd51, OPCi32x4_ne                  , "i32x4.ne"                  , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
/*	visitOp(0xfd52, i64x2_ne                  , "i64x2.ne"                  , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )*/ \
	visitOp(0xfd53, OPCf32x4_ne                  , "f32x4.ne"                  , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd54, OPCf64x2_ne                  , "f64x2.ne"                  , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd55, OPCi8x16_lt_s                , "i8x16.lt_s"                , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd56, OPCi8x16_lt_u                , "i8x16.lt_u"                , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd57, OPCi16x8_lt_s                , "i16x8.lt_s"                , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd58, OPCi16x8_lt_u                , "i16x8.lt_u"                , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd59, OPCi32x4_lt_s                , "i32x4.lt_s"                , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd5a, OPCi32x4_lt_u                , "i32x4.lt_u"                , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
/*	visitOp(0xfd5b, i64x2_lt_s                , "i64x2.lt_u"                , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )*/ \
/*	visitOp(0xfd5c, i64x2_lt_u                , "i64x2.lt_u"                , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )*/ \
	visitOp(0xfd5d, OPCf32x4_lt                  , "f32x4.lt"                  , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd5e, OPCf64x2_lt                  , "f64x2.lt"                  , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd5f, OPCi8x16_le_s                , "i8x16.le_s"                , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd60, OPCi8x16_le_u                , "i8x16.le_u"                , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd61, OPCi16x8_le_s                , "i16x8.le_s"                , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd62, OPCi16x8_le_u                , "i16x8.le_u"                , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd63, OPCi32x4_le_s                , "i32x4.le_s"                , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd64, OPCi32x4_le_u                , "i32x4.le_u"                , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
/*	visitOp(0xfd65, i64x2_le_s                , "i64x2.le_u"                , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )*/ \
/*	visitOp(0xfd66, i64x2_le_u                , "i64x2.le_u"                , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )*/ \
	visitOp(0xfd67, OPCf32x4_le                  , "f32x4.le"                  , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd68, OPCf64x2_le                  , "f64x2.le"                  , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd69, OPCi8x16_gt_s                , "i8x16.gt_s"                , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd6a, OPCi8x16_gt_u                , "i8x16.gt_u"                , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd6b, OPCi16x8_gt_s                , "i16x8.gt_s"                , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd6c, OPCi16x8_gt_u                , "i16x8.gt_u"                , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd6d, OPCi32x4_gt_s                , "i32x4.gt_s"                , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd6e, OPCi32x4_gt_u                , "i32x4.gt_u"                , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
/*	visitOp(0xfd6f, i64x2_gt_s                , "i64x2.gt_u"                , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )*/ \
/*	visitOp(0xfd70, i64x2_gt_u                , "i64x2.gt_u"                , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )*/ \
	visitOp(0xfd71, OPCf32x4_gt                  , "f32x4.gt"                  , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd72, OPCf64x2_gt                  , "f64x2.gt"                  , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd73, OPCi8x16_ge_s                , "i8x16.ge_s"                , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd74, OPCi8x16_ge_u                , "i8x16.ge_u"                , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd75, OPCi16x8_ge_s                , "i16x8.ge_s"                , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd76, OPCi16x8_ge_u                , "i16x8.ge_u"                , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd77, OPCi32x4_ge_s                , "i32x4.ge_s"                , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd78, OPCi32x4_ge_u                , "i32x4.ge_u"                , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
/*	visitOp(0xfd79, i64x2_ge_s                , "i64x2.ge_u"                , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )*/ \
/*	visitOp(0xfd7a, i64x2_ge_u                , "i64x2.ge_u"                , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )*/ \
	visitOp(0xfd7b, OPCf32x4_ge                  , "f32x4.ge"                  , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd7c, OPCf64x2_ge                  , "f64x2.ge"                  , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
/* v128 floating-point arithmetic                                                                                                                        */ \
	visitOp(0xfd7d, OPCf32x4_neg                 , "f32x4.neg"                 , NoImm                     ,1    , UNARY(v128,v128)     , simd                   )   \
	visitOp(0xfd7e, OPCf64x2_neg                 , "f64x2.neg"                 , NoImm                     ,1    , UNARY(v128,v128)     , simd                   )   \
	visitOp(0xfd7f, OPCf32x4_abs                 , "f32x4.abs"                 , NoImm                     ,1    , UNARY(v128,v128)     , simd                   )   \
	visitOp(0xfd80, OPCf64x2_abs                 , "f64x2.abs"                 , NoImm                     ,1    , UNARY(v128,v128)     , simd                   )   \
	visitOp(0xfd81, OPCf32x4_min                 , "f32x4.min"                 , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd82, OPCf64x2_min                 , "f64x2.min"                 , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd83, OPCf32x4_max                 , "f32x4.max"                 , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd84, OPCf64x2_max                 , "f64x2.max"                 , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd85, OPCf32x4_add                 , "f32x4.add"                 , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd86, OPCf64x2_add                 , "f64x2.add"                 , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd87, OPCf32x4_sub                 , "f32x4.sub"                 , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd88, OPCf64x2_sub                 , "f64x2.sub"                 , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd89, OPCf32x4_div                 , "f32x4.div"                 , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd8a, OPCf64x2_div                 , "f64x2.div"                 , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd8b, OPCf32x4_mul                 , "f32x4.mul"                 , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd8c, OPCf64x2_mul                 , "f64x2.mul"                 , NoImm                     ,1    , BINARY(v128,v128)    , simd                   )   \
	visitOp(0xfd8d, OPCf32x4_sqrt                , "f32x4.sqrt"                , NoImm                     ,1    , UNARY(v128,v128)     , simd                   )   \
	visitOp(0xfd8e, OPCf64x2_sqrt                , "f64x2.sqrt"                , NoImm                     ,1    , UNARY(v128,v128)     , simd                   )   \
/* v128 conversions                                                                                                                                      */ \
	visitOp(0xfd8f, OPCf32x4_convert_s_i32x4     , "f32x4.convert_s/i32x4"     , NoImm                     ,1    , UNARY(v128,v128)     , simd                   )   \
	visitOp(0xfd90, OPCf32x4_convert_u_i32x4     , "f32x4.convert_u/i32x4"     , NoImm                     ,1    , UNARY(v128,v128)     , simd                   )   \
	visitOp(0xfd91, OPCf64x2_convert_s_i64x2     , "f64x2.convert_s/i64x2"     , NoImm                     ,1    , UNARY(v128,v128)     , simd                   )   \
	visitOp(0xfd92, OPCf64x2_convert_u_i64x2     , "f64x2.convert_u/i64x2"     , NoImm                     ,1    , UNARY(v128,v128)     , simd                   )   \
	visitOp(0xfd93, OPCi32x4_trunc_s_sat_f32x4   , "i32x4.trunc_s:sat/f32x4"   , NoImm                     ,1    , UNARY(v128,v128)     , simd                   )   \
	visitOp(0xfd94, OPCi32x4_trunc_u_sat_f32x4   , "i32x4.trunc_u:sat/f32x4"   , NoImm                     ,1    , UNARY(v128,v128)     , simd                   )   \
	visitOp(0xfd95, OPCi64x2_trunc_s_sat_f64x2   , "i64x2.trunc_s:sat/f64x2"   , NoImm                     ,1    , UNARY(v128,v128)     , simd                   )   \
	visitOp(0xfd96, OPCi64x2_trunc_u_sat_f64x2   , "i64x2.trunc_u:sat/f64x2"   , NoImm                     ,1    , UNARY(v128,v128)     , simd                   )   \
/* Atomic wait/wake                                                                                                                                      */ \
	visitOp(0xfe00, OPCatomic_wake               , "atomic.wake"               , AtomicLoadOrStoreImm      ,1    , BINARY(i32,i32)      , atomics                )   \
	visitOp(0xfe01, OPCi32_atomic_wait           , "i32.atomic.wait"           , AtomicLoadOrStoreImm      ,1    , WAIT(i32)            , atomics                )   \
	visitOp(0xfe02, OPCi64_atomic_wait           , "i64.atomic.wait"           , AtomicLoadOrStoreImm      ,1    , WAIT(i64)            , atomics                )   \
/* Atomic load/store                                                                                                                                     */ \
	visitOp(0xfe10, OPCi32_atomic_load           , "i32.atomic.load"           , AtomicLoadOrStoreImm      ,1    , LOAD(i32)            , atomics                )   \
	visitOp(0xfe11, OPCi64_atomic_load           , "i64.atomic.load"           , AtomicLoadOrStoreImm      ,1    , LOAD(i64)            , atomics                )   \
	visitOp(0xfe12, OPCi32_atomic_load8_u        , "i32.atomic.load8_u"        , AtomicLoadOrStoreImm      ,1    , LOAD(i32)            , atomics                )   \
	visitOp(0xfe13, OPCi32_atomic_load16_u       , "i32.atomic.load16_u"       , AtomicLoadOrStoreImm      ,1    , LOAD(i32)            , atomics                )   \
	visitOp(0xfe14, OPCi64_atomic_load8_u        , "i64.atomic.load8_u"        , AtomicLoadOrStoreImm      ,1    , LOAD(i64)            , atomics                )   \
	visitOp(0xfe15, OPCi64_atomic_load16_u       , "i64.atomic.load16_u"       , AtomicLoadOrStoreImm      ,1    , LOAD(i64)            , atomics                )   \
	visitOp(0xfe16, OPCi64_atomic_load32_u       , "i64.atomic.load32_u"       , AtomicLoadOrStoreImm      ,1    , LOAD(i64)            , atomics                )   \
	visitOp(0xfe17, OPCi32_atomic_store          , "i32.atomic.store"          , AtomicLoadOrStoreImm      ,1    , STORE(i32)           , atomics                )   \
	visitOp(0xfe18, OPCi64_atomic_store          , "i64.atomic.store"          , AtomicLoadOrStoreImm      ,1    , STORE(i64)           , atomics                )   \
	visitOp(0xfe19, OPCi32_atomic_store8         , "i32.atomic.store8"         , AtomicLoadOrStoreImm      ,1    , STORE(i32)           , atomics                )   \
	visitOp(0xfe1a, OPCi32_atomic_store16        , "i32.atomic.store16"        , AtomicLoadOrStoreImm      ,1    , STORE(i32)           , atomics                )   \
	visitOp(0xfe1b, OPCi64_atomic_store8         , "i64.atomic.store8"         , AtomicLoadOrStoreImm      ,1    , STORE(i64)           , atomics                )   \
	visitOp(0xfe1c, OPCi64_atomic_store16        , "i64.atomic.store16"        , AtomicLoadOrStoreImm      ,1    , STORE(i64)           , atomics                )   \
	visitOp(0xfe1d, OPCi64_atomic_store32        , "i64.atomic.store32"        , AtomicLoadOrStoreImm      ,1    , STORE(i64)           , atomics                )   \
/* Atomic read-modify-write                                                                                                                              */ \
	visitOp(0xfe1e, OPCi32_atomic_rmw_add        , "i32.atomic.rmw.add"        , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i32)       , atomics                )   \
	visitOp(0xfe1f, OPCi64_atomic_rmw_add        , "i64.atomic.rmw.add"        , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i64)       , atomics                )   \
	visitOp(0xfe20, OPCi32_atomic_rmw8_u_add     , "i32.atomic.rmw8_u.add"     , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i32)       , atomics                )   \
	visitOp(0xfe21, OPCi32_atomic_rmw16_u_add    , "i32.atomic.rmw16_u.add"    , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i32)       , atomics                )   \
	visitOp(0xfe22, OPCi64_atomic_rmw8_u_add     , "i64.atomic.rmw8_u.add"     , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i64)       , atomics                )   \
	visitOp(0xfe23, OPCi64_atomic_rmw16_u_add    , "i64.atomic.rmw16_u.add"    , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i64)       , atomics                )   \
	visitOp(0xfe24, OPCi64_atomic_rmw32_u_add    , "i64.atomic.rmw32_u.add"    , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i64)       , atomics                )   \
	visitOp(0xfe25, OPCi32_atomic_rmw_sub        , "i32.atomic.rmw.sub"        , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i32)       , atomics                )   \
	visitOp(0xfe26, OPCi64_atomic_rmw_sub        , "i64.atomic.rmw.sub"        , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i64)       , atomics                )   \
	visitOp(0xfe27, OPCi32_atomic_rmw8_u_sub     , "i32.atomic.rmw8_u.sub"     , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i32)       , atomics                )   \
	visitOp(0xfe28, OPCi32_atomic_rmw16_u_sub    , "i32.atomic.rmw16_u.sub"    , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i32)       , atomics                )   \
	visitOp(0xfe29, OPCi64_atomic_rmw8_u_sub     , "i64.atomic.rmw8_u.sub"     , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i64)       , atomics                )   \
	visitOp(0xfe2a, OPCi64_atomic_rmw16_u_sub    , "i64.atomic.rmw16_u.sub"    , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i64)       , atomics                )   \
	visitOp(0xfe2b, OPCi64_atomic_rmw32_u_sub    , "i64.atomic.rmw32_u.sub"    , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i64)       , atomics                )   \
	visitOp(0xfe2c, OPCi32_atomic_rmw_and        , "i32.atomic.rmw.and"        , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i32)       , atomics                )   \
	visitOp(0xfe2d, OPCi64_atomic_rmw_and        , "i64.atomic.rmw.and"        , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i64)       , atomics                )   \
	visitOp(0xfe2e, OPCi32_atomic_rmw8_u_and     , "i32.atomic.rmw8_u.and"     , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i32)       , atomics                )   \
	visitOp(0xfe2f, OPCi32_atomic_rmw16_u_and    , "i32.atomic.rmw16_u.and"    , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i32)       , atomics                )   \
	visitOp(0xfe30, OPCi64_atomic_rmw8_u_and     , "i64.atomic.rmw8_u.and"     , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i64)       , atomics                )   \
	visitOp(0xfe31, OPCi64_atomic_rmw16_u_and    , "i64.atomic.rmw16_u.and"    , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i64)       , atomics                )   \
	visitOp(0xfe32, OPCi64_atomic_rmw32_u_and    , "i64.atomic.rmw32_u.and"    , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i64)       , atomics                )   \
	visitOp(0xfe33, OPCi32_atomic_rmw_or         , "i32.atomic.rmw.or"         , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i32)       , atomics                )   \
	visitOp(0xfe34, OPCi64_atomic_rmw_or         , "i64.atomic.rmw.or"         , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i64)       , atomics                )   \
	visitOp(0xfe35, OPCi32_atomic_rmw8_u_or      , "i32.atomic.rmw8_u.or"      , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i32)       , atomics                )   \
	visitOp(0xfe36, OPCi32_atomic_rmw16_u_or     , "i32.atomic.rmw16_u.or"     , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i32)       , atomics                )   \
	visitOp(0xfe37, OPCi64_atomic_rmw8_u_or      , "i64.atomic.rmw8_u.or"      , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i64)       , atomics                )   \
	visitOp(0xfe38, OPCi64_atomic_rmw16_u_or     , "i64.atomic.rmw16_u.or"     , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i64)       , atomics                )   \
	visitOp(0xfe39, OPCi64_atomic_rmw32_u_or     , "i64.atomic.rmw32_u.or"     , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i64)       , atomics                )   \
	visitOp(0xfe3a, OPCi32_atomic_rmw_xor        , "i32.atomic.rmw.xor"        , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i32)       , atomics                )   \
	visitOp(0xfe3b, OPCi64_atomic_rmw_xor        , "i64.atomic.rmw.xor"        , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i64)       , atomics                )   \
	visitOp(0xfe3c, OPCi32_atomic_rmw8_u_xor     , "i32.atomic.rmw8_u.xor"     , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i32)       , atomics                )   \
	visitOp(0xfe3d, OPCi32_atomic_rmw16_u_xor    , "i32.atomic.rmw16_u.xor"    , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i32)       , atomics                )   \
	visitOp(0xfe3e, OPCi64_atomic_rmw8_u_xor     , "i64.atomic.rmw8_u.xor"     , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i64)       , atomics                )   \
	visitOp(0xfe3f, OPCi64_atomic_rmw16_u_xor    , "i64.atomic.rmw16_u.xor"    , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i64)       , atomics                )   \
	visitOp(0xfe40, OPCi64_atomic_rmw32_u_xor    , "i64.atomic.rmw32_u.xor"    , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i64)       , atomics                )   \
	visitOp(0xfe41, OPCi32_atomic_rmw_xchg       , "i32.atomic.rmw.xchg"       , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i32)       , atomics                )   \
	visitOp(0xfe42, OPCi64_atomic_rmw_xchg       , "i64.atomic.rmw.xchg"       , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i64)       , atomics                )   \
	visitOp(0xfe43, OPCi32_atomic_rmw8_u_xchg    , "i32.atomic.rmw8_u.xchg"    , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i32)       , atomics                )   \
	visitOp(0xfe44, OPCi32_atomic_rmw16_u_xchg   , "i32.atomic.rmw16_u.xchg"   , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i32)       , atomics                )   \
	visitOp(0xfe45, OPCi64_atomic_rmw8_u_xchg    , "i64.atomic.rmw8_u.xchg"    , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i64)       , atomics                )   \
	visitOp(0xfe46, OPCi64_atomic_rmw16_u_xchg   , "i64.atomic.rmw16_u.xchg"   , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i64)       , atomics                )   \
	visitOp(0xfe47, OPCi64_atomic_rmw32_u_xchg   , "i64.atomic.rmw32_u.xchg"   , AtomicLoadOrStoreImm      ,1    , ATOMICRMW(i64)       , atomics                )   \
	visitOp(0xfe48, OPCi32_atomic_rmw_cmpxchg    , "i32.atomic.rmw.cmpxchg"    , AtomicLoadOrStoreImm      ,1    , COMPAREEXCHANGE(i32) , atomics                )   \
	visitOp(0xfe49, OPCi64_atomic_rmw_cmpxchg    , "i64.atomic.rmw.cmpxchg"    , AtomicLoadOrStoreImm      ,1    , COMPAREEXCHANGE(i64) , atomics                )   \
	visitOp(0xfe4a, OPCi32_atomic_rmw8_u_cmpxchg , "i32.atomic.rmw8_u.cmpxchg" , AtomicLoadOrStoreImm      ,1    , COMPAREEXCHANGE(i32) , atomics                )   \
	visitOp(0xfe4b, OPCi32_atomic_rmw16_u_cmpxchg, "i32.atomic.rmw16_u.cmpxchg", AtomicLoadOrStoreImm      ,1    , COMPAREEXCHANGE(i32) , atomics                )   \
	visitOp(0xfe4c, OPCi64_atomic_rmw8_u_cmpxchg , "i64.atomic.rmw8_u.cmpxchg" , AtomicLoadOrStoreImm      ,1    , COMPAREEXCHANGE(i64) , atomics                )   \
	visitOp(0xfe4d, OPCi64_atomic_rmw16_u_cmpxchg, "i64.atomic.rmw16_u.cmpxchg", AtomicLoadOrStoreImm      ,1    , COMPAREEXCHANGE(i64) , atomics                )   \
	visitOp(0xfe4e, OPCi64_atomic_rmw32_u_cmpxchg, "i64.atomic.rmw32_u.cmpxchg", AtomicLoadOrStoreImm      ,1    , COMPAREEXCHANGE(i64) , atomics                )


#define ENUM_NONCONTROL_OPERATORS(visitOp)                                                         \
	ENUM_PARAMETRIC_OPERATORS(visitOp)                                                             \
	ENUM_NONCONTROL_NONPARAMETRIC_OPERATORS(visitOp)

#define ENUM_OPERATORS(visitOp)                                                                    \
	ENUM_NONCONTROL_OPERATORS(visitOp)                                                             \
	ENUM_CONTROL_OPERATORS(visitOp)
