import pygame
pygame.init()
import API_interact
_images = {}
def getimage( path ):
    if path in _images:
        return _images[path]
    else:
        new_image = pygame.image.load(path).convert()
        _images[path] = new_image
        return new_image
import pdb
_font = {}
def getfont( path , size):
    if (path,size) in _font:
        return _font(path,size)
    else:
        new_font = pygame.font.Font(path, size)
        _font[(path,size)] = new_font
        return new_font
def printtext(screen, string , font = getfont('FreeSansBold.ttf',32)):
    text = font.render(string,True, (0,255,0), (255,0,0))
    screen.blit(text,(0,0))



class GameTable():
    _n = 6
    _m = 7
    mat = [[0 for i in range(7)] for j in range(6)]
    gID = ""
    playerID = 0
    def __init__ (self):
        self.gID = API_interact.startGame()
        self.playerID = 1
        print(self.gID)
    def __del__(self):
        API_interact.deleteGame(self.gID)
    def _input_matrix(self):
        with open("m.txt") as fin:
            for i in range(self._n):
                text = fin.readline()
                self.mat[i] =[int(x) for x in text.split()]

    def update_matrix(self):
        self.mat = API_interact.readGameTable(self.gID)
    def draw_matrix(self , screen):
        radius = 40
        row_incr = 150
        col_incr = 150
        start_pos = [50, 50]
        screen.fill((0,128,255))

        cur_pos = [50,50]
        for row in self.mat:
            for element in row:
                if element == 1:
                    pygame.draw.circle(screen, (255, 0, 0), cur_pos, radius)
                if element == 2:
                    pygame.draw.circle(screen, (255, 255, 0), cur_pos, radius)
                if element == 0:
                    pygame.draw.circle(screen, (211, 211, 211), cur_pos, radius)
                cur_pos[0] += col_incr

            cur_pos[0] = start_pos[0]
            cur_pos[1] += row_incr
        pygame.display.flip()

    def Move(self, col):
        if(col == 0):
            return
        try:
            r = API_interact.makeMove(self.gID, col, self.playerID)
        except:
            pdb.set_trace()
        return int(r)


def GetColumn(event):
    if(event.key == pygame.K_1):
        return 1
    if (event.key == pygame.K_2):
        return 2
    if (event.key == pygame.K_3):
        return 3
    if (event.key == pygame.K_4):
        return 4
    if (event.key == pygame.K_5):
        return 5
    if (event.key == pygame.K_6):
        return 6
    if (event.key == pygame.K_7):
        return 7
    return 0

def getPlayerId():
    return input("Player ID:")
def main():

    screen = pygame.display.set_mode((1000,950))
    screen.fill((0,128,255))
    clock = pygame.time.Clock()
    finish = False
    table = GameTable()
    table.update_matrix()
    game_screen = pygame.Surface((1000,850))
    last_time = pygame.time.get_ticks()
    cur_move = 0
    while not finish:
        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                finish = True
            if event.type == pygame.KEYDOWN :
                cur_move = GetColumn(event)
                if cur_move == 0:
                     continue
        if(pygame.time.get_ticks() - last_time >= 2000):
            table.Move(cur_move)
            cur_move = 0
            table.update_matrix()
            table.draw_matrix(game_screen)
            last_time = pygame.time.get_ticks()
        screen.blit(game_screen,(0,100))
        printtext(screen,'muie')
        pygame.display.flip()
        clock.tick(60)
    del table

main()