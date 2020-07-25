import tkinter
from tkinter.constants import *
from displaymatrix import main as rungame
def getGameID():
    entryID.get()
def MainScreen():
    root = tkinter.Tk()

    screen = tkinter.Canvas(root,width = 400, height = 300)
    screen.pack()

    button2 = tkinter.Button(screen, bg = "light blue",fg = "white", font = ('FreeSansBold.ttf',20),   text="Start Game", command=rungame)
    screen.create_window(200,100, window = button2)

    entryID = tkinter.Entry(root)
    screen.create_window(200,280,window = entryID)
    button1 = tkinter.Button(screen,bg = "light green",fg = "white",  font = ('FreeSansBold.ttf',20),  text="Join Game",  command=lambda : rungame(entryID.get()))
    screen.create_window(200,220, window = button1)
    screen.mainloop()


if __name__ == '__main__':
    MainScreen()
