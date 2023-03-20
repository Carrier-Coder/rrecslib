# RRRECSulator

The [RRECSulator](https://rrecsulator.com) helps USPS rural carriers calculate their pay check in RRECS.

## Source

This repository contains the source code for RRECS calculations.

Many eyes find many bugs.  Please report here, or email to [rrecsulator.com@gmail.com](mailto:rrecsulator.com@gmail.com)

## History

The calculator was developed from discussion in the [Rural Mail Talk Forum](https://www.ruralmailtalk.com/forums/rrecs-detailed-calculations-by-c.31/)

## Implementation Details

The server is written in go.  Each folder roughly corresonds to a different component of a carriers evaluation (e.g `flat` is for dealing with flats, like newspapers, magazines, etc.)  There are a few specialized folders:

- `standards` --> numeric constants are defined (like verifying letter addresses).
- `dataSets` --> glues together all of the other folders.
- `fullRRECS` --> shows running a calcuation and generating the report.

# License
GPL v3
