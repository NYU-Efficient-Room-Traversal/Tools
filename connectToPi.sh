#!/bin/bash
if  [ -z $(sudo arp -na | grep -i b8:27:eb | awk -F"[()]" 'NR==1{print $2}') ]; then
    echo "Refreshing ARP"
    sudo nmap -sn 192.168.1.0/24
    ssh pi@$(sudo arp -na | grep -i b8:27:eb | awk -F"[()]" 'NR==1{print $2}')
else
    ssh pi@$(sudo arp -na | grep -i b8:27:eb | awk -F"[()]" 'NR==1{print $2}')
fi
