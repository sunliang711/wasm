#include "imm.h"
package IR
//Note this file is created by makeTypes.sh,Don't modify this file directly

#define visit_imm(imm) \
        type OpcodeAndImm ##_## imm struct{ \
            Opcode; \
            Imm imm; \
        };
    enum_imms(visit_imm)
#undef visit_imm
