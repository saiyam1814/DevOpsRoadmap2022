How to login to a remote linux machine ?

- `ssh <user-name>@<ip-address>` then it will ask you for a password. Usually when you are launching an instance on cloud then they will not ask for password. Rather they ask for public key.

- `ssh -i <public-key> <user-name>@<ip-address>` After downloading the public key you can do this to get access to the instance.

- `exit` to get out of the cloud VM.

- `who` shows all the logged in users on the machine
- `w` also shows all the processes that are running on the system and it also shows the logged in users.