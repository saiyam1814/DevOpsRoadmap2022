Sometimes we want to do some specific task at a pre-determined time in future and for that we use crontabs.
- It's also helpful in managing your CI/CD workflows. When we want to run some test at a particular time then we can use crontabs.
- It's helpful in managing your github action workflows and jenkins pipelines.
- Also if you want not to forget someone's birthday you can make a remainder using this. 

## Commands
- to look at the system cron jobs use `cat /etc/crontab`
- to add a new cron job use `crontab -e` 
- to remove all the crontab use `crontab -r` for the active user



Common time frames:

    ┌───────────── minute (0 - 59)
    │ ┌───────────── hour (0 - 23)
    │ │ ┌───────────── day of month (1 - 31)
    │ │ │ ┌───────────── month (1 - 12)
    │ │ │ │ ┌───────────── day of week (0 - 6) (Sunday to Saturday;
    │ │ │ │ │                                       7 is also Sunday on some systems)
    │ │ │ │ │
    │ │ │ │ │
    * * * * *  command_to_execute

### Arguments

    *	any value
    ,	value list separator (to seperate two entries)
    -	range of values
    /	step values
    @yearly	(non-standard)
    @annually	(non-standard)
    @monthly	(non-standard)
    @weekly	(non-standard)
    @daily	(non-standard)
    @hourly	(non-standard)
    @reboot	(non-standard)

## Examples
- `*/10 * * * * echo "hello world" >> /home/kranurag7/demo.md` this will append hello world every 10 minutes to the file demo.md
- `30 6 1,15 * * echo "hello world" >> /home/kranurag7/hello.txt` this will append only on 1st and 15th of the month at 6:30 in the morning
- Make sure that you leave comment along with your crontab in case you forget it in future.