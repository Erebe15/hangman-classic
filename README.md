# **Hangman Game**
<img src="https://media.giphy.com/media/riqiausX2TY6Lk7WhF/giphy.gif" width="5000" height=""/>

## **Summary**

 *Hangman game classic.*<br />

The goal of the game is to find a randomly selected word letter by letter.<br />
If the word contains the chosen letter, it will be displayed in the word. If not, the hangman draw will progress by one step.<br />
There are only 10 steps, so you have to choose the letter wisely if you dont want *José* to get hanged.<br /><br />

If you think you've found the word, you can guess the entire word instead of just one letter. But be careful... If this is not right word the draw will progress by 2 steps. 

In this version you can also:

+ Choose between different languages: English, French and even Russian.
+ Select difficulty for more challenge. (maybe to much)
+ Save your game progression and continue later.

## 🛠️ Installation and requirements ⚠

### With Gitea

```bash
git clone https://ytrack.learn.ynov.com/git/pthomas/hangman-classic.git
cd hangman-classic/
```
### Then with GitBash ⚠
⚠ _For an optimal experience set the font size at 8pt (the minimum) with a black background_<br />
⚠ _To access the menu, launch the game with a terminal window **at least** 156 wide by 47 high._<br />
note that you can change the size of the window at any moment in the game!

## Start

Hangman can be run from Git Bash with the basic command `go run .`, it will create a new game by default in english with the easiest difficulty.<br />
(You can modify those later in the Main Menu → Options).<br /><br />
Also, you have the possibility to choose the language of the words to find by launching the command `go run . mots.txt`, `go run . words.txt` or `go run . слова.txt`.<br />
The game will automatically adapt the language to the list given.<br /><br />
As an example, if you launch the command:
```bash
go run . mots.txt
```
all the game will be in **French**.<br /><br />
## Saves :
You can stop and save a game at any time by typing "stop" instead of choosing a letter.<br />
You can then give the name of the save and the game will quit.<br /><br />
To launch the save you can select "saves" in the main menu and then select your save.<br />
*Be careful, by selecting a save and playing, it will delete the save file. If you want to keep it you will have to type "stop" again*<br /><br />
You can also launch a save directly from Bash with the flag `-startWith` or `-s` followed by the save name.<br /><br />
_Examples_ :
```bash
go run . --startWith <NameOfTheSave> 
```
```bash
go run . -s <NameOfTheSave> 
```
```bash
go run . -startWith=<NameOfTheSave> 
```

## Created by :

- [ANDRIEUX Rodolphe](https://www.youtube.com/watch?v=dQw4w9WgXcQ)
- [PIET Thomas](https://www.youtube.com/watch?v=dQw4w9WgXcQ)


