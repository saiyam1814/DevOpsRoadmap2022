# Printing

To manage printers and print directly from a computer or across a networked environment, you need to know how to configure and install a printer. Printing itself requires software that converts information from the application you are using to a language your printer can understand. The Linux standard for printing software is the Common UNIX Printing System (CUPS).

Modern Linux desktop systems make installing and administering printers pretty simple and intuitive, and not unlike how it is done on other operating systems.

CUPS is the underlying software almost all Linux systems use to print from applications like a web browser or LibreOffice. It converts page descriptions produced by your application (put a paragraph here, draw a line there, and so forth) and then sends the information to the printer. It acts as a print server for both local and network printers.

Printers manufactured by different companies may use their own particular print languages and formats. CUPS uses a modular printing system which accommodates a wide variety of printers and also processes various data formats. This makes the printing process simpler; you can concentrate more on printing and less on how to print.

Generally, the only time you should need to configure your printer is when you use it for the first time. In fact, CUPS often figures things out on its own by detecting and configuring any printers it locates.

**How does CUPS work?**

CUPS carries out the printing process with the help of its various components:

- Configuration files
- Scheduler
- Job files
- Log files
- Filter
- Printer drivers
- Backend.

**Scheduler**

CUPS is designed around a print scheduler that manages print jobs, handles administrative commands, allows users to query the printer status, and manages the flow of data through all CUPS components.

**Configuration FIles**

The print scheduler reads server settings from several configuration files, the two most important of which are **cupsd.conf** and **printers.conf**. These and all other CUPS related configuration files are stored under the **/etc/cups/** directory.

**cupsd.conf** is where most system-wide settings are located; it does not contain any printer-specific details. Most of the settings available in this file relate to network security, i.e. which systems can access CUPS network capabilities, how printers are advertised on the local network, what management features are offered, and so on.

**printers.conf** is where you will find the printer-specific settings. For every printer connected to the system, a corresponding section describes the printer’s status and capabilities. This file is generated or modified only after adding a printer to the system, and should not be modified by hand.

You can view the full list of configuration files by typing **ls -lF /etc/cups**.

**Job Files**

CUPS stores print requests as files under the **/var/spool/cups** directory (these can actually be accessed before a document is sent to a printer). Data files are prefixed with the letter **d** while control files are prefixed with the letter **c**.

After a printer successfully handles a job, data files are automatically removed. These data files belong to what is commonly known as ****the **print queue**.

**Log Files**

Log files are placed in **/var/log/cups** and are used by the scheduler to record activities that have taken place. These files include access, error, and page records.

To view what log files exist, type:**$ sudo ls -l /var/log/cups**

**Filters, Printer Drivers, and Backends**

