# blackshield

defensive surface tooling for controlled environments

focused on ingress pressure, authentication abuse, and packet-level anomalies

---

## components

### ingress (go)
connection pressure tracking with strike escalation

- per-source tracking
- adaptive strike model
- decay-based memory cleanup

### audit (python)
auth log correlation and brute-force detection

- multi-pattern parsing
- stateful tracking
- hashed source fingerprinting

### wire (rust)
packet-rate anomaly detection

- low-overhead UDP telemetry
- time-window pruning
- fast anomaly signaling

---

## build

```bash
bash ops/init.sh
