# exe-link
Simple Binary link maker for Windows, especially MINGW.

## Installation

``` go get github.com/wirekang/exe-link ```

You need **Go** because this program generates lightweight go code and builds it.


## Usage

```exe-link <source> <destination>```

Similar with ```ln -s <source> <destination>```  

It creates ```<destination>.exe``` file that executes ```<source>``` file. You can input both arguments in relative path, the program will convert source path to absolute path.

## Example
```exe-link Neovim/bin/nvim.exe vim```
Output filename is **vim.exe**. Now you can execute nvim.exe by typing ```vim``` in **cmd, powershell, MINGW** and so on with any arguments.

## Motivation
In unix system, we can set alias of commands via **ln**. For example, we can execute neovim by type vim or nvim after link **vim** to **nvim**

``` ln -s /usr/local/nvim /usr/local/vim ```

In windows, There is similar command **mklink**, but they don't work in CLI.
Excuting link from gui(explorer) work as fine, but meaningless. Shortcuts go without saying.  

There is one cmd-only way. we can make **.bat** or **.cmd** files to execute certain file with arguments.

%PATH%\ **crm.bat**
```
"C:\ProgramFiles....chrome.exe" %*
```  

Now ```crm``` works fine(open chrome) in **cmd or powershell** Because they infer **crm.bat** or **crm.cmd** or **crm.exe** from **crm**.
But in MINGW, they don't infer .bat or .cmd automatically unlike **.exe**. We should specify extensions on every execution.
```
$ crm
bash: crm: command not found

$ crm.bat
C:\Users\dmhsk\workspace\exe-link>"C:\Program Files\Google\Chrome\Application\chrome.exe"
    // works
```
  
In my case, I want to use neovim in IntelliJ Terminal by **vim** command. I could do that by renaming nvim.exe to vim.exe, but that's not a programmer's way. So I made this simple program.

