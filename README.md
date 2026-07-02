# Go QFS(Quick File Share)

A lightweight desktop file-sharing application built in Go. Run the binary in any directory to instantly share its contents with any device on the same network. no installation, no configuration.

## Usage
1. Download the latest release from the [releases page](https://github.com/y3eet/go-qfs/releases)
2. Place the executable file in the directory you want to share
3. Run it
4. A local HTTP server starts and the app displays a QR code and connection address
5. Connect from any device on the same network and scan the QR code or enter the address in a browser
   > [!NOTE]
   > You may need to expose your port if you are using a firewall [See this guide](#firewall-configuration)
6. Upload or download files directly from that shared directory

## Firewall Configuration Guide

## Windows
 
### Option A: PowerShell (fastest)
 
Open PowerShell **as Administrator**, then:
 
**Open the port**
```powershell
New-NetFirewallRule -DisplayName "Allow Port 8080" -Direction Inbound -LocalPort 8080 -Protocol TCP -Action Allow
```
 
**Close the port**
```powershell
Remove-NetFirewallRule -DisplayName "Allow Port 8080"
```
 
**Verify**
```powershell
Get-NetFirewallRule -DisplayName "Allow Port 8080"
```
 
### Option B: Windows Defender Firewall (GUI)
 
1. Press `Win`, search for **Windows Defender Firewall with Advanced Security**, and open it.
2. Click **Inbound Rules** on the left, then **New Rule...** on the right.
3. Select **Port** → **Next**.
4. Select **TCP**, choose **Specific local ports**, type `8080` → **Next**.
5. Select **Allow the connection** → **Next**.
6. Leave all profiles checked (or pick as needed) → **Next**.
7. Give it a name like `Allow Port 8080` → **Finish**.
To close it later: go to **Inbound Rules**, find the rule, right-click → **Delete**.
 
---
 
## Linux
 
### UFW
 
**Open the port**
```bash
sudo ufw allow 8080/tcp
```
 
**Close the port**
```bash
sudo ufw delete allow 8080/tcp
```
*(Or use `sudo ufw deny 8080/tcp` if you want to explicitly block it rather than just removing the rule.)*
 
**Verify**
```bash
sudo ufw status
```
 
### firewalld
 
**Open the port**
```bash
sudo firewall-cmd --permanent --add-port=8080/tcp
sudo firewall-cmd --reload
```
 
**Close the port**
```bash
sudo firewall-cmd --permanent --remove-port=8080/tcp
sudo firewall-cmd --reload
```
 
**Verify**
```bash
sudo firewall-cmd --list-ports
```
 
> `--permanent` saves the rule so it survives a reboot; `--reload` applies it immediately. Skip `--permanent` if you only need it open until the next reboot.
 
### iptables
 
**Open the port**
```bash
sudo iptables -A INPUT -p tcp --dport 8080 -j ACCEPT
```
 
**Close the port**
```bash
sudo iptables -D INPUT -p tcp --dport 8080 -j ACCEPT
```
 
**Verify**
```bash
sudo iptables -L -n | grep 8080
```
 
> iptables rules disappear after a reboot unless saved. On Debian/Ubuntu: `sudo apt install iptables-persistent && sudo netfilter-persistent save`. On RHEL-based systems it's usually simpler to just use `firewalld` instead.


## MakeFile

Run build make command with tests

```bash
make all
```

Build the application

```bash
make build
```

Run the application

```bash
make run
```

Live reload the application:

```bash
make watch
```

Run the test suite:

```bash
make test
```

Clean up binary from the last build:

```bash
make clean
```
