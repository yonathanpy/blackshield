<!-- Anvil (forging) -->
<img src="https://img.icons8.com/ios-filled/500/000000/anvil.png" width="200"/>

# blackshield

<p align="center">
════════════════════════════════════════════════════════════
</p>

BlackShield is a **layered defensive toolkit** for high-sensitivity infrastructure.  
It provides **connection tracking**, **authentication abuse detection**, **packet-rate anomaly monitoring**, and **kernel-level enforcement**.  

No external dependencies. Deterministic, fail-closed, and bounded memory design.

<p align="center">
════════════════════════════════════════════════════════════
</p>

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

**Core logic (abstracted):**
```go
if limit exceeded:
    increase strike level
    enforce drop at threshold
```

Per-source accounting  
Strike escalation model  
Time-decay cleanup  

---

### 2. Audit (Python)

Correlates authentication logs and detects brute-force attempts.

**Pattern logic (abstracted):**
```python
detect failed authentication events
track failures per source
```

Multi-pattern parsing  
Stateful tracking  
Deterministic fingerprinting  

---

### 3. Wire (Rust)

Monitors packet flows and identifies anomalies.

**Detection logic (abstracted):**
```rust
if traffic rate abnormal:
    trigger anomaly event
```

Low-latency intake  
Per-source rate tracking  
Time-window pruning  

---

### 4. FluxGuard (eBPF / XDP)

Kernel-level enforcement drops malicious packets before reaching userspace.

**Kernel decision (abstracted):**
```c
if threshold reached:
    drop packet
```

XDP hook for zero-copy evaluation  
Kernel memory tracking  
Immediate enforcement  

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
Simple and auditable  

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

Detach:
```bash
sudo ip link set dev eth0 xdp off
```

Run:
```bash
./ingress/ingress
python3 audit/core.py
./wire/wire
```

---

## Execution Flow

Packet → Kernel evaluation → Drop or pass  
Ingress tracking → Audit correlation → Wire analysis  
Enforcement + alerts  

---

## Key Features

Kernel-first enforcement  
Deterministic tracking  
Bounded memory design  
No external dependencies  
High-performance defense  

---

## Operational Scope

Intended for:

Perimeter nodes  
Authentication gateways  
Segmented networks  
Sensitive infrastructure  

Not designed for:

Forensics  
Distributed intelligence  

---

## Extension Points

nftables / ipset  
State propagation  
Protocol-aware parsing  
Adaptive thresholds  

---

## Summary

BlackShield observes → classifies → enforces → discards.

<p align="center">
════════════════════════════════════════════════════════════
</p>