CUPS uses **filters** to convert job file formats to printable formats. **Printer drivers** contain descriptions for currently connected and configured printers, and are usually stored under **/etc/cups/ppd/**. The print data is then sent to the printer through a filter, and via a **backend** that helps to locate devices connected to the system.

So, in short, when you execute a print command, the scheduler validates the command and processes the print job, creating job files according to the settings specified in the configuration files. Simultaneously, the scheduler records activities in the log files. Job files are processed with the help of the filter, printer driver, and backend, and then sent to the printer.

**Managing CUPS**

Assuming CUPS has been installed you'll need to start and manage the CUPS daemon so that CUPS is ready for configuring a printer. Managing the CUPS daemon is simple; all management features can be done with the **systemctl** utility:

**$ systemctl status cups**

**$ sudo systemctl [enable|disable] cups**

**$ sudo systemctl [start|stop|restart] cups**

**Adding Printers from the CUPS Web Interface**

A fact that few people know is that CUPS also comes with its own web server, which makes a configuration interface available via a set of CGI scripts.

The CUPS web interface is available on your browser at: [http://localhost:631](http://localhost:631/).

**Printing from the Command-Line Interface**

**lp** and **lpr** accept ****command line options that help you perform all operations that the GUI can accomplish. **lp** is typically used with a file name as an argument.

| Command | Usage |
| --- | --- |
| lp <filename> | To print the file to default printer |
| lp -d printer <filename> | To print to a specific printer (useful if multiple printers are available) |
| program | lpecho string | lp | To print the output of a program |
| lp -n number <filename> | To print multiple copies |
| lpoptions -d printer | To set the default printer |
| lpq -a | To show the queue status |
| lpadmin | To configure printer queues |
| lpstat -p -d | To get a list of available printers, along with their status |
| lpstat -a | To check the status of all connected printers, including job numbers |
| cancel job-id OR lprm job-id | To cancel a print job |
| lpmove job-id newprinter | To move a print job to new printer |

**Managing Print Jobs**

In Linux, command-line print job management commands allow you to monitor the job state as well as managing the listing of all printers and checking their status, and canceling or moving print jobs to another printer.

## Manipulating Postscript and PDF Files

PostScript is a standard page description language. It effectively manages scaling of fonts and vector graphics to provide quality printouts.

Postscript has been for the most part superseded by the PDF format (Portable Document Format) which produces far smaller files in a compressed format for which support has been integrated into many applications. However, one still has to deal with postscript documents, often as an intermediate format on the way to producing final documents.

**Working with enscript**

**enscript** is a tool that is used to convert a text file to PostScript and other formats. It also supports Rich Text Format (RTF) and HyperText Markup Language (HTML). For example, you can convert a text file to two columns (**-2**) formatted PostScript using the command:

**$ enscript -2 -r -p psfile.ps textfile.txt**

This command will also rotate (**-r**) the output to print so the width of the paper is greater than the height (aka landscape mode) thereby reducing the number of pages required for printing.

| Command | Usage |
| --- | --- |
| enscript -p psfile.ps textfile.txt | Convert a text file to PostScript (saved to psfile.ps) |
| enscript -n -p psfile.ps textfile.txt | Convert a text file to n columns where n=1-9 (saved in psfile.ps) |
| enscript textfile.txt | Print a text file directly to the default printer |

**Converting between PostScript and PDF**

| Command | Usage |
| --- | --- |
| pdf2ps file.pdf | Converts file.pdf to file.ps |
| ps2pdf file.ps | Converts file.ps to file.pdf |
| pstopdf input.ps output.pdf | Converts input.ps to output.pdf |
| pdftops input.pdf output.ps | Converts input.pdf to output.ps |
| convert input.ps output.pdf | Converts input.ps to output.pdf |
| convert input.pdf output.ps | Converts input.pdf to output.ps |

**Using qpdf**

| Command | Usage |
| --- | --- |
| qpdf --empty --pages 1.pdf 2.pdf -- 12.pdf | Merge the two documents 1.pdf and 2.pdf. The output will be saved to 12.pdf. |
| qpdf --empty --pages 1.pdf 1-2 -- new.pdf | Write only pages 1 and 2 of 1.pdf. The output will be saved to new.pdf. |
| qpdf --rotate=+90:1 1.pdf 1r.pdf | Rotate page 1 of 1.pdf 90 degrees clockwise and save to 1r.pdf. |
| qpdf --encrypt mypw mypw 128 -- public.pdf private.pdf | Encrypt with 128 bits public.pdf using as the passwd mypw with output as private.pdf. |
| qpdf --decrypt --password=mypw private.pdf file-decrypted.pdf | Decrypt private.pdf with output as file-decrypted.pdf. |
| qpdf --rotate=+90:1-z 1.pdf 1r-all.pdf | Rotate all pages of 1.pdf 90 degrees clockwise and save to 1r-all.pdf |

**Using Ghostscript**

Ghostscript is widely available as an interpreter for the Postscript and PDF languages. The executable program associated with it is abbreviated to gs.

This utility can do most of the operations pdftk can, as well as many others; see man gs for details. Use is somewhat complicated by the rather long nature of the options. For example:

- Combine three PDF files into one:
    
    **$ gs -dBATCH -dNOPAUSE -q -sDEVICE=pdfwrite  -sOutputFile=all.pdf file1.pdf file2.pdf file3.pdf**
    
- Split pages 10 to 20 out of a PDF file:
    
    **$ gs -sDEVICE=pdfwrite -dNOPAUSE -dBATCH -dDOPDFMARKS=false -dFirstPage=10 -dLastPage=20\-sOutputFile=split.pdf file.pdf**
