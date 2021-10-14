#!/bin/bash

#Bash params
set -e
set -o errexit
set -o pipefail

# Default values
ARGS=("$*")
RED="\e[91m"
GREEN="\e[92m"
ENDCOLOR="\e[0m"
re_url='^https:\/\/github\.com\/[a-zA-Z0-9]{3,39}\/[a-zA-Z0-9\-]{3,40}\/{0,1}+$'
re_num='^[0-9]+$'
re_pname='^[0-9a-zA-Z_]+$'
re_oname='^[a-zA-Z_]+$'
flags_array=("h" "s" "p")
## help info function
# show info about application
helpmessage() {
  echo ""
  echo "Options:"
  echo " -h               help information, you already here"
  echo " -s               silent mode. Removs questions about unset'd\wrong values"
  echo " -p <url>         git hub url, example https://github.com/rubenlagus/TelegramBots"
  echo ""
  echo "examples"
  echo ""
  echo -e "sudo ./$(basename "${BASH_SOURCE[0]}") -p chrome -c 5 -b n -n a -w Organization "
  echo -e "sudo ./$(basename "${BASH_SOURCE[0]}") -p chrome -s "
  echo -e "sudo ./$(basename "${BASH_SOURCE[0]}") -p chrome "
  echo "TNX | P.S> Created by Juri Gogolev"
  exit 0
}
## Checking process option function
# if proces option not set, ask to set new one or setting default value
# depends on silent mode option
url_checker() {
  if [[ $SILENT ]]; then
    echo '[Warning] URL option not set, setting default value "https://github.com/rubenlagus/TelegramBots"'
    URL='https://github.com/rubenlagus/TelegramBots'
  else
    while ! [[ $URL =~ $re_url ]]; do
      read -p "[MSG] Enter github url or press [ENTER] for default value 'https://github.com/rubenlagus/TelegramBots': " URL
      if ! [ "$URL" ]; then
        URL='https://github.com/rubenlagus/TelegramBots'
      elif ! [[ $URL =~ $re_url ]]; then
        echo "[Warning] Cant set URL as '$URL', must be url"
      fi
    done
    echo "[INFO] URL setted to $URL"
  fi
}

###Checking sudo rights | sudo recomended
if [[ $EUID -ne 0 ]]; then
  echo -e "Rights     | [FAIL] | ${RED}No Sudo${ENDCOLOR}"
else
  echo -e "Rights     |  [OK]  | ${GREEN}Sudo${ENDCOLOR}"
  SUDOSTATE=1
fi
###Chechinkg Operation System | must be linux
unamer="$(uname -s)"
if [[ $unamer == Linux* ]]; then
  echo -e "OS         |  [OK]  | ${GREEN}$unamer${ENDCOLOR}"
else
  echo -e "OS         | [FAIL] | ${RED}$unamer${ENDCOLOR}"
  exit 1
fi
### Tools check [ (netstat | ss)!important | whois ]
if [ -z "$(which netstat)" ] && [ -z "$(which ss)" ]; then
  echo -e "Tools      | [FAIL] | ${RED}netstat and ss not installed${ENDCOLOR}"
  exit 1
elif [ -z "$(which netstat)" ]; then
  echo -e "Tools      |  [OK]  | ${GREEN}ss${ENDCOLOR}"
  UTILSS=1
elif [ -z "$(which ss)" ]; then
  echo -e "Tools      |  [OK]  | ${GREEN}netstat${ENDCOLOR}"
  UTILNETSTAT=1
else
  echo -e "Tools      |  [OK]  | ${GREEN}netstat${ENDCOLOR} & ${GREEN}ss${ENDCOLOR}"
  UTILSS=1
  UTILNETSTAT=1
fi
if [ -z "$(which whois)" ]; then
  echo -e "Whois      | [FAIL] | ${RED}not installed ${ENDCOLOR}"
