<img align="right" src="https://github.com/piqoni/hn-text/actions/workflows/test.yml/badge.svg">
<h1 align="center"> HN-text </h1>
<p align="center"> A fast, easy-to-use and distraction-free Hacker News terminal client.</p>

<div align="center"> <img width="500" src="https://github.com/piqoni/hn-text/assets/3144671/339fe90f-29e8-4e58-b185-dfa9ce86464d"/> </div>

# Motivations: 
 - Easy to use (arrow keys or hjkl navigation should be enough for the client to be fully usable)
 - Distraction Free: articles, and comments are converted to simple readable text. 
 - Fast Navigation and Responsivity

# Current Features / Screenshots
 - Navigation and opening pages (text-version): ←↓↑→ arrow keys (or hjkl) will navigate from the HN Frontpage → Comments Page → Article's Text and back.
 - Open article in default's browser (**SPACE** key), Comment page ('**c**' key). 
 - Append "best" as argument if you want to see Hacker News Best page, instead of the default frontpage. 
   
 ## Frontpage
<img width="641" alt="image" src="https://github.com/piqoni/hn-text/assets/3144671/92beba8d-1a44-400a-8f0c-a3372a221d58">
 
 ## Comments
 <img width="787" alt="image" src="https://github.com/piqoni/hn-text/assets/3144671/fca7672a-d7a5-4e70-a636-95595b58d5ba">

 ## Article
 <img width="940" alt="image" src="https://github.com/piqoni/hn-text/assets/3144671/c4a6d098-7f79-4c81-8cd7-0506fe6aab23">
 
# Keymaps
<div align="center">
    <table >
     <tr>
        <td><b>Key</b></td>
        <td><b>Functionaly</b></td>
     </tr>
          <tr>
       <td> Down Arrow (↓) or `j` </td>
       <td> Down on the Frontpage Article List</td>
     </tr>
               <tr>
       <td> UP Arrow (↑) or `k` </td>
       <td> Up on the Frontpage Article List</td>
     </tr>
     <tr>
       <td> Right Arrow (→) or `l` </td>
       <td>Open Comment Page (while on frontpage) - Pressing again would open the article</td>
     </tr>
      <tr>
       <td> Left Arrow (←) or `h` </td>
       <td>Go Back</td>
     </tr>
      <tr>
       <td> SPACE </td>
       <td>Open Article on Browser (if for some reason not satisfied with text rendered version)</td>
     </tr>
      <tr>
       <td> `c` </td>
       <td>Open Comments page on Browser</td>
     </tr>
           <tr>
       <td> `q` </td>
       <td>Quit App</td>
     </tr>
    </table>
    </div>
    
# Installation
## Binaries
Download binaries for your OS at [release page](https://github.com/piqoni/hn-text/releases), and chmod +x the file to allow execution. 

## Using GO INSTALL
If you use GO, you can install it directly:
```
go install github.com/piqoni/hn-text@latest
```
