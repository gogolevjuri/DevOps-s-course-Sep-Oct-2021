### Netstat/ss script

##Task
![image](https://r.elite.ovh/dev/s1.jpg)
***
All parts of task done ... I hope :D
***
```
Options:
 -h               help information, you already here
 -s               silent mode. Removs questions about unset'd\wrong values
 -p <process>     process_name or process identificator (PID). Default process is 'firefox'.
 -c <number>      number of lines to display. Default value is 5.
 -d               disabling whois option
 -b <ss|netstat>  switch 'netstat' \  'cc'. Default backend is 'netstat'.
 -w <string>      select desired object from 'whois' output. Default object is 'Organization'.
 -n <l|e|a>       types of connections would you like to see. e -  'ESTABLISHED',l -  'LISTENING', a- 'ALL

examples

sudo ./main.sh -p chrome -c 5 -b n -n a -w Organization 
sudo ./main.sh -p chrome -s 
sudo ./main.sh -p chrome 
sudo ./main.sh -p firefox -c 5 -b n -n a -w Organization 
sudo ./main.sh -p firefox -s 
sudo ./main.sh -p firefox
```
