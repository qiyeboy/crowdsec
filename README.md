

<p align="center"> :warning: <b>Crowdsec BETA </b> :warning: </p>

<p align="center">
<img src="docs/assets/images/crowdsec_logo1.png" alt="CrowdSec" title="CrowdSec" />
</p>


<p align="center">
<img src="https://github.com/crowdsecurity/crowdsec/workflows/tests/badge.svg">
<img src="https://github.com/crowdsecurity/crowdsec/workflows/build/badge.svg">
<a href='https://coveralls.io/github/crowdsecurity/crowdsec?branch=master'><img src='https://coveralls.io/repos/github/crowdsecurity/crowdsec/badge.svg?branch=master' alt='Coverage Status' /></a>
<img src="https://goreportcard.com/badge/github.com/crowdsecurity/crowdsec">
<img src="https://img.shields.io/github/license/crowdsecurity/crowdsec">
<img src="https://github.com/crowdsecurity/crowdsec/workflows/Hub-CI/badge.svg">
</p>

<p align="center">
:books: <a href="https://doc.crowdsec.net">Documentation</a>
:diamond_shape_with_a_dot_inside: <a href="https://hub.crowdsec.net">Hub</a>
:speech_balloon: <a href="https://discourse.crowdsec.net">Discourse Forum</a>
:speech_balloon: <a href="https://gitter.im/crowdsec-project/community?utm_source=share-link&utm_medium=link&utm_campaign=share-link">Gitter Chat</a>
</p>

## <TL;DR>

A modern behavior detection system, written in Golang. It stacks on Fail2ban's philosophy, but uses Grok patterns & YAML grammar to analyse logs, a modern decoupled approach (detect here, remedy there) for Cloud/Containers/VM based infrastructures. Once detected you can remedy threats with various bouncers (block, 403, Captchas, etc.) and the blocked IPs are shared among all users to further improve their security.

## About the crowdsec project

Crowdsec is an open-source, lightweight software, detecting peers with aggressive behaviors to prevent them from accessing your systems. Its user friendly design and assistance offers a low technical barrier of entry and nevertheless a high security gain.

Processing is done in 5 steps:
 1. Read Data sources (log files, streams, trails, messages ...), normalize and enrich signals
 2. Matching those signals to behavior patterns, aka scenarios (*)
 3. If an unwanted behavior is detected, deal with it through a [bouncer](https://hub.crowdsec.net/browse/#bouncers) : a software component integrated into your applicative stack that supports various remediations such as block, return 403, and soon captcha, 2FA, etc.
 4. *(ONLY)* The aggressive IP, the scenario name triggered and a timestamp is then sent to our curation platform (to avoid poisoning & false positives)
 5. If verified, this IP is then integrated to the block list continuously distributed to all CrowdSec clients (which is used as an enrichment source in step1)

By detecting, blocking and sharing the threat they faced, all clients are reinforcing each-others (hence the name Crowd-Security). Crowdsec is designed for modern infrastructures, with its "*Detect Here, Remedy There*" approach, letting you analyse logs coming from several sources in one place and block threats at various levels (applicative, system, infrastructural) of your stack.

(*) CrowdSec ships by default with scenario (brute force, port scan, web scan, etc.) adapted for most context, but you can easily extend it by picking more of them from the [hub](https://hub.crowdsec.net). It is also very easy to adapt an existing one or create one yourself.

## What it is not

CrowdSec is not a SIEM, storing your logs (neither locally nor remotely).

Your data stay in your premices and are only analyzed and forgotten.

Signals sent to the curation platform are extremely limited (IP, Scenario, Timestamp), and are only there to allow the system to rule out false positives or poisoning attemps.


## Install it !

Find the [latest release](https://github.com/crowdsecurity/crowdsec/releases/latest)

Ensure you have dependencies :
<details open>
  <summary>for Debian based distributions</summary>

```bash
apt-get install bash gettext whiptail curl wget
```
</details>

<details>
  <summary>for RedHat based distributions</summary>

```bash
yum install bash gettext newt curl wget
 ```
</details>

```bash
curl -s https://api.github.com/repos/crowdsecurity/crowdsec/releases/latest | grep browser_download_url| cut -d '"' -f 4  | wget -i -
tar xvzf crowdsec-release.tgz
cd crowdsec-v*
sudo ./wizard.sh -i
```



## Key points

### Fast assisted installation, no technical barrier

<details open>
  <summary>User is assisted during setup, providing functional out-of-the-box setup</summary>
  <img src="https://github.com/crowdsecurity/crowdsec/blob/master/docs/assets/images/crowdsec_install.gif">
</details>

### Out of the box detection

<details>
  <summary>Baseline detection is effective out-of-the-box, no fine-tuning required (click to expand)</summary>
  <img src="https://github.com/crowdsecurity/crowdsec/blob/master/docs/assets/images/out-of-the-box-protection.gif">
</details>

### Easy bouncer deployment

<details>
  <summary>It's trivial to add bouncers to enforce decisions of crowdsec (click to expand)</summary>
  <img src="https://github.com/crowdsecurity/crowdsec/blob/master/docs/assets/images/blocker-installation.gif">
</details>

### Easy dashboard access

<details>
  <summary>It's easy to deploy a metabase interface to view your data simply with cscli (click to expand)</summary>
  <img src="https://github.com/crowdsecurity/crowdsec/blob/master/docs/assets/images/cscli-metabase.gif">
</details>

## About this repository

This repository contains the code for the two main components of crowdsec :
 - `crowdsec` : the daemon a-la-fail2ban that can read, parse, enrich and apply heuristis to logs. This is the component in charge of "detecting" the attacks
 - `cscli` : the cli tool mainly used to interact with crowdsec : ban/unban/view current bans, enable/disable parsers and scenarios.

## :warning: Beta version

Please note that crowdsec is currently in beta version, use with caution !


