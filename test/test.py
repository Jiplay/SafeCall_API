import requests, json


def register():
    url = "http://localhost:8080/register"

    payload = json.dumps({
        "Login": "testguy01",
        "Password": "safePasswo0d*",
        "Email": "jujujumail@gmail.cim",
    })

    headers = {
        'Content-Type': 'application/json'
    }

    response = requests.request("POST", url, headers=headers, data=payload)

    return response.text


def registerFriends():
    emails = ['jujujumail@gmail.cim', 'jujujumail@gmail.com']
    logins = ['testguy01', 'testguy02']
    i = 0

    while i != 2:
        url = "http://localhost:8080/register"
        payload = json.dumps({
            "Login": logins[i],
            "Password": "safePasswo0d*",
            "Email": emails[i],
            })

        headers = {
        'Content-Type': 'application/json'
        }
        requests.request("POST", url, headers=headers, data=payload)
        i += 1


def deleteFriends():
    logins = ['testguy01', 'testguy02']
    i = 0

    while i != 2:
        url = "http://localhost:8080/delete/"
        payload = json.dumps({
            "UserID": logins[i],
            })

        headers = {
        'Content-Type': 'application/json'
        }
        requests.request("POST", url, headers=headers, data=payload)
        i += 1


def login():
    url = "http://localhost:8080/login/testguy01/safePasswo0d*"

    payload = {}
    headers = {}

    response = requests.request("GET", url, headers=headers, data=payload)

    return response.text


def newIdslogin():
    url = "http://localhost:8080/login/testguy01/PassSaf3à*"

    payload = {}
    headers = {}

    response = requests.request("GET", url, headers=headers, data=payload)

    return response.text


def getProfile():
    url = "http://localhost:8080/profile/testguy01/"

    payload = {}
    headers = {}

    response = requests.request("GET", url, headers=headers, data=payload)

    return response.text


def delete_user():
    url = "http://localhost:8080/delete"
    payload = json.dumps({
        "UserID": "testguy01",
    })

    headers = {
        'Content-Type': 'application/json'
    }

    response = requests.request("POST", url, headers=headers, data=payload)

    return response.text  


def update_profile():
    endpoints = ['profileDescription', 'profileFullName', 'profilePhoneNB', 'profileEmail']
    data = ['MyTestDescription', 'MyFullNameTest', 'TestPhoneNB', 'myTest@email.com']
    i = 0
    while i < 4:
        url = "http://localhost:8080/"+endpoints[i]

        payload = json.dumps({
            "UserID": "testguy01",
            "Data": data[i]
        })

        headers = {
        'Content-Type': 'application/json'
        }

        requests.request("POST", url, headers=headers, data=payload)
        i += 1


def addFriend(action):
    url = "http://localhost:8080/manageFriend"

    payload = json.dumps({
        "UserID": "testguy01",
        "Friend": "testguy02",
        "Action": action
    })

    headers = {
    'Content-Type': 'application/json'
    }

    response = requests.request("POST", url, headers=headers, data=payload)

    return response.text


def replyFriend(action):
    url = "http://localhost:8080/replyFriend"

    payload = json.dumps({
        "UserID": "testguy02",
        "Friend": "testguy01",
        "Action": action
    })

    headers = {
    'Content-Type': 'application/json'
    }
    response = requests.request("POST", url, headers=headers, data=payload)

    return response.text


def listFriends(user):
    url = "http://localhost:8080/listFriends/"+ user

    payload = {}
    headers = {}

    response = requests.request("GET", url, headers=headers, data=payload)

    return response.text


def addEvent():
    url = "http://localhost:8080/addEvent"

    payload = json.dumps({
        "Guest1": "testguy01",
        "Guest2": "testguy02",
        "Subject": "bicyle",
        "Date": "Demain soir"
    })
    headers = {
        'Content-Type': 'application/json'
    }

    response = requests.request("POST", url, headers=headers, data=payload)

    return response.text


def listEvent(user):
    url = "http://localhost:8080/listEvent/"+ user

    payload = {}
    headers = {}

    response = requests.request("GET", url, headers=headers, data=payload)

    return response.text


def editPassword():
    url = "http://localhost:8080/editPassword"

    payload = json.dumps({
        "UserID": "testguy01",
        "PasswordOld": "safePasswo0d*",
        "PasswordNew": "PassSaf3à*"
    })
    headers = {
        'Content-Type': 'application/json'
    }

    response = requests.request("POST", url, headers=headers, data=payload)

    return response.text


def addNotification():
    url = "http://localhost:8080/AddNotification"

    payload = {
        "UserID": "testguy01",
        "Title": "bienvenue",
        "Content": "bienvenue",
        "Status": "false"
    }
    headers = {}

    response = requests.request("POST", url, headers=headers, data=payload)

    return response.text

def delNotification():
    url = "http://localhost:8080/DelNotification"

    payload = {
        "UserID": "testguy01",
        "Title": "bienvenue",
    }
    headers = {}

    response = requests.request("POST", url, headers=headers, data=payload)

    return response.text

def listNotification(user):
    url = "http://localhost:8080/notification/"+ user

    payload = {}
    headers = {}

    response = requests.request("GET", url, headers=headers, data=payload)

    return response.text

# Useful functions

