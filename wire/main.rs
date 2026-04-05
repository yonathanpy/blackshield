use std::collections::HashMap;
use std::net::UdpSocket;
use std::time::{Duration, Instant};

struct Entry {
    count: u64,
    last: Instant,
}

fn main() {
    let sock = UdpSocket::bind("0.0.0.0:9999").expect("bind failed");
    let mut buf = [0u8; 1500];

    let mut table: HashMap<String, Entry> = HashMap::new();
    let threshold: u64 = 200;

    loop {
        if let Ok((_, src)) = sock.recv_from(&mut buf) {
            let key = src.ip().to_string();

            let e = table.entry(key.clone()).or_insert(Entry {
                count: 0,
                last: Instant::now(),
            });

            e.count += 1;
            e.last = Instant::now();

            if e.count > threshold {
                println!("event=rate_anomaly src={} count={}", key, e.count);
            }
        }

        table.retain(|_, v| v.last.elapsed() < Duration::from_secs(240));
    }
}
