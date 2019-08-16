# zabbix-agent-extension-php-fpm

zabbix-agent-extension-php-fpm - this extension for monitoring PHP-FPM.

### Installation

```sh
# build the executable and install to /usr/bin/
go build -o /usr/bin/zabbix-agent-extension-php-fpm

# install the config file
cp zabbix-agent-extension-php-fpm.conf /etc/zabbix/zabbix_agentd.conf.d/
```

### Arguments

examples:

```
# full arguments
php-fpm.stats[listen queue,tcp,127.0.0.1:9000,/stats]

# partial arguments
php-fpm.stats[listen queue]
php-fpm.stats[listen queue,unix,/var/run/php-fpm.sock]
```

### Dependencies

zabbix-agent-extension-php-fpm requires [zabbix-agent](http://www.zabbix.com/download) v2.4+ to run.
