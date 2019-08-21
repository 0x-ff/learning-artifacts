
extern crate pnet;
extern crate clap;

use clap::{Arg, App};
use std::net::{Ipv4Addr, AddrParseError, IpAddr};
use pnet::packet::icmp::{MutableIcmpPacket, IcmpPacket, IcmpType, IcmpCode};
use pnet::packet::ip::{IpNextHeaderProtocols};
use pnet::transport::{transport_channel, icmp_packet_iter};
use pnet::transport::TransportProtocol::Ipv4;
use pnet::transport::TransportChannelType::Layer4;
use pnet::packet::Packet;

const ICMP_PAYLOAD_LENGTH :usize = 60;
const ICMP_HEADER_LENGTH :usize = 4;

fn cook_icmp_echo_request(packet: &mut MutableIcmpPacket)
{
	packet.set_icmp_type(IcmpType::new(8)); // ICMP Echo Request
	packet.set_icmp_code(IcmpCode::new(0));
	
	let payload: [u8; ICMP_PAYLOAD_LENGTH] = [
		0x00, 0xA0, // Здесь лучше передавать PID - Id процесса, чтобы не ловить icmp-пакеты от другого процесса ping на тот же хост
		0x00, 0x00, 
		0x5a, 0x4f, 0xe5, 0xad, 0x00, 0x00,
		0x16, 0xab, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15,
		0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f, 0x20, 0x21, 0x22, 0x23, 0x24, 0x25,
		0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35,
		0x36, 0x37
	];
	packet.set_payload(&payload);
	let icmp_checksum;
	{
		let icmp: IcmpPacket = packet.to_immutable();
		icmp_checksum = pnet::packet::icmp::checksum(&icmp);
	}
	packet.set_checksum(icmp_checksum);
}

fn main()
{	
	let matches = App::new("ether")
		.arg(Arg::with_name("ip")
		.long("ip")
		.required(true)
		.takes_value(true)
		.help("ip address"))
		.get_matches();
		
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
	
	let mut icmp_buffer = [0u8; ICMP_HEADER_LENGTH + ICMP_PAYLOAD_LENGTH]; 
	let mut packet = MutableIcmpPacket::new(&mut icmp_buffer).unwrap();
	cook_icmp_echo_request(&mut packet);

	let protocol = Layer4(Ipv4(IpNextHeaderProtocols::Icmp));

    let (mut tx, mut rx) = match transport_channel(4096, protocol) {
        Ok((tx, rx)) => (tx, rx),
        Err(e) => {
            panic!("An error occurred when creating the transport channel: {}", e)
        }
    };
    
    let result = tx.send_to(packet, IpAddr::V4(search_ip));
    match result {
		Ok(len) => {
			println!("Send icmp echo request {:?} bytes to {:?}", len, search_ip);
		},
		Err(e) => {
            panic!("An error occurred when send icmp echo request: {}", e)
        }
	};
    
    let mut iter = icmp_packet_iter(&mut rx);
    loop {
        match iter.next() {
            Ok((packet, addr)) => {
				let payload = packet.payload();
				if addr == search_ip && payload[0] == 0x00 && payload[1] == 0xA0 {
					println!("Received icmp echo reply from {:?}", search_ip);
					break;
				}
            }
            Err(e) => {
                panic!("An error occurred while reading: {}", e);
            }
        }
    }
}
