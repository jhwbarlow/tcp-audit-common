# tcp-audit-common

This module holds packages which are imported by both the main `tcp-audit` module and the modules containing the Eventer and Sinker plugins.

These common packages must be in a separate module due to [Go issue #27751](https://github.com/golang/go/issues/27751).
