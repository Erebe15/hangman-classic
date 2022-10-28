# **Hangman Game**
<img src="https://media.giphy.com/media/riqiausX2TY6Lk7WhF/giphy.gif" width="5000" height=""/>

## **Summary**

 *Hangman game classic.*<br />

The goal of the game is to find all the letter of the word.<br />
If the letter you choose is correct, it will be added to the word. If not, the hangman will progress.

In this version you can :

+ Choose between different languages: English, French, Russian.
+ Select difficulty for more challenge.
+ Save your game progression by writing "stop".

_For a better experience, launch the game in full screen._<br />
## 🛠️ Installation

### With Gitea

```bash
git clone https://ytrack.learn.ynov.com/git/pthomas/hangman-classic.git
cd hangman-classic/
```

## Quick Start

Hangman can be run from Git Bash, it will create a new game.<br />
_CLI Example_ :

```bash
go run . words.txt
```
You can also launch the game from a save.
_CLI Example_ :
````bash
go run . --startwith <NameOfTheSave>
````
## Created by :

- ANDRIEUX Rodolphe
- PIET Thomas


