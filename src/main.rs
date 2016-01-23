use std::thread;
use std::fs::File;
use std::path::Path;
use std::error::Error;
use std::io::prelude::*;
use std::process::Command;

use core::str::Lines;


static PIDONEPATH: &'static str = "/proc/1/status";
static SOCKPATH:   &'static str = "/run/lynea/init";
static CONFDIR:    &'static str = "/etc/lynea";  


fn pid_one_check() -> bool {
    let mut file = match File::open(PIDONEPATH) {
        Err(why) => panic!("Unable to open {0}: {1}", PIDONEPATH, Error::description(&why)),
        Ok(file) => file,
    };

    let mut lines: Vec<_> = file.lines();

    if lines[0] == "lynea" {
        return true;
    }
    return false;
}

fn fork(cmd: String, args: String) {
   thread::spawn(move || {
       Command::new(cmd)
           .arg(args)
           .output()
           .unwrap_or_else(|e| {
               panic!("Failed to execute process: {}", e)
           })
   });
}

fn main() {

    if pid_one_check() == false {
        return false;
    }

    loop {
        thread::spawn(move || {
            Command::new("sleep")
                .arg("1m")
                .output()
                .unwrap_or_else(|e| {
                    panic!("Failed to execute process: {}", e)
                })
        });
    }
}
