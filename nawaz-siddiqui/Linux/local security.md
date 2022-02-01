# Local Security Principles

**Types of Accounts**

By default, Linux distinguishes between several account types in order to isolate processes and workloads. Linux has four types of accounts:

- root
- System
- Normal
- Network

**Comparing sudo and su**

| su | sudo |
| --- | --- |
| When elevating privilege, you need to enter the root password. Giving the root password to a normal user should never, ever be done. | When elevating privilege, you need to enter the user’s password and not the root password. |
| Once a user elevates to the root account using su, the user can do anything that the root user can do for as long as the user wants, without being asked again for a password.  | Offers more features and is considered more secure and more configurable. Exactly what the user is allowed to do can be precisely controlled and limited. By default the user will either always have to keep giving their password to do further operations with sudo, or can avoid doing so for a configurable time interval. |
| The command has limited logging features. | The command has detailed logging features. |

**sudo Features**

**sudo** has the ability to keep track of unsuccessful attempts at gaining root access. Users' authorization for using **sudo** is based on configuration information stored in the **/etc/sudoers** file and in the **/etc/sudoers.d** directory.

## Working with Passwords

**How Passwords Are Stored**

Originally, encrypted passwords were stored in the **/etc/passwd** file, which was readable by everyone. This made it rather easy for passwords to be cracked.

On modern systems, passwords are actually stored in an encrypted format in a secondary file named **/etc/shadow**. Only those with root access can read or modify this file.

**Good Password Practices**

IT professionals follow several good practices for securing the data and the password of every user.

- Password aging is a method to ensure that users get prompts that remind them to create a new password after a specific period. This can ensure that passwords, if cracked, will only be usable for a limited amount of time. This feature is implemented using **chage**, which configures the password expiry information for a user.
- Another method is to force users to set strong passwords using **P**luggable **A**uthentication **M**odules (**PAM**). PAM can be configured to automatically verify that a password created or modified using the **passwd** utility is sufficiently strong. PAM configuration is implemented using a library called **pam_cracklib.so**, which can also be replaced by **pam_passwdqc.so** to take advantage of more options.
- One can also install password cracking programs, such as [John The Ripper](http://www.openwall.com/john/), to secure the password file and detect weak password entries. It is recommended that written authorization be obtained before installing such tools on any system that you do not own.

**Requiring Boot Loader Passwords**

For the older GRUB 1 boot method, it was relatively easy to set a password for **grub**. ****However, for the GRUB 2 version, things became more complicated. However, you have more flexibility.

Furthermore, you never edit **grub.cfg** directly; instead, you can modify the configuration files in **/etc/grub.d** and **/etc/defaults/grub**, and then run **update-grub**, or **grub2-mkconfig** and save the new configuration file.