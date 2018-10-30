# lynea

## DO NOT USE

When I had the idea of writing an init system in Go or Rust, it was because I didn't see any other project doing it and
I thought that it would be a fun exercise, especially given the hubbubb around systemd. Thanks to various major life
events, I was never really able to dedicate the time I needed to learn all of the relevant concepts well enough to
confidently "blaze the path" that I saw in my head.

Now, after roughly 3 years of inactivity in this project and with the idea being mostly unoriginal, I'm comfortable
"sunsetting" this project.

I may get the wild bug to make the attempt again, I may not. Mark this as a lesson: Big ideas are great, but if you
don't execute all you have is a dream.

I won't be removing this repo, instead, I want to keep it as a reminder to myself that an idea without proper planning 
and the will to dedicate the time (for whatever valid reasons) has failed before it began.

-- Brian

[![Join the chat at https://gitter.im/darthlukan/lynea](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/darthlukan/lynea?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

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


## Current State of Development

> Lynea is still in the very early stages of development.  It is still in the beginning stages of design.  This means 
> that currently the code is mostly a collection of comments as notes surrounded by incomplete functions that may or 
> may not be used.  It is not runnable by any means for use as an init system or anything else for that matter.

> It goes without saying, but you shouldn't attempt to run this as is.

## License

> GPL3, see LICENSE file
