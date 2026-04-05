# blackshield

<p align="center">
  <img src="https://img.icons8.com/ios-filled/500/000000/bug.png" width="90"/>
  <img src="https://img.icons8.com/ios-filled/500/000000/activity-history.png" width="90"/>
  <img src="https://img.icons8.com/ios-filled/500/000000/hammer.png" width="90"/>
  <img src="https://img.icons8.com/ios-filled/500/000000/anvil.png" width="90"/>
  <img src="https://img.icons8.com/ios-filled/500/000000/fire.png" width="90"/>
</p>

<p align="center">
────────────────────────────────────────────
</p>

BlackShield is a **layered defensive toolkit** for high-sensitivity infrastructure.  
It provides **connection tracking**, **authentication abuse detection**, **packet-rate anomaly monitoring**, and **kernel-level enforcement**.  

No external dependencies. Deterministic, fail-closed, and bounded memory design.

---

## Architecture Overview

```
blackshield/
│
├── ingress/ # Go — connection pressure & strike enforcement
├── audit/ # Python — authentication log correlation
├── wire/ # Rust — packet telemetry & anomaly detection
├── kernel/
│ └── fluxguard/ # eBPF/XDP kernel-level enforcement
├── policy/
│ └── baseline.yaml
└── ops/
└── init.sh
```

---

## Components

### 1. Ingress (Go)
Tracks TCP connection attempts with **strike-based enforcement**.  

**Core logic snippet (abstracted):**
```go
if hits exceed limit:
    increment strike counter
    enforce drop when threshold reached
```

Per-source accounting  
Strike escalation model  
Time-decay cleanup  

---

### 2. Audit (Python)

Correlates authentication logs and detects brute-force attempts.

Example pattern matching (simplified):

```python
extract source IP from failed authentication events
increment failure counters per source
```

Multi-pattern regex parsing  
Stateful per-source tracking  
Deterministic fingerprinting  

---

### 3. Wire (Rust)

Monitors packet flows and identifies anomalies.

Packet-rate check (abstracted):

```rust
if rate exceeds threshold:
    emit anomaly event
```

Low-latency intake  
Per-source rate tracking  
Time-window pruning  

---

### 4. FluxGuard (eBPF / XDP)

Kernel-level enforcement drops malicious packets before reaching userspace.

Decision logic (abstracted):

```c
if threshold exceeded:
    drop packet at kernel level
```

XDP hook for zero-copy evaluation  
Per-source packet accounting in kernel memory  
Immediate drop on threshold breach  

---

## Policy Configuration

File: `policy/baseline.yaml`

```yaml
ingress:
  hit_limit: 100
  strike_limit: 4

audit:
  fail_threshold: 7

wire:
  packet_threshold: 200

fluxguard:
  kernel_limit: 250

response:
  drop: true
  alert: true
```

Central control for thresholds and enforcement  
Simple, auditable, and human-readable  

---

## Build & Deployment

Userland:

```bash
bash ops/init.sh
```

Kernel Layer:

```bash
cd kernel/fluxguard
make
sudo ./fluxguard eth0 fluxguard.o
```

Detach kernel program:

```bash
sudo ip link set dev eth0 xdp off
```

Run userland modules:

```bash
./ingress/ingress
python3 audit/core.py
./wire/wire
```

---

## Execution Flow

Packet hits interface → FluxGuard evaluates in kernel  
High-rate sources → dropped instantly  
Surviving connections → Ingress monitors TCP attempts  
Authentication events → Audit correlates brute-force activity  
Wire monitors packet flow anomalies  
Alerts emitted across all modules  

---

## Key Features

Kernel-first enforcement  
Deterministic userland tracking (Go/Python/Rust)  
Fail-closed, bounded memory design  
No external telemetry or cloud services  
High-performance, real defensive posture  

---

## Operational Scope

Intended for:

Perimeter nodes  
Authentication gateways  
Segmented internal networks  
High-sensitivity infrastructure  

Not designed for:

Forensic analysis  
Distributed intelligence aggregation  

---

## Extension Points

nftables/ipset integration  
Distributed state propagation  
Protocol-aware parsing at eBPF level  
Adaptive thresholds based on baseline learning  

---

## Summary

BlackShield observes → classifies → enforces → discards.  
Early suppression is prioritized over late interpretation, delivering a robust, expert-level defensive toolkit.
