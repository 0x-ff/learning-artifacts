
extern crate pnet;
extern crate clap;

use pnet::datalink::{self};

fn main() 
{
	let interfaces = datalink::interfaces();
	for interface in interfaces {
		println!("{}: flags={}<...> mtu ...", interface.name, interface.flags);
		match interface.mac {
			None => println!("\tether <none>"),
			Some(mac) => println!("\tether {}", mac),
		};
	}
}
