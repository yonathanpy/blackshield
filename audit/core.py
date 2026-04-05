import re
import time
import hashlib
from collections import defaultdict

LOG = "/var/log/auth.log"

patterns = {
    "fail": re.compile(r"Failed password for .* from (\d+\.\d+\.\d+\.\d+)"),
    "ok": re.compile(r"Accepted password for .* from (\d+\.\d+\.\d+\.\d+)"),
    "invalid": re.compile(r"Invalid user .* from (\d+\.\d+\.\d+\.\d+)")
}

state = defaultdict(lambda: {"fail":0, "ok":0, "invalid":0})

FAIL_THRESHOLD = 7

def sig(ip):
    return hashlib.blake2b(ip.encode(), digest_size=6).hexdigest()

def analyze(line):
    for k, p in patterns.items():
        m = p.search(line)
        if m:
            ip = m.group(1)
            state[ip][k] += 1

            if state[ip]["fail"] >= FAIL_THRESHOLD:
                print(f"alert type=bruteforce src={ip} sig={sig(ip)} count={state[ip]['fail']}")

def follow():
    with open(LOG, "r") as f:
        f.seek(0, 2)
        while True:
            line = f.readline()
            if not line:
                time.sleep(0.2)
                continue
            analyze(line)

if __name__ == "__main__":
    follow()
