#include "imm.h"
package types

#define visit_imm(imm) \
        type OpcodeAndImm ##_## imm struct{ \
            Opcode; \
            Imm imm; \
        };
    enum_imms(visit_imm)
#undef visit_imm
