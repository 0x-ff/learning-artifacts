
extern crate pnet;
extern crate clap;

use clap::{Arg, App};
use pnet::packet::Packet;
use std::net::{Ipv4Addr, AddrParseError, IpAddr};
use pnet::packet::ipv4::checksum;
use pnet::datalink::{self, NetworkInterface};
use pnet::datalink::Channel;
use pnet::packet::ethernet::MutableEthernetPacket;
use pnet::packet::ethernet::EtherTypes;
use pnet::util::MacAddr;
use pnet::packet::ipv4::{MutableIpv4Packet, Ipv4Packet};

fn get_interface_by_name(name: &str) -> NetworkInterface
{
	let interfaces = datalink::interfaces();
	for interface in interfaces {
		if interface.name == name {
			return interface;
		}
	}
	println!("Unknown interface {}", name);
	std::process::exit(-1)
}

fn get_interface_ipv4_addr(interface: &NetworkInterface) -> Option<Ipv4Addr>
{
	let mut ip: Option<Ipv4Addr> = None;
	let interface_copy = interface.clone();
	for ip_net in interface_copy.ips {
		match ip_net.ip() {
			IpAddr::V4(addr) => {
				ip = Some(addr);
				break;
			},
			_ => {}
		};
	}
	return ip;
}

fn main()
{	
	let matches = App::new("ether")
		.arg(Arg::with_name("i")
		.long("i")
		.required(true)
		.takes_value(true)
		.help("interface"))
		.arg(Arg::with_name("ip")
		.long("ip")
		.required(true)
		.takes_value(true)
		.help("ip address"))
		.get_matches();
		
	let name: &str = &matches.value_of("i").unwrap();
	let ip_address: &str = matches.value_of("ip").unwrap();
	
	let ip_result: Result<Ipv4Addr, AddrParseError> = ip_address.trim().parse();
	let search_ip: Ipv4Addr;
	match ip_result {
		Ok(ipaddr) => { 
			search_ip = ipaddr 
		},
		Err(e) => { 
			println!("ip = {}", e);
			return;
		}
	};
	
	let interface = get_interface_by_name(name);
	let my_ip: Option<Ipv4Addr> = get_interface_ipv4_addr(&interface);

	if my_ip == None {
		println!("interface {:?} doesn't have ipv4 address", interface.name);
		return;
	}

	let(mut tx, _) = match pnet::datalink::channel(&interface, Default::default()) {
        Ok(Channel::Ethernet(tx, rx)) => (tx, rx),
        Ok(_) => panic!("Unknown channel type"),
        Err(e) => panic!("Error happened {}", e),
    };
    
	let mut buffer = [0u8; 42];
    let mut frame = MutableEthernetPacket::new(&mut buffer).unwrap();
	
	let dest_mac = MacAddr::new(0, 0, 0, 0, 0, 0);
	frame.set_destination(dest_mac);
	frame.set_source(interface.mac.unwrap());
	frame.set_ethertype(EtherTypes::Ipv4);
	
	let mut ip_buffer = [0u8; 40];
	let mut datagram = MutableIpv4Packet::new(&mut ip_buffer).unwrap();

	datagram.set_version(4);
	datagram.set_header_length(5);
	datagram.set_total_length(28);
	datagram.set_identification(1);
	datagram.set_ttl(100);
	
	datagram.set_source(my_ip.unwrap());
	datagram.set_destination(search_ip);

	let payload: [u8; 8] = [0, 0, 0, 0, 0, 0, 0, 0 ];
	datagram.set_payload(&payload);
	
	let chksum;
	{
		let ipv4: Ipv4Packet = datagram.to_immutable();
		chksum = checksum(&ipv4);
	}
	datagram.set_checksum(chksum);

	println!("datagram = {:?}", datagram);
	frame.set_payload(datagram.packet());
	
	for _i in 0..100 {
		tx.send_to(frame.packet(), Some(interface.clone()));
	}
}
