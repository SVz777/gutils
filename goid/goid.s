#include "textflag.h"

TEXT ·GetGoID(SB),NOSPLIT,$0-8
    MOVQ (TLS), AX
    LEAQ (+152)(AX),BX
    MOVQ BX,g+0(FP)
    RET
