# lynea

> Author: Brian Tomlinson <brian.tomlinson@linux.com>


## Description

> My wife's name is Lynea.  My life didn't really start until I met her, now, your computer won't really
> start without her either.

> Lynea is meant to be a very simple init system that does the following:

* Take over PID 1
* Mount necessary (virtual and real) filesystems
* Start, Stop, and Restart services
* Handle reboot, poweroff, halt
* Handle signals as appropriate for an init system

> Anything else that can be handled by another tool should probably be handled by that tool, so anything beyond
> the functionality provided by the five points above should not be expected. Ever.

## License

> GPL3, see LICENSE file
