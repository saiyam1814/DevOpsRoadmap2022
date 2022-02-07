# File Transfer

## FTP

**F**ile **T**ransfer **P**rotocol (**FTP**) is a well-known and popular method for transferring files between computers using the Internet. This method is built on a client-server model. FTP can be used within a browser or with stand-alone client programs. It is considered inadequate for modern needs, as well as being intrinsically insecure.

## FTP Clients

FTP clients enable you to transfer files with remote computers using the FTP protocol. All web browsers support FTP, all you have to do is give a URL like **ftp://ftp.kernel.org** where the usual **http://** becomes **ftp://**.

Some command line FTP clients are:

- ftp
- sftp
- ncftp
- yafc

FTP has fallen into disfavor on modern systems, as it is intrinsically insecure. As an alternative, **sftp** is a very secure mode of connection, which uses the Secure Shell (**ssh**) protocol.

## SSH

Secure Shell (SSH) is a cryptographic network protocol used for secure data communication. To login to a remote system using your same user name you can just type ssh some_system and press Enter. ssh then prompts you for the remote password.

If you want to run as another user, you can do either **ssh -l someone some_system** or **ssh someone@some_system**. To run a command on a remote system via SSH, at the command prompt, you can type **ssh some_system my_command**.

## Copying files securely with scp

We can also move files securely using Secure Copy (scp) between two networked hosts. scp uses the SSH protocol for transferring data.

To copy a local file to a remote system, at the command prompt, type scp <localfile> [user@remotesystem](mailto:user@remotesystem):/home/user/ and press Enter.
