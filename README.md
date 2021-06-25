# GoFxell
Fxell server written in Golang

# Feature List
## Footprint
[x] Footprint Meters on each node. Footprint is mesaured as a percentage and is a multiplier
for the other teams detection actions. 100% footprint results in the implant the team is using being exposed in the proclist.

Operational Security: Starts at 100 goes down to zero? Side affects?

Triage Logs?
Certain nodes provide multipliers for miners

Final Objectives
[x]Singple plant - a node is denoted as the final target node. Identify the node and plant a package on the node
    Deface
    Exfil
    Ddos**
Muli plant - several adjeacent nodes must be identified and planted upon. More stealth less time per plant.
    Exfil Loot
    Ddos
Distributed Final Actions
    Cryptolocker - Use entropy to cryptolock as many devices as possible. This makes opposing objectives difficult. Rewards based on completion?
    Botnet miners

## Defense Tools
[x]Clean Logs
Burn/Evacuate
Filesystem Monitor - scans for additions on disc
Network Monitor - Alerts on significant (leveled) footprint changes on adjacent nodes (those in routes)
procmon - looks for new processes in process list
Redirect/Proxy Traffic - redirect traffic from one node to another. Killable, logged, negated with a hop
[x]Firewall - prevents connections to a machine. Removed with dos or ddos but can be reinforced/held with resources
Honeypots - Create a honeypot node that logs ALL actions. Deployed while on a target and added to the targets routing table. Service and platform customizable.

## Hybrid Tools
Keylogging
Logging Stack and Monitoring Multiple Boxes
Retrieve/get-file/Download
Persistence - Service/Scheduled Tasks

## Offensive
[x]Kill
[x]DoS
DDoS
Analysis (After a Retrieve?)
Tool Takeover

## Final Objective Options
Plant Artifacts
Passive Monitoring
CryptoLocker
Deface
Steal Files
Botnet Mining

## Infrastructure
IoT
SCADA
Mobile
Wireless

## Exploits
Masquarade