##############################################################################################
##############################################################################################
##############################################################################################

# Scenarii

def account_creation():
    if login() != '{"failed":"404"}':
        print("Not OK, the account shoudn't exist before registration")
        print("But I will try to delete the accounts")
        delete_user()
        exit(1)
    if register() != '{"success":"200"}':
        print("ERROR : Registration")
        return 1
    if login() != '{"success":"testguy01"}':
        print("ERROR : Login")
        return 1
    print("SUCCESS : Account registration ok")
    delete_user()
    

def profile_edition():
    register()
    update_profile()
    if getProfile() != '{"profile":{"FullName":"MyFullNameTest","Description":"MyTestDescription","PhoneNb":"TestPhoneNB","Email":"myTest@email.com"}}':
        print("FAILED : Profile personalisation : ko")
        return 1
    delete_user()
    print("SUCCESS : Profile personalisation ok")


def friendship():
    registerFriends()
    # listFriends('testguy01')
    # listFriends('testguy02')
    addFriend('add')
    replyFriend('accept')
    a = listFriends('testguy01')
    b = listFriends('testguy02')
    deleteFriends()

    if a != '{"fetched":["testguy02"]}' or b != '{"fetched":["testguy01"]}':
        print("FAILED : Friendship : add : ko")
        print(a)
        print(b)
        return 1

    print('Success : Friendship : add : ok')


def refuse_friendship():
    registerFriends()
    addFriend('add')
    replyFriend('deny')
    a = listFriends('testguy01')
    b = listFriends('testguy02')
    deleteFriends()

    if '{"fetched":[]}' != a and a != b:
        print("FAILED : Refusing friendship : ko")
        return 1
    print('Success : Friendship : Deny : ok')


def delete_friendship():
    registerFriends()
    addFriend('add')
    replyFriend('accept')
    a = listFriends('testguy01')
    b = listFriends('testguy02')
    addFriend('rm')
    c = listFriends('testguy01')
    d = listFriends('testguy02')
    deleteFriends()

    if a != '{"fetched":["testguy02"]}' or b != '{"fetched":["testguy01"]}':
        print("FAILED : Friendship : add : ko")
        return 1
    if '{"fetched":[]}' != c and d != c:
        print("FAILED : Removing friendship : ko")
        return 1
    print("Success : Removing Frienship : ok")


def event_creation():
    registerFriends()
    addEvent()
    a = listEvent('testguy01')
    b = listEvent('testguy02')

    if a != b or a != '{"Success ":[{"Guests":"testguy01+testguy02","Date":"Demain soir","Subject":"bicyle","Confirmed":false}]}':
        print("FAILED : Event creation : ko")
        return 1
    deleteFriends()
    print("Success : Event creation : ok")


def password_check():
    register()
    login()
    c = editPassword()
    a = login()
    b = newIdslogin()

    delete_user()

    if a != '{"failed":"404"}' or b != '{"success":"testguy01"}' or c != '{"success":"200"}':
        print("FAILED : Error edit password : ko")
        return 1
    print("SUCCESS : edit password : ok")

def addnotification():
    register()
    a = addNotification()
    print(a)
    if not ("Success") in a:
        print("FAILED : Error add notification : ko")
        return 1
    print("SUCCESS : add notification : ok")
    delete_user()

def delnotification():
    register()
    addNotification()
    a = delNotification()
    if not ("Success") in a:
        print("FAILED : Error del notification : ko")
        return 1
    print("SUCCESS : del notification : ok")
    delete_user()

def getnotification():
    register()
    addNotification()
    a = listNotification("testguy01")
    if not ("Success") in a:
        print("FAILED : Error get notification : ko")
        return 1
    print("SUCCESS : get notification : ok")
    delete_user()


def addFeedback():
    url = "http://localhost:8080/feedback"

    payload = json.dumps({
        "Username": "testguy01",
        "Message": "J aime bien votre projet",
        "Date": "En ce jour"
    })
    headers = {
    'Content-Type': 'application/json'
    }

    response = requests.request("POST", url, headers=headers, data=payload)

    return response.text

def DelFeedback():
    url = "http://localhost:8080/delFeedback"

    payload = json.dumps({
        "Username": "testguy01",
        "Message": "J aime bien votre projet",
    })
    headers = {
    'Content-Type': 'application/json'
    }

    response = requests.request("POST", url, headers=headers, data=payload)

    return response.text


def getFeedback():
    url = "http://localhost:8080/feedback"

    payload = {}
    headers = {}

    response = requests.request("GET", url, headers=headers, data=payload)

    return response.text


def feedback():
    ExpectedOutput = '{"Username":"testguy01","Date":"En ce jour","Message":"J aime bien votre projet"}'
    addFeedback()
    resp = getFeedback()
    if ExpectedOutput in resp:
        print("YES")
        DelFeedback()
        pass
    else:
        print("FAILED : Feedback : ko")
        return 1
    # print(resp)



def test_all_scenari():
    account_creation()
    profile_edition()
    friendship()
    refuse_friendship()
    delete_friendship()
    event_creation()
    password_check()
    addnotification()
    # delnotification()
    # getnotification()
    feedback()
    pass



if __name__ == '__main__':
    test_all_scenari()  # Here should be a folder name
