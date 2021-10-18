## Flask app with Ansible deployment

### Tasks to do:
![image](https://r.elite.ovh/dev/s2.jpg)
> more info about task https://r.elite.ovh/dev/s3.jpg

Copy this repo with this task.
This task consists of:
```
.
├── animal_app.py                         ### Main python file, if u want check code
├── ansible                               ### Ansible homework dir
│   ├── ansible.cfg                       ### Ansible config file
│   ├── deploy.sh                         ### Starter script
│   ├── inventory 
│   │   └── hosts                         ### Inventory\conf file
│   ├── main.yml                          ### Main Playbook 
│   ├── README.md                         ### Readme here
│   └── roles
│       ├── config_app
│       │   └── tasks
│       │       └── main.yml              ### Requirements install Role
│       ├── config_daemon
│       │   ├── files
│       │   │   ├── projectproxy.service  ### Service file for proxy scipt (https_proxy.py)
│       │   │   └── project.service       ### Service file for main scipt (animal_app.py)   
│       │   └── tasks
│       │       └── main.yml              
│       ├── config_firewall
│       │   └── tasks
│       │       └── main.yml              ### Creating services Role
│       ├── config_ssl
│       │   ├── tasks
│       │   │   └── main.yml              ### Create self signed ssl cert Role
│       │   └── templates
│       │       └── project.j2            ### Template for ssl
│       ├── install_app
│       │   └── tasks
│       │       └── main.yml              ### Install app From REPO Role
│       ├── install_python
│       │   └── tasks
│       │       └── main.yml              ### Install python,libssl,git,etc. Role
│       └── update_system
│           └── tasks
│               └── main.yml              ### Set timezone, Update hosts, Update system Role
├── https_proxy.py                        ### My own proxy script
├── README.md                             ###README.MD u're here now
└── templates
    ├── 404                               ### Html template for 404 error of animal_app.py
    └── game                              ### Browser main page template of animal_app.py



20 directories, 20 files 
```
***
My app work ( animal_app.py ) created with Flask and emoji support. The service listens at least on port 80 and accepts GET and POST methods.
Also I created second app ( https_proxy.py ) what listens at port 443, with self signed cert ( u need create `ssl` dir, and place here private.pem && cert.crt )
> Starting this both apps allows u to use 443(ssl) and 80 port in same time [Sudo recomended.]

> curl -XPOST -d'{"animal":"cow", "sound":"moooo", "count": 3}' http://myvm.localhost/
***
If u want check ansible playbook + app, just run deploy.sh
***
