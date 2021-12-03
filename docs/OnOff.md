 
# On-Off

- If your code works well after checking, you can add `#define TRICE_OFF` just before the `#include "trice.h"` line and no *trice* code is generated anymore for that file, so no need to delete or comment out `TRICE` macros.
- No runtime On-Off switch is implemented for  several reasons:
  - Would need a control channel to the target.
  - Would add little performance and code overhead.
  - Would sligtly change target timing (testing).
  - User can add its own switches anywhere.
  - The short `TRICE` macro code negligible.
  - The trice output is encryptable, if needed.
  - The PC **trice** tool offers command line switches to `-pick` or `-ban` several *trice* channels for the runtime display.