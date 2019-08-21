
extern crate pnet;
extern crate clap;

use std::net::{Ipv4Addr, AddrParseError, IpAddr};

use clap::{Arg, App};

use pnet::util::MacAddr;
use pnet::packet::{Packet, MutablePacket};
use pnet::packet::ethernet::{EtherTypes, MutableEthernetPacket};
use pnet::packet::arp::{ArpHardwareTypes, ArpOperation, MutableArpPacket};
use pnet::datalink::{self, Channel, NetworkInterface};

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
	frame.set_ethertype(EtherTypes::Arp);
	
	let mut arp_buffer = [0u8; 28];
	let mut datagram = MutableArpPacket::new(&mut arp_buffer).unwrap();

	datagram.set_hardware_type(ArpHardwareTypes::Ethernet);
	datagram.set_protocol_type(EtherTypes::Ipv4);
	datagram.set_hw_addr_len(6);
	datagram.set_proto_addr_len(4);
	datagram.set_operation(ArpOperation::new(2));
	
	datagram.set_sender_hw_addr(interface.mac.unwrap());
	datagram.set_sender_proto_addr(search_ip);
	
	datagram.set_target_hw_addr(MacAddr::new(255, 255, 255, 255, 255, 255));
	datagram.set_target_proto_addr(Ipv4Addr::new(255, 255, 255, 255));

	frame.set_payload(datagram.packet_mut());
	
	tx.send_to(frame.packet(), Some(interface.clone()));
}
