import requests, json


def register():
    url = "http://20.234.168.103:8080/register/testguy01/safePasswo0d*/email@gmail.cam"

    payload = {}
    headers = {}

    response = requests.request("POST", url, headers=headers, data=payload)

    return response.text


def registerFriends():
    emails = ['email@gmail.cam', 'email@gmail.com']
    logins = ['testguy01', 'testguy02']
    i = 0

    while i != 2:
        url = "http://20.234.168.103:8080/register/"+logins[i]+"/safePasswo0d*/"+emails[i]
        payload = {}
        headers = {}
        requests.request("POST", url, headers=headers, data=payload)
        i += 1


def deleteFriends():
    logins = ['testguy01', 'testguy02']
    i = 0

    while i != 2:
        url = "http://20.234.168.103:8080/delete/"+logins[i]
        payload = {}
        headers = {}
        requests.request("POST", url, headers=headers, data=payload)
        i += 1


def login():
    url = "http://20.234.168.103:8080/login/testguy01/safePasswo0d*"

    payload = {}
    headers = {}

    response = requests.request("GET", url, headers=headers, data=payload)

    return response.text


def newIdslogin():
    url = "http://20.234.168.103:8080/login/testguy01/PassSaf3à*"

    payload = {}
    headers = {}

    response = requests.request("GET", url, headers=headers, data=payload)

    return response.text


def getProfile():
    url = "http://20.234.168.103:8080/profile/testguy01/"

    payload = {}
    headers = {}

    response = requests.request("GET", url, headers=headers, data=payload)

    return response.text


def delete_user():
    url = "http://20.234.168.103:8080/delete/testguy01"

    payload = {}
    headers = {}

    response = requests.request("POST", url, headers=headers, data=payload)

    return response.text  


def update_profile():
    endpoints = ['profileDescription', 'profileFullName', 'profilePhoneNB', 'profileEmail']
    data = ['MyTestDescription', 'MyFullNameTest', 'TestPhoneNB', 'myTest@email.com']
    i = 0
    while i < 4:
        url = "http://20.234.168.103:8080/"+endpoints[i]+"/testguy01/"+ data[i]

        payload = {}
        headers = {}

        requests.request("POST", url, headers=headers, data=payload)
        i += 1


def addFriend(action):
    url = "http://20.234.168.103:8080/manageFriend/testguy01/testguy02/"+action

    payload = {}
    headers = {}

    response = requests.request("POST", url, headers=headers, data=payload)

    return response.text


def replyFriend(action):
    url = "http://20.234.168.103:8080/replyFriend/testguy02/testguy01/"+ action

    payload = {}
    headers = {}

    response = requests.request("POST", url, headers=headers, data=payload)

    return response.text


def listFriends(user):
    url = "http://20.234.168.103:8080/listFriends/"+ user

    payload = {}
    headers = {}

    response = requests.request("GET", url, headers=headers, data=payload)

    return response.text


def addEvent():
    url = "http://20.234.168.103:8080/addEvent"

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
    url = "http://20.234.168.103:8080/listEvent/"+ user

    payload = {}
    headers = {}

    response = requests.request("GET", url, headers=headers, data=payload)

    return response.text


def editPassword():
    url = "http://20.234.168.103:8080/editPassword/testguy01/safePasswo0d*/PassSaf3à*"

    payload = {}
    headers = {}

    response = requests.request("POST", url, headers=headers, data=payload)

    return response.text


def addNotification():
    url = "http://20.234.168.103:8080/notification/testguy01/bienvenue/bienvenue/false"

    payload = {}
    headers = {}

    response = requests.request("POST", url, headers=headers, data=payload)

    return response.text

def delNotification():
    url = "http://20.234.168.103:8080/notification/testguy01/bienvenue/"

    payload = {}
    headers = {}

    response = requests.request("POST", url, headers=headers, data=payload)

    return response.text

def listNotification(user):
    url = "http://20.234.168.103:8080/notification/"+ user

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
        return 1
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
    listFriends('testguy01')
    listFriends('testguy02')
    addFriend('add')
    replyFriend('accept')
    a = listFriends('testguy01')
    b = listFriends('testguy02')
    deleteFriends()

    if a != '{"fetched":["testguy02"]}' or b != '{"fetched":["testguy01"]}':
        print("FAILED : Friendship : add : ko")
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

def test_all_scenari():
    # account_creation()
    # profile_edition()
    # friendship()
    # refuse_friendship()
    # delete_friendship()
    # event_creation()
    # password_check()
    # addnotification()
    # delnotification()
    # getnotification()
    pass



if __name__ == '__main__':
    test_all_scenari()  # Here should be a folder name
