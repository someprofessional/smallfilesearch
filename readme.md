**If you feel the strange urge to try it, please add a input and a output at the root of the project.**



# The purpose of this project is to learn document manipulation.

## By that i mean :
    - read / create / update and delete

## What i want to do is the following :
    1 - input some documents in a folder --done
    2 - have the programm do it's job --done
    3 - be able to search words or sentences --next 
    4 - add/remove the searched stored results
    5 - make a pdf including the fields that i want :
            -- authors / titles / docs / searched words|sentences / occurences / ...

## Other :
I want this app to be a cli app that do print into txt or pdf with a good interface so i'll make the visual with charm.


----------------------------------------------------------

### See inputs documents :  
    alice.pdf
    other.txt
    random paper.txt

### params:
    results width = 10 words

### Search :
    word or sentence here

### Action : 
    print to pdf / txt
    title :

    clear stored results

### Results :

    Title :                             Author :
    Date of creation :                  Type :

    Here is the //passage// that you searched .... [line 22] +
    Here is the //other// one .................... [line 750] +

    /////////////////////////////////////////////////////////
    Title :                             Author :
    Date of creation :                  Type :

    Here is the //passage// that you searched .... [line 22]
    Here is the //other// one .................... [line 750]

    /////////////////////////////////////////////////////////
    Title :                             Author :
    Date of creation :                  Type :

    Here is the //passage// that you searched .... [line 22]
    Here is the //other// one .................... [line 750] +

----------------------------------------------------------