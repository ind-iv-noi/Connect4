import requests

API_URL = "http://86.121.150.134:8800/api"


def startGame():
    try:
        r = requests.get(url=API_URL + "/start")
    except TimeoutError as err:
        print("Server is closed")
        raise err
    data = r.json()
    return data["Response"][0]

def readGameTable(gID):
    Params={"gid":gID}
    try:
        r = requests.get(url=API_URL + "/read",params=Params)
    except TimeoutError as err:
        print("Server is closed")
        raise err
    data = r.json()
    return data["Table"]

def makeMove(gID,col,player):
    Params={"gid":gID,"move":col,"player":player}
    try:
        r = requests.post(url=API_URL + "/move",params=Params)
    except TimeoutError as err:
        print("Server is closed")
        raise err
    data = r.json()
    print(data)
    return data["Response"][1]

def getTurn(gID):
    Params = {"gid": gID}
    try :
        r = requests.get(url = API_URL + "/turn", params  = Params)
    except TimeoutError as err:
        print("Server is closed")
        raise err
    data = r.json()
    return data["Response"][0]
def deleteGame(gID):
    Params = {"gid": gID}
    try:
        r = requests.post(url=API_URL + "/delete", params=Params)
    except TimeoutError as err:
        print("Server is closed")
        raise err
def restartGame(gID):
    Params = {"gid": gID}
    try:
        r = requests.post(url=API_URL + "/restart", params=Params)
    except TimeoutError as err:
        print("Server is closed")
        raise err
if __name__ == '__main__' :
    gID=startGame()
    print(gID)

    print(readGameTable(gID))
    print(makeMove(gID,1,1))

    print(readGameTable(gID))