else
  echo -e "Whois      |  [OK]  | ${GREEN}installed${ENDCOLOR}"
  UTILWHOIS=1
fi

#### checking flags of script
#      flag               |        descr
# -s silent mode          | show only errors if exists and result
# -h help                 | show help information
# -p url              | select proces by name or id [nubmer or string]
for item in ${flags_array[*]}; do
  TESTED=1
  unset OPTIND
  unset TMPOPTARG
  while getopts ":p:h:s:d:phsd" opt; do
    if [[ $opt == $item || ($opt == ':' && $OPTARG == $item) ]] && ! [[ $OPTARG =~ ^-[p/s/d/h]$ ]]; then
      if [[ $OPTARG && $opt != ':' ]]; then
        TESTED=3
        TMPOPTARG=$OPTARG
      else
        TESTED=2
      fi
    elif [[ ("-$item" == $OPTARG || ($opt == ':' && $OPTARG == $item)) && ($item == 's' || $item == 'h') ]]; then
      #    elif [[ ("-$item" == $OPTARG || "$item" == $OPTARG || $opt == $item) && ($item == 's' || $item == 'b' || $item == 'h') ]]; then
      TESTED=2
    fi
  done
  if [[ $item == 's' && ! $SUDOSTATE ]]; then
    echo '[Warning] Without sudo you will not get full info, recomended to start script using'
    if [ ! $SILENT ]; then
      echo -e "sudo ./$(basename "${BASH_SOURCE[0]}") ${ARGS[0]} "
    fi
  fi
  case "$item-$TESTED" in
  h-3 | h-2) helpmessage ;;
  s-3 | s-2)
    echo '[INGO] SILENT MODE IS SET'
    SILENT=1
    ;;
  p-1 | p-2) url_checker ;;
  p-3)
    if ! [[ $TMPOPTARG =~ $re_url ]]; then
      echo "[Warning] Cant set URL as '$TMPOPTARG', must url, look help"
      url_checker
    else
      URL=$TMPOPTARG
    fi
    ;;

  esac
done
#### Main code...
## I'm a litle tired now
MAINCOMMAND="$($(echo $USEDUTIL -tunap))"
if [ -z "${CONTYPE2}" ]; then
  CONTYPE2=$CONTYPE
fi
if [ $USEDUTIL == 'netstat' ]; then
  IPLIST="$(echo "$MAINCOMMAND" | grep -e $CONTYPE -e $CONTYPE2 | awk '/'"$PROCESS"/' {print $5}' | cut -d: -f1)"
else
  IPLIST="$(echo "$MAINCOMMAND" | grep -e $CONTYPE -e $CONTYPE2 | awk '/'"$PROCESS"/' {print $6}' | cut -d: -f1)"
fi
if [ -z "${IPLIST}" ]; then
  echo "Cant find any connections on $PROCESS process."
  exit 1
fi
echo -e "========================================================================"
if [ $UTILWHOIS ]; then
  echo -e "     IP\t\t|      Count""\t| Whois filter"
  echo -e "========================================================================"
else
  echo -e "     IP\t\t|      Count"
  echo -e "========================================================================"
fi
NETUTILRESULT="$(echo "$IPLIST" | cut -d: -f1 | sort | uniq -c | sort | tail -n$LIMIT)"
while IFS=', ' read -r str; do
  IP=$(echo $str | awk '{print $2}')
  COUNTER=$(echo $str | awk '{print $1}')
  if [ $UTILWHOIS ]; then
    ORG_NAME=$(whois $IP | awk -F':' '/'"$WHOISF"/' {print $2}')
    if [ -z "${ORG_NAME}" ]; then
      ORG_NAME='   Unknown'
    fi
    echo -e "$IP\t|\t$COUNTER\t| $ORG_NAME"
  else
    echo -e "$IP\t|\t$COUNTER"
  fi
done <<<"$NETUTILRESULT"
exit 1

