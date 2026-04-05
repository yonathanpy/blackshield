#!/bin/bash
set -e

echo "[*] build ingress"
cd ingress && go build -o ingress

echo "[*] prepare audit"
chmod +x ../audit/core.py

echo "[*] build wire"
cd ../wire && rustc main.rs -O -o wire

echo "[+] ready"
